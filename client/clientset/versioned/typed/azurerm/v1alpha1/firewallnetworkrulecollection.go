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

// FirewallNetworkRuleCollectionsGetter has a method to return a FirewallNetworkRuleCollectionInterface.
// A group's client should implement this interface.
type FirewallNetworkRuleCollectionsGetter interface {
	FirewallNetworkRuleCollections() FirewallNetworkRuleCollectionInterface
}

// FirewallNetworkRuleCollectionInterface has methods to work with FirewallNetworkRuleCollection resources.
type FirewallNetworkRuleCollectionInterface interface {
	Create(*v1alpha1.FirewallNetworkRuleCollection) (*v1alpha1.FirewallNetworkRuleCollection, error)
	Update(*v1alpha1.FirewallNetworkRuleCollection) (*v1alpha1.FirewallNetworkRuleCollection, error)
	UpdateStatus(*v1alpha1.FirewallNetworkRuleCollection) (*v1alpha1.FirewallNetworkRuleCollection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FirewallNetworkRuleCollection, error)
	List(opts v1.ListOptions) (*v1alpha1.FirewallNetworkRuleCollectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FirewallNetworkRuleCollection, err error)
	FirewallNetworkRuleCollectionExpansion
}

// firewallNetworkRuleCollections implements FirewallNetworkRuleCollectionInterface
type firewallNetworkRuleCollections struct {
	client rest.Interface
}

// newFirewallNetworkRuleCollections returns a FirewallNetworkRuleCollections
func newFirewallNetworkRuleCollections(c *AzurermV1alpha1Client) *firewallNetworkRuleCollections {
	return &firewallNetworkRuleCollections{
		client: c.RESTClient(),
	}
}

// Get takes name of the firewallNetworkRuleCollection, and returns the corresponding firewallNetworkRuleCollection object, and an error if there is any.
func (c *firewallNetworkRuleCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.FirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.FirewallNetworkRuleCollection{}
	err = c.client.Get().
		Resource("firewallnetworkrulecollections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FirewallNetworkRuleCollections that match those selectors.
func (c *firewallNetworkRuleCollections) List(opts v1.ListOptions) (result *v1alpha1.FirewallNetworkRuleCollectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FirewallNetworkRuleCollectionList{}
	err = c.client.Get().
		Resource("firewallnetworkrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested firewallNetworkRuleCollections.
func (c *firewallNetworkRuleCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("firewallnetworkrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a firewallNetworkRuleCollection and creates it.  Returns the server's representation of the firewallNetworkRuleCollection, and an error, if there is any.
func (c *firewallNetworkRuleCollections) Create(firewallNetworkRuleCollection *v1alpha1.FirewallNetworkRuleCollection) (result *v1alpha1.FirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.FirewallNetworkRuleCollection{}
	err = c.client.Post().
		Resource("firewallnetworkrulecollections").
		Body(firewallNetworkRuleCollection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a firewallNetworkRuleCollection and updates it. Returns the server's representation of the firewallNetworkRuleCollection, and an error, if there is any.
func (c *firewallNetworkRuleCollections) Update(firewallNetworkRuleCollection *v1alpha1.FirewallNetworkRuleCollection) (result *v1alpha1.FirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.FirewallNetworkRuleCollection{}
	err = c.client.Put().
		Resource("firewallnetworkrulecollections").
		Name(firewallNetworkRuleCollection.Name).
		Body(firewallNetworkRuleCollection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *firewallNetworkRuleCollections) UpdateStatus(firewallNetworkRuleCollection *v1alpha1.FirewallNetworkRuleCollection) (result *v1alpha1.FirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.FirewallNetworkRuleCollection{}
	err = c.client.Put().
		Resource("firewallnetworkrulecollections").
		Name(firewallNetworkRuleCollection.Name).
		SubResource("status").
		Body(firewallNetworkRuleCollection).
		Do().
		Into(result)
	return
}

// Delete takes name of the firewallNetworkRuleCollection and deletes it. Returns an error if one occurs.
func (c *firewallNetworkRuleCollections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("firewallnetworkrulecollections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *firewallNetworkRuleCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("firewallnetworkrulecollections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched firewallNetworkRuleCollection.
func (c *firewallNetworkRuleCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FirewallNetworkRuleCollection, err error) {
	result = &v1alpha1.FirewallNetworkRuleCollection{}
	err = c.client.Patch(pt).
		Resource("firewallnetworkrulecollections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
