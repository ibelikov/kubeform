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

// IothubsGetter has a method to return a IothubInterface.
// A group's client should implement this interface.
type IothubsGetter interface {
	Iothubs() IothubInterface
}

// IothubInterface has methods to work with Iothub resources.
type IothubInterface interface {
	Create(*v1alpha1.Iothub) (*v1alpha1.Iothub, error)
	Update(*v1alpha1.Iothub) (*v1alpha1.Iothub, error)
	UpdateStatus(*v1alpha1.Iothub) (*v1alpha1.Iothub, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Iothub, error)
	List(opts v1.ListOptions) (*v1alpha1.IothubList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Iothub, err error)
	IothubExpansion
}

// iothubs implements IothubInterface
type iothubs struct {
	client rest.Interface
}

// newIothubs returns a Iothubs
func newIothubs(c *AzurermV1alpha1Client) *iothubs {
	return &iothubs{
		client: c.RESTClient(),
	}
}

// Get takes name of the iothub, and returns the corresponding iothub object, and an error if there is any.
func (c *iothubs) Get(name string, options v1.GetOptions) (result *v1alpha1.Iothub, err error) {
	result = &v1alpha1.Iothub{}
	err = c.client.Get().
		Resource("iothubs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Iothubs that match those selectors.
func (c *iothubs) List(opts v1.ListOptions) (result *v1alpha1.IothubList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IothubList{}
	err = c.client.Get().
		Resource("iothubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iothubs.
func (c *iothubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("iothubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iothub and creates it.  Returns the server's representation of the iothub, and an error, if there is any.
func (c *iothubs) Create(iothub *v1alpha1.Iothub) (result *v1alpha1.Iothub, err error) {
	result = &v1alpha1.Iothub{}
	err = c.client.Post().
		Resource("iothubs").
		Body(iothub).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iothub and updates it. Returns the server's representation of the iothub, and an error, if there is any.
func (c *iothubs) Update(iothub *v1alpha1.Iothub) (result *v1alpha1.Iothub, err error) {
	result = &v1alpha1.Iothub{}
	err = c.client.Put().
		Resource("iothubs").
		Name(iothub.Name).
		Body(iothub).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iothubs) UpdateStatus(iothub *v1alpha1.Iothub) (result *v1alpha1.Iothub, err error) {
	result = &v1alpha1.Iothub{}
	err = c.client.Put().
		Resource("iothubs").
		Name(iothub.Name).
		SubResource("status").
		Body(iothub).
		Do().
		Into(result)
	return
}

// Delete takes name of the iothub and deletes it. Returns an error if one occurs.
func (c *iothubs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("iothubs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iothubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("iothubs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iothub.
func (c *iothubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Iothub, err error) {
	result = &v1alpha1.Iothub{}
	err = c.client.Patch(pt).
		Resource("iothubs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
