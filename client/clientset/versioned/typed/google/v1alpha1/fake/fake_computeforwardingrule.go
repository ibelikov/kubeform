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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeComputeForwardingRules implements ComputeForwardingRuleInterface
type FakeComputeForwardingRules struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeforwardingrulesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeforwardingrules"}

var computeforwardingrulesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeForwardingRule"}

// Get takes name of the computeForwardingRule, and returns the corresponding computeForwardingRule object, and an error if there is any.
func (c *FakeComputeForwardingRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeforwardingrulesResource, c.ns, name), &v1alpha1.ComputeForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeForwardingRule), err
}

// List takes label and field selectors, and returns the list of ComputeForwardingRules that match those selectors.
func (c *FakeComputeForwardingRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeForwardingRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeforwardingrulesResource, computeforwardingrulesKind, c.ns, opts), &v1alpha1.ComputeForwardingRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeForwardingRuleList{ListMeta: obj.(*v1alpha1.ComputeForwardingRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeForwardingRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeForwardingRules.
func (c *FakeComputeForwardingRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeforwardingrulesResource, c.ns, opts))

}

// Create takes the representation of a computeForwardingRule and creates it.  Returns the server's representation of the computeForwardingRule, and an error, if there is any.
func (c *FakeComputeForwardingRules) Create(ctx context.Context, computeForwardingRule *v1alpha1.ComputeForwardingRule, opts v1.CreateOptions) (result *v1alpha1.ComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeforwardingrulesResource, c.ns, computeForwardingRule), &v1alpha1.ComputeForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeForwardingRule), err
}

// Update takes the representation of a computeForwardingRule and updates it. Returns the server's representation of the computeForwardingRule, and an error, if there is any.
func (c *FakeComputeForwardingRules) Update(ctx context.Context, computeForwardingRule *v1alpha1.ComputeForwardingRule, opts v1.UpdateOptions) (result *v1alpha1.ComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeforwardingrulesResource, c.ns, computeForwardingRule), &v1alpha1.ComputeForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeForwardingRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeForwardingRules) UpdateStatus(ctx context.Context, computeForwardingRule *v1alpha1.ComputeForwardingRule, opts v1.UpdateOptions) (*v1alpha1.ComputeForwardingRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeforwardingrulesResource, "status", c.ns, computeForwardingRule), &v1alpha1.ComputeForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeForwardingRule), err
}

// Delete takes name of the computeForwardingRule and deletes it. Returns an error if one occurs.
func (c *FakeComputeForwardingRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeforwardingrulesResource, c.ns, name), &v1alpha1.ComputeForwardingRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeForwardingRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeforwardingrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeForwardingRuleList{})
	return err
}

// Patch applies the patch and returns the patched computeForwardingRule.
func (c *FakeComputeForwardingRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeForwardingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeforwardingrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeForwardingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeForwardingRule), err
}
