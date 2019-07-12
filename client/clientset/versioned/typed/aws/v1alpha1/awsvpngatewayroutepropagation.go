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

// AwsVpnGatewayRoutePropagationsGetter has a method to return a AwsVpnGatewayRoutePropagationInterface.
// A group's client should implement this interface.
type AwsVpnGatewayRoutePropagationsGetter interface {
	AwsVpnGatewayRoutePropagations() AwsVpnGatewayRoutePropagationInterface
}

// AwsVpnGatewayRoutePropagationInterface has methods to work with AwsVpnGatewayRoutePropagation resources.
type AwsVpnGatewayRoutePropagationInterface interface {
	Create(*v1alpha1.AwsVpnGatewayRoutePropagation) (*v1alpha1.AwsVpnGatewayRoutePropagation, error)
	Update(*v1alpha1.AwsVpnGatewayRoutePropagation) (*v1alpha1.AwsVpnGatewayRoutePropagation, error)
	UpdateStatus(*v1alpha1.AwsVpnGatewayRoutePropagation) (*v1alpha1.AwsVpnGatewayRoutePropagation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsVpnGatewayRoutePropagation, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsVpnGatewayRoutePropagationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpnGatewayRoutePropagation, err error)
	AwsVpnGatewayRoutePropagationExpansion
}

// awsVpnGatewayRoutePropagations implements AwsVpnGatewayRoutePropagationInterface
type awsVpnGatewayRoutePropagations struct {
	client rest.Interface
}

// newAwsVpnGatewayRoutePropagations returns a AwsVpnGatewayRoutePropagations
func newAwsVpnGatewayRoutePropagations(c *AwsV1alpha1Client) *awsVpnGatewayRoutePropagations {
	return &awsVpnGatewayRoutePropagations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsVpnGatewayRoutePropagation, and returns the corresponding awsVpnGatewayRoutePropagation object, and an error if there is any.
func (c *awsVpnGatewayRoutePropagations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpnGatewayRoutePropagation, err error) {
	result = &v1alpha1.AwsVpnGatewayRoutePropagation{}
	err = c.client.Get().
		Resource("awsvpngatewayroutepropagations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpnGatewayRoutePropagations that match those selectors.
func (c *awsVpnGatewayRoutePropagations) List(opts v1.ListOptions) (result *v1alpha1.AwsVpnGatewayRoutePropagationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsVpnGatewayRoutePropagationList{}
	err = c.client.Get().
		Resource("awsvpngatewayroutepropagations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpnGatewayRoutePropagations.
func (c *awsVpnGatewayRoutePropagations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsvpngatewayroutepropagations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsVpnGatewayRoutePropagation and creates it.  Returns the server's representation of the awsVpnGatewayRoutePropagation, and an error, if there is any.
func (c *awsVpnGatewayRoutePropagations) Create(awsVpnGatewayRoutePropagation *v1alpha1.AwsVpnGatewayRoutePropagation) (result *v1alpha1.AwsVpnGatewayRoutePropagation, err error) {
	result = &v1alpha1.AwsVpnGatewayRoutePropagation{}
	err = c.client.Post().
		Resource("awsvpngatewayroutepropagations").
		Body(awsVpnGatewayRoutePropagation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpnGatewayRoutePropagation and updates it. Returns the server's representation of the awsVpnGatewayRoutePropagation, and an error, if there is any.
func (c *awsVpnGatewayRoutePropagations) Update(awsVpnGatewayRoutePropagation *v1alpha1.AwsVpnGatewayRoutePropagation) (result *v1alpha1.AwsVpnGatewayRoutePropagation, err error) {
	result = &v1alpha1.AwsVpnGatewayRoutePropagation{}
	err = c.client.Put().
		Resource("awsvpngatewayroutepropagations").
		Name(awsVpnGatewayRoutePropagation.Name).
		Body(awsVpnGatewayRoutePropagation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsVpnGatewayRoutePropagations) UpdateStatus(awsVpnGatewayRoutePropagation *v1alpha1.AwsVpnGatewayRoutePropagation) (result *v1alpha1.AwsVpnGatewayRoutePropagation, err error) {
	result = &v1alpha1.AwsVpnGatewayRoutePropagation{}
	err = c.client.Put().
		Resource("awsvpngatewayroutepropagations").
		Name(awsVpnGatewayRoutePropagation.Name).
		SubResource("status").
		Body(awsVpnGatewayRoutePropagation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpnGatewayRoutePropagation and deletes it. Returns an error if one occurs.
func (c *awsVpnGatewayRoutePropagations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsvpngatewayroutepropagations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpnGatewayRoutePropagations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsvpngatewayroutepropagations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpnGatewayRoutePropagation.
func (c *awsVpnGatewayRoutePropagations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpnGatewayRoutePropagation, err error) {
	result = &v1alpha1.AwsVpnGatewayRoutePropagation{}
	err = c.client.Patch(pt).
		Resource("awsvpngatewayroutepropagations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
