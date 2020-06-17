/*
Copyright The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AvailabilitySetsGetter has a method to return a AvailabilitySetInterface.
// A group's client should implement this interface.
type AvailabilitySetsGetter interface {
	AvailabilitySets(namespace string) AvailabilitySetInterface
}

// AvailabilitySetInterface has methods to work with AvailabilitySet resources.
type AvailabilitySetInterface interface {
	Create(*v1alpha1.AvailabilitySet) (*v1alpha1.AvailabilitySet, error)
	Update(*v1alpha1.AvailabilitySet) (*v1alpha1.AvailabilitySet, error)
	UpdateStatus(*v1alpha1.AvailabilitySet) (*v1alpha1.AvailabilitySet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AvailabilitySet, error)
	List(opts v1.ListOptions) (*v1alpha1.AvailabilitySetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AvailabilitySet, err error)
	AvailabilitySetExpansion
}

// availabilitySets implements AvailabilitySetInterface
type availabilitySets struct {
	client rest.Interface
	ns     string
}

// newAvailabilitySets returns a AvailabilitySets
func newAvailabilitySets(c *AzurermV1alpha1Client, namespace string) *availabilitySets {
	return &availabilitySets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the availabilitySet, and returns the corresponding availabilitySet object, and an error if there is any.
func (c *availabilitySets) Get(name string, options v1.GetOptions) (result *v1alpha1.AvailabilitySet, err error) {
	result = &v1alpha1.AvailabilitySet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("availabilitysets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AvailabilitySets that match those selectors.
func (c *availabilitySets) List(opts v1.ListOptions) (result *v1alpha1.AvailabilitySetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AvailabilitySetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("availabilitysets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested availabilitySets.
func (c *availabilitySets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("availabilitysets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a availabilitySet and creates it.  Returns the server's representation of the availabilitySet, and an error, if there is any.
func (c *availabilitySets) Create(availabilitySet *v1alpha1.AvailabilitySet) (result *v1alpha1.AvailabilitySet, err error) {
	result = &v1alpha1.AvailabilitySet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("availabilitysets").
		Body(availabilitySet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a availabilitySet and updates it. Returns the server's representation of the availabilitySet, and an error, if there is any.
func (c *availabilitySets) Update(availabilitySet *v1alpha1.AvailabilitySet) (result *v1alpha1.AvailabilitySet, err error) {
	result = &v1alpha1.AvailabilitySet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("availabilitysets").
		Name(availabilitySet.Name).
		Body(availabilitySet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *availabilitySets) UpdateStatus(availabilitySet *v1alpha1.AvailabilitySet) (result *v1alpha1.AvailabilitySet, err error) {
	result = &v1alpha1.AvailabilitySet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("availabilitysets").
		Name(availabilitySet.Name).
		SubResource("status").
		Body(availabilitySet).
		Do().
		Into(result)
	return
}

// Delete takes name of the availabilitySet and deletes it. Returns an error if one occurs.
func (c *availabilitySets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("availabilitysets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *availabilitySets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("availabilitysets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched availabilitySet.
func (c *availabilitySets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AvailabilitySet, err error) {
	result = &v1alpha1.AvailabilitySet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("availabilitysets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
