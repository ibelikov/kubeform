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

// AzurermUserAssignedIdentitiesGetter has a method to return a AzurermUserAssignedIdentityInterface.
// A group's client should implement this interface.
type AzurermUserAssignedIdentitiesGetter interface {
	AzurermUserAssignedIdentities() AzurermUserAssignedIdentityInterface
}

// AzurermUserAssignedIdentityInterface has methods to work with AzurermUserAssignedIdentity resources.
type AzurermUserAssignedIdentityInterface interface {
	Create(*v1alpha1.AzurermUserAssignedIdentity) (*v1alpha1.AzurermUserAssignedIdentity, error)
	Update(*v1alpha1.AzurermUserAssignedIdentity) (*v1alpha1.AzurermUserAssignedIdentity, error)
	UpdateStatus(*v1alpha1.AzurermUserAssignedIdentity) (*v1alpha1.AzurermUserAssignedIdentity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermUserAssignedIdentity, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermUserAssignedIdentityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermUserAssignedIdentity, err error)
	AzurermUserAssignedIdentityExpansion
}

// azurermUserAssignedIdentities implements AzurermUserAssignedIdentityInterface
type azurermUserAssignedIdentities struct {
	client rest.Interface
}

// newAzurermUserAssignedIdentities returns a AzurermUserAssignedIdentities
func newAzurermUserAssignedIdentities(c *AzurermV1alpha1Client) *azurermUserAssignedIdentities {
	return &azurermUserAssignedIdentities{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermUserAssignedIdentity, and returns the corresponding azurermUserAssignedIdentity object, and an error if there is any.
func (c *azurermUserAssignedIdentities) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermUserAssignedIdentity, err error) {
	result = &v1alpha1.AzurermUserAssignedIdentity{}
	err = c.client.Get().
		Resource("azurermuserassignedidentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermUserAssignedIdentities that match those selectors.
func (c *azurermUserAssignedIdentities) List(opts v1.ListOptions) (result *v1alpha1.AzurermUserAssignedIdentityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermUserAssignedIdentityList{}
	err = c.client.Get().
		Resource("azurermuserassignedidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermUserAssignedIdentities.
func (c *azurermUserAssignedIdentities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermuserassignedidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermUserAssignedIdentity and creates it.  Returns the server's representation of the azurermUserAssignedIdentity, and an error, if there is any.
func (c *azurermUserAssignedIdentities) Create(azurermUserAssignedIdentity *v1alpha1.AzurermUserAssignedIdentity) (result *v1alpha1.AzurermUserAssignedIdentity, err error) {
	result = &v1alpha1.AzurermUserAssignedIdentity{}
	err = c.client.Post().
		Resource("azurermuserassignedidentities").
		Body(azurermUserAssignedIdentity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermUserAssignedIdentity and updates it. Returns the server's representation of the azurermUserAssignedIdentity, and an error, if there is any.
func (c *azurermUserAssignedIdentities) Update(azurermUserAssignedIdentity *v1alpha1.AzurermUserAssignedIdentity) (result *v1alpha1.AzurermUserAssignedIdentity, err error) {
	result = &v1alpha1.AzurermUserAssignedIdentity{}
	err = c.client.Put().
		Resource("azurermuserassignedidentities").
		Name(azurermUserAssignedIdentity.Name).
		Body(azurermUserAssignedIdentity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermUserAssignedIdentities) UpdateStatus(azurermUserAssignedIdentity *v1alpha1.AzurermUserAssignedIdentity) (result *v1alpha1.AzurermUserAssignedIdentity, err error) {
	result = &v1alpha1.AzurermUserAssignedIdentity{}
	err = c.client.Put().
		Resource("azurermuserassignedidentities").
		Name(azurermUserAssignedIdentity.Name).
		SubResource("status").
		Body(azurermUserAssignedIdentity).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermUserAssignedIdentity and deletes it. Returns an error if one occurs.
func (c *azurermUserAssignedIdentities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermuserassignedidentities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermUserAssignedIdentities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermuserassignedidentities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermUserAssignedIdentity.
func (c *azurermUserAssignedIdentities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermUserAssignedIdentity, err error) {
	result = &v1alpha1.AzurermUserAssignedIdentity{}
	err = c.client.Patch(pt).
		Resource("azurermuserassignedidentities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
