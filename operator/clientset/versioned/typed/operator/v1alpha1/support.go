// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SupportsGetter has a method to return a SupportInterface.
// A group's client should implement this interface.
type SupportsGetter interface {
	Supports() SupportInterface
}

// SupportInterface has methods to work with Support resources.
type SupportInterface interface {
	Create(*v1alpha1.Support) (*v1alpha1.Support, error)
	Update(*v1alpha1.Support) (*v1alpha1.Support, error)
	UpdateStatus(*v1alpha1.Support) (*v1alpha1.Support, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Support, error)
	List(opts v1.ListOptions) (*v1alpha1.SupportList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Support, err error)
	SupportExpansion
}

// supports implements SupportInterface
type supports struct {
	client rest.Interface
}

// newSupports returns a Supports
func newSupports(c *OperatorV1alpha1Client) *supports {
	return &supports{
		client: c.RESTClient(),
	}
}

// Get takes name of the support, and returns the corresponding support object, and an error if there is any.
func (c *supports) Get(name string, options v1.GetOptions) (result *v1alpha1.Support, err error) {
	result = &v1alpha1.Support{}
	err = c.client.Get().
		Resource("supports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Supports that match those selectors.
func (c *supports) List(opts v1.ListOptions) (result *v1alpha1.SupportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SupportList{}
	err = c.client.Get().
		Resource("supports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested supports.
func (c *supports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("supports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a support and creates it.  Returns the server's representation of the support, and an error, if there is any.
func (c *supports) Create(support *v1alpha1.Support) (result *v1alpha1.Support, err error) {
	result = &v1alpha1.Support{}
	err = c.client.Post().
		Resource("supports").
		Body(support).
		Do().
		Into(result)
	return
}

// Update takes the representation of a support and updates it. Returns the server's representation of the support, and an error, if there is any.
func (c *supports) Update(support *v1alpha1.Support) (result *v1alpha1.Support, err error) {
	result = &v1alpha1.Support{}
	err = c.client.Put().
		Resource("supports").
		Name(support.Name).
		Body(support).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *supports) UpdateStatus(support *v1alpha1.Support) (result *v1alpha1.Support, err error) {
	result = &v1alpha1.Support{}
	err = c.client.Put().
		Resource("supports").
		Name(support.Name).
		SubResource("status").
		Body(support).
		Do().
		Into(result)
	return
}

// Delete takes name of the support and deletes it. Returns an error if one occurs.
func (c *supports) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("supports").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *supports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("supports").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched support.
func (c *supports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Support, err error) {
	result = &v1alpha1.Support{}
	err = c.client.Patch(pt).
		Resource("supports").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
