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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeIotTopicRules implements IotTopicRuleInterface
type FakeIotTopicRules struct {
	Fake *FakeAwsV1alpha1
}

var iottopicrulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iottopicrules"}

var iottopicrulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IotTopicRule"}

// Get takes name of the iotTopicRule, and returns the corresponding iotTopicRule object, and an error if there is any.
func (c *FakeIotTopicRules) Get(name string, options v1.GetOptions) (result *v1alpha1.IotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(iottopicrulesResource, name), &v1alpha1.IotTopicRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotTopicRule), err
}

// List takes label and field selectors, and returns the list of IotTopicRules that match those selectors.
func (c *FakeIotTopicRules) List(opts v1.ListOptions) (result *v1alpha1.IotTopicRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(iottopicrulesResource, iottopicrulesKind, opts), &v1alpha1.IotTopicRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IotTopicRuleList{ListMeta: obj.(*v1alpha1.IotTopicRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.IotTopicRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iotTopicRules.
func (c *FakeIotTopicRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(iottopicrulesResource, opts))
}

// Create takes the representation of a iotTopicRule and creates it.  Returns the server's representation of the iotTopicRule, and an error, if there is any.
func (c *FakeIotTopicRules) Create(iotTopicRule *v1alpha1.IotTopicRule) (result *v1alpha1.IotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(iottopicrulesResource, iotTopicRule), &v1alpha1.IotTopicRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotTopicRule), err
}

// Update takes the representation of a iotTopicRule and updates it. Returns the server's representation of the iotTopicRule, and an error, if there is any.
func (c *FakeIotTopicRules) Update(iotTopicRule *v1alpha1.IotTopicRule) (result *v1alpha1.IotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(iottopicrulesResource, iotTopicRule), &v1alpha1.IotTopicRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotTopicRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIotTopicRules) UpdateStatus(iotTopicRule *v1alpha1.IotTopicRule) (*v1alpha1.IotTopicRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(iottopicrulesResource, "status", iotTopicRule), &v1alpha1.IotTopicRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotTopicRule), err
}

// Delete takes name of the iotTopicRule and deletes it. Returns an error if one occurs.
func (c *FakeIotTopicRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(iottopicrulesResource, name), &v1alpha1.IotTopicRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIotTopicRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(iottopicrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IotTopicRuleList{})
	return err
}

// Patch applies the patch and returns the patched iotTopicRule.
func (c *FakeIotTopicRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotTopicRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(iottopicrulesResource, name, pt, data, subresources...), &v1alpha1.IotTopicRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotTopicRule), err
}
