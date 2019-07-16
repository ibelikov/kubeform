/*
Copyright 2019 The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// SpannerInstanceIamBindingsGetter has a method to return a SpannerInstanceIamBindingInterface.
// A group's client should implement this interface.
type SpannerInstanceIamBindingsGetter interface {
	SpannerInstanceIamBindings() SpannerInstanceIamBindingInterface
}

// SpannerInstanceIamBindingInterface has methods to work with SpannerInstanceIamBinding resources.
type SpannerInstanceIamBindingInterface interface {
	Create(*v1alpha1.SpannerInstanceIamBinding) (*v1alpha1.SpannerInstanceIamBinding, error)
	Update(*v1alpha1.SpannerInstanceIamBinding) (*v1alpha1.SpannerInstanceIamBinding, error)
	UpdateStatus(*v1alpha1.SpannerInstanceIamBinding) (*v1alpha1.SpannerInstanceIamBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SpannerInstanceIamBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.SpannerInstanceIamBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerInstanceIamBinding, err error)
	SpannerInstanceIamBindingExpansion
}

// spannerInstanceIamBindings implements SpannerInstanceIamBindingInterface
type spannerInstanceIamBindings struct {
	client rest.Interface
}

// newSpannerInstanceIamBindings returns a SpannerInstanceIamBindings
func newSpannerInstanceIamBindings(c *GoogleV1alpha1Client) *spannerInstanceIamBindings {
	return &spannerInstanceIamBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the spannerInstanceIamBinding, and returns the corresponding spannerInstanceIamBinding object, and an error if there is any.
func (c *spannerInstanceIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.SpannerInstanceIamBinding, err error) {
	result = &v1alpha1.SpannerInstanceIamBinding{}
	err = c.client.Get().
		Resource("spannerinstanceiambindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpannerInstanceIamBindings that match those selectors.
func (c *spannerInstanceIamBindings) List(opts v1.ListOptions) (result *v1alpha1.SpannerInstanceIamBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpannerInstanceIamBindingList{}
	err = c.client.Get().
		Resource("spannerinstanceiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spannerInstanceIamBindings.
func (c *spannerInstanceIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("spannerinstanceiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a spannerInstanceIamBinding and creates it.  Returns the server's representation of the spannerInstanceIamBinding, and an error, if there is any.
func (c *spannerInstanceIamBindings) Create(spannerInstanceIamBinding *v1alpha1.SpannerInstanceIamBinding) (result *v1alpha1.SpannerInstanceIamBinding, err error) {
	result = &v1alpha1.SpannerInstanceIamBinding{}
	err = c.client.Post().
		Resource("spannerinstanceiambindings").
		Body(spannerInstanceIamBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a spannerInstanceIamBinding and updates it. Returns the server's representation of the spannerInstanceIamBinding, and an error, if there is any.
func (c *spannerInstanceIamBindings) Update(spannerInstanceIamBinding *v1alpha1.SpannerInstanceIamBinding) (result *v1alpha1.SpannerInstanceIamBinding, err error) {
	result = &v1alpha1.SpannerInstanceIamBinding{}
	err = c.client.Put().
		Resource("spannerinstanceiambindings").
		Name(spannerInstanceIamBinding.Name).
		Body(spannerInstanceIamBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *spannerInstanceIamBindings) UpdateStatus(spannerInstanceIamBinding *v1alpha1.SpannerInstanceIamBinding) (result *v1alpha1.SpannerInstanceIamBinding, err error) {
	result = &v1alpha1.SpannerInstanceIamBinding{}
	err = c.client.Put().
		Resource("spannerinstanceiambindings").
		Name(spannerInstanceIamBinding.Name).
		SubResource("status").
		Body(spannerInstanceIamBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the spannerInstanceIamBinding and deletes it. Returns an error if one occurs.
func (c *spannerInstanceIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("spannerinstanceiambindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spannerInstanceIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("spannerinstanceiambindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched spannerInstanceIamBinding.
func (c *spannerInstanceIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerInstanceIamBinding, err error) {
	result = &v1alpha1.SpannerInstanceIamBinding{}
	err = c.client.Patch(pt).
		Resource("spannerinstanceiambindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}