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

// AzurermFirewallNetworkRuleCollectionsGetter has a method to return a AzurermFirewallNetworkRuleCollectionInterface.
// A group's client should implement this interface.
type AzurermFirewallNetworkRuleCollectionsGetter interface {
	AzurermFirewallNetworkRuleCollections() AzurermFirewallNetworkRuleCollectionInterface
}

// AzurermFirewallNetworkRuleCollectionInterface has methods to work with AzurermFirewallNetworkRuleCollection resources.
type AzurermFirewallNetworkRuleCollectionInterface interface {
	Create(*v1alpha1.AzurermFirewallNetworkRuleCollection) (*v1alpha1.AzurermFirewallNetworkRuleCollection, error)
	Update(*v1alpha1.AzurermFirewallNetworkRuleCollection) (*v1alpha1.AzurermFirewallNetworkRuleCollection, error)
	UpdateStatus(*v1alpha1.AzurermFirewallNetworkRuleCollection) (*v1alpha1.AzurermFirewallNetworkRuleCollection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermFirewallNetworkRuleCollection, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermFirewallNetworkRuleCollectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermFirewallNetworkRuleCollection, err error)
	AzurermFirewallNetworkRuleCollectionExpansion
}

// azurermFirewallNetworkRuleCollections implements AzurermFirewallNetworkRuleCollectionInterface
type azurermFirewallNetworkRuleCollections struct {
	client rest.Interface
}

// newAzurermFirewallNetworkRuleCollections returns a AzurermFirewallNetworkRuleCollections
func newAzurermFirewallNetworkRuleCollections(c *AzurermV1alpha1Client) *azurermFirewallNetworkRuleCollections {
	return &azurermFirewallNetworkRuleCollections{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermFirewallNetworkRuleCollection, and returns the corresponding azurermFirewallNetworkRuleCollection object, and an error if there is any.
func (c *azurermFirewallNetworkRuleCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermFirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallNetworkRuleCollection{}
	err = c.client.Get().
		Resource("azurermfirewallnetworkrulecollections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermFirewallNetworkRuleCollections that match those selectors.
func (c *azurermFirewallNetworkRuleCollections) List(opts v1.ListOptions) (result *v1alpha1.AzurermFirewallNetworkRuleCollectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermFirewallNetworkRuleCollectionList{}
	err = c.client.Get().
		Resource("azurermfirewallnetworkrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermFirewallNetworkRuleCollections.
func (c *azurermFirewallNetworkRuleCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermfirewallnetworkrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermFirewallNetworkRuleCollection and creates it.  Returns the server's representation of the azurermFirewallNetworkRuleCollection, and an error, if there is any.
func (c *azurermFirewallNetworkRuleCollections) Create(azurermFirewallNetworkRuleCollection *v1alpha1.AzurermFirewallNetworkRuleCollection) (result *v1alpha1.AzurermFirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallNetworkRuleCollection{}
	err = c.client.Post().
		Resource("azurermfirewallnetworkrulecollections").
		Body(azurermFirewallNetworkRuleCollection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermFirewallNetworkRuleCollection and updates it. Returns the server's representation of the azurermFirewallNetworkRuleCollection, and an error, if there is any.
func (c *azurermFirewallNetworkRuleCollections) Update(azurermFirewallNetworkRuleCollection *v1alpha1.AzurermFirewallNetworkRuleCollection) (result *v1alpha1.AzurermFirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallNetworkRuleCollection{}
	err = c.client.Put().
		Resource("azurermfirewallnetworkrulecollections").
		Name(azurermFirewallNetworkRuleCollection.Name).
		Body(azurermFirewallNetworkRuleCollection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermFirewallNetworkRuleCollections) UpdateStatus(azurermFirewallNetworkRuleCollection *v1alpha1.AzurermFirewallNetworkRuleCollection) (result *v1alpha1.AzurermFirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallNetworkRuleCollection{}
	err = c.client.Put().
		Resource("azurermfirewallnetworkrulecollections").
		Name(azurermFirewallNetworkRuleCollection.Name).
		SubResource("status").
		Body(azurermFirewallNetworkRuleCollection).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermFirewallNetworkRuleCollection and deletes it. Returns an error if one occurs.
func (c *azurermFirewallNetworkRuleCollections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermfirewallnetworkrulecollections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermFirewallNetworkRuleCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermfirewallnetworkrulecollections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermFirewallNetworkRuleCollection.
func (c *azurermFirewallNetworkRuleCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermFirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.AzurermFirewallNetworkRuleCollection{}
	err = c.client.Patch(pt).
		Resource("azurermfirewallnetworkrulecollections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
