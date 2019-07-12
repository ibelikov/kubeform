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

// AzurermApplicationGatewaysGetter has a method to return a AzurermApplicationGatewayInterface.
// A group's client should implement this interface.
type AzurermApplicationGatewaysGetter interface {
	AzurermApplicationGateways() AzurermApplicationGatewayInterface
}

// AzurermApplicationGatewayInterface has methods to work with AzurermApplicationGateway resources.
type AzurermApplicationGatewayInterface interface {
	Create(*v1alpha1.AzurermApplicationGateway) (*v1alpha1.AzurermApplicationGateway, error)
	Update(*v1alpha1.AzurermApplicationGateway) (*v1alpha1.AzurermApplicationGateway, error)
	UpdateStatus(*v1alpha1.AzurermApplicationGateway) (*v1alpha1.AzurermApplicationGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermApplicationGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermApplicationGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApplicationGateway, err error)
	AzurermApplicationGatewayExpansion
}

// azurermApplicationGateways implements AzurermApplicationGatewayInterface
type azurermApplicationGateways struct {
	client rest.Interface
}

// newAzurermApplicationGateways returns a AzurermApplicationGateways
func newAzurermApplicationGateways(c *AzurermV1alpha1Client) *azurermApplicationGateways {
	return &azurermApplicationGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermApplicationGateway, and returns the corresponding azurermApplicationGateway object, and an error if there is any.
func (c *azurermApplicationGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApplicationGateway, err error) {
	result = &v1alpha1.AzurermApplicationGateway{}
	err = c.client.Get().
		Resource("azurermapplicationgateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermApplicationGateways that match those selectors.
func (c *azurermApplicationGateways) List(opts v1.ListOptions) (result *v1alpha1.AzurermApplicationGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermApplicationGatewayList{}
	err = c.client.Get().
		Resource("azurermapplicationgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermApplicationGateways.
func (c *azurermApplicationGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermapplicationgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermApplicationGateway and creates it.  Returns the server's representation of the azurermApplicationGateway, and an error, if there is any.
func (c *azurermApplicationGateways) Create(azurermApplicationGateway *v1alpha1.AzurermApplicationGateway) (result *v1alpha1.AzurermApplicationGateway, err error) {
	result = &v1alpha1.AzurermApplicationGateway{}
	err = c.client.Post().
		Resource("azurermapplicationgateways").
		Body(azurermApplicationGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermApplicationGateway and updates it. Returns the server's representation of the azurermApplicationGateway, and an error, if there is any.
func (c *azurermApplicationGateways) Update(azurermApplicationGateway *v1alpha1.AzurermApplicationGateway) (result *v1alpha1.AzurermApplicationGateway, err error) {
	result = &v1alpha1.AzurermApplicationGateway{}
	err = c.client.Put().
		Resource("azurermapplicationgateways").
		Name(azurermApplicationGateway.Name).
		Body(azurermApplicationGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermApplicationGateways) UpdateStatus(azurermApplicationGateway *v1alpha1.AzurermApplicationGateway) (result *v1alpha1.AzurermApplicationGateway, err error) {
	result = &v1alpha1.AzurermApplicationGateway{}
	err = c.client.Put().
		Resource("azurermapplicationgateways").
		Name(azurermApplicationGateway.Name).
		SubResource("status").
		Body(azurermApplicationGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermApplicationGateway and deletes it. Returns an error if one occurs.
func (c *azurermApplicationGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermapplicationgateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermApplicationGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermapplicationgateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermApplicationGateway.
func (c *azurermApplicationGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApplicationGateway, err error) {
	result = &v1alpha1.AzurermApplicationGateway{}
	err = c.client.Patch(pt).
		Resource("azurermapplicationgateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
