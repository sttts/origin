package testclient

import (
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	core "k8s.io/client-go/testing"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

// FakePolicyBindings implements PolicyBindingInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakePolicyBindings struct {
	Fake      *Fake
	Namespace string
}

var policyBindingsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "policybindings"}

func (c *FakePolicyBindings) Get(name string, options metav1.GetOptions) (*authorizationapi.PolicyBinding, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(policyBindingsResource, c.Namespace, name), &authorizationapi.PolicyBinding{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.PolicyBinding), err
}

func (c *FakePolicyBindings) List(opts metainternal.ListOptions) (*authorizationapi.PolicyBindingList, error) {
	obj, err := c.Fake.Invokes(core.NewListAction(policyBindingsResource, c.Namespace, opts), &authorizationapi.PolicyBindingList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.PolicyBindingList), err
}

func (c *FakePolicyBindings) Create(inObj *authorizationapi.PolicyBinding) (*authorizationapi.PolicyBinding, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(policyBindingsResource, c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.PolicyBinding), err
}

func (c *FakePolicyBindings) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewDeleteAction(policyBindingsResource, c.Namespace, name), &authorizationapi.PolicyBinding{})
	return err
}

func (c *FakePolicyBindings) Watch(opts metainternal.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(core.NewWatchAction(policyBindingsResource, c.Namespace, opts))
}
