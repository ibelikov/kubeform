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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermSchedulerJobCollectionsGetter has a method to return a AzurermSchedulerJobCollectionInterface.
// A group's client should implement this interface.
type AzurermSchedulerJobCollectionsGetter interface {
	AzurermSchedulerJobCollections() AzurermSchedulerJobCollectionInterface
}

// AzurermSchedulerJobCollectionInterface has methods to work with AzurermSchedulerJobCollection resources.
type AzurermSchedulerJobCollectionInterface interface {
	Create(*v1alpha1.AzurermSchedulerJobCollection) (*v1alpha1.AzurermSchedulerJobCollection, error)
	Update(*v1alpha1.AzurermSchedulerJobCollection) (*v1alpha1.AzurermSchedulerJobCollection, error)
	UpdateStatus(*v1alpha1.AzurermSchedulerJobCollection) (*v1alpha1.AzurermSchedulerJobCollection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermSchedulerJobCollection, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermSchedulerJobCollectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSchedulerJobCollection, err error)
	AzurermSchedulerJobCollectionExpansion
}

// azurermSchedulerJobCollections implements AzurermSchedulerJobCollectionInterface
type azurermSchedulerJobCollections struct {
	client rest.Interface
}

// newAzurermSchedulerJobCollections returns a AzurermSchedulerJobCollections
func newAzurermSchedulerJobCollections(c *AzurermV1alpha1Client) *azurermSchedulerJobCollections {
	return &azurermSchedulerJobCollections{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermSchedulerJobCollection, and returns the corresponding azurermSchedulerJobCollection object, and an error if there is any.
func (c *azurermSchedulerJobCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermSchedulerJobCollection, err error) {
	result = &v1alpha1.AzurermSchedulerJobCollection{}
	err = c.client.Get().
		Resource("azurermschedulerjobcollections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermSchedulerJobCollections that match those selectors.
func (c *azurermSchedulerJobCollections) List(opts v1.ListOptions) (result *v1alpha1.AzurermSchedulerJobCollectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermSchedulerJobCollectionList{}
	err = c.client.Get().
		Resource("azurermschedulerjobcollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermSchedulerJobCollections.
func (c *azurermSchedulerJobCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermschedulerjobcollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermSchedulerJobCollection and creates it.  Returns the server's representation of the azurermSchedulerJobCollection, and an error, if there is any.
func (c *azurermSchedulerJobCollections) Create(azurermSchedulerJobCollection *v1alpha1.AzurermSchedulerJobCollection) (result *v1alpha1.AzurermSchedulerJobCollection, err error) {
	result = &v1alpha1.AzurermSchedulerJobCollection{}
	err = c.client.Post().
		Resource("azurermschedulerjobcollections").
		Body(azurermSchedulerJobCollection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermSchedulerJobCollection and updates it. Returns the server's representation of the azurermSchedulerJobCollection, and an error, if there is any.
func (c *azurermSchedulerJobCollections) Update(azurermSchedulerJobCollection *v1alpha1.AzurermSchedulerJobCollection) (result *v1alpha1.AzurermSchedulerJobCollection, err error) {
	result = &v1alpha1.AzurermSchedulerJobCollection{}
	err = c.client.Put().
		Resource("azurermschedulerjobcollections").
		Name(azurermSchedulerJobCollection.Name).
		Body(azurermSchedulerJobCollection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermSchedulerJobCollections) UpdateStatus(azurermSchedulerJobCollection *v1alpha1.AzurermSchedulerJobCollection) (result *v1alpha1.AzurermSchedulerJobCollection, err error) {
	result = &v1alpha1.AzurermSchedulerJobCollection{}
	err = c.client.Put().
		Resource("azurermschedulerjobcollections").
		Name(azurermSchedulerJobCollection.Name).
		SubResource("status").
		Body(azurermSchedulerJobCollection).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermSchedulerJobCollection and deletes it. Returns an error if one occurs.
func (c *azurermSchedulerJobCollections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermschedulerjobcollections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermSchedulerJobCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermschedulerjobcollections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermSchedulerJobCollection.
func (c *azurermSchedulerJobCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSchedulerJobCollection, err error) {
	result = &v1alpha1.AzurermSchedulerJobCollection{}
	err = c.client.Patch(pt).
		Resource("azurermschedulerjobcollections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
