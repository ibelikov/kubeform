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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DefaultNetworkACLsGetter has a method to return a DefaultNetworkACLInterface.
// A group's client should implement this interface.
type DefaultNetworkACLsGetter interface {
	DefaultNetworkACLs(namespace string) DefaultNetworkACLInterface
}

// DefaultNetworkACLInterface has methods to work with DefaultNetworkACL resources.
type DefaultNetworkACLInterface interface {
	Create(*v1alpha1.DefaultNetworkACL) (*v1alpha1.DefaultNetworkACL, error)
	Update(*v1alpha1.DefaultNetworkACL) (*v1alpha1.DefaultNetworkACL, error)
	UpdateStatus(*v1alpha1.DefaultNetworkACL) (*v1alpha1.DefaultNetworkACL, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DefaultNetworkACL, error)
	List(opts v1.ListOptions) (*v1alpha1.DefaultNetworkACLList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultNetworkACL, err error)
	DefaultNetworkACLExpansion
}

// defaultNetworkACLs implements DefaultNetworkACLInterface
type defaultNetworkACLs struct {
	client rest.Interface
	ns     string
}

// newDefaultNetworkACLs returns a DefaultNetworkACLs
func newDefaultNetworkACLs(c *AwsV1alpha1Client, namespace string) *defaultNetworkACLs {
	return &defaultNetworkACLs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the defaultNetworkACL, and returns the corresponding defaultNetworkACL object, and an error if there is any.
func (c *defaultNetworkACLs) Get(name string, options v1.GetOptions) (result *v1alpha1.DefaultNetworkACL, err error) {
	result = &v1alpha1.DefaultNetworkACL{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DefaultNetworkACLs that match those selectors.
func (c *defaultNetworkACLs) List(opts v1.ListOptions) (result *v1alpha1.DefaultNetworkACLList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DefaultNetworkACLList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested defaultNetworkACLs.
func (c *defaultNetworkACLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a defaultNetworkACL and creates it.  Returns the server's representation of the defaultNetworkACL, and an error, if there is any.
func (c *defaultNetworkACLs) Create(defaultNetworkACL *v1alpha1.DefaultNetworkACL) (result *v1alpha1.DefaultNetworkACL, err error) {
	result = &v1alpha1.DefaultNetworkACL{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		Body(defaultNetworkACL).
		Do().
		Into(result)
	return
}

// Update takes the representation of a defaultNetworkACL and updates it. Returns the server's representation of the defaultNetworkACL, and an error, if there is any.
func (c *defaultNetworkACLs) Update(defaultNetworkACL *v1alpha1.DefaultNetworkACL) (result *v1alpha1.DefaultNetworkACL, err error) {
	result = &v1alpha1.DefaultNetworkACL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		Name(defaultNetworkACL.Name).
		Body(defaultNetworkACL).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *defaultNetworkACLs) UpdateStatus(defaultNetworkACL *v1alpha1.DefaultNetworkACL) (result *v1alpha1.DefaultNetworkACL, err error) {
	result = &v1alpha1.DefaultNetworkACL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		Name(defaultNetworkACL.Name).
		SubResource("status").
		Body(defaultNetworkACL).
		Do().
		Into(result)
	return
}

// Delete takes name of the defaultNetworkACL and deletes it. Returns an error if one occurs.
func (c *defaultNetworkACLs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *defaultNetworkACLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched defaultNetworkACL.
func (c *defaultNetworkACLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultNetworkACL, err error) {
	result = &v1alpha1.DefaultNetworkACL{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("defaultnetworkacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
