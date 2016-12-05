package testclient

import (
	oauthapi "github.com/openshift/origin/pkg/oauth/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeOAuthAuthorizeTokens struct {
	Fake *Fake
}

func (c *FakeOAuthAuthorizeTokens) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewRootDeleteAction(oauthapi.SchemeGroupVersion.WithResource("oauthauthorizetokens"), name), &oauthapi.OAuthAuthorizeToken{})
	return err
}

func (c *FakeOAuthAuthorizeTokens) Create(inObj *oauthapi.OAuthAuthorizeToken) (*oauthapi.OAuthAuthorizeToken, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(oauthapi.SchemeGroupVersion.WithResource("oauthauthorizetokens"), inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*oauthapi.OAuthAuthorizeToken), err
}
