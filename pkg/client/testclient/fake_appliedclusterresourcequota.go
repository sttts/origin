package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/testing/core"

	quotaapi "github.com/openshift/origin/pkg/quota/api"
)

type FakeAppliedClusterResourceQuotas struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeAppliedClusterResourceQuotas) Get(name string) (*quotaapi.AppliedClusterResourceQuota, error) {
	obj, err := c.Fake.Invokes(core.NewGetAction(quotaapi.SchemeGroupVersion.WithResource("appliedclusterresourcequotas"), c.Namespace, name), &quotaapi.AppliedClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}

	return obj.(*quotaapi.AppliedClusterResourceQuota), err
}

func (c *FakeAppliedClusterResourceQuotas) List(opts kapi.ListOptions) (*quotaapi.AppliedClusterResourceQuotaList, error) {
	obj, err := c.Fake.Invokes(core.NewListAction(quotaapi.SchemeGroupVersion.WithResource("appliedclusterresourcequotas"), c.Namespace, opts), &quotaapi.AppliedClusterResourceQuotaList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*quotaapi.AppliedClusterResourceQuotaList), err
}
