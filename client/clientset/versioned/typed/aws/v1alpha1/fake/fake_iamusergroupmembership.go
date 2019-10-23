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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIamUserGroupMemberships implements IamUserGroupMembershipInterface
type FakeIamUserGroupMemberships struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamusergroupmembershipsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamusergroupmemberships"}

var iamusergroupmembershipsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamUserGroupMembership"}

// Get takes name of the iamUserGroupMembership, and returns the corresponding iamUserGroupMembership object, and an error if there is any.
func (c *FakeIamUserGroupMemberships) Get(name string, options v1.GetOptions) (result *v1alpha1.IamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamusergroupmembershipsResource, c.ns, name), &v1alpha1.IamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserGroupMembership), err
}

// List takes label and field selectors, and returns the list of IamUserGroupMemberships that match those selectors.
func (c *FakeIamUserGroupMemberships) List(opts v1.ListOptions) (result *v1alpha1.IamUserGroupMembershipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamusergroupmembershipsResource, iamusergroupmembershipsKind, c.ns, opts), &v1alpha1.IamUserGroupMembershipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamUserGroupMembershipList{ListMeta: obj.(*v1alpha1.IamUserGroupMembershipList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamUserGroupMembershipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamUserGroupMemberships.
func (c *FakeIamUserGroupMemberships) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamusergroupmembershipsResource, c.ns, opts))

}

// Create takes the representation of a iamUserGroupMembership and creates it.  Returns the server's representation of the iamUserGroupMembership, and an error, if there is any.
func (c *FakeIamUserGroupMemberships) Create(iamUserGroupMembership *v1alpha1.IamUserGroupMembership) (result *v1alpha1.IamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamusergroupmembershipsResource, c.ns, iamUserGroupMembership), &v1alpha1.IamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserGroupMembership), err
}

// Update takes the representation of a iamUserGroupMembership and updates it. Returns the server's representation of the iamUserGroupMembership, and an error, if there is any.
func (c *FakeIamUserGroupMemberships) Update(iamUserGroupMembership *v1alpha1.IamUserGroupMembership) (result *v1alpha1.IamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamusergroupmembershipsResource, c.ns, iamUserGroupMembership), &v1alpha1.IamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserGroupMembership), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamUserGroupMemberships) UpdateStatus(iamUserGroupMembership *v1alpha1.IamUserGroupMembership) (*v1alpha1.IamUserGroupMembership, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamusergroupmembershipsResource, "status", c.ns, iamUserGroupMembership), &v1alpha1.IamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserGroupMembership), err
}

// Delete takes name of the iamUserGroupMembership and deletes it. Returns an error if one occurs.
func (c *FakeIamUserGroupMemberships) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamusergroupmembershipsResource, c.ns, name), &v1alpha1.IamUserGroupMembership{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamUserGroupMemberships) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamusergroupmembershipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamUserGroupMembershipList{})
	return err
}

// Patch applies the patch and returns the patched iamUserGroupMembership.
func (c *FakeIamUserGroupMemberships) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserGroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamusergroupmembershipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamUserGroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserGroupMembership), err
}
