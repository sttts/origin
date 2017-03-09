package fake

import (
	internalversion "github.com/openshift/origin/pkg/template/client/clientset_generated/internalclientset/typed/core/internalversion"
	restclient "k8s.io/client-go/rest"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

type FakeCore struct {
	*core.Fake
}

func (c *FakeCore) Templates(namespace string) internalversion.TemplateInterface {
	return &FakeTemplates{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCore) RESTClient() restclient.Interface {
	var ret *restclient.RESTClient
	return ret
}
