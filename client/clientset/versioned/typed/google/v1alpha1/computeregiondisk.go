/*
Copyright The Kubeform Authors.

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

// ComputeRegionDisksGetter has a method to return a ComputeRegionDiskInterface.
// A group's client should implement this interface.
type ComputeRegionDisksGetter interface {
	ComputeRegionDisks(namespace string) ComputeRegionDiskInterface
}

// ComputeRegionDiskInterface has methods to work with ComputeRegionDisk resources.
type ComputeRegionDiskInterface interface {
	Create(*v1alpha1.ComputeRegionDisk) (*v1alpha1.ComputeRegionDisk, error)
	Update(*v1alpha1.ComputeRegionDisk) (*v1alpha1.ComputeRegionDisk, error)
	UpdateStatus(*v1alpha1.ComputeRegionDisk) (*v1alpha1.ComputeRegionDisk, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeRegionDisk, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeRegionDiskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRegionDisk, err error)
	ComputeRegionDiskExpansion
}

// computeRegionDisks implements ComputeRegionDiskInterface
type computeRegionDisks struct {
	client rest.Interface
	ns     string
}

// newComputeRegionDisks returns a ComputeRegionDisks
func newComputeRegionDisks(c *GoogleV1alpha1Client, namespace string) *computeRegionDisks {
	return &computeRegionDisks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeRegionDisk, and returns the corresponding computeRegionDisk object, and an error if there is any.
func (c *computeRegionDisks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRegionDisk, err error) {
	result = &v1alpha1.ComputeRegionDisk{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeregiondisks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeRegionDisks that match those selectors.
func (c *computeRegionDisks) List(opts v1.ListOptions) (result *v1alpha1.ComputeRegionDiskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeRegionDiskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeregiondisks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeRegionDisks.
func (c *computeRegionDisks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeregiondisks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeRegionDisk and creates it.  Returns the server's representation of the computeRegionDisk, and an error, if there is any.
func (c *computeRegionDisks) Create(computeRegionDisk *v1alpha1.ComputeRegionDisk) (result *v1alpha1.ComputeRegionDisk, err error) {
	result = &v1alpha1.ComputeRegionDisk{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeregiondisks").
		Body(computeRegionDisk).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeRegionDisk and updates it. Returns the server's representation of the computeRegionDisk, and an error, if there is any.
func (c *computeRegionDisks) Update(computeRegionDisk *v1alpha1.ComputeRegionDisk) (result *v1alpha1.ComputeRegionDisk, err error) {
	result = &v1alpha1.ComputeRegionDisk{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeregiondisks").
		Name(computeRegionDisk.Name).
		Body(computeRegionDisk).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeRegionDisks) UpdateStatus(computeRegionDisk *v1alpha1.ComputeRegionDisk) (result *v1alpha1.ComputeRegionDisk, err error) {
	result = &v1alpha1.ComputeRegionDisk{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeregiondisks").
		Name(computeRegionDisk.Name).
		SubResource("status").
		Body(computeRegionDisk).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeRegionDisk and deletes it. Returns an error if one occurs.
func (c *computeRegionDisks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeregiondisks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeRegionDisks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeregiondisks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeRegionDisk.
func (c *computeRegionDisks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRegionDisk, err error) {
	result = &v1alpha1.ComputeRegionDisk{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeregiondisks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
