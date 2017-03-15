package testclient

import (
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	"k8s.io/apimachinery/pkg/runtime/schema"
	core "k8s.io/client-go/testing"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

// FakeRoles implements RoleInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeRoles struct {
	Fake      *Fake
	Namespace string
}

var rolesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "roles"}

func (c *FakeRoles) Get(name string) (*authorizationapi.Role, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(rolesResource, c.Namespace, name), &authorizationapi.Role{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.Role), err
}

func (c *FakeRoles) List(opts metainternal.ListOptions) (*authorizationapi.RoleList, error) {
	obj, err := c.Fake.Invokes(core.NewListAction(rolesResource, c.Namespace, opts), &authorizationapi.RoleList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.RoleList), err
}

func (c *FakeRoles) Create(inObj *authorizationapi.Role) (*authorizationapi.Role, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(rolesResource, c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.Role), err
}

func (c *FakeRoles) Update(inObj *authorizationapi.Role) (*authorizationapi.Role, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(rolesResource, c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.Role), err
}

func (c *FakeRoles) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewDeleteAction(rolesResource, c.Namespace, name), &authorizationapi.Role{})
	return err
}
