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

// FakeWafRuleGroups implements WafRuleGroupInterface
type FakeWafRuleGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var wafrulegroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "wafrulegroups"}

var wafrulegroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WafRuleGroup"}

// Get takes name of the wafRuleGroup, and returns the corresponding wafRuleGroup object, and an error if there is any.
func (c *FakeWafRuleGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.WafRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(wafrulegroupsResource, c.ns, name), &v1alpha1.WafRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRuleGroup), err
}

// List takes label and field selectors, and returns the list of WafRuleGroups that match those selectors.
func (c *FakeWafRuleGroups) List(opts v1.ListOptions) (result *v1alpha1.WafRuleGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(wafrulegroupsResource, wafrulegroupsKind, c.ns, opts), &v1alpha1.WafRuleGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafRuleGroupList{ListMeta: obj.(*v1alpha1.WafRuleGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafRuleGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafRuleGroups.
func (c *FakeWafRuleGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(wafrulegroupsResource, c.ns, opts))

}

// Create takes the representation of a wafRuleGroup and creates it.  Returns the server's representation of the wafRuleGroup, and an error, if there is any.
func (c *FakeWafRuleGroups) Create(wafRuleGroup *v1alpha1.WafRuleGroup) (result *v1alpha1.WafRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(wafrulegroupsResource, c.ns, wafRuleGroup), &v1alpha1.WafRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRuleGroup), err
}

// Update takes the representation of a wafRuleGroup and updates it. Returns the server's representation of the wafRuleGroup, and an error, if there is any.
func (c *FakeWafRuleGroups) Update(wafRuleGroup *v1alpha1.WafRuleGroup) (result *v1alpha1.WafRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(wafrulegroupsResource, c.ns, wafRuleGroup), &v1alpha1.WafRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRuleGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafRuleGroups) UpdateStatus(wafRuleGroup *v1alpha1.WafRuleGroup) (*v1alpha1.WafRuleGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(wafrulegroupsResource, "status", c.ns, wafRuleGroup), &v1alpha1.WafRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRuleGroup), err
}

// Delete takes name of the wafRuleGroup and deletes it. Returns an error if one occurs.
func (c *FakeWafRuleGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(wafrulegroupsResource, c.ns, name), &v1alpha1.WafRuleGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafRuleGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(wafrulegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafRuleGroupList{})
	return err
}

// Patch applies the patch and returns the patched wafRuleGroup.
func (c *FakeWafRuleGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(wafrulegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WafRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRuleGroup), err
}
