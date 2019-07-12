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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermDataLakeStoreFirewallRules implements AzurermDataLakeStoreFirewallRuleInterface
type FakeAzurermDataLakeStoreFirewallRules struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdatalakestorefirewallrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdatalakestorefirewallrules"}

var azurermdatalakestorefirewallrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDataLakeStoreFirewallRule"}

// Get takes name of the azurermDataLakeStoreFirewallRule, and returns the corresponding azurermDataLakeStoreFirewallRule object, and an error if there is any.
func (c *FakeAzurermDataLakeStoreFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdatalakestorefirewallrulesResource, name), &v1alpha1.AzurermDataLakeStoreFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeStoreFirewallRule), err
}

// List takes label and field selectors, and returns the list of AzurermDataLakeStoreFirewallRules that match those selectors.
func (c *FakeAzurermDataLakeStoreFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermDataLakeStoreFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdatalakestorefirewallrulesResource, azurermdatalakestorefirewallrulesKind, opts), &v1alpha1.AzurermDataLakeStoreFirewallRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDataLakeStoreFirewallRuleList{ListMeta: obj.(*v1alpha1.AzurermDataLakeStoreFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDataLakeStoreFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDataLakeStoreFirewallRules.
func (c *FakeAzurermDataLakeStoreFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdatalakestorefirewallrulesResource, opts))
}

// Create takes the representation of a azurermDataLakeStoreFirewallRule and creates it.  Returns the server's representation of the azurermDataLakeStoreFirewallRule, and an error, if there is any.
func (c *FakeAzurermDataLakeStoreFirewallRules) Create(azurermDataLakeStoreFirewallRule *v1alpha1.AzurermDataLakeStoreFirewallRule) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdatalakestorefirewallrulesResource, azurermDataLakeStoreFirewallRule), &v1alpha1.AzurermDataLakeStoreFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeStoreFirewallRule), err
}

// Update takes the representation of a azurermDataLakeStoreFirewallRule and updates it. Returns the server's representation of the azurermDataLakeStoreFirewallRule, and an error, if there is any.
func (c *FakeAzurermDataLakeStoreFirewallRules) Update(azurermDataLakeStoreFirewallRule *v1alpha1.AzurermDataLakeStoreFirewallRule) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdatalakestorefirewallrulesResource, azurermDataLakeStoreFirewallRule), &v1alpha1.AzurermDataLakeStoreFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeStoreFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDataLakeStoreFirewallRules) UpdateStatus(azurermDataLakeStoreFirewallRule *v1alpha1.AzurermDataLakeStoreFirewallRule) (*v1alpha1.AzurermDataLakeStoreFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdatalakestorefirewallrulesResource, "status", azurermDataLakeStoreFirewallRule), &v1alpha1.AzurermDataLakeStoreFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeStoreFirewallRule), err
}

// Delete takes name of the azurermDataLakeStoreFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDataLakeStoreFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdatalakestorefirewallrulesResource, name), &v1alpha1.AzurermDataLakeStoreFirewallRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDataLakeStoreFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdatalakestorefirewallrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDataLakeStoreFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched azurermDataLakeStoreFirewallRule.
func (c *FakeAzurermDataLakeStoreFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDataLakeStoreFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdatalakestorefirewallrulesResource, name, pt, data, subresources...), &v1alpha1.AzurermDataLakeStoreFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDataLakeStoreFirewallRule), err
}
