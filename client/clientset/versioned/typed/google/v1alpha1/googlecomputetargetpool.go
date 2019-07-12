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

// GoogleComputeTargetPoolsGetter has a method to return a GoogleComputeTargetPoolInterface.
// A group's client should implement this interface.
type GoogleComputeTargetPoolsGetter interface {
	GoogleComputeTargetPools() GoogleComputeTargetPoolInterface
}

// GoogleComputeTargetPoolInterface has methods to work with GoogleComputeTargetPool resources.
type GoogleComputeTargetPoolInterface interface {
	Create(*v1alpha1.GoogleComputeTargetPool) (*v1alpha1.GoogleComputeTargetPool, error)
	Update(*v1alpha1.GoogleComputeTargetPool) (*v1alpha1.GoogleComputeTargetPool, error)
	UpdateStatus(*v1alpha1.GoogleComputeTargetPool) (*v1alpha1.GoogleComputeTargetPool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleComputeTargetPool, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleComputeTargetPoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeTargetPool, err error)
	GoogleComputeTargetPoolExpansion
}

// googleComputeTargetPools implements GoogleComputeTargetPoolInterface
type googleComputeTargetPools struct {
	client rest.Interface
}

// newGoogleComputeTargetPools returns a GoogleComputeTargetPools
func newGoogleComputeTargetPools(c *GoogleV1alpha1Client) *googleComputeTargetPools {
	return &googleComputeTargetPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleComputeTargetPool, and returns the corresponding googleComputeTargetPool object, and an error if there is any.
func (c *googleComputeTargetPools) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeTargetPool, err error) {
	result = &v1alpha1.GoogleComputeTargetPool{}
	err = c.client.Get().
		Resource("googlecomputetargetpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleComputeTargetPools that match those selectors.
func (c *googleComputeTargetPools) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeTargetPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleComputeTargetPoolList{}
	err = c.client.Get().
		Resource("googlecomputetargetpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleComputeTargetPools.
func (c *googleComputeTargetPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlecomputetargetpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleComputeTargetPool and creates it.  Returns the server's representation of the googleComputeTargetPool, and an error, if there is any.
func (c *googleComputeTargetPools) Create(googleComputeTargetPool *v1alpha1.GoogleComputeTargetPool) (result *v1alpha1.GoogleComputeTargetPool, err error) {
	result = &v1alpha1.GoogleComputeTargetPool{}
	err = c.client.Post().
		Resource("googlecomputetargetpools").
		Body(googleComputeTargetPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleComputeTargetPool and updates it. Returns the server's representation of the googleComputeTargetPool, and an error, if there is any.
func (c *googleComputeTargetPools) Update(googleComputeTargetPool *v1alpha1.GoogleComputeTargetPool) (result *v1alpha1.GoogleComputeTargetPool, err error) {
	result = &v1alpha1.GoogleComputeTargetPool{}
	err = c.client.Put().
		Resource("googlecomputetargetpools").
		Name(googleComputeTargetPool.Name).
		Body(googleComputeTargetPool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleComputeTargetPools) UpdateStatus(googleComputeTargetPool *v1alpha1.GoogleComputeTargetPool) (result *v1alpha1.GoogleComputeTargetPool, err error) {
	result = &v1alpha1.GoogleComputeTargetPool{}
	err = c.client.Put().
		Resource("googlecomputetargetpools").
		Name(googleComputeTargetPool.Name).
		SubResource("status").
		Body(googleComputeTargetPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleComputeTargetPool and deletes it. Returns an error if one occurs.
func (c *googleComputeTargetPools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlecomputetargetpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleComputeTargetPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlecomputetargetpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleComputeTargetPool.
func (c *googleComputeTargetPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeTargetPool, err error) {
	result = &v1alpha1.GoogleComputeTargetPool{}
	err = c.client.Patch(pt).
		Resource("googlecomputetargetpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
