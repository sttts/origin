package etcd

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"

	"github.com/openshift/origin/pkg/template/api"
	"github.com/openshift/origin/pkg/template/registry/template"
	"github.com/openshift/origin/pkg/util/restoptions"
)

// REST implements a RESTStorage for templates against etcd
type REST struct {
	*registry.Store
}

// NewREST returns a RESTStorage object that will work against templates.
func NewREST(optsGetter restoptions.Getter) (*REST, error) {
	store := &registry.Store{
		NewFunc:           func() runtime.Object { return &api.Template{} },
		NewListFunc:       func() runtime.Object { return &api.TemplateList{} },
		PredicateFunc:     template.Matcher,
		QualifiedResource: api.Resource("templates"),

		CreateStrategy: template.Strategy,
		UpdateStrategy: template.Strategy,

		ReturnDeletedObject: true,
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: tregistry.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}

	return &REST{store}, nil
}
