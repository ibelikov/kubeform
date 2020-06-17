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

// FakeKeyVaultKeys implements KeyVaultKeyInterface
type FakeKeyVaultKeys struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var keyvaultkeysResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "keyvaultkeys"}

var keyvaultkeysKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "KeyVaultKey"}

// Get takes name of the keyVaultKey, and returns the corresponding keyVaultKey object, and an error if there is any.
func (c *FakeKeyVaultKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.KeyVaultKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keyvaultkeysResource, c.ns, name), &v1alpha1.KeyVaultKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultKey), err
}

// List takes label and field selectors, and returns the list of KeyVaultKeys that match those selectors.
func (c *FakeKeyVaultKeys) List(opts v1.ListOptions) (result *v1alpha1.KeyVaultKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keyvaultkeysResource, keyvaultkeysKind, c.ns, opts), &v1alpha1.KeyVaultKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeyVaultKeyList{ListMeta: obj.(*v1alpha1.KeyVaultKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeyVaultKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keyVaultKeys.
func (c *FakeKeyVaultKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keyvaultkeysResource, c.ns, opts))

}

// Create takes the representation of a keyVaultKey and creates it.  Returns the server's representation of the keyVaultKey, and an error, if there is any.
func (c *FakeKeyVaultKeys) Create(keyVaultKey *v1alpha1.KeyVaultKey) (result *v1alpha1.KeyVaultKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keyvaultkeysResource, c.ns, keyVaultKey), &v1alpha1.KeyVaultKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultKey), err
}

// Update takes the representation of a keyVaultKey and updates it. Returns the server's representation of the keyVaultKey, and an error, if there is any.
func (c *FakeKeyVaultKeys) Update(keyVaultKey *v1alpha1.KeyVaultKey) (result *v1alpha1.KeyVaultKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keyvaultkeysResource, c.ns, keyVaultKey), &v1alpha1.KeyVaultKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeyVaultKeys) UpdateStatus(keyVaultKey *v1alpha1.KeyVaultKey) (*v1alpha1.KeyVaultKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keyvaultkeysResource, "status", c.ns, keyVaultKey), &v1alpha1.KeyVaultKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultKey), err
}

// Delete takes name of the keyVaultKey and deletes it. Returns an error if one occurs.
func (c *FakeKeyVaultKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keyvaultkeysResource, c.ns, name), &v1alpha1.KeyVaultKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeyVaultKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keyvaultkeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeyVaultKeyList{})
	return err
}

// Patch applies the patch and returns the patched keyVaultKey.
func (c *FakeKeyVaultKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KeyVaultKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keyvaultkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.KeyVaultKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultKey), err
}
