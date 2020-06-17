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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeBillingAccountIamMembers implements BillingAccountIamMemberInterface
type FakeBillingAccountIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var billingaccountiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "billingaccountiammembers"}

var billingaccountiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BillingAccountIamMember"}

// Get takes name of the billingAccountIamMember, and returns the corresponding billingAccountIamMember object, and an error if there is any.
func (c *FakeBillingAccountIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.BillingAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(billingaccountiammembersResource, c.ns, name), &v1alpha1.BillingAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamMember), err
}

// List takes label and field selectors, and returns the list of BillingAccountIamMembers that match those selectors.
func (c *FakeBillingAccountIamMembers) List(opts v1.ListOptions) (result *v1alpha1.BillingAccountIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(billingaccountiammembersResource, billingaccountiammembersKind, c.ns, opts), &v1alpha1.BillingAccountIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BillingAccountIamMemberList{ListMeta: obj.(*v1alpha1.BillingAccountIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.BillingAccountIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested billingAccountIamMembers.
func (c *FakeBillingAccountIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(billingaccountiammembersResource, c.ns, opts))

}

// Create takes the representation of a billingAccountIamMember and creates it.  Returns the server's representation of the billingAccountIamMember, and an error, if there is any.
func (c *FakeBillingAccountIamMembers) Create(billingAccountIamMember *v1alpha1.BillingAccountIamMember) (result *v1alpha1.BillingAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(billingaccountiammembersResource, c.ns, billingAccountIamMember), &v1alpha1.BillingAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamMember), err
}

// Update takes the representation of a billingAccountIamMember and updates it. Returns the server's representation of the billingAccountIamMember, and an error, if there is any.
func (c *FakeBillingAccountIamMembers) Update(billingAccountIamMember *v1alpha1.BillingAccountIamMember) (result *v1alpha1.BillingAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(billingaccountiammembersResource, c.ns, billingAccountIamMember), &v1alpha1.BillingAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBillingAccountIamMembers) UpdateStatus(billingAccountIamMember *v1alpha1.BillingAccountIamMember) (*v1alpha1.BillingAccountIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(billingaccountiammembersResource, "status", c.ns, billingAccountIamMember), &v1alpha1.BillingAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamMember), err
}

// Delete takes name of the billingAccountIamMember and deletes it. Returns an error if one occurs.
func (c *FakeBillingAccountIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(billingaccountiammembersResource, c.ns, name), &v1alpha1.BillingAccountIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBillingAccountIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(billingaccountiammembersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BillingAccountIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched billingAccountIamMember.
func (c *FakeBillingAccountIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BillingAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(billingaccountiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.BillingAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamMember), err
}
