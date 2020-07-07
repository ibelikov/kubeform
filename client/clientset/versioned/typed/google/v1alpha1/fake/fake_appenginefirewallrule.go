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

// FakeAppEngineFirewallRules implements AppEngineFirewallRuleInterface
type FakeAppEngineFirewallRules struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var appenginefirewallrulesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "appenginefirewallrules"}

var appenginefirewallrulesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "AppEngineFirewallRule"}

// Get takes name of the appEngineFirewallRule, and returns the corresponding appEngineFirewallRule object, and an error if there is any.
func (c *FakeAppEngineFirewallRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppEngineFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appenginefirewallrulesResource, c.ns, name), &v1alpha1.AppEngineFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFirewallRule), err
}

// List takes label and field selectors, and returns the list of AppEngineFirewallRules that match those selectors.
func (c *FakeAppEngineFirewallRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppEngineFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appenginefirewallrulesResource, appenginefirewallrulesKind, c.ns, opts), &v1alpha1.AppEngineFirewallRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppEngineFirewallRuleList{ListMeta: obj.(*v1alpha1.AppEngineFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppEngineFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appEngineFirewallRules.
func (c *FakeAppEngineFirewallRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appenginefirewallrulesResource, c.ns, opts))

}

// Create takes the representation of a appEngineFirewallRule and creates it.  Returns the server's representation of the appEngineFirewallRule, and an error, if there is any.
func (c *FakeAppEngineFirewallRules) Create(ctx context.Context, appEngineFirewallRule *v1alpha1.AppEngineFirewallRule, opts v1.CreateOptions) (result *v1alpha1.AppEngineFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appenginefirewallrulesResource, c.ns, appEngineFirewallRule), &v1alpha1.AppEngineFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFirewallRule), err
}

// Update takes the representation of a appEngineFirewallRule and updates it. Returns the server's representation of the appEngineFirewallRule, and an error, if there is any.
func (c *FakeAppEngineFirewallRules) Update(ctx context.Context, appEngineFirewallRule *v1alpha1.AppEngineFirewallRule, opts v1.UpdateOptions) (result *v1alpha1.AppEngineFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appenginefirewallrulesResource, c.ns, appEngineFirewallRule), &v1alpha1.AppEngineFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppEngineFirewallRules) UpdateStatus(ctx context.Context, appEngineFirewallRule *v1alpha1.AppEngineFirewallRule, opts v1.UpdateOptions) (*v1alpha1.AppEngineFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appenginefirewallrulesResource, "status", c.ns, appEngineFirewallRule), &v1alpha1.AppEngineFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFirewallRule), err
}

// Delete takes name of the appEngineFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeAppEngineFirewallRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appenginefirewallrulesResource, c.ns, name), &v1alpha1.AppEngineFirewallRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppEngineFirewallRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appenginefirewallrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppEngineFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched appEngineFirewallRule.
func (c *FakeAppEngineFirewallRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppEngineFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appenginefirewallrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppEngineFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFirewallRule), err
}
