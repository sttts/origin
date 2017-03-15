package testclient

import (
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	core "k8s.io/client-go/testing"

	projectapi "github.com/openshift/origin/pkg/project/api"
)

// FakeProjects implements ProjectInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeProjects struct {
	Fake *Fake
}

var projectsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "projects"}

func (c *FakeProjects) Get(name string, options metav1.GetOptions) (*projectapi.Project, error) {
	obj, err := c.Fake.Invokes(core.NewRootGetAction(projectsResource, name), &projectapi.Project{})
	if obj == nil {
		return nil, err
	}

	return obj.(*projectapi.Project), err
}

func (c *FakeProjects) List(opts metainternal.ListOptions) (*projectapi.ProjectList, error) {
	obj, err := c.Fake.Invokes(core.NewRootListAction(projectsResource, opts), &projectapi.ProjectList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*projectapi.ProjectList), err
}

func (c *FakeProjects) Create(inObj *projectapi.Project) (*projectapi.Project, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(projectsResource, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*projectapi.Project), err
}

func (c *FakeProjects) Update(inObj *projectapi.Project) (*projectapi.Project, error) {
	obj, err := c.Fake.Invokes(core.NewRootUpdateAction(projectsResource, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*projectapi.Project), err
}

func (c *FakeProjects) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewRootDeleteAction(projectsResource, name), &projectapi.Project{})
	return err
}

func (c *FakeProjects) Watch(opts metainternal.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(core.NewRootWatchAction(projectsResource, opts))
}
