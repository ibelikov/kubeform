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

// FakeIamGroupPolicies implements IamGroupPolicyInterface
type FakeIamGroupPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamgrouppoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamgrouppolicies"}

var iamgrouppoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamGroupPolicy"}

// Get takes name of the iamGroupPolicy, and returns the corresponding iamGroupPolicy object, and an error if there is any.
func (c *FakeIamGroupPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamgrouppoliciesResource, c.ns, name), &v1alpha1.IamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicy), err
}

// List takes label and field selectors, and returns the list of IamGroupPolicies that match those selectors.
func (c *FakeIamGroupPolicies) List(opts v1.ListOptions) (result *v1alpha1.IamGroupPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamgrouppoliciesResource, iamgrouppoliciesKind, c.ns, opts), &v1alpha1.IamGroupPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamGroupPolicyList{ListMeta: obj.(*v1alpha1.IamGroupPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamGroupPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamGroupPolicies.
func (c *FakeIamGroupPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamgrouppoliciesResource, c.ns, opts))

}

// Create takes the representation of a iamGroupPolicy and creates it.  Returns the server's representation of the iamGroupPolicy, and an error, if there is any.
func (c *FakeIamGroupPolicies) Create(iamGroupPolicy *v1alpha1.IamGroupPolicy) (result *v1alpha1.IamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamgrouppoliciesResource, c.ns, iamGroupPolicy), &v1alpha1.IamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicy), err
}

// Update takes the representation of a iamGroupPolicy and updates it. Returns the server's representation of the iamGroupPolicy, and an error, if there is any.
func (c *FakeIamGroupPolicies) Update(iamGroupPolicy *v1alpha1.IamGroupPolicy) (result *v1alpha1.IamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamgrouppoliciesResource, c.ns, iamGroupPolicy), &v1alpha1.IamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamGroupPolicies) UpdateStatus(iamGroupPolicy *v1alpha1.IamGroupPolicy) (*v1alpha1.IamGroupPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamgrouppoliciesResource, "status", c.ns, iamGroupPolicy), &v1alpha1.IamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicy), err
}

// Delete takes name of the iamGroupPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIamGroupPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamgrouppoliciesResource, c.ns, name), &v1alpha1.IamGroupPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamGroupPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamgrouppoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamGroupPolicyList{})
	return err
}

// Patch applies the patch and returns the patched iamGroupPolicy.
func (c *FakeIamGroupPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamgrouppoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamGroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicy), err
}
