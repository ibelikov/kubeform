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

// AwsDxHostedPrivateVirtualInterfaceAcceptersGetter has a method to return a AwsDxHostedPrivateVirtualInterfaceAccepterInterface.
// A group's client should implement this interface.
type AwsDxHostedPrivateVirtualInterfaceAcceptersGetter interface {
	AwsDxHostedPrivateVirtualInterfaceAccepters() AwsDxHostedPrivateVirtualInterfaceAccepterInterface
}

// AwsDxHostedPrivateVirtualInterfaceAccepterInterface has methods to work with AwsDxHostedPrivateVirtualInterfaceAccepter resources.
type AwsDxHostedPrivateVirtualInterfaceAccepterInterface interface {
	Create(*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter) (*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, error)
	Update(*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter) (*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, error)
	UpdateStatus(*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter) (*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, err error)
	AwsDxHostedPrivateVirtualInterfaceAccepterExpansion
}

// awsDxHostedPrivateVirtualInterfaceAccepters implements AwsDxHostedPrivateVirtualInterfaceAccepterInterface
type awsDxHostedPrivateVirtualInterfaceAccepters struct {
	client rest.Interface
}

// newAwsDxHostedPrivateVirtualInterfaceAccepters returns a AwsDxHostedPrivateVirtualInterfaceAccepters
func newAwsDxHostedPrivateVirtualInterfaceAccepters(c *AwsV1alpha1Client) *awsDxHostedPrivateVirtualInterfaceAccepters {
	return &awsDxHostedPrivateVirtualInterfaceAccepters{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDxHostedPrivateVirtualInterfaceAccepter, and returns the corresponding awsDxHostedPrivateVirtualInterfaceAccepter object, and an error if there is any.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, err error) {
	result = &v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter{}
	err = c.client.Get().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDxHostedPrivateVirtualInterfaceAccepters that match those selectors.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) List(opts v1.ListOptions) (result *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepterList{}
	err = c.client.Get().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDxHostedPrivateVirtualInterfaceAccepters.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDxHostedPrivateVirtualInterfaceAccepter and creates it.  Returns the server's representation of the awsDxHostedPrivateVirtualInterfaceAccepter, and an error, if there is any.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) Create(awsDxHostedPrivateVirtualInterfaceAccepter *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter) (result *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, err error) {
	result = &v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter{}
	err = c.client.Post().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		Body(awsDxHostedPrivateVirtualInterfaceAccepter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDxHostedPrivateVirtualInterfaceAccepter and updates it. Returns the server's representation of the awsDxHostedPrivateVirtualInterfaceAccepter, and an error, if there is any.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) Update(awsDxHostedPrivateVirtualInterfaceAccepter *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter) (result *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, err error) {
	result = &v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter{}
	err = c.client.Put().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		Name(awsDxHostedPrivateVirtualInterfaceAccepter.Name).
		Body(awsDxHostedPrivateVirtualInterfaceAccepter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDxHostedPrivateVirtualInterfaceAccepters) UpdateStatus(awsDxHostedPrivateVirtualInterfaceAccepter *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter) (result *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, err error) {
	result = &v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter{}
	err = c.client.Put().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		Name(awsDxHostedPrivateVirtualInterfaceAccepter.Name).
		SubResource("status").
		Body(awsDxHostedPrivateVirtualInterfaceAccepter).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDxHostedPrivateVirtualInterfaceAccepter and deletes it. Returns an error if one occurs.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDxHostedPrivateVirtualInterfaceAccepter.
func (c *awsDxHostedPrivateVirtualInterfaceAccepters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter, err error) {
	result = &v1alpha1.AwsDxHostedPrivateVirtualInterfaceAccepter{}
	err = c.client.Patch(pt).
		Resource("awsdxhostedprivatevirtualinterfaceaccepters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
