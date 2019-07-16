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

// CognitoIdentityProvidersGetter has a method to return a CognitoIdentityProviderInterface.
// A group's client should implement this interface.
type CognitoIdentityProvidersGetter interface {
	CognitoIdentityProviders() CognitoIdentityProviderInterface
}

// CognitoIdentityProviderInterface has methods to work with CognitoIdentityProvider resources.
type CognitoIdentityProviderInterface interface {
	Create(*v1alpha1.CognitoIdentityProvider) (*v1alpha1.CognitoIdentityProvider, error)
	Update(*v1alpha1.CognitoIdentityProvider) (*v1alpha1.CognitoIdentityProvider, error)
	UpdateStatus(*v1alpha1.CognitoIdentityProvider) (*v1alpha1.CognitoIdentityProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CognitoIdentityProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.CognitoIdentityProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoIdentityProvider, err error)
	CognitoIdentityProviderExpansion
}

// cognitoIdentityProviders implements CognitoIdentityProviderInterface
type cognitoIdentityProviders struct {
	client rest.Interface
}

// newCognitoIdentityProviders returns a CognitoIdentityProviders
func newCognitoIdentityProviders(c *AwsV1alpha1Client) *cognitoIdentityProviders {
	return &cognitoIdentityProviders{
		client: c.RESTClient(),
	}
}

// Get takes name of the cognitoIdentityProvider, and returns the corresponding cognitoIdentityProvider object, and an error if there is any.
func (c *cognitoIdentityProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitoIdentityProvider, err error) {
	result = &v1alpha1.CognitoIdentityProvider{}
	err = c.client.Get().
		Resource("cognitoidentityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CognitoIdentityProviders that match those selectors.
func (c *cognitoIdentityProviders) List(opts v1.ListOptions) (result *v1alpha1.CognitoIdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CognitoIdentityProviderList{}
	err = c.client.Get().
		Resource("cognitoidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cognitoIdentityProviders.
func (c *cognitoIdentityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cognitoidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cognitoIdentityProvider and creates it.  Returns the server's representation of the cognitoIdentityProvider, and an error, if there is any.
func (c *cognitoIdentityProviders) Create(cognitoIdentityProvider *v1alpha1.CognitoIdentityProvider) (result *v1alpha1.CognitoIdentityProvider, err error) {
	result = &v1alpha1.CognitoIdentityProvider{}
	err = c.client.Post().
		Resource("cognitoidentityproviders").
		Body(cognitoIdentityProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cognitoIdentityProvider and updates it. Returns the server's representation of the cognitoIdentityProvider, and an error, if there is any.
func (c *cognitoIdentityProviders) Update(cognitoIdentityProvider *v1alpha1.CognitoIdentityProvider) (result *v1alpha1.CognitoIdentityProvider, err error) {
	result = &v1alpha1.CognitoIdentityProvider{}
	err = c.client.Put().
		Resource("cognitoidentityproviders").
		Name(cognitoIdentityProvider.Name).
		Body(cognitoIdentityProvider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cognitoIdentityProviders) UpdateStatus(cognitoIdentityProvider *v1alpha1.CognitoIdentityProvider) (result *v1alpha1.CognitoIdentityProvider, err error) {
	result = &v1alpha1.CognitoIdentityProvider{}
	err = c.client.Put().
		Resource("cognitoidentityproviders").
		Name(cognitoIdentityProvider.Name).
		SubResource("status").
		Body(cognitoIdentityProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the cognitoIdentityProvider and deletes it. Returns an error if one occurs.
func (c *cognitoIdentityProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cognitoidentityproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cognitoIdentityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cognitoidentityproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cognitoIdentityProvider.
func (c *cognitoIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoIdentityProvider, err error) {
	result = &v1alpha1.CognitoIdentityProvider{}
	err = c.client.Patch(pt).
		Resource("cognitoidentityproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
