package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/testing/core"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

// FakeRoles implements RoleInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeRoles struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeRoles) Get(name string) (*authorizationapi.Role, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(authorizationapi.SchemeGroupVersion.WithResource("roles"), c.Namespace, name), &authorizationapi.Role{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.Role), err
}

func (c *FakeRoles) List(opts kapi.ListOptions) (*authorizationapi.RoleList, error) {
	obj, err := c.Fake.Invokes(core.NewListAction(authorizationapi.SchemeGroupVersion.WithResource("roles"), c.Namespace, opts), &authorizationapi.RoleList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.RoleList), err
}

func (c *FakeRoles) Create(inObj *authorizationapi.Role) (*authorizationapi.Role, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(authorizationapi.SchemeGroupVersion.WithResource("roles"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.Role), err
}

func (c *FakeRoles) Update(inObj *authorizationapi.Role) (*authorizationapi.Role, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(authorizationapi.SchemeGroupVersion.WithResource("roles"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.Role), err
}

func (c *FakeRoles) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewDeleteAction(authorizationapi.SchemeGroupVersion.WithResource("roles"), c.Namespace, name), &authorizationapi.Role{})
	return err
}
