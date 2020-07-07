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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAutomationCredentials implements AutomationCredentialInterface
type FakeAutomationCredentials struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var automationcredentialsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "automationcredentials"}

var automationcredentialsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AutomationCredential"}

// Get takes name of the automationCredential, and returns the corresponding automationCredential object, and an error if there is any.
func (c *FakeAutomationCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(automationcredentialsResource, c.ns, name), &v1alpha1.AutomationCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationCredential), err
}

// List takes label and field selectors, and returns the list of AutomationCredentials that match those selectors.
func (c *FakeAutomationCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AutomationCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(automationcredentialsResource, automationcredentialsKind, c.ns, opts), &v1alpha1.AutomationCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutomationCredentialList{ListMeta: obj.(*v1alpha1.AutomationCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutomationCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested automationCredentials.
func (c *FakeAutomationCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(automationcredentialsResource, c.ns, opts))

}

// Create takes the representation of a automationCredential and creates it.  Returns the server's representation of the automationCredential, and an error, if there is any.
func (c *FakeAutomationCredentials) Create(ctx context.Context, automationCredential *v1alpha1.AutomationCredential, opts v1.CreateOptions) (result *v1alpha1.AutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(automationcredentialsResource, c.ns, automationCredential), &v1alpha1.AutomationCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationCredential), err
}

// Update takes the representation of a automationCredential and updates it. Returns the server's representation of the automationCredential, and an error, if there is any.
func (c *FakeAutomationCredentials) Update(ctx context.Context, automationCredential *v1alpha1.AutomationCredential, opts v1.UpdateOptions) (result *v1alpha1.AutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(automationcredentialsResource, c.ns, automationCredential), &v1alpha1.AutomationCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutomationCredentials) UpdateStatus(ctx context.Context, automationCredential *v1alpha1.AutomationCredential, opts v1.UpdateOptions) (*v1alpha1.AutomationCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(automationcredentialsResource, "status", c.ns, automationCredential), &v1alpha1.AutomationCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationCredential), err
}

// Delete takes name of the automationCredential and deletes it. Returns an error if one occurs.
func (c *FakeAutomationCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(automationcredentialsResource, c.ns, name), &v1alpha1.AutomationCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutomationCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(automationcredentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutomationCredentialList{})
	return err
}

// Patch applies the patch and returns the patched automationCredential.
func (c *FakeAutomationCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(automationcredentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutomationCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationCredential), err
}
