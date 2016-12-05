package testclient

import (
	"github.com/openshift/origin/pkg/client"
	imageapi "github.com/openshift/origin/pkg/image/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
)

// FakeImageStreamTags implements ImageStreamTagInterface. Meant to be
// embedded into a struct to get a default implementation. This makes faking
// out just the methods you want to test easier.
type FakeImageStreamTags struct {
	Fake      *Fake
	Namespace string
}

var _ client.ImageStreamTagInterface = &FakeImageStreamTags{}

func (c *FakeImageStreamTags) Get(name, tag string) (*imageapi.ImageStreamTag, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(imageapi.SchemeGroupVersion.WithResource("imagestreamtags"), c.Namespace, imageapi.JoinImageStreamTag(name, tag)), &imageapi.ImageStreamTag{})
	if obj == nil {
		return nil, err
	}

	return obj.(*imageapi.ImageStreamTag), err
}

func (c *FakeImageStreamTags) Update(inObj *imageapi.ImageStreamTag) (*imageapi.ImageStreamTag, error) {
	obj, err := c.Fake.Invokes(core.NewUpdateAction(imageapi.SchemeGroupVersion.WithResource("imagestreamtags"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*imageapi.ImageStreamTag), err
}

func (c *FakeImageStreamTags) Create(inObj *imageapi.ImageStreamTag) (*imageapi.ImageStreamTag, error) {
	obj, err := c.Fake.Invokes(core.NewCreateAction(imageapi.SchemeGroupVersion.WithResource("imagestreamtags"), c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*imageapi.ImageStreamTag), err
}

func (c *FakeImageStreamTags) Delete(name, tag string) error {
	_, err := c.Fake.Invokes(core.NewDeleteAction(imageapi.SchemeGroupVersion.WithResource("imagestreamtags"), c.Namespace, imageapi.JoinImageStreamTag(name, tag)), &imageapi.ImageStreamTag{})
	return err
}
