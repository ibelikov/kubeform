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

// VpnGatewaysGetter has a method to return a VpnGatewayInterface.
// A group's client should implement this interface.
type VpnGatewaysGetter interface {
	VpnGateways() VpnGatewayInterface
}

// VpnGatewayInterface has methods to work with VpnGateway resources.
type VpnGatewayInterface interface {
	Create(*v1alpha1.VpnGateway) (*v1alpha1.VpnGateway, error)
	Update(*v1alpha1.VpnGateway) (*v1alpha1.VpnGateway, error)
	UpdateStatus(*v1alpha1.VpnGateway) (*v1alpha1.VpnGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VpnGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.VpnGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnGateway, err error)
	VpnGatewayExpansion
}

// vpnGateways implements VpnGatewayInterface
type vpnGateways struct {
	client rest.Interface
}

// newVpnGateways returns a VpnGateways
func newVpnGateways(c *AwsV1alpha1Client) *vpnGateways {
	return &vpnGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the vpnGateway, and returns the corresponding vpnGateway object, and an error if there is any.
func (c *vpnGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.VpnGateway, err error) {
	result = &v1alpha1.VpnGateway{}
	err = c.client.Get().
		Resource("vpngateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpnGateways that match those selectors.
func (c *vpnGateways) List(opts v1.ListOptions) (result *v1alpha1.VpnGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpnGatewayList{}
	err = c.client.Get().
		Resource("vpngateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpnGateways.
func (c *vpnGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("vpngateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vpnGateway and creates it.  Returns the server's representation of the vpnGateway, and an error, if there is any.
func (c *vpnGateways) Create(vpnGateway *v1alpha1.VpnGateway) (result *v1alpha1.VpnGateway, err error) {
	result = &v1alpha1.VpnGateway{}
	err = c.client.Post().
		Resource("vpngateways").
		Body(vpnGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vpnGateway and updates it. Returns the server's representation of the vpnGateway, and an error, if there is any.
func (c *vpnGateways) Update(vpnGateway *v1alpha1.VpnGateway) (result *v1alpha1.VpnGateway, err error) {
	result = &v1alpha1.VpnGateway{}
	err = c.client.Put().
		Resource("vpngateways").
		Name(vpnGateway.Name).
		Body(vpnGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vpnGateways) UpdateStatus(vpnGateway *v1alpha1.VpnGateway) (result *v1alpha1.VpnGateway, err error) {
	result = &v1alpha1.VpnGateway{}
	err = c.client.Put().
		Resource("vpngateways").
		Name(vpnGateway.Name).
		SubResource("status").
		Body(vpnGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the vpnGateway and deletes it. Returns an error if one occurs.
func (c *vpnGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("vpngateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpnGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("vpngateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vpnGateway.
func (c *vpnGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnGateway, err error) {
	result = &v1alpha1.VpnGateway{}
	err = c.client.Patch(pt).
		Resource("vpngateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
