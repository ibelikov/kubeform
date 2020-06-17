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

// FakeCosmosdbAccounts implements CosmosdbAccountInterface
type FakeCosmosdbAccounts struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var cosmosdbaccountsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "cosmosdbaccounts"}

var cosmosdbaccountsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "CosmosdbAccount"}

// Get takes name of the cosmosdbAccount, and returns the corresponding cosmosdbAccount object, and an error if there is any.
func (c *FakeCosmosdbAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.CosmosdbAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cosmosdbaccountsResource, c.ns, name), &v1alpha1.CosmosdbAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbAccount), err
}

// List takes label and field selectors, and returns the list of CosmosdbAccounts that match those selectors.
func (c *FakeCosmosdbAccounts) List(opts v1.ListOptions) (result *v1alpha1.CosmosdbAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cosmosdbaccountsResource, cosmosdbaccountsKind, c.ns, opts), &v1alpha1.CosmosdbAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CosmosdbAccountList{ListMeta: obj.(*v1alpha1.CosmosdbAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.CosmosdbAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cosmosdbAccounts.
func (c *FakeCosmosdbAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cosmosdbaccountsResource, c.ns, opts))

}

// Create takes the representation of a cosmosdbAccount and creates it.  Returns the server's representation of the cosmosdbAccount, and an error, if there is any.
func (c *FakeCosmosdbAccounts) Create(cosmosdbAccount *v1alpha1.CosmosdbAccount) (result *v1alpha1.CosmosdbAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cosmosdbaccountsResource, c.ns, cosmosdbAccount), &v1alpha1.CosmosdbAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbAccount), err
}

// Update takes the representation of a cosmosdbAccount and updates it. Returns the server's representation of the cosmosdbAccount, and an error, if there is any.
func (c *FakeCosmosdbAccounts) Update(cosmosdbAccount *v1alpha1.CosmosdbAccount) (result *v1alpha1.CosmosdbAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cosmosdbaccountsResource, c.ns, cosmosdbAccount), &v1alpha1.CosmosdbAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCosmosdbAccounts) UpdateStatus(cosmosdbAccount *v1alpha1.CosmosdbAccount) (*v1alpha1.CosmosdbAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cosmosdbaccountsResource, "status", c.ns, cosmosdbAccount), &v1alpha1.CosmosdbAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbAccount), err
}

// Delete takes name of the cosmosdbAccount and deletes it. Returns an error if one occurs.
func (c *FakeCosmosdbAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cosmosdbaccountsResource, c.ns, name), &v1alpha1.CosmosdbAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCosmosdbAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cosmosdbaccountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CosmosdbAccountList{})
	return err
}

// Patch applies the patch and returns the patched cosmosdbAccount.
func (c *FakeCosmosdbAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CosmosdbAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cosmosdbaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CosmosdbAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbAccount), err
}
