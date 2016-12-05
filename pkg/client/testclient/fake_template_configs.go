package testclient

import (
	templateapi "github.com/openshift/origin/pkg/template/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

// FakeTemplateConfigs implements TemplateConfigsInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeTemplateConfigs struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeTemplateConfigs) Create(inObj *templateapi.Template) (*templateapi.Template, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(templateapi.SchemeGroupVersion.WithResource("templateconfigs"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*templateapi.Template), err
}
