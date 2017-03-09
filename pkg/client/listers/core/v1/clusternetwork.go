// This file was automatically generated by lister-gen with arguments: --input-dirs=[github.com/openshift/origin/pkg/authorization/api,github.com/openshift/origin/pkg/authorization/api/v1,github.com/openshift/origin/pkg/build/api,github.com/openshift/origin/pkg/build/api/v1,github.com/openshift/origin/pkg/deploy/api,github.com/openshift/origin/pkg/deploy/api/v1,github.com/openshift/origin/pkg/image/api,github.com/openshift/origin/pkg/image/api/v1,github.com/openshift/origin/pkg/oauth/api,github.com/openshift/origin/pkg/oauth/api/v1,github.com/openshift/origin/pkg/project/api,github.com/openshift/origin/pkg/project/api/v1,github.com/openshift/origin/pkg/route/api,github.com/openshift/origin/pkg/route/api/v1,github.com/openshift/origin/pkg/sdn/api,github.com/openshift/origin/pkg/sdn/api/v1,github.com/openshift/origin/pkg/template/api,github.com/openshift/origin/pkg/template/api/v1,github.com/openshift/origin/pkg/user/api,github.com/openshift/origin/pkg/user/api/v1] --logtostderr=true

package v1

import (
	api "github.com/openshift/origin/pkg/sdn/api"
	v1 "github.com/openshift/origin/pkg/sdn/api/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// ClusterNetworkLister helps list ClusterNetworks.
type ClusterNetworkLister interface {
	// List lists all ClusterNetworks in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterNetwork, err error)
	// ClusterNetworks returns an object that can list and get ClusterNetworks.
	ClusterNetworks(namespace string) ClusterNetworkNamespaceLister
	ClusterNetworkListerExpansion
}

// clusterNetworkLister implements the ClusterNetworkLister interface.
type clusterNetworkLister struct {
	indexer cache.Indexer
}

// NewClusterNetworkLister returns a new ClusterNetworkLister.
func NewClusterNetworkLister(indexer cache.Indexer) ClusterNetworkLister {
	return &clusterNetworkLister{indexer: indexer}
}

// List lists all ClusterNetworks in the indexer.
func (s *clusterNetworkLister) List(selector labels.Selector) (ret []*v1.ClusterNetwork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterNetwork))
	})
	return ret, err
}

// ClusterNetworks returns an object that can list and get ClusterNetworks.
func (s *clusterNetworkLister) ClusterNetworks(namespace string) ClusterNetworkNamespaceLister {
	return clusterNetworkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterNetworkNamespaceLister helps list and get ClusterNetworks.
type ClusterNetworkNamespaceLister interface {
	// List lists all ClusterNetworks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ClusterNetwork, err error)
	// Get retrieves the ClusterNetwork from the indexer for a given namespace and name.
	Get(name string) (*v1.ClusterNetwork, error)
	ClusterNetworkNamespaceListerExpansion
}

// clusterNetworkNamespaceLister implements the ClusterNetworkNamespaceLister
// interface.
type clusterNetworkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterNetworks in the indexer for a given namespace.
func (s clusterNetworkNamespaceLister) List(selector labels.Selector) (ret []*v1.ClusterNetwork, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterNetwork))
	})
	return ret, err
}

// Get retrieves the ClusterNetwork from the indexer for a given namespace and name.
func (s clusterNetworkNamespaceLister) Get(name string) (*v1.ClusterNetwork, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("clusternetwork"), name)
	}
	return obj.(*v1.ClusterNetwork), nil
}
