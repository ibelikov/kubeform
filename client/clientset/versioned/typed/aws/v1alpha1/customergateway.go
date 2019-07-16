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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// CustomerGatewaysGetter has a method to return a CustomerGatewayInterface.
// A group's client should implement this interface.
type CustomerGatewaysGetter interface {
	CustomerGateways() CustomerGatewayInterface
}

// CustomerGatewayInterface has methods to work with CustomerGateway resources.
type CustomerGatewayInterface interface {
	Create(*v1alpha1.CustomerGateway) (*v1alpha1.CustomerGateway, error)
	Update(*v1alpha1.CustomerGateway) (*v1alpha1.CustomerGateway, error)
	UpdateStatus(*v1alpha1.CustomerGateway) (*v1alpha1.CustomerGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CustomerGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.CustomerGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CustomerGateway, err error)
	CustomerGatewayExpansion
}

// customerGateways implements CustomerGatewayInterface
type customerGateways struct {
	client rest.Interface
}

// newCustomerGateways returns a CustomerGateways
func newCustomerGateways(c *AwsV1alpha1Client) *customerGateways {
	return &customerGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the customerGateway, and returns the corresponding customerGateway object, and an error if there is any.
func (c *customerGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.CustomerGateway, err error) {
	result = &v1alpha1.CustomerGateway{}
	err = c.client.Get().
		Resource("customergateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CustomerGateways that match those selectors.
func (c *customerGateways) List(opts v1.ListOptions) (result *v1alpha1.CustomerGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CustomerGatewayList{}
	err = c.client.Get().
		Resource("customergateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested customerGateways.
func (c *customerGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("customergateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a customerGateway and creates it.  Returns the server's representation of the customerGateway, and an error, if there is any.
func (c *customerGateways) Create(customerGateway *v1alpha1.CustomerGateway) (result *v1alpha1.CustomerGateway, err error) {
	result = &v1alpha1.CustomerGateway{}
	err = c.client.Post().
		Resource("customergateways").
		Body(customerGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a customerGateway and updates it. Returns the server's representation of the customerGateway, and an error, if there is any.
func (c *customerGateways) Update(customerGateway *v1alpha1.CustomerGateway) (result *v1alpha1.CustomerGateway, err error) {
	result = &v1alpha1.CustomerGateway{}
	err = c.client.Put().
		Resource("customergateways").
		Name(customerGateway.Name).
		Body(customerGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *customerGateways) UpdateStatus(customerGateway *v1alpha1.CustomerGateway) (result *v1alpha1.CustomerGateway, err error) {
	result = &v1alpha1.CustomerGateway{}
	err = c.client.Put().
		Resource("customergateways").
		Name(customerGateway.Name).
		SubResource("status").
		Body(customerGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the customerGateway and deletes it. Returns an error if one occurs.
func (c *customerGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("customergateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *customerGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("customergateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched customerGateway.
func (c *customerGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CustomerGateway, err error) {
	result = &v1alpha1.CustomerGateway{}
	err = c.client.Patch(pt).
		Resource("customergateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
