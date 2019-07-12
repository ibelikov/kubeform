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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleComputeForwardingRules implements GoogleComputeForwardingRuleInterface
type FakeGoogleComputeForwardingRules struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputeforwardingrulesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputeforwardingrules"}

var googlecomputeforwardingrulesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeForwardingRule"}

// Get takes name of the googleComputeForwardingRule, and returns the corresponding googleComputeForwardingRule object, and an error if there is any.
func (c *FakeGoogleComputeForwardingRules) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputeforwardingrulesResource, name), &v1alpha1.GoogleComputeForwardingRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeForwardingRule), err
}

// List takes label and field selectors, and returns the list of GoogleComputeForwardingRules that match those selectors.
func (c *FakeGoogleComputeForwardingRules) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeForwardingRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputeforwardingrulesResource, googlecomputeforwardingrulesKind, opts), &v1alpha1.GoogleComputeForwardingRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeForwardingRuleList{ListMeta: obj.(*v1alpha1.GoogleComputeForwardingRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeForwardingRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeForwardingRules.
func (c *FakeGoogleComputeForwardingRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputeforwardingrulesResource, opts))
}

// Create takes the representation of a googleComputeForwardingRule and creates it.  Returns the server's representation of the googleComputeForwardingRule, and an error, if there is any.
func (c *FakeGoogleComputeForwardingRules) Create(googleComputeForwardingRule *v1alpha1.GoogleComputeForwardingRule) (result *v1alpha1.GoogleComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputeforwardingrulesResource, googleComputeForwardingRule), &v1alpha1.GoogleComputeForwardingRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeForwardingRule), err
}

// Update takes the representation of a googleComputeForwardingRule and updates it. Returns the server's representation of the googleComputeForwardingRule, and an error, if there is any.
func (c *FakeGoogleComputeForwardingRules) Update(googleComputeForwardingRule *v1alpha1.GoogleComputeForwardingRule) (result *v1alpha1.GoogleComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputeforwardingrulesResource, googleComputeForwardingRule), &v1alpha1.GoogleComputeForwardingRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeForwardingRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeForwardingRules) UpdateStatus(googleComputeForwardingRule *v1alpha1.GoogleComputeForwardingRule) (*v1alpha1.GoogleComputeForwardingRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputeforwardingrulesResource, "status", googleComputeForwardingRule), &v1alpha1.GoogleComputeForwardingRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeForwardingRule), err
}

// Delete takes name of the googleComputeForwardingRule and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeForwardingRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputeforwardingrulesResource, name), &v1alpha1.GoogleComputeForwardingRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeForwardingRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputeforwardingrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeForwardingRuleList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeForwardingRule.
func (c *FakeGoogleComputeForwardingRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputeforwardingrulesResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeForwardingRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeForwardingRule), err
}
