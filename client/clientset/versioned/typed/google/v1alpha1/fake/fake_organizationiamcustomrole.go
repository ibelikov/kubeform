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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOrganizationIamCustomRoles implements OrganizationIamCustomRoleInterface
type FakeOrganizationIamCustomRoles struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var organizationiamcustomrolesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "organizationiamcustomroles"}

var organizationiamcustomrolesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "OrganizationIamCustomRole"}

// Get takes name of the organizationIamCustomRole, and returns the corresponding organizationIamCustomRole object, and an error if there is any.
func (c *FakeOrganizationIamCustomRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.OrganizationIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(organizationiamcustomrolesResource, c.ns, name), &v1alpha1.OrganizationIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamCustomRole), err
}

// List takes label and field selectors, and returns the list of OrganizationIamCustomRoles that match those selectors.
func (c *FakeOrganizationIamCustomRoles) List(opts v1.ListOptions) (result *v1alpha1.OrganizationIamCustomRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(organizationiamcustomrolesResource, organizationiamcustomrolesKind, c.ns, opts), &v1alpha1.OrganizationIamCustomRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationIamCustomRoleList{ListMeta: obj.(*v1alpha1.OrganizationIamCustomRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationIamCustomRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationIamCustomRoles.
func (c *FakeOrganizationIamCustomRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(organizationiamcustomrolesResource, c.ns, opts))

}

// Create takes the representation of a organizationIamCustomRole and creates it.  Returns the server's representation of the organizationIamCustomRole, and an error, if there is any.
func (c *FakeOrganizationIamCustomRoles) Create(organizationIamCustomRole *v1alpha1.OrganizationIamCustomRole) (result *v1alpha1.OrganizationIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(organizationiamcustomrolesResource, c.ns, organizationIamCustomRole), &v1alpha1.OrganizationIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamCustomRole), err
}

// Update takes the representation of a organizationIamCustomRole and updates it. Returns the server's representation of the organizationIamCustomRole, and an error, if there is any.
func (c *FakeOrganizationIamCustomRoles) Update(organizationIamCustomRole *v1alpha1.OrganizationIamCustomRole) (result *v1alpha1.OrganizationIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(organizationiamcustomrolesResource, c.ns, organizationIamCustomRole), &v1alpha1.OrganizationIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamCustomRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationIamCustomRoles) UpdateStatus(organizationIamCustomRole *v1alpha1.OrganizationIamCustomRole) (*v1alpha1.OrganizationIamCustomRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(organizationiamcustomrolesResource, "status", c.ns, organizationIamCustomRole), &v1alpha1.OrganizationIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamCustomRole), err
}

// Delete takes name of the organizationIamCustomRole and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationIamCustomRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(organizationiamcustomrolesResource, c.ns, name), &v1alpha1.OrganizationIamCustomRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationIamCustomRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(organizationiamcustomrolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationIamCustomRoleList{})
	return err
}

// Patch applies the patch and returns the patched organizationIamCustomRole.
func (c *FakeOrganizationIamCustomRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(organizationiamcustomrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrganizationIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationIamCustomRole), err
}
