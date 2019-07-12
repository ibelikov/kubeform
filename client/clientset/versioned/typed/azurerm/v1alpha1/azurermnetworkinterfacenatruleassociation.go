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

// AzurermNetworkInterfaceNatRuleAssociationsGetter has a method to return a AzurermNetworkInterfaceNatRuleAssociationInterface.
// A group's client should implement this interface.
type AzurermNetworkInterfaceNatRuleAssociationsGetter interface {
	AzurermNetworkInterfaceNatRuleAssociations() AzurermNetworkInterfaceNatRuleAssociationInterface
}

// AzurermNetworkInterfaceNatRuleAssociationInterface has methods to work with AzurermNetworkInterfaceNatRuleAssociation resources.
type AzurermNetworkInterfaceNatRuleAssociationInterface interface {
	Create(*v1alpha1.AzurermNetworkInterfaceNatRuleAssociation) (*v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, error)
	Update(*v1alpha1.AzurermNetworkInterfaceNatRuleAssociation) (*v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, error)
	UpdateStatus(*v1alpha1.AzurermNetworkInterfaceNatRuleAssociation) (*v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermNetworkInterfaceNatRuleAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, err error)
	AzurermNetworkInterfaceNatRuleAssociationExpansion
}

// azurermNetworkInterfaceNatRuleAssociations implements AzurermNetworkInterfaceNatRuleAssociationInterface
type azurermNetworkInterfaceNatRuleAssociations struct {
	client rest.Interface
}

// newAzurermNetworkInterfaceNatRuleAssociations returns a AzurermNetworkInterfaceNatRuleAssociations
func newAzurermNetworkInterfaceNatRuleAssociations(c *AzurermV1alpha1Client) *azurermNetworkInterfaceNatRuleAssociations {
	return &azurermNetworkInterfaceNatRuleAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermNetworkInterfaceNatRuleAssociation, and returns the corresponding azurermNetworkInterfaceNatRuleAssociation object, and an error if there is any.
func (c *azurermNetworkInterfaceNatRuleAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, err error) {
	result = &v1alpha1.AzurermNetworkInterfaceNatRuleAssociation{}
	err = c.client.Get().
		Resource("azurermnetworkinterfacenatruleassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermNetworkInterfaceNatRuleAssociations that match those selectors.
func (c *azurermNetworkInterfaceNatRuleAssociations) List(opts v1.ListOptions) (result *v1alpha1.AzurermNetworkInterfaceNatRuleAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermNetworkInterfaceNatRuleAssociationList{}
	err = c.client.Get().
		Resource("azurermnetworkinterfacenatruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermNetworkInterfaceNatRuleAssociations.
func (c *azurermNetworkInterfaceNatRuleAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermnetworkinterfacenatruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermNetworkInterfaceNatRuleAssociation and creates it.  Returns the server's representation of the azurermNetworkInterfaceNatRuleAssociation, and an error, if there is any.
func (c *azurermNetworkInterfaceNatRuleAssociations) Create(azurermNetworkInterfaceNatRuleAssociation *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation) (result *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, err error) {
	result = &v1alpha1.AzurermNetworkInterfaceNatRuleAssociation{}
	err = c.client.Post().
		Resource("azurermnetworkinterfacenatruleassociations").
		Body(azurermNetworkInterfaceNatRuleAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermNetworkInterfaceNatRuleAssociation and updates it. Returns the server's representation of the azurermNetworkInterfaceNatRuleAssociation, and an error, if there is any.
func (c *azurermNetworkInterfaceNatRuleAssociations) Update(azurermNetworkInterfaceNatRuleAssociation *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation) (result *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, err error) {
	result = &v1alpha1.AzurermNetworkInterfaceNatRuleAssociation{}
	err = c.client.Put().
		Resource("azurermnetworkinterfacenatruleassociations").
		Name(azurermNetworkInterfaceNatRuleAssociation.Name).
		Body(azurermNetworkInterfaceNatRuleAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermNetworkInterfaceNatRuleAssociations) UpdateStatus(azurermNetworkInterfaceNatRuleAssociation *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation) (result *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, err error) {
	result = &v1alpha1.AzurermNetworkInterfaceNatRuleAssociation{}
	err = c.client.Put().
		Resource("azurermnetworkinterfacenatruleassociations").
		Name(azurermNetworkInterfaceNatRuleAssociation.Name).
		SubResource("status").
		Body(azurermNetworkInterfaceNatRuleAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermNetworkInterfaceNatRuleAssociation and deletes it. Returns an error if one occurs.
func (c *azurermNetworkInterfaceNatRuleAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermnetworkinterfacenatruleassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermNetworkInterfaceNatRuleAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermnetworkinterfacenatruleassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermNetworkInterfaceNatRuleAssociation.
func (c *azurermNetworkInterfaceNatRuleAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermNetworkInterfaceNatRuleAssociation, err error) {
	result = &v1alpha1.AzurermNetworkInterfaceNatRuleAssociation{}
	err = c.client.Patch(pt).
		Resource("azurermnetworkinterfacenatruleassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
