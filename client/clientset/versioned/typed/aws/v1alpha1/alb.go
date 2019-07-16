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

// AlbsGetter has a method to return a AlbInterface.
// A group's client should implement this interface.
type AlbsGetter interface {
	Albs() AlbInterface
}

// AlbInterface has methods to work with Alb resources.
type AlbInterface interface {
	Create(*v1alpha1.Alb) (*v1alpha1.Alb, error)
	Update(*v1alpha1.Alb) (*v1alpha1.Alb, error)
	UpdateStatus(*v1alpha1.Alb) (*v1alpha1.Alb, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Alb, error)
	List(opts v1.ListOptions) (*v1alpha1.AlbList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Alb, err error)
	AlbExpansion
}

// albs implements AlbInterface
type albs struct {
	client rest.Interface
}

// newAlbs returns a Albs
func newAlbs(c *AwsV1alpha1Client) *albs {
	return &albs{
		client: c.RESTClient(),
	}
}

// Get takes name of the alb, and returns the corresponding alb object, and an error if there is any.
func (c *albs) Get(name string, options v1.GetOptions) (result *v1alpha1.Alb, err error) {
	result = &v1alpha1.Alb{}
	err = c.client.Get().
		Resource("albs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Albs that match those selectors.
func (c *albs) List(opts v1.ListOptions) (result *v1alpha1.AlbList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AlbList{}
	err = c.client.Get().
		Resource("albs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested albs.
func (c *albs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("albs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a alb and creates it.  Returns the server's representation of the alb, and an error, if there is any.
func (c *albs) Create(alb *v1alpha1.Alb) (result *v1alpha1.Alb, err error) {
	result = &v1alpha1.Alb{}
	err = c.client.Post().
		Resource("albs").
		Body(alb).
		Do().
		Into(result)
	return
}

// Update takes the representation of a alb and updates it. Returns the server's representation of the alb, and an error, if there is any.
func (c *albs) Update(alb *v1alpha1.Alb) (result *v1alpha1.Alb, err error) {
	result = &v1alpha1.Alb{}
	err = c.client.Put().
		Resource("albs").
		Name(alb.Name).
		Body(alb).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *albs) UpdateStatus(alb *v1alpha1.Alb) (result *v1alpha1.Alb, err error) {
	result = &v1alpha1.Alb{}
	err = c.client.Put().
		Resource("albs").
		Name(alb.Name).
		SubResource("status").
		Body(alb).
		Do().
		Into(result)
	return
}

// Delete takes name of the alb and deletes it. Returns an error if one occurs.
func (c *albs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("albs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *albs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("albs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched alb.
func (c *albs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Alb, err error) {
	result = &v1alpha1.Alb{}
	err = c.client.Patch(pt).
		Resource("albs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
