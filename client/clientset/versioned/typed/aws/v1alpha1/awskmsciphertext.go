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

// AwsKmsCiphertextsGetter has a method to return a AwsKmsCiphertextInterface.
// A group's client should implement this interface.
type AwsKmsCiphertextsGetter interface {
	AwsKmsCiphertexts() AwsKmsCiphertextInterface
}

// AwsKmsCiphertextInterface has methods to work with AwsKmsCiphertext resources.
type AwsKmsCiphertextInterface interface {
	Create(*v1alpha1.AwsKmsCiphertext) (*v1alpha1.AwsKmsCiphertext, error)
	Update(*v1alpha1.AwsKmsCiphertext) (*v1alpha1.AwsKmsCiphertext, error)
	UpdateStatus(*v1alpha1.AwsKmsCiphertext) (*v1alpha1.AwsKmsCiphertext, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsKmsCiphertext, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsKmsCiphertextList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKmsCiphertext, err error)
	AwsKmsCiphertextExpansion
}

// awsKmsCiphertexts implements AwsKmsCiphertextInterface
type awsKmsCiphertexts struct {
	client rest.Interface
}

// newAwsKmsCiphertexts returns a AwsKmsCiphertexts
func newAwsKmsCiphertexts(c *AwsV1alpha1Client) *awsKmsCiphertexts {
	return &awsKmsCiphertexts{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsKmsCiphertext, and returns the corresponding awsKmsCiphertext object, and an error if there is any.
func (c *awsKmsCiphertexts) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsKmsCiphertext, err error) {
	result = &v1alpha1.AwsKmsCiphertext{}
	err = c.client.Get().
		Resource("awskmsciphertexts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsKmsCiphertexts that match those selectors.
func (c *awsKmsCiphertexts) List(opts v1.ListOptions) (result *v1alpha1.AwsKmsCiphertextList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsKmsCiphertextList{}
	err = c.client.Get().
		Resource("awskmsciphertexts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsKmsCiphertexts.
func (c *awsKmsCiphertexts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awskmsciphertexts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsKmsCiphertext and creates it.  Returns the server's representation of the awsKmsCiphertext, and an error, if there is any.
func (c *awsKmsCiphertexts) Create(awsKmsCiphertext *v1alpha1.AwsKmsCiphertext) (result *v1alpha1.AwsKmsCiphertext, err error) {
	result = &v1alpha1.AwsKmsCiphertext{}
	err = c.client.Post().
		Resource("awskmsciphertexts").
		Body(awsKmsCiphertext).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsKmsCiphertext and updates it. Returns the server's representation of the awsKmsCiphertext, and an error, if there is any.
func (c *awsKmsCiphertexts) Update(awsKmsCiphertext *v1alpha1.AwsKmsCiphertext) (result *v1alpha1.AwsKmsCiphertext, err error) {
	result = &v1alpha1.AwsKmsCiphertext{}
	err = c.client.Put().
		Resource("awskmsciphertexts").
		Name(awsKmsCiphertext.Name).
		Body(awsKmsCiphertext).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsKmsCiphertexts) UpdateStatus(awsKmsCiphertext *v1alpha1.AwsKmsCiphertext) (result *v1alpha1.AwsKmsCiphertext, err error) {
	result = &v1alpha1.AwsKmsCiphertext{}
	err = c.client.Put().
		Resource("awskmsciphertexts").
		Name(awsKmsCiphertext.Name).
		SubResource("status").
		Body(awsKmsCiphertext).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsKmsCiphertext and deletes it. Returns an error if one occurs.
func (c *awsKmsCiphertexts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awskmsciphertexts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsKmsCiphertexts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awskmsciphertexts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsKmsCiphertext.
func (c *awsKmsCiphertexts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsKmsCiphertext, err error) {
	result = &v1alpha1.AwsKmsCiphertext{}
	err = c.client.Patch(pt).
		Resource("awskmsciphertexts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
