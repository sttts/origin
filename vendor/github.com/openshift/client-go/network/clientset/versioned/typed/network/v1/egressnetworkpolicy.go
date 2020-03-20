// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openshift/api/network/v1"
	scheme "github.com/openshift/client-go/network/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EgressNetworkPoliciesGetter has a method to return a EgressNetworkPolicyInterface.
// A group's client should implement this interface.
type EgressNetworkPoliciesGetter interface {
	EgressNetworkPolicies(namespace string) EgressNetworkPolicyInterface
}

// EgressNetworkPolicyInterface has methods to work with EgressNetworkPolicy resources.
type EgressNetworkPolicyInterface interface {
	Create(ctx context.Context, egressNetworkPolicy *v1.EgressNetworkPolicy, opts metav1.CreateOptions) (*v1.EgressNetworkPolicy, error)
	Update(ctx context.Context, egressNetworkPolicy *v1.EgressNetworkPolicy, opts metav1.UpdateOptions) (*v1.EgressNetworkPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.EgressNetworkPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.EgressNetworkPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EgressNetworkPolicy, err error)
	EgressNetworkPolicyExpansion
}

// egressNetworkPolicies implements EgressNetworkPolicyInterface
type egressNetworkPolicies struct {
	client rest.Interface
	ns     string
}

// newEgressNetworkPolicies returns a EgressNetworkPolicies
func newEgressNetworkPolicies(c *NetworkV1Client, namespace string) *egressNetworkPolicies {
	return &egressNetworkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the egressNetworkPolicy, and returns the corresponding egressNetworkPolicy object, and an error if there is any.
func (c *egressNetworkPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EgressNetworkPolicies that match those selectors.
func (c *egressNetworkPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EgressNetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EgressNetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested egressNetworkPolicies.
func (c *egressNetworkPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a egressNetworkPolicy and creates it.  Returns the server's representation of the egressNetworkPolicy, and an error, if there is any.
func (c *egressNetworkPolicies) Create(ctx context.Context, egressNetworkPolicy *v1.EgressNetworkPolicy, opts metav1.CreateOptions) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(egressNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a egressNetworkPolicy and updates it. Returns the server's representation of the egressNetworkPolicy, and an error, if there is any.
func (c *egressNetworkPolicies) Update(ctx context.Context, egressNetworkPolicy *v1.EgressNetworkPolicy, opts metav1.UpdateOptions) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Name(egressNetworkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(egressNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the egressNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *egressNetworkPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *egressNetworkPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched egressNetworkPolicy.
func (c *egressNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
