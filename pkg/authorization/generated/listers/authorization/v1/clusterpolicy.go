// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterPolicyLister helps list ClusterPolicies.
type ClusterPolicyLister interface {
	// List lists all ClusterPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterPolicy, err error)
	// Get retrieves the ClusterPolicy from the index for a given name.
	Get(name string) (*v1.ClusterPolicy, error)
	ClusterPolicyListerExpansion
}

// clusterPolicyLister implements the ClusterPolicyLister interface.
type clusterPolicyLister struct {
	indexer cache.Indexer
}

// NewClusterPolicyLister returns a new ClusterPolicyLister.
func NewClusterPolicyLister(indexer cache.Indexer) ClusterPolicyLister {
	return &clusterPolicyLister{indexer: indexer}
}

// List lists all ClusterPolicies in the indexer.
func (s *clusterPolicyLister) List(selector labels.Selector) (ret []*v1.ClusterPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterPolicy))
	})
	return ret, err
}

// Get retrieves the ClusterPolicy from the index for a given name.
func (s *clusterPolicyLister) Get(name string) (*v1.ClusterPolicy, error) {
	key := &v1.ClusterPolicy{ObjectMeta: meta_v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterpolicy"), name)
	}
	return obj.(*v1.ClusterPolicy), nil
}
