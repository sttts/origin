// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/openshift/origin/pkg/oauth/generated/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// OAuthAccessTokens returns a OAuthAccessTokenInformer.
	OAuthAccessTokens() OAuthAccessTokenInformer
	// OAuthAuthorizeTokens returns a OAuthAuthorizeTokenInformer.
	OAuthAuthorizeTokens() OAuthAuthorizeTokenInformer
	// OAuthClients returns a OAuthClientInformer.
	OAuthClients() OAuthClientInformer
	// OAuthClientAuthorizations returns a OAuthClientAuthorizationInformer.
	OAuthClientAuthorizations() OAuthClientAuthorizationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// OAuthAccessTokens returns a OAuthAccessTokenInformer.
func (v *version) OAuthAccessTokens() OAuthAccessTokenInformer {
	return &oAuthAccessTokenInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OAuthAuthorizeTokens returns a OAuthAuthorizeTokenInformer.
func (v *version) OAuthAuthorizeTokens() OAuthAuthorizeTokenInformer {
	return &oAuthAuthorizeTokenInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OAuthClients returns a OAuthClientInformer.
func (v *version) OAuthClients() OAuthClientInformer {
	return &oAuthClientInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OAuthClientAuthorizations returns a OAuthClientAuthorizationInformer.
func (v *version) OAuthClientAuthorizations() OAuthClientAuthorizationInformer {
	return &oAuthClientAuthorizationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
