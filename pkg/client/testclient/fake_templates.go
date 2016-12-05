package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
	"k8s.io/kubernetes/pkg/watch"

	templateapi "github.com/openshift/origin/pkg/template/api"
)

// FakeTemplates implements TemplateInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeTemplates struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeTemplates) Get(name string) (*templateapi.Template, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(templateapi.SchemeGroupVersion.WithResource("templates"), c.Namespace, name), &templateapi.Template{})
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.Template), err
}

func (c *FakeTemplates) List(opts kapi.ListOptions) (*templateapi.TemplateList, error) {
	obj, err := c.Fake.Invokes(core.NewListAction(templateapi.SchemeGroupVersion.WithResource("templates"), c.Namespace, opts), &templateapi.TemplateList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.TemplateList), err
}

func (c *FakeTemplates) Create(inObj *templateapi.Template) (*templateapi.Template, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(templateapi.SchemeGroupVersion.WithResource("templates"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.Template), err
}

func (c *FakeTemplates) Update(inObj *templateapi.Template) (*templateapi.Template, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(templateapi.SchemeGroupVersion.WithResource("templates"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.Template), err
}

func (c *FakeTemplates) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewDeleteAction(templateapi.SchemeGroupVersion.WithResource("templates"), c.Namespace, name), &templateapi.Template{})
	return err
}

func (c *FakeTemplates) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(core.NewWatchAction(templateapi.SchemeGroupVersion.WithResource("templates"), c.Namespace, opts))
}
