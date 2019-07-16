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

// SesConfigurationSetsGetter has a method to return a SesConfigurationSetInterface.
// A group's client should implement this interface.
type SesConfigurationSetsGetter interface {
	SesConfigurationSets() SesConfigurationSetInterface
}

// SesConfigurationSetInterface has methods to work with SesConfigurationSet resources.
type SesConfigurationSetInterface interface {
	Create(*v1alpha1.SesConfigurationSet) (*v1alpha1.SesConfigurationSet, error)
	Update(*v1alpha1.SesConfigurationSet) (*v1alpha1.SesConfigurationSet, error)
	UpdateStatus(*v1alpha1.SesConfigurationSet) (*v1alpha1.SesConfigurationSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesConfigurationSet, error)
	List(opts v1.ListOptions) (*v1alpha1.SesConfigurationSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesConfigurationSet, err error)
	SesConfigurationSetExpansion
}

// sesConfigurationSets implements SesConfigurationSetInterface
type sesConfigurationSets struct {
	client rest.Interface
}

// newSesConfigurationSets returns a SesConfigurationSets
func newSesConfigurationSets(c *AwsV1alpha1Client) *sesConfigurationSets {
	return &sesConfigurationSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the sesConfigurationSet, and returns the corresponding sesConfigurationSet object, and an error if there is any.
func (c *sesConfigurationSets) Get(name string, options v1.GetOptions) (result *v1alpha1.SesConfigurationSet, err error) {
	result = &v1alpha1.SesConfigurationSet{}
	err = c.client.Get().
		Resource("sesconfigurationsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesConfigurationSets that match those selectors.
func (c *sesConfigurationSets) List(opts v1.ListOptions) (result *v1alpha1.SesConfigurationSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesConfigurationSetList{}
	err = c.client.Get().
		Resource("sesconfigurationsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesConfigurationSets.
func (c *sesConfigurationSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sesconfigurationsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesConfigurationSet and creates it.  Returns the server's representation of the sesConfigurationSet, and an error, if there is any.
func (c *sesConfigurationSets) Create(sesConfigurationSet *v1alpha1.SesConfigurationSet) (result *v1alpha1.SesConfigurationSet, err error) {
	result = &v1alpha1.SesConfigurationSet{}
	err = c.client.Post().
		Resource("sesconfigurationsets").
		Body(sesConfigurationSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesConfigurationSet and updates it. Returns the server's representation of the sesConfigurationSet, and an error, if there is any.
func (c *sesConfigurationSets) Update(sesConfigurationSet *v1alpha1.SesConfigurationSet) (result *v1alpha1.SesConfigurationSet, err error) {
	result = &v1alpha1.SesConfigurationSet{}
	err = c.client.Put().
		Resource("sesconfigurationsets").
		Name(sesConfigurationSet.Name).
		Body(sesConfigurationSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesConfigurationSets) UpdateStatus(sesConfigurationSet *v1alpha1.SesConfigurationSet) (result *v1alpha1.SesConfigurationSet, err error) {
	result = &v1alpha1.SesConfigurationSet{}
	err = c.client.Put().
		Resource("sesconfigurationsets").
		Name(sesConfigurationSet.Name).
		SubResource("status").
		Body(sesConfigurationSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesConfigurationSet and deletes it. Returns an error if one occurs.
func (c *sesConfigurationSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sesconfigurationsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesConfigurationSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sesconfigurationsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesConfigurationSet.
func (c *sesConfigurationSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesConfigurationSet, err error) {
	result = &v1alpha1.SesConfigurationSet{}
	err = c.client.Patch(pt).
		Resource("sesconfigurationsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
