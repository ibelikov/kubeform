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

// ComputeHealthChecksGetter has a method to return a ComputeHealthCheckInterface.
// A group's client should implement this interface.
type ComputeHealthChecksGetter interface {
	ComputeHealthChecks() ComputeHealthCheckInterface
}

// ComputeHealthCheckInterface has methods to work with ComputeHealthCheck resources.
type ComputeHealthCheckInterface interface {
	Create(*v1alpha1.ComputeHealthCheck) (*v1alpha1.ComputeHealthCheck, error)
	Update(*v1alpha1.ComputeHealthCheck) (*v1alpha1.ComputeHealthCheck, error)
	UpdateStatus(*v1alpha1.ComputeHealthCheck) (*v1alpha1.ComputeHealthCheck, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeHealthCheck, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeHealthCheckList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeHealthCheck, err error)
	ComputeHealthCheckExpansion
}

// computeHealthChecks implements ComputeHealthCheckInterface
type computeHealthChecks struct {
	client rest.Interface
}

// newComputeHealthChecks returns a ComputeHealthChecks
func newComputeHealthChecks(c *GoogleV1alpha1Client) *computeHealthChecks {
	return &computeHealthChecks{
		client: c.RESTClient(),
	}
}

// Get takes name of the computeHealthCheck, and returns the corresponding computeHealthCheck object, and an error if there is any.
func (c *computeHealthChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeHealthCheck, err error) {
	result = &v1alpha1.ComputeHealthCheck{}
	err = c.client.Get().
		Resource("computehealthchecks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeHealthChecks that match those selectors.
func (c *computeHealthChecks) List(opts v1.ListOptions) (result *v1alpha1.ComputeHealthCheckList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeHealthCheckList{}
	err = c.client.Get().
		Resource("computehealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeHealthChecks.
func (c *computeHealthChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("computehealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeHealthCheck and creates it.  Returns the server's representation of the computeHealthCheck, and an error, if there is any.
func (c *computeHealthChecks) Create(computeHealthCheck *v1alpha1.ComputeHealthCheck) (result *v1alpha1.ComputeHealthCheck, err error) {
	result = &v1alpha1.ComputeHealthCheck{}
	err = c.client.Post().
		Resource("computehealthchecks").
		Body(computeHealthCheck).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeHealthCheck and updates it. Returns the server's representation of the computeHealthCheck, and an error, if there is any.
func (c *computeHealthChecks) Update(computeHealthCheck *v1alpha1.ComputeHealthCheck) (result *v1alpha1.ComputeHealthCheck, err error) {
	result = &v1alpha1.ComputeHealthCheck{}
	err = c.client.Put().
		Resource("computehealthchecks").
		Name(computeHealthCheck.Name).
		Body(computeHealthCheck).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeHealthChecks) UpdateStatus(computeHealthCheck *v1alpha1.ComputeHealthCheck) (result *v1alpha1.ComputeHealthCheck, err error) {
	result = &v1alpha1.ComputeHealthCheck{}
	err = c.client.Put().
		Resource("computehealthchecks").
		Name(computeHealthCheck.Name).
		SubResource("status").
		Body(computeHealthCheck).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeHealthCheck and deletes it. Returns an error if one occurs.
func (c *computeHealthChecks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("computehealthchecks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeHealthChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("computehealthchecks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeHealthCheck.
func (c *computeHealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeHealthCheck, err error) {
	result = &v1alpha1.ComputeHealthCheck{}
	err = c.client.Patch(pt).
		Resource("computehealthchecks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
