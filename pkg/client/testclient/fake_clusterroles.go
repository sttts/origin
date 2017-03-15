package testclient

import (
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	core "k8s.io/client-go/testing"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

// FakeClusterRoles implements ClusterRoleInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeClusterRoles struct {
	Fake *Fake
}

var clusterRolesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "clusterroles"}

func (c *FakeClusterRoles) Get(name string, options metav1.GetOptions) (*authorizationapi.ClusterRole, error) {
	obj, err := c.Fake.Invokes(core.NewRootGetAction(clusterRolesResource, name), &authorizationapi.ClusterRole{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRole), err
}

func (c *FakeClusterRoles) List(opts metainternal.ListOptions) (*authorizationapi.ClusterRoleList, error) {
	obj, err := c.Fake.Invokes(core.NewRootListAction(clusterRolesResource, opts), &authorizationapi.ClusterRoleList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRoleList), err
}

func (c *FakeClusterRoles) Create(inObj *authorizationapi.ClusterRole) (*authorizationapi.ClusterRole, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(clusterRolesResource, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRole), err
}

func (c *FakeClusterRoles) Update(inObj *authorizationapi.ClusterRole) (*authorizationapi.ClusterRole, error) {
	obj, err := c.Fake.Invokes(core.NewRootUpdateAction(clusterRolesResource, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*authorizationapi.ClusterRole), err
}

func (c *FakeClusterRoles) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewRootDeleteAction(clusterRolesResource, name), &authorizationapi.ClusterRole{})
	return err
}
