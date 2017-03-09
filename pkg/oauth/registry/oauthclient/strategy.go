package oauthclient

import (
	"fmt"

	"github.com/openshift/origin/pkg/oauth/api"
	"github.com/openshift/origin/pkg/oauth/api/validation"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/labels"
	kstorage "k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

// strategy implements behavior for OAuthClient objects
type strategy struct {
	runtime.ObjectTyper
}

// Strategy is the default logic that applies when creating or updating OAuthClient objects
// objects via the REST API.
var Strategy = strategy{kapi.Scheme}

func (strategy) PrepareForUpdate(ctx kapi.Context, obj, old runtime.Object) {}

// NamespaceScoped is false for OAuth objects
func (strategy) NamespaceScoped() bool {
	return false
}

func (strategy) GenerateName(base string) string {
	return base
}

func (strategy) PrepareForCreate(ctx kapi.Context, obj runtime.Object) {
}

// Canonicalize normalizes the object after validation.
func (strategy) Canonicalize(obj runtime.Object) {
}

// Validate validates a new client
func (strategy) Validate(ctx kapi.Context, obj runtime.Object) field.ErrorList {
	token := obj.(*api.OAuthClient)
	return validation.ValidateClient(token)
}

// ValidateUpdate validates a client update
func (strategy) ValidateUpdate(ctx kapi.Context, obj runtime.Object, old runtime.Object) field.ErrorList {
	client := obj.(*api.OAuthClient)
	oldClient := old.(*api.OAuthClient)
	return validation.ValidateClientUpdate(client, oldClient)
}

// AllowCreateOnUpdate is false for OAuth objects
func (strategy) AllowCreateOnUpdate() bool {
	return false
}

func (strategy) AllowUnconditionalUpdate() bool {
	return false
}

// GetAttrs returns labels and fields of a given object for filtering purposes
func GetAttrs(o runtime.Object) (labels.Set, fields.Set, error) {
	obj, ok := o.(*api.OAuthClient)
	if !ok {
		return nil, nil, fmt.Errorf("not a OAuthClient")
	}
	return labels.Set(obj.Labels), SelectableFields(obj), nil
}

// Matcher returns a generic matcher for a given label and field selector.
func Matcher(label labels.Selector, field fields.Selector) kstorage.SelectionPredicate {
	return kstorage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that can be used for filter selection
func SelectableFields(obj *api.OAuthClient) fields.Set {
	return api.OAuthClientToSelectableFields(obj)
}
