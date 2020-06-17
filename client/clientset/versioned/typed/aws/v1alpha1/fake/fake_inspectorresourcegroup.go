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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeInspectorResourceGroups implements InspectorResourceGroupInterface
type FakeInspectorResourceGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var inspectorresourcegroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "inspectorresourcegroups"}

var inspectorresourcegroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "InspectorResourceGroup"}

// Get takes name of the inspectorResourceGroup, and returns the corresponding inspectorResourceGroup object, and an error if there is any.
func (c *FakeInspectorResourceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.InspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(inspectorresourcegroupsResource, c.ns, name), &v1alpha1.InspectorResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorResourceGroup), err
}

// List takes label and field selectors, and returns the list of InspectorResourceGroups that match those selectors.
func (c *FakeInspectorResourceGroups) List(opts v1.ListOptions) (result *v1alpha1.InspectorResourceGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(inspectorresourcegroupsResource, inspectorresourcegroupsKind, c.ns, opts), &v1alpha1.InspectorResourceGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InspectorResourceGroupList{ListMeta: obj.(*v1alpha1.InspectorResourceGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.InspectorResourceGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inspectorResourceGroups.
func (c *FakeInspectorResourceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(inspectorresourcegroupsResource, c.ns, opts))

}

// Create takes the representation of a inspectorResourceGroup and creates it.  Returns the server's representation of the inspectorResourceGroup, and an error, if there is any.
func (c *FakeInspectorResourceGroups) Create(inspectorResourceGroup *v1alpha1.InspectorResourceGroup) (result *v1alpha1.InspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(inspectorresourcegroupsResource, c.ns, inspectorResourceGroup), &v1alpha1.InspectorResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorResourceGroup), err
}

// Update takes the representation of a inspectorResourceGroup and updates it. Returns the server's representation of the inspectorResourceGroup, and an error, if there is any.
func (c *FakeInspectorResourceGroups) Update(inspectorResourceGroup *v1alpha1.InspectorResourceGroup) (result *v1alpha1.InspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(inspectorresourcegroupsResource, c.ns, inspectorResourceGroup), &v1alpha1.InspectorResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorResourceGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInspectorResourceGroups) UpdateStatus(inspectorResourceGroup *v1alpha1.InspectorResourceGroup) (*v1alpha1.InspectorResourceGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(inspectorresourcegroupsResource, "status", c.ns, inspectorResourceGroup), &v1alpha1.InspectorResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorResourceGroup), err
}

// Delete takes name of the inspectorResourceGroup and deletes it. Returns an error if one occurs.
func (c *FakeInspectorResourceGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(inspectorresourcegroupsResource, c.ns, name), &v1alpha1.InspectorResourceGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInspectorResourceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(inspectorresourcegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InspectorResourceGroupList{})
	return err
}

// Patch applies the patch and returns the patched inspectorResourceGroup.
func (c *FakeInspectorResourceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InspectorResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(inspectorresourcegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.InspectorResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorResourceGroup), err
}
