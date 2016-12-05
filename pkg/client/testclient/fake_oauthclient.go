package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
	"k8s.io/kubernetes/pkg/watch"

	oauthapi "github.com/openshift/origin/pkg/oauth/api"
)

type FakeOAuthClient struct {
	Fake *Fake
}

func (c *FakeOAuthClient) Get(name string) (*oauthapi.OAuthClient, error) {
	obj, err := c.Fake.Invokes(core.NewRootGetAction(oauthapi.SchemeGroupVersion.WithResource("oauthclients"), name), &oauthapi.OAuthClient{})
	if obj == nil {
		return nil, err
	}

	return obj.(*oauthapi.OAuthClient), err
}

func (c *FakeOAuthClient) List(opts kapi.ListOptions) (*oauthapi.OAuthClientList, error) {
	obj, err := c.Fake.Invokes(core.NewRootListAction(oauthapi.SchemeGroupVersion.WithResource("oauthclients"), opts), &oauthapi.OAuthClientList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*oauthapi.OAuthClientList), err
}

func (c *FakeOAuthClient) Create(inObj *oauthapi.OAuthClient) (*oauthapi.OAuthClient, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(oauthapi.SchemeGroupVersion.WithResource("oauthclients"), inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*oauthapi.OAuthClient), err
}

func (c *FakeOAuthClient) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewRootDeleteAction(oauthapi.SchemeGroupVersion.WithResource("oauthclients"), name), &oauthapi.OAuthClient{})
	return err
}

func (c *FakeOAuthClient) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(core.NewRootWatchAction(oauthapi.SchemeGroupVersion.WithResource("oauthclients"), opts))
}

func (c *FakeOAuthClient) Update(client *oauthapi.OAuthClient) (*oauthapi.OAuthClient, error) {
	obj, err := c.Fake.Invokes(core.NewRootUpdateAction(oauthapi.SchemeGroupVersion.WithResource("oauthclients"), client), &oauthapi.OAuthClient{})
	if obj == nil {
		return nil, err
	}

	return obj.(*oauthapi.OAuthClient), err
}
