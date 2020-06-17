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

// FakeSesReceiptRules implements SesReceiptRuleInterface
type FakeSesReceiptRules struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var sesreceiptrulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "sesreceiptrules"}

var sesreceiptrulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SesReceiptRule"}

// Get takes name of the sesReceiptRule, and returns the corresponding sesReceiptRule object, and an error if there is any.
func (c *FakeSesReceiptRules) Get(name string, options v1.GetOptions) (result *v1alpha1.SesReceiptRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sesreceiptrulesResource, c.ns, name), &v1alpha1.SesReceiptRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesReceiptRule), err
}

// List takes label and field selectors, and returns the list of SesReceiptRules that match those selectors.
func (c *FakeSesReceiptRules) List(opts v1.ListOptions) (result *v1alpha1.SesReceiptRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sesreceiptrulesResource, sesreceiptrulesKind, c.ns, opts), &v1alpha1.SesReceiptRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SesReceiptRuleList{ListMeta: obj.(*v1alpha1.SesReceiptRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.SesReceiptRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sesReceiptRules.
func (c *FakeSesReceiptRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sesreceiptrulesResource, c.ns, opts))

}

// Create takes the representation of a sesReceiptRule and creates it.  Returns the server's representation of the sesReceiptRule, and an error, if there is any.
func (c *FakeSesReceiptRules) Create(sesReceiptRule *v1alpha1.SesReceiptRule) (result *v1alpha1.SesReceiptRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sesreceiptrulesResource, c.ns, sesReceiptRule), &v1alpha1.SesReceiptRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesReceiptRule), err
}

// Update takes the representation of a sesReceiptRule and updates it. Returns the server's representation of the sesReceiptRule, and an error, if there is any.
func (c *FakeSesReceiptRules) Update(sesReceiptRule *v1alpha1.SesReceiptRule) (result *v1alpha1.SesReceiptRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sesreceiptrulesResource, c.ns, sesReceiptRule), &v1alpha1.SesReceiptRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesReceiptRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSesReceiptRules) UpdateStatus(sesReceiptRule *v1alpha1.SesReceiptRule) (*v1alpha1.SesReceiptRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sesreceiptrulesResource, "status", c.ns, sesReceiptRule), &v1alpha1.SesReceiptRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesReceiptRule), err
}

// Delete takes name of the sesReceiptRule and deletes it. Returns an error if one occurs.
func (c *FakeSesReceiptRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sesreceiptrulesResource, c.ns, name), &v1alpha1.SesReceiptRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSesReceiptRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sesreceiptrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SesReceiptRuleList{})
	return err
}

// Patch applies the patch and returns the patched sesReceiptRule.
func (c *FakeSesReceiptRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesReceiptRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sesreceiptrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SesReceiptRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesReceiptRule), err
}
