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

// FakeComputeSubnetworkIamPolicies implements ComputeSubnetworkIamPolicyInterface
type FakeComputeSubnetworkIamPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computesubnetworkiampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computesubnetworkiampolicies"}

var computesubnetworkiampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeSubnetworkIamPolicy"}

// Get takes name of the computeSubnetworkIamPolicy, and returns the corresponding computeSubnetworkIamPolicy object, and an error if there is any.
func (c *FakeComputeSubnetworkIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computesubnetworkiampoliciesResource, c.ns, name), &v1alpha1.ComputeSubnetworkIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetworkIamPolicy), err
}

// List takes label and field selectors, and returns the list of ComputeSubnetworkIamPolicies that match those selectors.
func (c *FakeComputeSubnetworkIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.ComputeSubnetworkIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computesubnetworkiampoliciesResource, computesubnetworkiampoliciesKind, c.ns, opts), &v1alpha1.ComputeSubnetworkIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeSubnetworkIamPolicyList{ListMeta: obj.(*v1alpha1.ComputeSubnetworkIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeSubnetworkIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeSubnetworkIamPolicies.
func (c *FakeComputeSubnetworkIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computesubnetworkiampoliciesResource, c.ns, opts))

}

// Create takes the representation of a computeSubnetworkIamPolicy and creates it.  Returns the server's representation of the computeSubnetworkIamPolicy, and an error, if there is any.
func (c *FakeComputeSubnetworkIamPolicies) Create(computeSubnetworkIamPolicy *v1alpha1.ComputeSubnetworkIamPolicy) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computesubnetworkiampoliciesResource, c.ns, computeSubnetworkIamPolicy), &v1alpha1.ComputeSubnetworkIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetworkIamPolicy), err
}

// Update takes the representation of a computeSubnetworkIamPolicy and updates it. Returns the server's representation of the computeSubnetworkIamPolicy, and an error, if there is any.
func (c *FakeComputeSubnetworkIamPolicies) Update(computeSubnetworkIamPolicy *v1alpha1.ComputeSubnetworkIamPolicy) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computesubnetworkiampoliciesResource, c.ns, computeSubnetworkIamPolicy), &v1alpha1.ComputeSubnetworkIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetworkIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeSubnetworkIamPolicies) UpdateStatus(computeSubnetworkIamPolicy *v1alpha1.ComputeSubnetworkIamPolicy) (*v1alpha1.ComputeSubnetworkIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computesubnetworkiampoliciesResource, "status", c.ns, computeSubnetworkIamPolicy), &v1alpha1.ComputeSubnetworkIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetworkIamPolicy), err
}

// Delete takes name of the computeSubnetworkIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeComputeSubnetworkIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computesubnetworkiampoliciesResource, c.ns, name), &v1alpha1.ComputeSubnetworkIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeSubnetworkIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computesubnetworkiampoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeSubnetworkIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched computeSubnetworkIamPolicy.
func (c *FakeComputeSubnetworkIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computesubnetworkiampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeSubnetworkIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetworkIamPolicy), err
}
