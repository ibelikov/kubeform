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

// AwsAmiCopiesGetter has a method to return a AwsAmiCopyInterface.
// A group's client should implement this interface.
type AwsAmiCopiesGetter interface {
	AwsAmiCopies() AwsAmiCopyInterface
}

// AwsAmiCopyInterface has methods to work with AwsAmiCopy resources.
type AwsAmiCopyInterface interface {
	Create(*v1alpha1.AwsAmiCopy) (*v1alpha1.AwsAmiCopy, error)
	Update(*v1alpha1.AwsAmiCopy) (*v1alpha1.AwsAmiCopy, error)
	UpdateStatus(*v1alpha1.AwsAmiCopy) (*v1alpha1.AwsAmiCopy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAmiCopy, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAmiCopyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmiCopy, err error)
	AwsAmiCopyExpansion
}

// awsAmiCopies implements AwsAmiCopyInterface
type awsAmiCopies struct {
	client rest.Interface
}

// newAwsAmiCopies returns a AwsAmiCopies
func newAwsAmiCopies(c *AwsV1alpha1Client) *awsAmiCopies {
	return &awsAmiCopies{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsAmiCopy, and returns the corresponding awsAmiCopy object, and an error if there is any.
func (c *awsAmiCopies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAmiCopy, err error) {
	result = &v1alpha1.AwsAmiCopy{}
	err = c.client.Get().
		Resource("awsamicopies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAmiCopies that match those selectors.
func (c *awsAmiCopies) List(opts v1.ListOptions) (result *v1alpha1.AwsAmiCopyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsAmiCopyList{}
	err = c.client.Get().
		Resource("awsamicopies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAmiCopies.
func (c *awsAmiCopies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsamicopies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsAmiCopy and creates it.  Returns the server's representation of the awsAmiCopy, and an error, if there is any.
func (c *awsAmiCopies) Create(awsAmiCopy *v1alpha1.AwsAmiCopy) (result *v1alpha1.AwsAmiCopy, err error) {
	result = &v1alpha1.AwsAmiCopy{}
	err = c.client.Post().
		Resource("awsamicopies").
		Body(awsAmiCopy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAmiCopy and updates it. Returns the server's representation of the awsAmiCopy, and an error, if there is any.
func (c *awsAmiCopies) Update(awsAmiCopy *v1alpha1.AwsAmiCopy) (result *v1alpha1.AwsAmiCopy, err error) {
	result = &v1alpha1.AwsAmiCopy{}
	err = c.client.Put().
		Resource("awsamicopies").
		Name(awsAmiCopy.Name).
		Body(awsAmiCopy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsAmiCopies) UpdateStatus(awsAmiCopy *v1alpha1.AwsAmiCopy) (result *v1alpha1.AwsAmiCopy, err error) {
	result = &v1alpha1.AwsAmiCopy{}
	err = c.client.Put().
		Resource("awsamicopies").
		Name(awsAmiCopy.Name).
		SubResource("status").
		Body(awsAmiCopy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAmiCopy and deletes it. Returns an error if one occurs.
func (c *awsAmiCopies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsamicopies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAmiCopies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsamicopies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAmiCopy.
func (c *awsAmiCopies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAmiCopy, err error) {
	result = &v1alpha1.AwsAmiCopy{}
	err = c.client.Patch(pt).
		Resource("awsamicopies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
