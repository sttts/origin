package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
	"k8s.io/kubernetes/pkg/watch"

	buildapi "github.com/openshift/origin/pkg/build/api"
)

// FakeBuilds implements BuildInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeBuilds struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeBuilds) Get(name string) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(buildapi.SchemeGroupVersion.WithResource("builds"), c.Namespace, name), &buildapi.Build{})
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) List(opts kapi.ListOptions) (*buildapi.BuildList, error) {
	obj, err := c.Fake.Invokes(core.NewListAction(buildapi.SchemeGroupVersion.WithResource("builds"), c.Namespace, opts), &buildapi.BuildList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.BuildList), err
}

func (c *FakeBuilds) Create(inObj *buildapi.Build) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(buildapi.SchemeGroupVersion.WithResource("builds"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) Update(inObj *buildapi.Build) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(buildapi.SchemeGroupVersion.WithResource("builds"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewDeleteAction(buildapi.SchemeGroupVersion.WithResource("builds"), c.Namespace, name), &buildapi.Build{})
	return err
}

func (c *FakeBuilds) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(core.NewWatchAction(buildapi.SchemeGroupVersion.WithResource("builds"), c.Namespace, opts))
}

func (c *FakeBuilds) Clone(request *buildapi.BuildRequest) (result *buildapi.Build, err error) {
	action := core.NewCreateAction(buildapi.SchemeGroupVersion.WithResource("builds"), c.Namespace, request)
	action.Subresource = "clone"
	obj, err := c.Fake.Invokes(action, &buildapi.Build{})
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) UpdateDetails(inObj *buildapi.Build) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(buildapi.SchemeGroupVersion.WithResource("builds/details"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}
