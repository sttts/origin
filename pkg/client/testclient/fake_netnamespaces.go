package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/testing/core"
	"k8s.io/kubernetes/pkg/watch"

	sdnapi "github.com/openshift/origin/pkg/sdn/api"
)

// FakeNetNamespace implements NetNamespaceInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeNetNamespace struct {
	Fake *Fake
}

func (c *FakeNetNamespace) Get(name string) (*sdnapi.NetNamespace, error) {
	obj, err := c.Fake.Invokes(core.NewRootGetAction(sdnapi.SchemeGroupVersion.WithResource("netnamespaces"), name), &sdnapi.NetNamespace{})
	if obj == nil {
		return nil, err
	}

	return obj.(*sdnapi.NetNamespace), err
}

func (c *FakeNetNamespace) List(opts kapi.ListOptions) (*sdnapi.NetNamespaceList, error) {
	obj, err := c.Fake.Invokes(core.NewRootListAction(sdnapi.SchemeGroupVersion.WithResource("netnamespaces"), opts), &sdnapi.NetNamespaceList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*sdnapi.NetNamespaceList), err
}

func (c *FakeNetNamespace) Create(inObj *sdnapi.NetNamespace) (*sdnapi.NetNamespace, error) {
	obj, err := c.Fake.Invokes(core.NewRootCreateAction(sdnapi.SchemeGroupVersion.WithResource("netnamespaces"), inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*sdnapi.NetNamespace), err
}

func (c *FakeNetNamespace) Update(inObj *sdnapi.NetNamespace) (*sdnapi.NetNamespace, error) {
	obj, err := c.Fake.Invokes(core.NewRootUpdateAction(sdnapi.SchemeGroupVersion.WithResource("netnamespaces"), inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*sdnapi.NetNamespace), err
}

func (c *FakeNetNamespace) Delete(name string) error {
	_, err := c.Fake.Invokes(core.NewRootDeleteAction(sdnapi.SchemeGroupVersion.WithResource("netnamespaces"), name), &sdnapi.NetNamespace{})
	return err
}

func (c *FakeNetNamespace) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(core.NewRootWatchAction(sdnapi.SchemeGroupVersion.WithResource("netnamespaces"), opts))
}
