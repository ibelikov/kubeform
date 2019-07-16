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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzureadServicePrincipalsGetter has a method to return a AzureadServicePrincipalInterface.
// A group's client should implement this interface.
type AzureadServicePrincipalsGetter interface {
	AzureadServicePrincipals() AzureadServicePrincipalInterface
}

// AzureadServicePrincipalInterface has methods to work with AzureadServicePrincipal resources.
type AzureadServicePrincipalInterface interface {
	Create(*v1alpha1.AzureadServicePrincipal) (*v1alpha1.AzureadServicePrincipal, error)
	Update(*v1alpha1.AzureadServicePrincipal) (*v1alpha1.AzureadServicePrincipal, error)
	UpdateStatus(*v1alpha1.AzureadServicePrincipal) (*v1alpha1.AzureadServicePrincipal, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzureadServicePrincipal, error)
	List(opts v1.ListOptions) (*v1alpha1.AzureadServicePrincipalList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureadServicePrincipal, err error)
	AzureadServicePrincipalExpansion
}

// azureadServicePrincipals implements AzureadServicePrincipalInterface
type azureadServicePrincipals struct {
	client rest.Interface
}

// newAzureadServicePrincipals returns a AzureadServicePrincipals
func newAzureadServicePrincipals(c *AzurermV1alpha1Client) *azureadServicePrincipals {
	return &azureadServicePrincipals{
		client: c.RESTClient(),
	}
}

// Get takes name of the azureadServicePrincipal, and returns the corresponding azureadServicePrincipal object, and an error if there is any.
func (c *azureadServicePrincipals) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureadServicePrincipal, err error) {
	result = &v1alpha1.AzureadServicePrincipal{}
	err = c.client.Get().
		Resource("azureadserviceprincipals").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureadServicePrincipals that match those selectors.
func (c *azureadServicePrincipals) List(opts v1.ListOptions) (result *v1alpha1.AzureadServicePrincipalList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzureadServicePrincipalList{}
	err = c.client.Get().
		Resource("azureadserviceprincipals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureadServicePrincipals.
func (c *azureadServicePrincipals) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azureadserviceprincipals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azureadServicePrincipal and creates it.  Returns the server's representation of the azureadServicePrincipal, and an error, if there is any.
func (c *azureadServicePrincipals) Create(azureadServicePrincipal *v1alpha1.AzureadServicePrincipal) (result *v1alpha1.AzureadServicePrincipal, err error) {
	result = &v1alpha1.AzureadServicePrincipal{}
	err = c.client.Post().
		Resource("azureadserviceprincipals").
		Body(azureadServicePrincipal).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azureadServicePrincipal and updates it. Returns the server's representation of the azureadServicePrincipal, and an error, if there is any.
func (c *azureadServicePrincipals) Update(azureadServicePrincipal *v1alpha1.AzureadServicePrincipal) (result *v1alpha1.AzureadServicePrincipal, err error) {
	result = &v1alpha1.AzureadServicePrincipal{}
	err = c.client.Put().
		Resource("azureadserviceprincipals").
		Name(azureadServicePrincipal.Name).
		Body(azureadServicePrincipal).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azureadServicePrincipals) UpdateStatus(azureadServicePrincipal *v1alpha1.AzureadServicePrincipal) (result *v1alpha1.AzureadServicePrincipal, err error) {
	result = &v1alpha1.AzureadServicePrincipal{}
	err = c.client.Put().
		Resource("azureadserviceprincipals").
		Name(azureadServicePrincipal.Name).
		SubResource("status").
		Body(azureadServicePrincipal).
		Do().
		Into(result)
	return
}

// Delete takes name of the azureadServicePrincipal and deletes it. Returns an error if one occurs.
func (c *azureadServicePrincipals) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azureadserviceprincipals").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureadServicePrincipals) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azureadserviceprincipals").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azureadServicePrincipal.
func (c *azureadServicePrincipals) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureadServicePrincipal, err error) {
	result = &v1alpha1.AzureadServicePrincipal{}
	err = c.client.Patch(pt).
		Resource("azureadserviceprincipals").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
