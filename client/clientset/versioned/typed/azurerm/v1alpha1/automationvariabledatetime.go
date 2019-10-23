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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AutomationVariableDatetimesGetter has a method to return a AutomationVariableDatetimeInterface.
// A group's client should implement this interface.
type AutomationVariableDatetimesGetter interface {
	AutomationVariableDatetimes(namespace string) AutomationVariableDatetimeInterface
}

// AutomationVariableDatetimeInterface has methods to work with AutomationVariableDatetime resources.
type AutomationVariableDatetimeInterface interface {
	Create(*v1alpha1.AutomationVariableDatetime) (*v1alpha1.AutomationVariableDatetime, error)
	Update(*v1alpha1.AutomationVariableDatetime) (*v1alpha1.AutomationVariableDatetime, error)
	UpdateStatus(*v1alpha1.AutomationVariableDatetime) (*v1alpha1.AutomationVariableDatetime, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutomationVariableDatetime, error)
	List(opts v1.ListOptions) (*v1alpha1.AutomationVariableDatetimeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationVariableDatetime, err error)
	AutomationVariableDatetimeExpansion
}

// automationVariableDatetimes implements AutomationVariableDatetimeInterface
type automationVariableDatetimes struct {
	client rest.Interface
	ns     string
}

// newAutomationVariableDatetimes returns a AutomationVariableDatetimes
func newAutomationVariableDatetimes(c *AzurermV1alpha1Client, namespace string) *automationVariableDatetimes {
	return &automationVariableDatetimes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the automationVariableDatetime, and returns the corresponding automationVariableDatetime object, and an error if there is any.
func (c *automationVariableDatetimes) Get(name string, options v1.GetOptions) (result *v1alpha1.AutomationVariableDatetime, err error) {
	result = &v1alpha1.AutomationVariableDatetime{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutomationVariableDatetimes that match those selectors.
func (c *automationVariableDatetimes) List(opts v1.ListOptions) (result *v1alpha1.AutomationVariableDatetimeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutomationVariableDatetimeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested automationVariableDatetimes.
func (c *automationVariableDatetimes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a automationVariableDatetime and creates it.  Returns the server's representation of the automationVariableDatetime, and an error, if there is any.
func (c *automationVariableDatetimes) Create(automationVariableDatetime *v1alpha1.AutomationVariableDatetime) (result *v1alpha1.AutomationVariableDatetime, err error) {
	result = &v1alpha1.AutomationVariableDatetime{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		Body(automationVariableDatetime).
		Do().
		Into(result)
	return
}

// Update takes the representation of a automationVariableDatetime and updates it. Returns the server's representation of the automationVariableDatetime, and an error, if there is any.
func (c *automationVariableDatetimes) Update(automationVariableDatetime *v1alpha1.AutomationVariableDatetime) (result *v1alpha1.AutomationVariableDatetime, err error) {
	result = &v1alpha1.AutomationVariableDatetime{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		Name(automationVariableDatetime.Name).
		Body(automationVariableDatetime).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *automationVariableDatetimes) UpdateStatus(automationVariableDatetime *v1alpha1.AutomationVariableDatetime) (result *v1alpha1.AutomationVariableDatetime, err error) {
	result = &v1alpha1.AutomationVariableDatetime{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		Name(automationVariableDatetime.Name).
		SubResource("status").
		Body(automationVariableDatetime).
		Do().
		Into(result)
	return
}

// Delete takes name of the automationVariableDatetime and deletes it. Returns an error if one occurs.
func (c *automationVariableDatetimes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *automationVariableDatetimes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched automationVariableDatetime.
func (c *automationVariableDatetimes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationVariableDatetime, err error) {
	result = &v1alpha1.AutomationVariableDatetime{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("automationvariabledatetimes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
