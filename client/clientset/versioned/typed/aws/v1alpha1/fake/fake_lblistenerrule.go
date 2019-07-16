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

// FakeLbListenerRules implements LbListenerRuleInterface
type FakeLbListenerRules struct {
	Fake *FakeAwsV1alpha1
}

var lblistenerrulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lblistenerrules"}

var lblistenerrulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LbListenerRule"}

// Get takes name of the lbListenerRule, and returns the corresponding lbListenerRule object, and an error if there is any.
func (c *FakeLbListenerRules) Get(name string, options v1.GetOptions) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(lblistenerrulesResource, name), &v1alpha1.LbListenerRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// List takes label and field selectors, and returns the list of LbListenerRules that match those selectors.
func (c *FakeLbListenerRules) List(opts v1.ListOptions) (result *v1alpha1.LbListenerRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(lblistenerrulesResource, lblistenerrulesKind, opts), &v1alpha1.LbListenerRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbListenerRuleList{ListMeta: obj.(*v1alpha1.LbListenerRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbListenerRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbListenerRules.
func (c *FakeLbListenerRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(lblistenerrulesResource, opts))
}

// Create takes the representation of a lbListenerRule and creates it.  Returns the server's representation of the lbListenerRule, and an error, if there is any.
func (c *FakeLbListenerRules) Create(lbListenerRule *v1alpha1.LbListenerRule) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(lblistenerrulesResource, lbListenerRule), &v1alpha1.LbListenerRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// Update takes the representation of a lbListenerRule and updates it. Returns the server's representation of the lbListenerRule, and an error, if there is any.
func (c *FakeLbListenerRules) Update(lbListenerRule *v1alpha1.LbListenerRule) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(lblistenerrulesResource, lbListenerRule), &v1alpha1.LbListenerRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbListenerRules) UpdateStatus(lbListenerRule *v1alpha1.LbListenerRule) (*v1alpha1.LbListenerRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(lblistenerrulesResource, "status", lbListenerRule), &v1alpha1.LbListenerRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// Delete takes name of the lbListenerRule and deletes it. Returns an error if one occurs.
func (c *FakeLbListenerRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(lblistenerrulesResource, name), &v1alpha1.LbListenerRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbListenerRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(lblistenerrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbListenerRuleList{})
	return err
}

// Patch applies the patch and returns the patched lbListenerRule.
func (c *FakeLbListenerRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(lblistenerrulesResource, name, pt, data, subresources...), &v1alpha1.LbListenerRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}
