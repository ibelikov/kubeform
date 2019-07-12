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

// AzurermApiManagementPropertiesGetter has a method to return a AzurermApiManagementPropertyInterface.
// A group's client should implement this interface.
type AzurermApiManagementPropertiesGetter interface {
	AzurermApiManagementProperties() AzurermApiManagementPropertyInterface
}

// AzurermApiManagementPropertyInterface has methods to work with AzurermApiManagementProperty resources.
type AzurermApiManagementPropertyInterface interface {
	Create(*v1alpha1.AzurermApiManagementProperty) (*v1alpha1.AzurermApiManagementProperty, error)
	Update(*v1alpha1.AzurermApiManagementProperty) (*v1alpha1.AzurermApiManagementProperty, error)
	UpdateStatus(*v1alpha1.AzurermApiManagementProperty) (*v1alpha1.AzurermApiManagementProperty, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermApiManagementProperty, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermApiManagementPropertyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementProperty, err error)
	AzurermApiManagementPropertyExpansion
}

// azurermApiManagementProperties implements AzurermApiManagementPropertyInterface
type azurermApiManagementProperties struct {
	client rest.Interface
}

// newAzurermApiManagementProperties returns a AzurermApiManagementProperties
func newAzurermApiManagementProperties(c *AzurermV1alpha1Client) *azurermApiManagementProperties {
	return &azurermApiManagementProperties{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermApiManagementProperty, and returns the corresponding azurermApiManagementProperty object, and an error if there is any.
func (c *azurermApiManagementProperties) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementProperty, err error) {
	result = &v1alpha1.AzurermApiManagementProperty{}
	err = c.client.Get().
		Resource("azurermapimanagementproperties").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermApiManagementProperties that match those selectors.
func (c *azurermApiManagementProperties) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementPropertyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermApiManagementPropertyList{}
	err = c.client.Get().
		Resource("azurermapimanagementproperties").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementProperties.
func (c *azurermApiManagementProperties) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermapimanagementproperties").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermApiManagementProperty and creates it.  Returns the server's representation of the azurermApiManagementProperty, and an error, if there is any.
func (c *azurermApiManagementProperties) Create(azurermApiManagementProperty *v1alpha1.AzurermApiManagementProperty) (result *v1alpha1.AzurermApiManagementProperty, err error) {
	result = &v1alpha1.AzurermApiManagementProperty{}
	err = c.client.Post().
		Resource("azurermapimanagementproperties").
		Body(azurermApiManagementProperty).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermApiManagementProperty and updates it. Returns the server's representation of the azurermApiManagementProperty, and an error, if there is any.
func (c *azurermApiManagementProperties) Update(azurermApiManagementProperty *v1alpha1.AzurermApiManagementProperty) (result *v1alpha1.AzurermApiManagementProperty, err error) {
	result = &v1alpha1.AzurermApiManagementProperty{}
	err = c.client.Put().
		Resource("azurermapimanagementproperties").
		Name(azurermApiManagementProperty.Name).
		Body(azurermApiManagementProperty).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermApiManagementProperties) UpdateStatus(azurermApiManagementProperty *v1alpha1.AzurermApiManagementProperty) (result *v1alpha1.AzurermApiManagementProperty, err error) {
	result = &v1alpha1.AzurermApiManagementProperty{}
	err = c.client.Put().
		Resource("azurermapimanagementproperties").
		Name(azurermApiManagementProperty.Name).
		SubResource("status").
		Body(azurermApiManagementProperty).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermApiManagementProperty and deletes it. Returns an error if one occurs.
func (c *azurermApiManagementProperties) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermapimanagementproperties").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermApiManagementProperties) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermapimanagementproperties").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermApiManagementProperty.
func (c *azurermApiManagementProperties) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementProperty, err error) {
	result = &v1alpha1.AzurermApiManagementProperty{}
	err = c.client.Patch(pt).
		Resource("azurermapimanagementproperties").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
