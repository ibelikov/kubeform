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

// FakeVirtualMachineScaleSets implements VirtualMachineScaleSetInterface
type FakeVirtualMachineScaleSets struct {
	Fake *FakeAzurermV1alpha1
}

var virtualmachinescalesetsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "virtualmachinescalesets"}

var virtualmachinescalesetsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "VirtualMachineScaleSet"}

// Get takes name of the virtualMachineScaleSet, and returns the corresponding virtualMachineScaleSet object, and an error if there is any.
func (c *FakeVirtualMachineScaleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(virtualmachinescalesetsResource, name), &v1alpha1.VirtualMachineScaleSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSet), err
}

// List takes label and field selectors, and returns the list of VirtualMachineScaleSets that match those selectors.
func (c *FakeVirtualMachineScaleSets) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineScaleSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(virtualmachinescalesetsResource, virtualmachinescalesetsKind, opts), &v1alpha1.VirtualMachineScaleSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineScaleSetList{ListMeta: obj.(*v1alpha1.VirtualMachineScaleSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineScaleSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineScaleSets.
func (c *FakeVirtualMachineScaleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(virtualmachinescalesetsResource, opts))
}

// Create takes the representation of a virtualMachineScaleSet and creates it.  Returns the server's representation of the virtualMachineScaleSet, and an error, if there is any.
func (c *FakeVirtualMachineScaleSets) Create(virtualMachineScaleSet *v1alpha1.VirtualMachineScaleSet) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(virtualmachinescalesetsResource, virtualMachineScaleSet), &v1alpha1.VirtualMachineScaleSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSet), err
}

// Update takes the representation of a virtualMachineScaleSet and updates it. Returns the server's representation of the virtualMachineScaleSet, and an error, if there is any.
func (c *FakeVirtualMachineScaleSets) Update(virtualMachineScaleSet *v1alpha1.VirtualMachineScaleSet) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(virtualmachinescalesetsResource, virtualMachineScaleSet), &v1alpha1.VirtualMachineScaleSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineScaleSets) UpdateStatus(virtualMachineScaleSet *v1alpha1.VirtualMachineScaleSet) (*v1alpha1.VirtualMachineScaleSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(virtualmachinescalesetsResource, "status", virtualMachineScaleSet), &v1alpha1.VirtualMachineScaleSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSet), err
}

// Delete takes name of the virtualMachineScaleSet and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineScaleSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(virtualmachinescalesetsResource, name), &v1alpha1.VirtualMachineScaleSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineScaleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(virtualmachinescalesetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineScaleSetList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineScaleSet.
func (c *FakeVirtualMachineScaleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(virtualmachinescalesetsResource, name, pt, data, subresources...), &v1alpha1.VirtualMachineScaleSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineScaleSet), err
}
