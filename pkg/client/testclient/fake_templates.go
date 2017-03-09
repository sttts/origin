package testclient

import (
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	clientgotesting "k8s.io/client-go/testing"

	templateapi "github.com/openshift/origin/pkg/template/api"
)

// FakeTemplates implements TemplateInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeTemplates struct {
	Fake      *Fake
	Namespace string
}

var templatesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "templates"}

func (c *FakeTemplates) Get(name string) (*templateapi.Template, error) {
	obj, err := c.Fake.Invokes(clientgotesting.NewGetAction(templatesResource, c.Namespace, name), &templateapi.Template{})
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.Template), err
}

func (c *FakeTemplates) List(opts metainternal.ListOptions) (*templateapi.TemplateList, error) {
	obj, err := c.Fake.Invokes(clientgotesting.NewListAction(templatesResource, c.Namespace, opts), &templateapi.TemplateList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.TemplateList), err
}

func (c *FakeTemplates) Create(inObj *templateapi.Template) (*templateapi.Template, error) {
	obj, err := c.Fake.Invokes(clientgotesting.NewCreateAction(templatesResource, c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.Template), err
}

func (c *FakeTemplates) Update(inObj *templateapi.Template) (*templateapi.Template, error) {
	obj, err := c.Fake.Invokes(clientgotesting.NewUpdateAction(templatesResource, c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.Template), err
}

func (c *FakeTemplates) Delete(name string) error {
	_, err := c.Fake.Invokes(clientgotesting.NewDeleteAction(templatesResource, c.Namespace, name), &templateapi.Template{})
	return err
}

func (c *FakeTemplates) Watch(opts metainternal.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(clientgotesting.NewWatchAction(templatesResource, c.Namespace, opts))
}
