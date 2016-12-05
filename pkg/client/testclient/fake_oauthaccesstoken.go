package testclient

import (
	oauthapi "github.com/openshift/origin/pkg/oauth/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

// FakeOAuthAccessTokens implements OAuthAccessTokenInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeOAuthAccessTokens struct {
	Fake *Fake
}

func (c *FakeOAuthAccessTokens) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewRootDeleteAction(oauthapi.SchemeGroupVersion.WithResource("oauthaccesstokens"), name), &oauthapi.OAuthAccessToken{})
	return err
}

func (c *FakeOAuthAccessTokens) Create(inObj *oauthapi.OAuthAccessToken) (*oauthapi.OAuthAccessToken, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(oauthapi.SchemeGroupVersion.WithResource("oauthaccesstokens"), inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*oauthapi.OAuthAccessToken), err
}

// Get returns information about a particular image and error if one occurs.
func (c *FakeOAuthAccessTokens) Get(name string) (*oauthapi.OAuthAccessToken, error) {
	obj, err := c.Fake.Invokes(core.NewRootGetAction(oauthapi.SchemeGroupVersion.WithResource("oauthaccesstokens"), name), &oauthapi.OAuthAccessToken{})
	if obj == nil {
		return nil, err
	}

	return obj.(*oauthapi.OAuthAccessToken), err
}
