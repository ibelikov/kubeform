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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// S3BucketInventoriesGetter has a method to return a S3BucketInventoryInterface.
// A group's client should implement this interface.
type S3BucketInventoriesGetter interface {
	S3BucketInventories(namespace string) S3BucketInventoryInterface
}

// S3BucketInventoryInterface has methods to work with S3BucketInventory resources.
type S3BucketInventoryInterface interface {
	Create(*v1alpha1.S3BucketInventory) (*v1alpha1.S3BucketInventory, error)
	Update(*v1alpha1.S3BucketInventory) (*v1alpha1.S3BucketInventory, error)
	UpdateStatus(*v1alpha1.S3BucketInventory) (*v1alpha1.S3BucketInventory, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.S3BucketInventory, error)
	List(opts v1.ListOptions) (*v1alpha1.S3BucketInventoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketInventory, err error)
	S3BucketInventoryExpansion
}

// s3BucketInventories implements S3BucketInventoryInterface
type s3BucketInventories struct {
	client rest.Interface
	ns     string
}

// newS3BucketInventories returns a S3BucketInventories
func newS3BucketInventories(c *AwsV1alpha1Client, namespace string) *s3BucketInventories {
	return &s3BucketInventories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3BucketInventory, and returns the corresponding s3BucketInventory object, and an error if there is any.
func (c *s3BucketInventories) Get(name string, options v1.GetOptions) (result *v1alpha1.S3BucketInventory, err error) {
	result = &v1alpha1.S3BucketInventory{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3BucketInventories that match those selectors.
func (c *s3BucketInventories) List(opts v1.ListOptions) (result *v1alpha1.S3BucketInventoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S3BucketInventoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3BucketInventories.
func (c *s3BucketInventories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a s3BucketInventory and creates it.  Returns the server's representation of the s3BucketInventory, and an error, if there is any.
func (c *s3BucketInventories) Create(s3BucketInventory *v1alpha1.S3BucketInventory) (result *v1alpha1.S3BucketInventory, err error) {
	result = &v1alpha1.S3BucketInventory{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		Body(s3BucketInventory).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s3BucketInventory and updates it. Returns the server's representation of the s3BucketInventory, and an error, if there is any.
func (c *s3BucketInventories) Update(s3BucketInventory *v1alpha1.S3BucketInventory) (result *v1alpha1.S3BucketInventory, err error) {
	result = &v1alpha1.S3BucketInventory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		Name(s3BucketInventory.Name).
		Body(s3BucketInventory).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *s3BucketInventories) UpdateStatus(s3BucketInventory *v1alpha1.S3BucketInventory) (result *v1alpha1.S3BucketInventory, err error) {
	result = &v1alpha1.S3BucketInventory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		Name(s3BucketInventory.Name).
		SubResource("status").
		Body(s3BucketInventory).
		Do().
		Into(result)
	return
}

// Delete takes name of the s3BucketInventory and deletes it. Returns an error if one occurs.
func (c *s3BucketInventories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3BucketInventories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketinventories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s3BucketInventory.
func (c *s3BucketInventories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketInventory, err error) {
	result = &v1alpha1.S3BucketInventory{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3bucketinventories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
