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

// FakeAzurermLogicAppActionCustoms implements AzurermLogicAppActionCustomInterface
type FakeAzurermLogicAppActionCustoms struct {
	Fake *FakeAzurermV1alpha1
}

var azurermlogicappactioncustomsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermlogicappactioncustoms"}

var azurermlogicappactioncustomsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermLogicAppActionCustom"}

// Get takes name of the azurermLogicAppActionCustom, and returns the corresponding azurermLogicAppActionCustom object, and an error if there is any.
func (c *FakeAzurermLogicAppActionCustoms) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermlogicappactioncustomsResource, name), &v1alpha1.AzurermLogicAppActionCustom{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppActionCustom), err
}

// List takes label and field selectors, and returns the list of AzurermLogicAppActionCustoms that match those selectors.
func (c *FakeAzurermLogicAppActionCustoms) List(opts v1.ListOptions) (result *v1alpha1.AzurermLogicAppActionCustomList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermlogicappactioncustomsResource, azurermlogicappactioncustomsKind, opts), &v1alpha1.AzurermLogicAppActionCustomList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermLogicAppActionCustomList{ListMeta: obj.(*v1alpha1.AzurermLogicAppActionCustomList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermLogicAppActionCustomList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermLogicAppActionCustoms.
func (c *FakeAzurermLogicAppActionCustoms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermlogicappactioncustomsResource, opts))
}

// Create takes the representation of a azurermLogicAppActionCustom and creates it.  Returns the server's representation of the azurermLogicAppActionCustom, and an error, if there is any.
func (c *FakeAzurermLogicAppActionCustoms) Create(azurermLogicAppActionCustom *v1alpha1.AzurermLogicAppActionCustom) (result *v1alpha1.AzurermLogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermlogicappactioncustomsResource, azurermLogicAppActionCustom), &v1alpha1.AzurermLogicAppActionCustom{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppActionCustom), err
}

// Update takes the representation of a azurermLogicAppActionCustom and updates it. Returns the server's representation of the azurermLogicAppActionCustom, and an error, if there is any.
func (c *FakeAzurermLogicAppActionCustoms) Update(azurermLogicAppActionCustom *v1alpha1.AzurermLogicAppActionCustom) (result *v1alpha1.AzurermLogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermlogicappactioncustomsResource, azurermLogicAppActionCustom), &v1alpha1.AzurermLogicAppActionCustom{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppActionCustom), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermLogicAppActionCustoms) UpdateStatus(azurermLogicAppActionCustom *v1alpha1.AzurermLogicAppActionCustom) (*v1alpha1.AzurermLogicAppActionCustom, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermlogicappactioncustomsResource, "status", azurermLogicAppActionCustom), &v1alpha1.AzurermLogicAppActionCustom{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppActionCustom), err
}

// Delete takes name of the azurermLogicAppActionCustom and deletes it. Returns an error if one occurs.
func (c *FakeAzurermLogicAppActionCustoms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermlogicappactioncustomsResource, name), &v1alpha1.AzurermLogicAppActionCustom{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermLogicAppActionCustoms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermlogicappactioncustomsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermLogicAppActionCustomList{})
	return err
}

// Patch applies the patch and returns the patched azurermLogicAppActionCustom.
func (c *FakeAzurermLogicAppActionCustoms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermlogicappactioncustomsResource, name, pt, data, subresources...), &v1alpha1.AzurermLogicAppActionCustom{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppActionCustom), err
}
