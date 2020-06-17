/*
Copyright The Kubeform Authors.

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

// FakeSqlVirtualNetworkRules implements SqlVirtualNetworkRuleInterface
type FakeSqlVirtualNetworkRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var sqlvirtualnetworkrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "sqlvirtualnetworkrules"}

var sqlvirtualnetworkrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SqlVirtualNetworkRule"}

// Get takes name of the sqlVirtualNetworkRule, and returns the corresponding sqlVirtualNetworkRule object, and an error if there is any.
func (c *FakeSqlVirtualNetworkRules) Get(name string, options v1.GetOptions) (result *v1alpha1.SqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqlvirtualnetworkrulesResource, c.ns, name), &v1alpha1.SqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlVirtualNetworkRule), err
}

// List takes label and field selectors, and returns the list of SqlVirtualNetworkRules that match those selectors.
func (c *FakeSqlVirtualNetworkRules) List(opts v1.ListOptions) (result *v1alpha1.SqlVirtualNetworkRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqlvirtualnetworkrulesResource, sqlvirtualnetworkrulesKind, c.ns, opts), &v1alpha1.SqlVirtualNetworkRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqlVirtualNetworkRuleList{ListMeta: obj.(*v1alpha1.SqlVirtualNetworkRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqlVirtualNetworkRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlVirtualNetworkRules.
func (c *FakeSqlVirtualNetworkRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqlvirtualnetworkrulesResource, c.ns, opts))

}

// Create takes the representation of a sqlVirtualNetworkRule and creates it.  Returns the server's representation of the sqlVirtualNetworkRule, and an error, if there is any.
func (c *FakeSqlVirtualNetworkRules) Create(sqlVirtualNetworkRule *v1alpha1.SqlVirtualNetworkRule) (result *v1alpha1.SqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqlvirtualnetworkrulesResource, c.ns, sqlVirtualNetworkRule), &v1alpha1.SqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlVirtualNetworkRule), err
}

// Update takes the representation of a sqlVirtualNetworkRule and updates it. Returns the server's representation of the sqlVirtualNetworkRule, and an error, if there is any.
func (c *FakeSqlVirtualNetworkRules) Update(sqlVirtualNetworkRule *v1alpha1.SqlVirtualNetworkRule) (result *v1alpha1.SqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqlvirtualnetworkrulesResource, c.ns, sqlVirtualNetworkRule), &v1alpha1.SqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlVirtualNetworkRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqlVirtualNetworkRules) UpdateStatus(sqlVirtualNetworkRule *v1alpha1.SqlVirtualNetworkRule) (*v1alpha1.SqlVirtualNetworkRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqlvirtualnetworkrulesResource, "status", c.ns, sqlVirtualNetworkRule), &v1alpha1.SqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlVirtualNetworkRule), err
}

// Delete takes name of the sqlVirtualNetworkRule and deletes it. Returns an error if one occurs.
func (c *FakeSqlVirtualNetworkRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqlvirtualnetworkrulesResource, c.ns, name), &v1alpha1.SqlVirtualNetworkRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlVirtualNetworkRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqlvirtualnetworkrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqlVirtualNetworkRuleList{})
	return err
}

// Patch applies the patch and returns the patched sqlVirtualNetworkRule.
func (c *FakeSqlVirtualNetworkRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqlvirtualnetworkrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlVirtualNetworkRule), err
}
