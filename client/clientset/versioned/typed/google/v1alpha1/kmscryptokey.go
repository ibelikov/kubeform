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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// KmsCryptoKeysGetter has a method to return a KmsCryptoKeyInterface.
// A group's client should implement this interface.
type KmsCryptoKeysGetter interface {
	KmsCryptoKeys() KmsCryptoKeyInterface
}

// KmsCryptoKeyInterface has methods to work with KmsCryptoKey resources.
type KmsCryptoKeyInterface interface {
	Create(*v1alpha1.KmsCryptoKey) (*v1alpha1.KmsCryptoKey, error)
	Update(*v1alpha1.KmsCryptoKey) (*v1alpha1.KmsCryptoKey, error)
	UpdateStatus(*v1alpha1.KmsCryptoKey) (*v1alpha1.KmsCryptoKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KmsCryptoKey, error)
	List(opts v1.ListOptions) (*v1alpha1.KmsCryptoKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCryptoKey, err error)
	KmsCryptoKeyExpansion
}

// kmsCryptoKeys implements KmsCryptoKeyInterface
type kmsCryptoKeys struct {
	client rest.Interface
}

// newKmsCryptoKeys returns a KmsCryptoKeys
func newKmsCryptoKeys(c *GoogleV1alpha1Client) *kmsCryptoKeys {
	return &kmsCryptoKeys{
		client: c.RESTClient(),
	}
}

// Get takes name of the kmsCryptoKey, and returns the corresponding kmsCryptoKey object, and an error if there is any.
func (c *kmsCryptoKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsCryptoKey, err error) {
	result = &v1alpha1.KmsCryptoKey{}
	err = c.client.Get().
		Resource("kmscryptokeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KmsCryptoKeys that match those selectors.
func (c *kmsCryptoKeys) List(opts v1.ListOptions) (result *v1alpha1.KmsCryptoKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KmsCryptoKeyList{}
	err = c.client.Get().
		Resource("kmscryptokeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kmsCryptoKeys.
func (c *kmsCryptoKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("kmscryptokeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kmsCryptoKey and creates it.  Returns the server's representation of the kmsCryptoKey, and an error, if there is any.
func (c *kmsCryptoKeys) Create(kmsCryptoKey *v1alpha1.KmsCryptoKey) (result *v1alpha1.KmsCryptoKey, err error) {
	result = &v1alpha1.KmsCryptoKey{}
	err = c.client.Post().
		Resource("kmscryptokeys").
		Body(kmsCryptoKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kmsCryptoKey and updates it. Returns the server's representation of the kmsCryptoKey, and an error, if there is any.
func (c *kmsCryptoKeys) Update(kmsCryptoKey *v1alpha1.KmsCryptoKey) (result *v1alpha1.KmsCryptoKey, err error) {
	result = &v1alpha1.KmsCryptoKey{}
	err = c.client.Put().
		Resource("kmscryptokeys").
		Name(kmsCryptoKey.Name).
		Body(kmsCryptoKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kmsCryptoKeys) UpdateStatus(kmsCryptoKey *v1alpha1.KmsCryptoKey) (result *v1alpha1.KmsCryptoKey, err error) {
	result = &v1alpha1.KmsCryptoKey{}
	err = c.client.Put().
		Resource("kmscryptokeys").
		Name(kmsCryptoKey.Name).
		SubResource("status").
		Body(kmsCryptoKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the kmsCryptoKey and deletes it. Returns an error if one occurs.
func (c *kmsCryptoKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("kmscryptokeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kmsCryptoKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("kmscryptokeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kmsCryptoKey.
func (c *kmsCryptoKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCryptoKey, err error) {
	result = &v1alpha1.KmsCryptoKey{}
	err = c.client.Patch(pt).
		Resource("kmscryptokeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
