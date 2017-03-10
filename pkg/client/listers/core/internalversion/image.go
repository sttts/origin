// This file was automatically generated by lister-gen with arguments: --input-dirs=[github.com/openshift/origin/pkg/authorization/api,github.com/openshift/origin/pkg/authorization/api/v1,github.com/openshift/origin/pkg/build/api,github.com/openshift/origin/pkg/build/api/v1,github.com/openshift/origin/pkg/deploy/api,github.com/openshift/origin/pkg/deploy/api/v1,github.com/openshift/origin/pkg/image/api,github.com/openshift/origin/pkg/image/api/v1,github.com/openshift/origin/pkg/oauth/api,github.com/openshift/origin/pkg/oauth/api/v1,github.com/openshift/origin/pkg/project/api,github.com/openshift/origin/pkg/project/api/v1,github.com/openshift/origin/pkg/route/api,github.com/openshift/origin/pkg/route/api/v1,github.com/openshift/origin/pkg/sdn/api,github.com/openshift/origin/pkg/sdn/api/v1,github.com/openshift/origin/pkg/template/api,github.com/openshift/origin/pkg/template/api/v1,github.com/openshift/origin/pkg/user/api,github.com/openshift/origin/pkg/user/api/v1] --logtostderr=true

package internalversion

import (
	api "github.com/openshift/origin/pkg/image/api"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/kubernetes/pkg/client/cache"
)

// ImageLister helps list Images.
type ImageLister interface {
	// List lists all Images in the indexer.
	List(selector labels.Selector) (ret []*api.Image, err error)
	// Images returns an object that can list and get Images.
	Images(namespace string) ImageNamespaceLister
	ImageListerExpansion
}

// imageLister implements the ImageLister interface.
type imageLister struct {
	indexer cache.Indexer
}

// NewImageLister returns a new ImageLister.
func NewImageLister(indexer cache.Indexer) ImageLister {
	return &imageLister{indexer: indexer}
}

// List lists all Images in the indexer.
func (s *imageLister) List(selector labels.Selector) (ret []*api.Image, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*api.Image))
	})
	return ret, err
}

// Images returns an object that can list and get Images.
func (s *imageLister) Images(namespace string) ImageNamespaceLister {
	return imageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImageNamespaceLister helps list and get Images.
type ImageNamespaceLister interface {
	// List lists all Images in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*api.Image, err error)
	// Get retrieves the Image from the indexer for a given namespace and name.
	Get(name string) (*api.Image, error)
	ImageNamespaceListerExpansion
}

// imageNamespaceLister implements the ImageNamespaceLister
// interface.
type imageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Images in the indexer for a given namespace.
func (s imageNamespaceLister) List(selector labels.Selector) (ret []*api.Image, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*api.Image))
	})
	return ret, err
}

// Get retrieves the Image from the indexer for a given namespace and name.
func (s imageNamespaceLister) Get(name string) (*api.Image, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("image"), name)
	}
	return obj.(*api.Image), nil
}
