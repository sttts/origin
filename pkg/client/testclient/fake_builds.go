package testclient

import (
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	core "k8s.io/client-go/testing"

	buildapi "github.com/openshift/origin/pkg/build/api"
)

// FakeBuilds implements BuildInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeBuilds struct {
	Fake      *Fake
	Namespace string
}

var buildsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "builds"}

func (c *FakeBuilds) Get(name string, options metav1.GetOptions) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(buildsResource, c.Namespace, name), &buildapi.Build{})
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) List(opts metainternal.ListOptions) (*buildapi.BuildList, error) {
	optsv1 := metav1.ListOptions{}
	err := metainternal.Convert_internalversion_ListOptions_To_v1_ListOptions(&opts, &optsv1, nil)
	if err != nil {
		return nil, err
	}
	obj, err := c.Fake.Invokes(core.NewListAction(buildsResource, c.Namespace, optsv1), &buildapi.BuildList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.BuildList), err
}

func (c *FakeBuilds) Create(inObj *buildapi.Build) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(buildsResource, c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) Update(inObj *buildapi.Build) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(buildsResource, c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewDeleteAction(buildsResource, c.Namespace, name), &buildapi.Build{})
	return err
}

func (c *FakeBuilds) Watch(opts metainternal.ListOptions) (watch.Interface, error) {
	optsv1 := metav1.ListOptions{}
	err := metainternal.Convert_internalversion_ListOptions_To_v1_ListOptions(&opts, &optsv1, nil)
	if err != nil {
		return nil, err
	}
	return c.Fake.InvokesWatch(core.NewWatchAction(buildsResource, c.Namespace, optsv1))
}

func (c *FakeBuilds) Clone(request *buildapi.BuildRequest) (result *buildapi.Build, err error) {
	action := core.NewCreateAction(buildsResource, c.Namespace, request)
	action.Subresource = "clone"
	obj, err := c.Fake.Invokes(action, &buildapi.Build{})
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}

func (c *FakeBuilds) UpdateDetails(inObj *buildapi.Build) (*buildapi.Build, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(buildapi.LegacySchemeGroupVersion.WithResource("builds/details"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*buildapi.Build), err
}
