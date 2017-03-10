package etcd

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/kubernetes/pkg/registry/generic/registry"

	"github.com/openshift/origin/pkg/sdn/api"
	"github.com/openshift/origin/pkg/sdn/registry/clusternetwork"
	"github.com/openshift/origin/pkg/util/restoptions"
)

// rest implements a RESTStorage for sdn against etcd
type REST struct {
	*registry.Store
}

// NewREST returns a RESTStorage object that will work against subnets
func NewREST(optsGetter restoptions.Getter) (*REST, error) {
	store := &registry.Store{
		NewFunc:           func() runtime.Object { return &api.ClusterNetwork{} },
		NewListFunc:       func() runtime.Object { return &api.ClusterNetworkList{} },
		PredicateFunc:     clusternetwork.Matcher,
		QualifiedResource: api.Resource("clusternetworks"),

		CreateStrategy: clusternetwork.Strategy,
		UpdateStrategy: clusternetwork.Strategy,
	}

	// TODO this will be uncommented after 1.6 rebase:
	// options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: user.GetAttrs}
	// if err := store.CompleteWithOptions(options); err != nil {
	if err := restoptions.ApplyOptions(optsGetter, store, storage.NoTriggerPublisher); err != nil {
		return nil, err
	}

	return &REST{store}, nil
}
