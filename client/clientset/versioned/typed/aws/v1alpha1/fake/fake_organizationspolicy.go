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

// FakeOrganizationsPolicies implements OrganizationsPolicyInterface
type FakeOrganizationsPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var organizationspoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "organizationspolicies"}

var organizationspoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OrganizationsPolicy"}

// Get takes name of the organizationsPolicy, and returns the corresponding organizationsPolicy object, and an error if there is any.
func (c *FakeOrganizationsPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.OrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(organizationspoliciesResource, c.ns, name), &v1alpha1.OrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicy), err
}

// List takes label and field selectors, and returns the list of OrganizationsPolicies that match those selectors.
func (c *FakeOrganizationsPolicies) List(opts v1.ListOptions) (result *v1alpha1.OrganizationsPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(organizationspoliciesResource, organizationspoliciesKind, c.ns, opts), &v1alpha1.OrganizationsPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationsPolicyList{ListMeta: obj.(*v1alpha1.OrganizationsPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationsPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationsPolicies.
func (c *FakeOrganizationsPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(organizationspoliciesResource, c.ns, opts))

}

// Create takes the representation of a organizationsPolicy and creates it.  Returns the server's representation of the organizationsPolicy, and an error, if there is any.
func (c *FakeOrganizationsPolicies) Create(organizationsPolicy *v1alpha1.OrganizationsPolicy) (result *v1alpha1.OrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(organizationspoliciesResource, c.ns, organizationsPolicy), &v1alpha1.OrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicy), err
}

// Update takes the representation of a organizationsPolicy and updates it. Returns the server's representation of the organizationsPolicy, and an error, if there is any.
func (c *FakeOrganizationsPolicies) Update(organizationsPolicy *v1alpha1.OrganizationsPolicy) (result *v1alpha1.OrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(organizationspoliciesResource, c.ns, organizationsPolicy), &v1alpha1.OrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationsPolicies) UpdateStatus(organizationsPolicy *v1alpha1.OrganizationsPolicy) (*v1alpha1.OrganizationsPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(organizationspoliciesResource, "status", c.ns, organizationsPolicy), &v1alpha1.OrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicy), err
}

// Delete takes name of the organizationsPolicy and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationsPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(organizationspoliciesResource, c.ns, name), &v1alpha1.OrganizationsPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationsPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(organizationspoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationsPolicyList{})
	return err
}

// Patch applies the patch and returns the patched organizationsPolicy.
func (c *FakeOrganizationsPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(organizationspoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicy), err
}
