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

// BatchComputeEnvironmentsGetter has a method to return a BatchComputeEnvironmentInterface.
// A group's client should implement this interface.
type BatchComputeEnvironmentsGetter interface {
	BatchComputeEnvironments(namespace string) BatchComputeEnvironmentInterface
}

// BatchComputeEnvironmentInterface has methods to work with BatchComputeEnvironment resources.
type BatchComputeEnvironmentInterface interface {
	Create(*v1alpha1.BatchComputeEnvironment) (*v1alpha1.BatchComputeEnvironment, error)
	Update(*v1alpha1.BatchComputeEnvironment) (*v1alpha1.BatchComputeEnvironment, error)
	UpdateStatus(*v1alpha1.BatchComputeEnvironment) (*v1alpha1.BatchComputeEnvironment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BatchComputeEnvironment, error)
	List(opts v1.ListOptions) (*v1alpha1.BatchComputeEnvironmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchComputeEnvironment, err error)
	BatchComputeEnvironmentExpansion
}

// batchComputeEnvironments implements BatchComputeEnvironmentInterface
type batchComputeEnvironments struct {
	client rest.Interface
	ns     string
}

// newBatchComputeEnvironments returns a BatchComputeEnvironments
func newBatchComputeEnvironments(c *AwsV1alpha1Client, namespace string) *batchComputeEnvironments {
	return &batchComputeEnvironments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the batchComputeEnvironment, and returns the corresponding batchComputeEnvironment object, and an error if there is any.
func (c *batchComputeEnvironments) Get(name string, options v1.GetOptions) (result *v1alpha1.BatchComputeEnvironment, err error) {
	result = &v1alpha1.BatchComputeEnvironment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BatchComputeEnvironments that match those selectors.
func (c *batchComputeEnvironments) List(opts v1.ListOptions) (result *v1alpha1.BatchComputeEnvironmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BatchComputeEnvironmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested batchComputeEnvironments.
func (c *batchComputeEnvironments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a batchComputeEnvironment and creates it.  Returns the server's representation of the batchComputeEnvironment, and an error, if there is any.
func (c *batchComputeEnvironments) Create(batchComputeEnvironment *v1alpha1.BatchComputeEnvironment) (result *v1alpha1.BatchComputeEnvironment, err error) {
	result = &v1alpha1.BatchComputeEnvironment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		Body(batchComputeEnvironment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a batchComputeEnvironment and updates it. Returns the server's representation of the batchComputeEnvironment, and an error, if there is any.
func (c *batchComputeEnvironments) Update(batchComputeEnvironment *v1alpha1.BatchComputeEnvironment) (result *v1alpha1.BatchComputeEnvironment, err error) {
	result = &v1alpha1.BatchComputeEnvironment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		Name(batchComputeEnvironment.Name).
		Body(batchComputeEnvironment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *batchComputeEnvironments) UpdateStatus(batchComputeEnvironment *v1alpha1.BatchComputeEnvironment) (result *v1alpha1.BatchComputeEnvironment, err error) {
	result = &v1alpha1.BatchComputeEnvironment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		Name(batchComputeEnvironment.Name).
		SubResource("status").
		Body(batchComputeEnvironment).
		Do().
		Into(result)
	return
}

// Delete takes name of the batchComputeEnvironment and deletes it. Returns an error if one occurs.
func (c *batchComputeEnvironments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *batchComputeEnvironments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched batchComputeEnvironment.
func (c *batchComputeEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchComputeEnvironment, err error) {
	result = &v1alpha1.BatchComputeEnvironment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("batchcomputeenvironments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
