# Proposal for [Bug 1512042](https://bugzilla.redhat.com/show_bug.cgi?id=1512042)

## Introduction
The issue is that when a custom namespace (not openshift) is used for holding the images the broker does not currently give access to that namespace when creating the service account for the sandbox. This cause's the resulting transient namespace to get stuck in an error pulling the image. 

# Problem Description
The problem is that we want to be able to control how the privileges get escalated into that namespace to see/pull the images. This means that we will need a robust solution, that is not a simple bug fix.

## Implementation Details
The proposed changes are adding a configuration option to the local_registry. This configuration option will be similar to the auto_escalate feature we already have. The option will allow for checking if the originating user has access to the custom namespace to pull images. This will be the default, and you can specify to auto grant access to the namespace. There will only be two options for the configuration value `CheckAccessToNamespace` and `AutoGrantAccessToNamespace`.

To grant access to sandbox's service account we will grant access to pull images from the custom namespace. 

There is also a third option that we should document, which is to allow everyone to access the images in that custom namespace. This is purelly an RBAC policy setting and will allow all the users to access the images. 

### Open Questions
1. I think this check should be in the `pkg/handler/handler.go` file. This is an authorization check in my mind and should be treated as such. This check is only needed when an APB is being launched and should follow the placement of the current role authorization check.
2. The runtime package should be used to determine the check, something like `CheckAccessToImage`.
3. We need to figure out how to handle if the adapter is any other adapter then local_openshift. 
    1. Add a field on the spec to state the registry/adapter that it is using?
    2. We will also need to track the namespace that it should check.
    3. Need the ability to turn off the check when certain specs from other adapters are the ones being used.
    4. Need this to be usable by other adapters if they want (ecr adapter could use this to look at IAM access of the originating origin user, gcr adapter could use their Cloud IAM). Or a calller could want to do their own authorization check using our built in auth model (basic auth example).

### Possible Solutions
1. Create new fields on the spec. These fields will keep track of the adapters/registries that added the spec. Add a function in the runtime package that will be called by the handler for authorization. This function will take the spec, the originating user structure, and some undefined structures as a variadic. Example:
```go
type RuntimeAuthorization func(apb.Spec, apb.UserInfo, interface{}...)

var RuntimeAuthorization RuntimeAuthorization
```
2.  Create new fields on the spec so that it can track the adapters/registries that added this specific spec. Add an interface Authorization for the runtime package. This will have a single method, Authorize that will take a UserInfo struct. This then can be chained, and the runtime package should manage this as this is supposed to be exposed so everyone can add one. We can have helper methods to add providers to the authorization on the fly.
```go
type UserInfo {
    .... openservicebroker userinfo
}
type Authorization interface {
        Authorize(UserInfo) bool
} 

var authz Authorization

func AddAuthorization(a Authorization) {
    authz = chainAuthz(authz, a) 
}

func chainAuthz(a, b Authorization){
    return AuthorizationFunc(func(u UserInfo){
        if a(u) {
            return b(u)
        }
        return false
    })
}
type AuthorizationFunc func(UserInfo)

func (a AuthorizationFunc) Authorize(u UserInfo) {
    a(u)
}
```
Using it would be:
```go
func addProviderAuth(provider auth.Provider) runtime.Authorization {
    return runtime.AuthorizationFunc(func(u UserInfo) bool {
        principal, err := provider.GetPrincipal(r)
        if principal != nil {
            log.Debug("We found one. HOORAY!")
            // we found our principal, stop looking
            return true
        }
        if err != nil {
            log.Errorf("Error - %v", err)
        }
        return false
    })
}
.... 
runtime.AddAuthorization(addProviderAuth(provider))
...
```