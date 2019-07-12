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

// AwsCognitoUserPoolsGetter has a method to return a AwsCognitoUserPoolInterface.
// A group's client should implement this interface.
type AwsCognitoUserPoolsGetter interface {
	AwsCognitoUserPools() AwsCognitoUserPoolInterface
}

// AwsCognitoUserPoolInterface has methods to work with AwsCognitoUserPool resources.
type AwsCognitoUserPoolInterface interface {
	Create(*v1alpha1.AwsCognitoUserPool) (*v1alpha1.AwsCognitoUserPool, error)
	Update(*v1alpha1.AwsCognitoUserPool) (*v1alpha1.AwsCognitoUserPool, error)
	UpdateStatus(*v1alpha1.AwsCognitoUserPool) (*v1alpha1.AwsCognitoUserPool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCognitoUserPool, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCognitoUserPoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoUserPool, err error)
	AwsCognitoUserPoolExpansion
}

// awsCognitoUserPools implements AwsCognitoUserPoolInterface
type awsCognitoUserPools struct {
	client rest.Interface
}

// newAwsCognitoUserPools returns a AwsCognitoUserPools
func newAwsCognitoUserPools(c *AwsV1alpha1Client) *awsCognitoUserPools {
	return &awsCognitoUserPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCognitoUserPool, and returns the corresponding awsCognitoUserPool object, and an error if there is any.
func (c *awsCognitoUserPools) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCognitoUserPool, err error) {
	result = &v1alpha1.AwsCognitoUserPool{}
	err = c.client.Get().
		Resource("awscognitouserpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCognitoUserPools that match those selectors.
func (c *awsCognitoUserPools) List(opts v1.ListOptions) (result *v1alpha1.AwsCognitoUserPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCognitoUserPoolList{}
	err = c.client.Get().
		Resource("awscognitouserpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCognitoUserPools.
func (c *awsCognitoUserPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscognitouserpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCognitoUserPool and creates it.  Returns the server's representation of the awsCognitoUserPool, and an error, if there is any.
func (c *awsCognitoUserPools) Create(awsCognitoUserPool *v1alpha1.AwsCognitoUserPool) (result *v1alpha1.AwsCognitoUserPool, err error) {
	result = &v1alpha1.AwsCognitoUserPool{}
	err = c.client.Post().
		Resource("awscognitouserpools").
		Body(awsCognitoUserPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCognitoUserPool and updates it. Returns the server's representation of the awsCognitoUserPool, and an error, if there is any.
func (c *awsCognitoUserPools) Update(awsCognitoUserPool *v1alpha1.AwsCognitoUserPool) (result *v1alpha1.AwsCognitoUserPool, err error) {
	result = &v1alpha1.AwsCognitoUserPool{}
	err = c.client.Put().
		Resource("awscognitouserpools").
		Name(awsCognitoUserPool.Name).
		Body(awsCognitoUserPool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCognitoUserPools) UpdateStatus(awsCognitoUserPool *v1alpha1.AwsCognitoUserPool) (result *v1alpha1.AwsCognitoUserPool, err error) {
	result = &v1alpha1.AwsCognitoUserPool{}
	err = c.client.Put().
		Resource("awscognitouserpools").
		Name(awsCognitoUserPool.Name).
		SubResource("status").
		Body(awsCognitoUserPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCognitoUserPool and deletes it. Returns an error if one occurs.
func (c *awsCognitoUserPools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscognitouserpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCognitoUserPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscognitouserpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCognitoUserPool.
func (c *awsCognitoUserPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoUserPool, err error) {
	result = &v1alpha1.AwsCognitoUserPool{}
	err = c.client.Patch(pt).
		Resource("awscognitouserpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
