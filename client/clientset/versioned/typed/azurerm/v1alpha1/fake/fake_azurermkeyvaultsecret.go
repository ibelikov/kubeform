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

// FakeAzurermKeyVaultSecrets implements AzurermKeyVaultSecretInterface
type FakeAzurermKeyVaultSecrets struct {
	Fake *FakeAzurermV1alpha1
}

var azurermkeyvaultsecretsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermkeyvaultsecrets"}

var azurermkeyvaultsecretsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermKeyVaultSecret"}

// Get takes name of the azurermKeyVaultSecret, and returns the corresponding azurermKeyVaultSecret object, and an error if there is any.
func (c *FakeAzurermKeyVaultSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermkeyvaultsecretsResource, name), &v1alpha1.AzurermKeyVaultSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKeyVaultSecret), err
}

// List takes label and field selectors, and returns the list of AzurermKeyVaultSecrets that match those selectors.
func (c *FakeAzurermKeyVaultSecrets) List(opts v1.ListOptions) (result *v1alpha1.AzurermKeyVaultSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermkeyvaultsecretsResource, azurermkeyvaultsecretsKind, opts), &v1alpha1.AzurermKeyVaultSecretList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermKeyVaultSecretList{ListMeta: obj.(*v1alpha1.AzurermKeyVaultSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermKeyVaultSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermKeyVaultSecrets.
func (c *FakeAzurermKeyVaultSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermkeyvaultsecretsResource, opts))
}

// Create takes the representation of a azurermKeyVaultSecret and creates it.  Returns the server's representation of the azurermKeyVaultSecret, and an error, if there is any.
func (c *FakeAzurermKeyVaultSecrets) Create(azurermKeyVaultSecret *v1alpha1.AzurermKeyVaultSecret) (result *v1alpha1.AzurermKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermkeyvaultsecretsResource, azurermKeyVaultSecret), &v1alpha1.AzurermKeyVaultSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKeyVaultSecret), err
}

// Update takes the representation of a azurermKeyVaultSecret and updates it. Returns the server's representation of the azurermKeyVaultSecret, and an error, if there is any.
func (c *FakeAzurermKeyVaultSecrets) Update(azurermKeyVaultSecret *v1alpha1.AzurermKeyVaultSecret) (result *v1alpha1.AzurermKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermkeyvaultsecretsResource, azurermKeyVaultSecret), &v1alpha1.AzurermKeyVaultSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKeyVaultSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermKeyVaultSecrets) UpdateStatus(azurermKeyVaultSecret *v1alpha1.AzurermKeyVaultSecret) (*v1alpha1.AzurermKeyVaultSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermkeyvaultsecretsResource, "status", azurermKeyVaultSecret), &v1alpha1.AzurermKeyVaultSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKeyVaultSecret), err
}

// Delete takes name of the azurermKeyVaultSecret and deletes it. Returns an error if one occurs.
func (c *FakeAzurermKeyVaultSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermkeyvaultsecretsResource, name), &v1alpha1.AzurermKeyVaultSecret{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermKeyVaultSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermkeyvaultsecretsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermKeyVaultSecretList{})
	return err
}

// Patch applies the patch and returns the patched azurermKeyVaultSecret.
func (c *FakeAzurermKeyVaultSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermkeyvaultsecretsResource, name, pt, data, subresources...), &v1alpha1.AzurermKeyVaultSecret{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermKeyVaultSecret), err
}
