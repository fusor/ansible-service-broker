package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"

	logging "github.com/op/go-logging"
	"github.com/openshift/ansible-service-broker/pkg/origin/copy/authorization"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/homedir"
	kapi "k8s.io/kubernetes/pkg/api"
	kapihelper "k8s.io/kubernetes/pkg/api/helper"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
)

/* Start of V1 Authorizaiont rules need for openshift rest call */
var oldAllowAllPolicyRule = PolicyRule{APIGroups: nil, Verbs: []string{"*"}, Resources: []string{"*"}}

// OpenshiftClient - Client to interact with openshift api
type OpenshiftClient struct {
	restClient rest.Interface
}

// SubjectRulesReview is a resource you can create to determine which actions another user can perform in a namespace
type SubjectRulesReview struct {
	metav1.TypeMeta `json:",inline"`

	// Spec adds information about how to conduct the check
	Spec SubjectRulesReviewSpec `json:"spec" protobuf:"bytes,1,opt,name=spec"`

	// Status is completed by the server to tell which permissions you have
	Status SubjectRulesReviewStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// SubjectRulesReviewSpec adds information about how to conduct the check
type SubjectRulesReviewSpec struct {
	// User is optional.  At least one of User and Groups must be specified.
	User string `json:"user" protobuf:"bytes,1,opt,name=user"`
	// Groups is optional.  Groups is the list of groups to which the User belongs.  At least one of User and Groups must be specified.
	Groups []string `json:"groups" protobuf:"bytes,2,rep,name=groups"`
	// Scopes to use for the evaluation.  Empty means "use the unscoped (full) permissions of the user/groups".
	Scopes []string `json:"scopes" protobuf:"bytes,3,opt,name=scopes"`
}

// SubjectRulesReviewStatus is contains the result of a rules check
type SubjectRulesReviewStatus struct {
	// Rules is the list of rules (no particular sort) that are allowed for the subject
	Rules []PolicyRule `json:"rules" protobuf:"bytes,1,rep,name=rules"`
	// EvaluationError can appear in combination with Rules.  It means some error happened during evaluation
	// that may have prevented additional rules from being populated.
	EvaluationError string `json:"evaluationError,omitempty" protobuf:"bytes,2,opt,name=evaluationError"`
}

// PolicyRule - the v1 Policy rule from openshift API
type PolicyRule struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
	Verbs []string `json:"verbs" protobuf:"bytes,1,rep,name=verbs"`
	// AttributeRestrictions will vary depending on what the Authorizer/AuthorizationAttributeBuilder pair supports.
	// If the Authorizer does not recognize how to handle the AttributeRestrictions, the Authorizer should report an error.
	AttributeRestrictions kruntime.RawExtension `json:"attributeRestrictions,omitempty" protobuf:"bytes,2,opt,name=attributeRestrictions"`
	// APIGroups is the name of the APIGroup that contains the resources.  If this field is empty, then both kubernetes and origin API groups are assumed.
	// That means that if an action is requested against one of the enumerated resources in either the kubernetes or the origin API group, the request
	// will be allowed
	APIGroups []string `json:"apiGroups" protobuf:"bytes,3,rep,name=apiGroups"`
	// Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
	Resources []string `json:"resources" protobuf:"bytes,4,rep,name=resources"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
	ResourceNames []string `json:"resourceNames,omitempty" protobuf:"bytes,5,rep,name=resourceNames"`
	// NonResourceURLsSlice is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path
	// This name is intentionally different than the internal type so that the DefaultConvert works nicely and because the ordering may be different.
	NonResourceURLsSlice []string `json:"nonResourceURLs,omitempty" protobuf:"bytes,6,rep,name=nonResourceURLs"`
}

// convert_runtime_RawExtension_To_runtime_Object attempts to convert an incoming object into the
// appropriate output type.
func convertruntimeRawExtensionToruntimeObject(c runtime.ObjectConvertor, in *runtime.RawExtension, out *runtime.Object, s conversion.Scope) error {
	if in == nil || in.Object == nil {
		return nil
	}

	switch in.Object.(type) {
	case *runtime.Unknown, *unstructured.Unstructured:
		*out = in.Object
		return nil
	}

	switch t := s.Meta().Context.(type) {
	case runtime.GroupVersioner:
		converted, err := c.ConvertToVersion(in.Object, t)
		if err != nil {
			return err
		}
		in.Object = converted
		*out = converted
	default:
		return fmt.Errorf("unrecognized conversion context for conversion to internal: %#v (%T)", t, t)
	}
	return nil
}

func convertPolicyRuleToAuthorizationPolicyRule(in *PolicyRule, out *authorization.PolicyRule, s conversion.Scope) error {
	setDefaultsPolicyRule(in)
	if err := convertruntimeRawExtensionToruntimeObject(kapi.Scheme, &in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}

	out.APIGroups = in.APIGroups

	out.Resources = sets.String{}
	out.Resources.Insert(in.Resources...)

	out.Verbs = sets.String{}
	out.Verbs.Insert(in.Verbs...)

	out.ResourceNames = sets.NewString(in.ResourceNames...)

	out.NonResourceURLs = sets.NewString(in.NonResourceURLsSlice...)

	return nil

}

func setDefaultsPolicyRule(obj *PolicyRule) {
	if obj == nil {
		return
	}

	// match the old allow all rule, but only if API groups is nil (not specified in the incoming JSON)
	oldAllowAllRule := obj.APIGroups == nil &&
		// avoid calling the very expensive DeepEqual by inlining specific checks
		len(obj.Verbs) == 1 && obj.Verbs[0] == "*" &&
		len(obj.Resources) == 1 && obj.Resources[0] == "*" &&
		len(obj.AttributeRestrictions.Raw) == 0 && len(obj.ResourceNames) == 0 &&
		len(obj.NonResourceURLsSlice) == 0 &&
		// semantic equalities will ignore nil vs empty for other fields as a safety
		// DO NOT REMOVE THIS CHECK unless you replace it with full equality comparisons
		kapihelper.Semantic.Equalities.DeepEqual(oldAllowAllPolicyRule, *obj)

	if oldAllowAllRule {
		obj.APIGroups = []string{"*"}
	}

	// if no groups are found, then we assume ""
	if len(obj.Resources) > 0 && len(obj.APIGroups) == 0 {
		obj.APIGroups = []string{""}
	}
}

func autoConvertv1PolicyRuleToauthorizationPolicyRule(in *PolicyRule, out *authorization.PolicyRule, s conversion.Scope) error {
	// WARNING: in.Verbs requires manual conversion: inconvertible types ([]string vs k8s.io/apimachinery/pkg/util/sets.String)
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	// WARNING: in.Resources requires manual conversion: inconvertible types ([]string vs k8s.io/apimachinery/pkg/util/sets.String)
	// WARNING: in.ResourceNames requires manual conversion: inconvertible types ([]string vs k8s.io/apimachinery/pkg/util/sets.String)
	// WARNING: in.NonResourceURLsSlice requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvertV1PolicyRulesToAuthorizationPolicyRules(in []PolicyRule) ([]authorization.PolicyRule, error) {
	var out []authorization.PolicyRule
	if in != nil {
		out = make([]authorization.PolicyRule, len(in))
		for i := range in {
			aRule := authorization.PolicyRule{}
			if err := convertPolicyRuleToAuthorizationPolicyRule(&in[i], &aRule, nil); err != nil {
				return nil, err
			}
			out[i] = aRule
		}
	} else {
		out = nil
	}
	return out, nil
}

/* End of v1 openshift rest calls that are need. */

// Openshift - Create a new openshift client if needed, returns reference
func Openshift(log *logging.Logger) (*OpenshiftClient, error) {
	errMsg := "Something went wrong while initializing kubernetes client!\n"
	once.Openshift.Do(func() {
		client, err := newOpenshift(log)
		if err != nil {
			log.Error(errMsg)
			// NOTE: Looking to leverage panic recovery to gracefully handle this
			// with things like retries or better intelligence, but the environment
			// is probably in a unrecoverable state as far as the broker is concerned,
			// and demands the attention of an operator.
			panic(err.Error())
		}
		instances.Openshift = client
	})
	if instances.Openshift == nil {
		return nil, errors.New("Kubernetes client instance is nil")
	}
	return instances.Openshift, nil
}

func newOpenshift(log *logging.Logger) (*OpenshiftClient, error) {
	// NOTE: Both the external and internal client object are using the same
	// clientset library. Internal clientset normally uses a different
	// library
	clientConfig, err := rest.InClusterConfig()
	if err != nil {
		log.Warning("Failed to create a InternalClientSet: %v.", err)

		log.Debug("Checking for a local Cluster Config")
		clientConfig, err = createClientConfigFromFile(homedir.HomeDir() + "/.kube/config")
		if err != nil {
			log.Error("Failed to create LocalClientSet")
			return nil, err
		}
	}

	clientset, err := newForConfig(clientConfig)
	if err != nil {
		log.Error("Failed to create LocalClientSet")
		return nil, err
	}

	return clientset, err
}

func newForConfig(c *rest.Config) (*OpenshiftClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &OpenshiftClient{restClient: client}, nil
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis/authorization.openshift.io"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// SubjectRulesReview - create and run a OpenShift Subject Rules Review
func (o OpenshiftClient) SubjectRulesReview(user, namespace string, log *logging.Logger) (result []rbac.PolicyRule, err error) {
	body := &SubjectRulesReview{
		Spec: SubjectRulesReviewSpec{
			User: user,
		},
	}
	body.Kind = "SubjectRulesReview"
	body.APIVersion = "authorization.openshift.io/v1"
	b, _ := json.Marshal(body)
	r := &SubjectRulesReview{}
	res, err := o.restClient.Post().
		Namespace(namespace).
		Resource("subjectrulesreviews").
		Body(b).
		DoRaw()
	err = json.Unmarshal(res, r)
	if err != nil {
		log.Errorf("error - %v\n unmarshall - %q", err, res)
		return
	}
	//Need to take the v1 Policy Rule and make it a Authorization Rule.
	pr, err := autoConvertV1PolicyRulesToAuthorizationPolicyRules(r.Status.Rules)
	if err != nil {
		return nil, err
	}
	return authorization.ConvertAPIPolicyRulesToRBACPolicyRules(pr), nil
}
