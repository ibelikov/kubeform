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

// AutomationDscNodeconfigurationsGetter has a method to return a AutomationDscNodeconfigurationInterface.
// A group's client should implement this interface.
type AutomationDscNodeconfigurationsGetter interface {
	AutomationDscNodeconfigurations(namespace string) AutomationDscNodeconfigurationInterface
}

// AutomationDscNodeconfigurationInterface has methods to work with AutomationDscNodeconfiguration resources.
type AutomationDscNodeconfigurationInterface interface {
	Create(*v1alpha1.AutomationDscNodeconfiguration) (*v1alpha1.AutomationDscNodeconfiguration, error)
	Update(*v1alpha1.AutomationDscNodeconfiguration) (*v1alpha1.AutomationDscNodeconfiguration, error)
	UpdateStatus(*v1alpha1.AutomationDscNodeconfiguration) (*v1alpha1.AutomationDscNodeconfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutomationDscNodeconfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.AutomationDscNodeconfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationDscNodeconfiguration, err error)
	AutomationDscNodeconfigurationExpansion
}

// automationDscNodeconfigurations implements AutomationDscNodeconfigurationInterface
type automationDscNodeconfigurations struct {
	client rest.Interface
	ns     string
}

// newAutomationDscNodeconfigurations returns a AutomationDscNodeconfigurations
func newAutomationDscNodeconfigurations(c *AzurermV1alpha1Client, namespace string) *automationDscNodeconfigurations {
	return &automationDscNodeconfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the automationDscNodeconfiguration, and returns the corresponding automationDscNodeconfiguration object, and an error if there is any.
func (c *automationDscNodeconfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	result = &v1alpha1.AutomationDscNodeconfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutomationDscNodeconfigurations that match those selectors.
func (c *automationDscNodeconfigurations) List(opts v1.ListOptions) (result *v1alpha1.AutomationDscNodeconfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutomationDscNodeconfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested automationDscNodeconfigurations.
func (c *automationDscNodeconfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a automationDscNodeconfiguration and creates it.  Returns the server's representation of the automationDscNodeconfiguration, and an error, if there is any.
func (c *automationDscNodeconfigurations) Create(automationDscNodeconfiguration *v1alpha1.AutomationDscNodeconfiguration) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	result = &v1alpha1.AutomationDscNodeconfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		Body(automationDscNodeconfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a automationDscNodeconfiguration and updates it. Returns the server's representation of the automationDscNodeconfiguration, and an error, if there is any.
func (c *automationDscNodeconfigurations) Update(automationDscNodeconfiguration *v1alpha1.AutomationDscNodeconfiguration) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	result = &v1alpha1.AutomationDscNodeconfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		Name(automationDscNodeconfiguration.Name).
		Body(automationDscNodeconfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *automationDscNodeconfigurations) UpdateStatus(automationDscNodeconfiguration *v1alpha1.AutomationDscNodeconfiguration) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	result = &v1alpha1.AutomationDscNodeconfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		Name(automationDscNodeconfiguration.Name).
		SubResource("status").
		Body(automationDscNodeconfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the automationDscNodeconfiguration and deletes it. Returns an error if one occurs.
func (c *automationDscNodeconfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *automationDscNodeconfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched automationDscNodeconfiguration.
func (c *automationDscNodeconfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	result = &v1alpha1.AutomationDscNodeconfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("automationdscnodeconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
