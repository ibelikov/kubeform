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

// FakeAzurermMariadbFirewallRules implements AzurermMariadbFirewallRuleInterface
type FakeAzurermMariadbFirewallRules struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmariadbfirewallrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmariadbfirewallrules"}

var azurermmariadbfirewallrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMariadbFirewallRule"}

// Get takes name of the azurermMariadbFirewallRule, and returns the corresponding azurermMariadbFirewallRule object, and an error if there is any.
func (c *FakeAzurermMariadbFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmariadbfirewallrulesResource, name), &v1alpha1.AzurermMariadbFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbFirewallRule), err
}

// List takes label and field selectors, and returns the list of AzurermMariadbFirewallRules that match those selectors.
func (c *FakeAzurermMariadbFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermMariadbFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmariadbfirewallrulesResource, azurermmariadbfirewallrulesKind, opts), &v1alpha1.AzurermMariadbFirewallRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMariadbFirewallRuleList{ListMeta: obj.(*v1alpha1.AzurermMariadbFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMariadbFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMariadbFirewallRules.
func (c *FakeAzurermMariadbFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmariadbfirewallrulesResource, opts))
}

// Create takes the representation of a azurermMariadbFirewallRule and creates it.  Returns the server's representation of the azurermMariadbFirewallRule, and an error, if there is any.
func (c *FakeAzurermMariadbFirewallRules) Create(azurermMariadbFirewallRule *v1alpha1.AzurermMariadbFirewallRule) (result *v1alpha1.AzurermMariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmariadbfirewallrulesResource, azurermMariadbFirewallRule), &v1alpha1.AzurermMariadbFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbFirewallRule), err
}

// Update takes the representation of a azurermMariadbFirewallRule and updates it. Returns the server's representation of the azurermMariadbFirewallRule, and an error, if there is any.
func (c *FakeAzurermMariadbFirewallRules) Update(azurermMariadbFirewallRule *v1alpha1.AzurermMariadbFirewallRule) (result *v1alpha1.AzurermMariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmariadbfirewallrulesResource, azurermMariadbFirewallRule), &v1alpha1.AzurermMariadbFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMariadbFirewallRules) UpdateStatus(azurermMariadbFirewallRule *v1alpha1.AzurermMariadbFirewallRule) (*v1alpha1.AzurermMariadbFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmariadbfirewallrulesResource, "status", azurermMariadbFirewallRule), &v1alpha1.AzurermMariadbFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbFirewallRule), err
}

// Delete takes name of the azurermMariadbFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMariadbFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmariadbfirewallrulesResource, name), &v1alpha1.AzurermMariadbFirewallRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMariadbFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmariadbfirewallrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMariadbFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched azurermMariadbFirewallRule.
func (c *FakeAzurermMariadbFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmariadbfirewallrulesResource, name, pt, data, subresources...), &v1alpha1.AzurermMariadbFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMariadbFirewallRule), err
}
