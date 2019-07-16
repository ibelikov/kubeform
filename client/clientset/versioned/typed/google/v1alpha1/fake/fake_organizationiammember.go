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

// FakeOrganizationIamMembers implements OrganizationIamMemberInterface
type FakeOrganizationIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var organizationiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "organizationiammembers"}

var organizationiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "OrganizationIamMember"}

// Get takes name of the organizationIamMember, and returns the corresponding organizationIamMember object, and an error if there is any.
func (c *FakeOrganizationIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.OrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(organizationiammembersResource, name), &v1alpha1.OrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamMember), err
}

// List takes label and field selectors, and returns the list of OrganizationIamMembers that match those selectors.
func (c *FakeOrganizationIamMembers) List(opts v1.ListOptions) (result *v1alpha1.OrganizationIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(organizationiammembersResource, organizationiammembersKind, opts), &v1alpha1.OrganizationIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationIamMemberList{ListMeta: obj.(*v1alpha1.OrganizationIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationIamMembers.
func (c *FakeOrganizationIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(organizationiammembersResource, opts))
}

// Create takes the representation of a organizationIamMember and creates it.  Returns the server's representation of the organizationIamMember, and an error, if there is any.
func (c *FakeOrganizationIamMembers) Create(organizationIamMember *v1alpha1.OrganizationIamMember) (result *v1alpha1.OrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(organizationiammembersResource, organizationIamMember), &v1alpha1.OrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamMember), err
}

// Update takes the representation of a organizationIamMember and updates it. Returns the server's representation of the organizationIamMember, and an error, if there is any.
func (c *FakeOrganizationIamMembers) Update(organizationIamMember *v1alpha1.OrganizationIamMember) (result *v1alpha1.OrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(organizationiammembersResource, organizationIamMember), &v1alpha1.OrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationIamMembers) UpdateStatus(organizationIamMember *v1alpha1.OrganizationIamMember) (*v1alpha1.OrganizationIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(organizationiammembersResource, "status", organizationIamMember), &v1alpha1.OrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamMember), err
}

// Delete takes name of the organizationIamMember and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(organizationiammembersResource, name), &v1alpha1.OrganizationIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(organizationiammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched organizationIamMember.
func (c *FakeOrganizationIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(organizationiammembersResource, name, pt, data, subresources...), &v1alpha1.OrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamMember), err
}
