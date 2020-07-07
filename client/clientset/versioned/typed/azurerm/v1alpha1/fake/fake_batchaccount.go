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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeBatchAccounts implements BatchAccountInterface
type FakeBatchAccounts struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var batchaccountsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "batchaccounts"}

var batchaccountsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BatchAccount"}

// Get takes name of the batchAccount, and returns the corresponding batchAccount object, and an error if there is any.
func (c *FakeBatchAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(batchaccountsResource, c.ns, name), &v1alpha1.BatchAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchAccount), err
}

// List takes label and field selectors, and returns the list of BatchAccounts that match those selectors.
func (c *FakeBatchAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BatchAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(batchaccountsResource, batchaccountsKind, c.ns, opts), &v1alpha1.BatchAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BatchAccountList{ListMeta: obj.(*v1alpha1.BatchAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.BatchAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested batchAccounts.
func (c *FakeBatchAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(batchaccountsResource, c.ns, opts))

}

// Create takes the representation of a batchAccount and creates it.  Returns the server's representation of the batchAccount, and an error, if there is any.
func (c *FakeBatchAccounts) Create(ctx context.Context, batchAccount *v1alpha1.BatchAccount, opts v1.CreateOptions) (result *v1alpha1.BatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(batchaccountsResource, c.ns, batchAccount), &v1alpha1.BatchAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchAccount), err
}

// Update takes the representation of a batchAccount and updates it. Returns the server's representation of the batchAccount, and an error, if there is any.
func (c *FakeBatchAccounts) Update(ctx context.Context, batchAccount *v1alpha1.BatchAccount, opts v1.UpdateOptions) (result *v1alpha1.BatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(batchaccountsResource, c.ns, batchAccount), &v1alpha1.BatchAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBatchAccounts) UpdateStatus(ctx context.Context, batchAccount *v1alpha1.BatchAccount, opts v1.UpdateOptions) (*v1alpha1.BatchAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(batchaccountsResource, "status", c.ns, batchAccount), &v1alpha1.BatchAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchAccount), err
}

// Delete takes name of the batchAccount and deletes it. Returns an error if one occurs.
func (c *FakeBatchAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(batchaccountsResource, c.ns, name), &v1alpha1.BatchAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBatchAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(batchaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BatchAccountList{})
	return err
}

// Patch applies the patch and returns the patched batchAccount.
func (c *FakeBatchAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BatchAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(batchaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BatchAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchAccount), err
}
