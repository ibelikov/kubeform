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

// AzurermStreamAnalyticsOutputServicebusQueuesGetter has a method to return a AzurermStreamAnalyticsOutputServicebusQueueInterface.
// A group's client should implement this interface.
type AzurermStreamAnalyticsOutputServicebusQueuesGetter interface {
	AzurermStreamAnalyticsOutputServicebusQueues() AzurermStreamAnalyticsOutputServicebusQueueInterface
}

// AzurermStreamAnalyticsOutputServicebusQueueInterface has methods to work with AzurermStreamAnalyticsOutputServicebusQueue resources.
type AzurermStreamAnalyticsOutputServicebusQueueInterface interface {
	Create(*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue) (*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, error)
	Update(*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue) (*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, error)
	UpdateStatus(*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue) (*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermStreamAnalyticsOutputServicebusQueueList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, err error)
	AzurermStreamAnalyticsOutputServicebusQueueExpansion
}

// azurermStreamAnalyticsOutputServicebusQueues implements AzurermStreamAnalyticsOutputServicebusQueueInterface
type azurermStreamAnalyticsOutputServicebusQueues struct {
	client rest.Interface
}

// newAzurermStreamAnalyticsOutputServicebusQueues returns a AzurermStreamAnalyticsOutputServicebusQueues
func newAzurermStreamAnalyticsOutputServicebusQueues(c *AzurermV1alpha1Client) *azurermStreamAnalyticsOutputServicebusQueues {
	return &azurermStreamAnalyticsOutputServicebusQueues{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermStreamAnalyticsOutputServicebusQueue, and returns the corresponding azurermStreamAnalyticsOutputServicebusQueue object, and an error if there is any.
func (c *azurermStreamAnalyticsOutputServicebusQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, err error) {
	result = &v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue{}
	err = c.client.Get().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermStreamAnalyticsOutputServicebusQueues that match those selectors.
func (c *azurermStreamAnalyticsOutputServicebusQueues) List(opts v1.ListOptions) (result *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueueList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermStreamAnalyticsOutputServicebusQueueList{}
	err = c.client.Get().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermStreamAnalyticsOutputServicebusQueues.
func (c *azurermStreamAnalyticsOutputServicebusQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermStreamAnalyticsOutputServicebusQueue and creates it.  Returns the server's representation of the azurermStreamAnalyticsOutputServicebusQueue, and an error, if there is any.
func (c *azurermStreamAnalyticsOutputServicebusQueues) Create(azurermStreamAnalyticsOutputServicebusQueue *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue) (result *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, err error) {
	result = &v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue{}
	err = c.client.Post().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		Body(azurermStreamAnalyticsOutputServicebusQueue).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermStreamAnalyticsOutputServicebusQueue and updates it. Returns the server's representation of the azurermStreamAnalyticsOutputServicebusQueue, and an error, if there is any.
func (c *azurermStreamAnalyticsOutputServicebusQueues) Update(azurermStreamAnalyticsOutputServicebusQueue *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue) (result *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, err error) {
	result = &v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue{}
	err = c.client.Put().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		Name(azurermStreamAnalyticsOutputServicebusQueue.Name).
		Body(azurermStreamAnalyticsOutputServicebusQueue).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermStreamAnalyticsOutputServicebusQueues) UpdateStatus(azurermStreamAnalyticsOutputServicebusQueue *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue) (result *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, err error) {
	result = &v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue{}
	err = c.client.Put().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		Name(azurermStreamAnalyticsOutputServicebusQueue.Name).
		SubResource("status").
		Body(azurermStreamAnalyticsOutputServicebusQueue).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermStreamAnalyticsOutputServicebusQueue and deletes it. Returns an error if one occurs.
func (c *azurermStreamAnalyticsOutputServicebusQueues) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermStreamAnalyticsOutputServicebusQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermStreamAnalyticsOutputServicebusQueue.
func (c *azurermStreamAnalyticsOutputServicebusQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue, err error) {
	result = &v1alpha1.AzurermStreamAnalyticsOutputServicebusQueue{}
	err = c.client.Patch(pt).
		Resource("azurermstreamanalyticsoutputservicebusqueues").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
