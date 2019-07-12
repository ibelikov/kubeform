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

// FakeAwsIamRolePolicies implements AwsIamRolePolicyInterface
type FakeAwsIamRolePolicies struct {
	Fake *FakeAwsV1alpha1
}

var awsiamrolepoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsiamrolepolicies"}

var awsiamrolepoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsIamRolePolicy"}

// Get takes name of the awsIamRolePolicy, and returns the corresponding awsIamRolePolicy object, and an error if there is any.
func (c *FakeAwsIamRolePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsiamrolepoliciesResource, name), &v1alpha1.AwsIamRolePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamRolePolicy), err
}

// List takes label and field selectors, and returns the list of AwsIamRolePolicies that match those selectors.
func (c *FakeAwsIamRolePolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsIamRolePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsiamrolepoliciesResource, awsiamrolepoliciesKind, opts), &v1alpha1.AwsIamRolePolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamRolePolicyList{ListMeta: obj.(*v1alpha1.AwsIamRolePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsIamRolePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamRolePolicies.
func (c *FakeAwsIamRolePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsiamrolepoliciesResource, opts))
}

// Create takes the representation of a awsIamRolePolicy and creates it.  Returns the server's representation of the awsIamRolePolicy, and an error, if there is any.
func (c *FakeAwsIamRolePolicies) Create(awsIamRolePolicy *v1alpha1.AwsIamRolePolicy) (result *v1alpha1.AwsIamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsiamrolepoliciesResource, awsIamRolePolicy), &v1alpha1.AwsIamRolePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamRolePolicy), err
}

// Update takes the representation of a awsIamRolePolicy and updates it. Returns the server's representation of the awsIamRolePolicy, and an error, if there is any.
func (c *FakeAwsIamRolePolicies) Update(awsIamRolePolicy *v1alpha1.AwsIamRolePolicy) (result *v1alpha1.AwsIamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsiamrolepoliciesResource, awsIamRolePolicy), &v1alpha1.AwsIamRolePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamRolePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsIamRolePolicies) UpdateStatus(awsIamRolePolicy *v1alpha1.AwsIamRolePolicy) (*v1alpha1.AwsIamRolePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsiamrolepoliciesResource, "status", awsIamRolePolicy), &v1alpha1.AwsIamRolePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamRolePolicy), err
}

// Delete takes name of the awsIamRolePolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamRolePolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsiamrolepoliciesResource, name), &v1alpha1.AwsIamRolePolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamRolePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsiamrolepoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamRolePolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsIamRolePolicy.
func (c *FakeAwsIamRolePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamRolePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsiamrolepoliciesResource, name, pt, data, subresources...), &v1alpha1.AwsIamRolePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamRolePolicy), err
}
