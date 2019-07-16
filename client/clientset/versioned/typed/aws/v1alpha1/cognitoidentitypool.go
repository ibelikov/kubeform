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

// CognitoIdentityPoolsGetter has a method to return a CognitoIdentityPoolInterface.
// A group's client should implement this interface.
type CognitoIdentityPoolsGetter interface {
	CognitoIdentityPools() CognitoIdentityPoolInterface
}

// CognitoIdentityPoolInterface has methods to work with CognitoIdentityPool resources.
type CognitoIdentityPoolInterface interface {
	Create(*v1alpha1.CognitoIdentityPool) (*v1alpha1.CognitoIdentityPool, error)
	Update(*v1alpha1.CognitoIdentityPool) (*v1alpha1.CognitoIdentityPool, error)
	UpdateStatus(*v1alpha1.CognitoIdentityPool) (*v1alpha1.CognitoIdentityPool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CognitoIdentityPool, error)
	List(opts v1.ListOptions) (*v1alpha1.CognitoIdentityPoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoIdentityPool, err error)
	CognitoIdentityPoolExpansion
}

// cognitoIdentityPools implements CognitoIdentityPoolInterface
type cognitoIdentityPools struct {
	client rest.Interface
}

// newCognitoIdentityPools returns a CognitoIdentityPools
func newCognitoIdentityPools(c *AwsV1alpha1Client) *cognitoIdentityPools {
	return &cognitoIdentityPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the cognitoIdentityPool, and returns the corresponding cognitoIdentityPool object, and an error if there is any.
func (c *cognitoIdentityPools) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitoIdentityPool, err error) {
	result = &v1alpha1.CognitoIdentityPool{}
	err = c.client.Get().
		Resource("cognitoidentitypools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CognitoIdentityPools that match those selectors.
func (c *cognitoIdentityPools) List(opts v1.ListOptions) (result *v1alpha1.CognitoIdentityPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CognitoIdentityPoolList{}
	err = c.client.Get().
		Resource("cognitoidentitypools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cognitoIdentityPools.
func (c *cognitoIdentityPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cognitoidentitypools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cognitoIdentityPool and creates it.  Returns the server's representation of the cognitoIdentityPool, and an error, if there is any.
func (c *cognitoIdentityPools) Create(cognitoIdentityPool *v1alpha1.CognitoIdentityPool) (result *v1alpha1.CognitoIdentityPool, err error) {
	result = &v1alpha1.CognitoIdentityPool{}
	err = c.client.Post().
		Resource("cognitoidentitypools").
		Body(cognitoIdentityPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cognitoIdentityPool and updates it. Returns the server's representation of the cognitoIdentityPool, and an error, if there is any.
func (c *cognitoIdentityPools) Update(cognitoIdentityPool *v1alpha1.CognitoIdentityPool) (result *v1alpha1.CognitoIdentityPool, err error) {
	result = &v1alpha1.CognitoIdentityPool{}
	err = c.client.Put().
		Resource("cognitoidentitypools").
		Name(cognitoIdentityPool.Name).
		Body(cognitoIdentityPool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cognitoIdentityPools) UpdateStatus(cognitoIdentityPool *v1alpha1.CognitoIdentityPool) (result *v1alpha1.CognitoIdentityPool, err error) {
	result = &v1alpha1.CognitoIdentityPool{}
	err = c.client.Put().
		Resource("cognitoidentitypools").
		Name(cognitoIdentityPool.Name).
		SubResource("status").
		Body(cognitoIdentityPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the cognitoIdentityPool and deletes it. Returns an error if one occurs.
func (c *cognitoIdentityPools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cognitoidentitypools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cognitoIdentityPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cognitoidentitypools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cognitoIdentityPool.
func (c *cognitoIdentityPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoIdentityPool, err error) {
	result = &v1alpha1.CognitoIdentityPool{}
	err = c.client.Patch(pt).
		Resource("cognitoidentitypools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
