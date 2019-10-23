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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLogicAppActionCustoms implements LogicAppActionCustomInterface
type FakeLogicAppActionCustoms struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var logicappactioncustomsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "logicappactioncustoms"}

var logicappactioncustomsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogicAppActionCustom"}

// Get takes name of the logicAppActionCustom, and returns the corresponding logicAppActionCustom object, and an error if there is any.
func (c *FakeLogicAppActionCustoms) Get(name string, options v1.GetOptions) (result *v1alpha1.LogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logicappactioncustomsResource, c.ns, name), &v1alpha1.LogicAppActionCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppActionCustom), err
}

// List takes label and field selectors, and returns the list of LogicAppActionCustoms that match those selectors.
func (c *FakeLogicAppActionCustoms) List(opts v1.ListOptions) (result *v1alpha1.LogicAppActionCustomList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logicappactioncustomsResource, logicappactioncustomsKind, c.ns, opts), &v1alpha1.LogicAppActionCustomList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogicAppActionCustomList{ListMeta: obj.(*v1alpha1.LogicAppActionCustomList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogicAppActionCustomList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logicAppActionCustoms.
func (c *FakeLogicAppActionCustoms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logicappactioncustomsResource, c.ns, opts))

}

// Create takes the representation of a logicAppActionCustom and creates it.  Returns the server's representation of the logicAppActionCustom, and an error, if there is any.
func (c *FakeLogicAppActionCustoms) Create(logicAppActionCustom *v1alpha1.LogicAppActionCustom) (result *v1alpha1.LogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logicappactioncustomsResource, c.ns, logicAppActionCustom), &v1alpha1.LogicAppActionCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppActionCustom), err
}

// Update takes the representation of a logicAppActionCustom and updates it. Returns the server's representation of the logicAppActionCustom, and an error, if there is any.
func (c *FakeLogicAppActionCustoms) Update(logicAppActionCustom *v1alpha1.LogicAppActionCustom) (result *v1alpha1.LogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logicappactioncustomsResource, c.ns, logicAppActionCustom), &v1alpha1.LogicAppActionCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppActionCustom), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogicAppActionCustoms) UpdateStatus(logicAppActionCustom *v1alpha1.LogicAppActionCustom) (*v1alpha1.LogicAppActionCustom, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logicappactioncustomsResource, "status", c.ns, logicAppActionCustom), &v1alpha1.LogicAppActionCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppActionCustom), err
}

// Delete takes name of the logicAppActionCustom and deletes it. Returns an error if one occurs.
func (c *FakeLogicAppActionCustoms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logicappactioncustomsResource, c.ns, name), &v1alpha1.LogicAppActionCustom{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogicAppActionCustoms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logicappactioncustomsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogicAppActionCustomList{})
	return err
}

// Patch applies the patch and returns the patched logicAppActionCustom.
func (c *FakeLogicAppActionCustoms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogicAppActionCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logicappactioncustomsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogicAppActionCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppActionCustom), err
}
