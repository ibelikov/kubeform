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

// Ec2ClientVpnEndpointsGetter has a method to return a Ec2ClientVpnEndpointInterface.
// A group's client should implement this interface.
type Ec2ClientVpnEndpointsGetter interface {
	Ec2ClientVpnEndpoints() Ec2ClientVpnEndpointInterface
}

// Ec2ClientVpnEndpointInterface has methods to work with Ec2ClientVpnEndpoint resources.
type Ec2ClientVpnEndpointInterface interface {
	Create(*v1alpha1.Ec2ClientVpnEndpoint) (*v1alpha1.Ec2ClientVpnEndpoint, error)
	Update(*v1alpha1.Ec2ClientVpnEndpoint) (*v1alpha1.Ec2ClientVpnEndpoint, error)
	UpdateStatus(*v1alpha1.Ec2ClientVpnEndpoint) (*v1alpha1.Ec2ClientVpnEndpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Ec2ClientVpnEndpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.Ec2ClientVpnEndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2ClientVpnEndpoint, err error)
	Ec2ClientVpnEndpointExpansion
}

// ec2ClientVpnEndpoints implements Ec2ClientVpnEndpointInterface
type ec2ClientVpnEndpoints struct {
	client rest.Interface
}

// newEc2ClientVpnEndpoints returns a Ec2ClientVpnEndpoints
func newEc2ClientVpnEndpoints(c *AwsV1alpha1Client) *ec2ClientVpnEndpoints {
	return &ec2ClientVpnEndpoints{
		client: c.RESTClient(),
	}
}

// Get takes name of the ec2ClientVpnEndpoint, and returns the corresponding ec2ClientVpnEndpoint object, and an error if there is any.
func (c *ec2ClientVpnEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.Ec2ClientVpnEndpoint, err error) {
	result = &v1alpha1.Ec2ClientVpnEndpoint{}
	err = c.client.Get().
		Resource("ec2clientvpnendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ec2ClientVpnEndpoints that match those selectors.
func (c *ec2ClientVpnEndpoints) List(opts v1.ListOptions) (result *v1alpha1.Ec2ClientVpnEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Ec2ClientVpnEndpointList{}
	err = c.client.Get().
		Resource("ec2clientvpnendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ec2ClientVpnEndpoints.
func (c *ec2ClientVpnEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ec2clientvpnendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ec2ClientVpnEndpoint and creates it.  Returns the server's representation of the ec2ClientVpnEndpoint, and an error, if there is any.
func (c *ec2ClientVpnEndpoints) Create(ec2ClientVpnEndpoint *v1alpha1.Ec2ClientVpnEndpoint) (result *v1alpha1.Ec2ClientVpnEndpoint, err error) {
	result = &v1alpha1.Ec2ClientVpnEndpoint{}
	err = c.client.Post().
		Resource("ec2clientvpnendpoints").
		Body(ec2ClientVpnEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ec2ClientVpnEndpoint and updates it. Returns the server's representation of the ec2ClientVpnEndpoint, and an error, if there is any.
func (c *ec2ClientVpnEndpoints) Update(ec2ClientVpnEndpoint *v1alpha1.Ec2ClientVpnEndpoint) (result *v1alpha1.Ec2ClientVpnEndpoint, err error) {
	result = &v1alpha1.Ec2ClientVpnEndpoint{}
	err = c.client.Put().
		Resource("ec2clientvpnendpoints").
		Name(ec2ClientVpnEndpoint.Name).
		Body(ec2ClientVpnEndpoint).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ec2ClientVpnEndpoints) UpdateStatus(ec2ClientVpnEndpoint *v1alpha1.Ec2ClientVpnEndpoint) (result *v1alpha1.Ec2ClientVpnEndpoint, err error) {
	result = &v1alpha1.Ec2ClientVpnEndpoint{}
	err = c.client.Put().
		Resource("ec2clientvpnendpoints").
		Name(ec2ClientVpnEndpoint.Name).
		SubResource("status").
		Body(ec2ClientVpnEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the ec2ClientVpnEndpoint and deletes it. Returns an error if one occurs.
func (c *ec2ClientVpnEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ec2clientvpnendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ec2ClientVpnEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ec2clientvpnendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ec2ClientVpnEndpoint.
func (c *ec2ClientVpnEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2ClientVpnEndpoint, err error) {
	result = &v1alpha1.Ec2ClientVpnEndpoint{}
	err = c.client.Patch(pt).
		Resource("ec2clientvpnendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}