// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package authorization

import (
	unsafe "unsafe"

	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	if in.Content == nil {
		out.Content = nil
	} else {
		out.Content = in.Content.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicy) DeepCopyInto(out *ClusterPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.LastModified.DeepCopyInto(&out.LastModified)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make(ClusterRolesByName, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(ClusterRole)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicy.
func (in *ClusterPolicy) DeepCopy() *ClusterPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPolicy) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicyBinding) DeepCopyInto(out *ClusterPolicyBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.LastModified.DeepCopyInto(&out.LastModified)
	out.PolicyRef = in.PolicyRef
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make(ClusterRoleBindingsByName, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(ClusterRoleBinding)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicyBinding.
func (in *ClusterPolicyBinding) DeepCopy() *ClusterPolicyBinding {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicyBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPolicyBinding) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicyBindingList) DeepCopyInto(out *ClusterPolicyBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterPolicyBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicyBindingList.
func (in *ClusterPolicyBindingList) DeepCopy() *ClusterPolicyBindingList {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicyBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPolicyBindingList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPolicyList) DeepCopyInto(out *ClusterPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPolicyList.
func (in *ClusterPolicyList) DeepCopy() *ClusterPolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPolicyList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRole) DeepCopyInto(out *ClusterRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRole.
func (in *ClusterRole) DeepCopy() *ClusterRole {
	if in == nil {
		return nil
	}
	out := new(ClusterRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRole) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleBinding) DeepCopyInto(out *ClusterRoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]core.ObjectReference, len(*in))
		copy(*out, *in)
	}
	out.RoleRef = in.RoleRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleBinding.
func (in *ClusterRoleBinding) DeepCopy() *ClusterRoleBinding {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRoleBinding) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleBindingList) DeepCopyInto(out *ClusterRoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleBindingList.
func (in *ClusterRoleBindingList) DeepCopy() *ClusterRoleBindingList {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRoleBindingList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleBindingsByName) DeepCopyInto(out *ClusterRoleBindingsByName) {
	{
		in := (*map[string]*ClusterRoleBinding)(unsafe.Pointer(in))
		out := (*map[string]*ClusterRoleBinding)(unsafe.Pointer(out))
		*out = make(map[string]*ClusterRoleBinding, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(ClusterRoleBinding)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleBindingsByName.
func (in *ClusterRoleBindingsByName) DeepCopy() *ClusterRoleBindingsByName {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleBindingsByName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleList) DeepCopyInto(out *ClusterRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleList.
func (in *ClusterRoleList) DeepCopy() *ClusterRoleList {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRoleList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRolesByName) DeepCopyInto(out *ClusterRolesByName) {
	{
		in := (*map[string]*ClusterRole)(unsafe.Pointer(in))
		out := (*map[string]*ClusterRole)(unsafe.Pointer(out))
		*out = make(map[string]*ClusterRole, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(ClusterRole)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRolesByName.
func (in *ClusterRolesByName) DeepCopy() *ClusterRolesByName {
	if in == nil {
		return nil
	}
	out := new(ClusterRolesByName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupRestriction) DeepCopyInto(out *GroupRestriction) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupRestriction.
func (in *GroupRestriction) DeepCopy() *GroupRestriction {
	if in == nil {
		return nil
	}
	out := new(GroupRestriction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IsPersonalSubjectAccessReview) DeepCopyInto(out *IsPersonalSubjectAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IsPersonalSubjectAccessReview.
func (in *IsPersonalSubjectAccessReview) DeepCopy() *IsPersonalSubjectAccessReview {
	if in == nil {
		return nil
	}
	out := new(IsPersonalSubjectAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IsPersonalSubjectAccessReview) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalResourceAccessReview) DeepCopyInto(out *LocalResourceAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Action.DeepCopyInto(&out.Action)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalResourceAccessReview.
func (in *LocalResourceAccessReview) DeepCopy() *LocalResourceAccessReview {
	if in == nil {
		return nil
	}
	out := new(LocalResourceAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalResourceAccessReview) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSubjectAccessReview) DeepCopyInto(out *LocalSubjectAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Action.DeepCopyInto(&out.Action)
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSubjectAccessReview.
func (in *LocalSubjectAccessReview) DeepCopy() *LocalSubjectAccessReview {
	if in == nil {
		return nil
	}
	out := new(LocalSubjectAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalSubjectAccessReview) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.LastModified.DeepCopyInto(&out.LastModified)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make(RolesByName, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(Role)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyBinding) DeepCopyInto(out *PolicyBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.LastModified.DeepCopyInto(&out.LastModified)
	out.PolicyRef = in.PolicyRef
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make(RoleBindingsByName, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(RoleBinding)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyBinding.
func (in *PolicyBinding) DeepCopy() *PolicyBinding {
	if in == nil {
		return nil
	}
	out := new(PolicyBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyBinding) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyBindingList) DeepCopyInto(out *PolicyBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyBindingList.
func (in *PolicyBindingList) DeepCopy() *PolicyBindingList {
	if in == nil {
		return nil
	}
	out := new(PolicyBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyBindingList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyBindingSorter) DeepCopyInto(out *PolicyBindingSorter) {
	{
		in := (*[]PolicyBinding)(unsafe.Pointer(in))
		out := (*[]PolicyBinding)(unsafe.Pointer(out))
		*out = make([]PolicyBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyBindingSorter.
func (in *PolicyBindingSorter) DeepCopy() *PolicyBindingSorter {
	if in == nil {
		return nil
	}
	out := new(PolicyBindingSorter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRule) DeepCopyInto(out *PolicyRule) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AttributeRestrictions == nil {
		out.AttributeRestrictions = nil
	} else {
		out.AttributeRestrictions = in.AttributeRestrictions.DeepCopyObject()
	}
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NonResourceURLs != nil {
		in, out := &in.NonResourceURLs, &out.NonResourceURLs
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRule.
func (in *PolicyRule) DeepCopy() *PolicyRule {
	if in == nil {
		return nil
	}
	out := new(PolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRuleBuilder) DeepCopyInto(out *PolicyRuleBuilder) {
	*out = *in
	in.PolicyRule.DeepCopyInto(&out.PolicyRule)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRuleBuilder.
func (in *PolicyRuleBuilder) DeepCopy() *PolicyRuleBuilder {
	if in == nil {
		return nil
	}
	out := new(PolicyRuleBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceAccessReview) DeepCopyInto(out *ResourceAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Action.DeepCopyInto(&out.Action)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceAccessReview.
func (in *ResourceAccessReview) DeepCopy() *ResourceAccessReview {
	if in == nil {
		return nil
	}
	out := new(ResourceAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceAccessReview) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceAccessReviewResponse) DeepCopyInto(out *ResourceAccessReviewResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceAccessReviewResponse.
func (in *ResourceAccessReviewResponse) DeepCopy() *ResourceAccessReviewResponse {
	if in == nil {
		return nil
	}
	out := new(ResourceAccessReviewResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceAccessReviewResponse) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Role) DeepCopyInto(out *Role) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Role.
func (in *Role) DeepCopy() *Role {
	if in == nil {
		return nil
	}
	out := new(Role)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Role) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBinding) DeepCopyInto(out *RoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]core.ObjectReference, len(*in))
		copy(*out, *in)
	}
	out.RoleRef = in.RoleRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBinding.
func (in *RoleBinding) DeepCopy() *RoleBinding {
	if in == nil {
		return nil
	}
	out := new(RoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleBinding) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingList) DeepCopyInto(out *RoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingList.
func (in *RoleBindingList) DeepCopy() *RoleBindingList {
	if in == nil {
		return nil
	}
	out := new(RoleBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleBindingList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingRestriction) DeepCopyInto(out *RoleBindingRestriction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingRestriction.
func (in *RoleBindingRestriction) DeepCopy() *RoleBindingRestriction {
	if in == nil {
		return nil
	}
	out := new(RoleBindingRestriction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleBindingRestriction) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingRestrictionList) DeepCopyInto(out *RoleBindingRestrictionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleBindingRestriction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingRestrictionList.
func (in *RoleBindingRestrictionList) DeepCopy() *RoleBindingRestrictionList {
	if in == nil {
		return nil
	}
	out := new(RoleBindingRestrictionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleBindingRestrictionList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingRestrictionSpec) DeepCopyInto(out *RoleBindingRestrictionSpec) {
	*out = *in
	if in.UserRestriction != nil {
		in, out := &in.UserRestriction, &out.UserRestriction
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserRestriction)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.GroupRestriction != nil {
		in, out := &in.GroupRestriction, &out.GroupRestriction
		if *in == nil {
			*out = nil
		} else {
			*out = new(GroupRestriction)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceAccountRestriction != nil {
		in, out := &in.ServiceAccountRestriction, &out.ServiceAccountRestriction
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceAccountRestriction)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingRestrictionSpec.
func (in *RoleBindingRestrictionSpec) DeepCopy() *RoleBindingRestrictionSpec {
	if in == nil {
		return nil
	}
	out := new(RoleBindingRestrictionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingSorter) DeepCopyInto(out *RoleBindingSorter) {
	{
		in := (*[]RoleBinding)(unsafe.Pointer(in))
		out := (*[]RoleBinding)(unsafe.Pointer(out))
		*out = make([]RoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingSorter.
func (in *RoleBindingSorter) DeepCopy() *RoleBindingSorter {
	if in == nil {
		return nil
	}
	out := new(RoleBindingSorter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingsByName) DeepCopyInto(out *RoleBindingsByName) {
	{
		in := (*map[string]*RoleBinding)(unsafe.Pointer(in))
		out := (*map[string]*RoleBinding)(unsafe.Pointer(out))
		*out = make(map[string]*RoleBinding, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(RoleBinding)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingsByName.
func (in *RoleBindingsByName) DeepCopy() *RoleBindingsByName {
	if in == nil {
		return nil
	}
	out := new(RoleBindingsByName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleList) DeepCopyInto(out *RoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Role, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleList.
func (in *RoleList) DeepCopy() *RoleList {
	if in == nil {
		return nil
	}
	out := new(RoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolesByName) DeepCopyInto(out *RolesByName) {
	{
		in := (*map[string]*Role)(unsafe.Pointer(in))
		out := (*map[string]*Role)(unsafe.Pointer(out))
		*out = make(map[string]*Role, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(Role)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolesByName.
func (in *RolesByName) DeepCopy() *RolesByName {
	if in == nil {
		return nil
	}
	out := new(RolesByName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectRulesReview) DeepCopyInto(out *SelfSubjectRulesReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectRulesReview.
func (in *SelfSubjectRulesReview) DeepCopy() *SelfSubjectRulesReview {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectRulesReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSubjectRulesReview) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectRulesReviewSpec) DeepCopyInto(out *SelfSubjectRulesReviewSpec) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectRulesReviewSpec.
func (in *SelfSubjectRulesReviewSpec) DeepCopy() *SelfSubjectRulesReviewSpec {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectRulesReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountReference) DeepCopyInto(out *ServiceAccountReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountReference.
func (in *ServiceAccountReference) DeepCopy() *ServiceAccountReference {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountRestriction) DeepCopyInto(out *ServiceAccountRestriction) {
	*out = *in
	if in.ServiceAccounts != nil {
		in, out := &in.ServiceAccounts, &out.ServiceAccounts
		*out = make([]ServiceAccountReference, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountRestriction.
func (in *ServiceAccountRestriction) DeepCopy() *ServiceAccountRestriction {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountRestriction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SortableRuleSlice) DeepCopyInto(out *SortableRuleSlice) {
	{
		in := (*[]PolicyRule)(unsafe.Pointer(in))
		out := (*[]PolicyRule)(unsafe.Pointer(out))
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SortableRuleSlice.
func (in *SortableRuleSlice) DeepCopy() *SortableRuleSlice {
	if in == nil {
		return nil
	}
	out := new(SortableRuleSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessReview) DeepCopyInto(out *SubjectAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Action.DeepCopyInto(&out.Action)
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessReview.
func (in *SubjectAccessReview) DeepCopy() *SubjectAccessReview {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubjectAccessReview) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessReviewResponse) DeepCopyInto(out *SubjectAccessReviewResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessReviewResponse.
func (in *SubjectAccessReviewResponse) DeepCopy() *SubjectAccessReviewResponse {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessReviewResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubjectAccessReviewResponse) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectRulesReview) DeepCopyInto(out *SubjectRulesReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectRulesReview.
func (in *SubjectRulesReview) DeepCopy() *SubjectRulesReview {
	if in == nil {
		return nil
	}
	out := new(SubjectRulesReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubjectRulesReview) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectRulesReviewSpec) DeepCopyInto(out *SubjectRulesReviewSpec) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectRulesReviewSpec.
func (in *SubjectRulesReviewSpec) DeepCopy() *SubjectRulesReviewSpec {
	if in == nil {
		return nil
	}
	out := new(SubjectRulesReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectRulesReviewStatus) DeepCopyInto(out *SubjectRulesReviewStatus) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectRulesReviewStatus.
func (in *SubjectRulesReviewStatus) DeepCopy() *SubjectRulesReviewStatus {
	if in == nil {
		return nil
	}
	out := new(SubjectRulesReviewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserRestriction) DeepCopyInto(out *UserRestriction) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserRestriction.
func (in *UserRestriction) DeepCopy() *UserRestriction {
	if in == nil {
		return nil
	}
	out := new(UserRestriction)
	in.DeepCopyInto(out)
	return out
}
