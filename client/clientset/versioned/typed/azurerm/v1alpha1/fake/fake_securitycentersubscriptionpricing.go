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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeSecurityCenterSubscriptionPricings implements SecurityCenterSubscriptionPricingInterface
type FakeSecurityCenterSubscriptionPricings struct {
	Fake *FakeAzurermV1alpha1
}

var securitycentersubscriptionpricingsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "securitycentersubscriptionpricings"}

var securitycentersubscriptionpricingsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SecurityCenterSubscriptionPricing"}

// Get takes name of the securityCenterSubscriptionPricing, and returns the corresponding securityCenterSubscriptionPricing object, and an error if there is any.
func (c *FakeSecurityCenterSubscriptionPricings) Get(name string, options v1.GetOptions) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(securitycentersubscriptionpricingsResource, name), &v1alpha1.SecurityCenterSubscriptionPricing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterSubscriptionPricing), err
}

// List takes label and field selectors, and returns the list of SecurityCenterSubscriptionPricings that match those selectors.
func (c *FakeSecurityCenterSubscriptionPricings) List(opts v1.ListOptions) (result *v1alpha1.SecurityCenterSubscriptionPricingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(securitycentersubscriptionpricingsResource, securitycentersubscriptionpricingsKind, opts), &v1alpha1.SecurityCenterSubscriptionPricingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecurityCenterSubscriptionPricingList{ListMeta: obj.(*v1alpha1.SecurityCenterSubscriptionPricingList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecurityCenterSubscriptionPricingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested securityCenterSubscriptionPricings.
func (c *FakeSecurityCenterSubscriptionPricings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(securitycentersubscriptionpricingsResource, opts))
}

// Create takes the representation of a securityCenterSubscriptionPricing and creates it.  Returns the server's representation of the securityCenterSubscriptionPricing, and an error, if there is any.
func (c *FakeSecurityCenterSubscriptionPricings) Create(securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(securitycentersubscriptionpricingsResource, securityCenterSubscriptionPricing), &v1alpha1.SecurityCenterSubscriptionPricing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterSubscriptionPricing), err
}

// Update takes the representation of a securityCenterSubscriptionPricing and updates it. Returns the server's representation of the securityCenterSubscriptionPricing, and an error, if there is any.
func (c *FakeSecurityCenterSubscriptionPricings) Update(securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(securitycentersubscriptionpricingsResource, securityCenterSubscriptionPricing), &v1alpha1.SecurityCenterSubscriptionPricing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterSubscriptionPricing), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecurityCenterSubscriptionPricings) UpdateStatus(securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing) (*v1alpha1.SecurityCenterSubscriptionPricing, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(securitycentersubscriptionpricingsResource, "status", securityCenterSubscriptionPricing), &v1alpha1.SecurityCenterSubscriptionPricing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterSubscriptionPricing), err
}

// Delete takes name of the securityCenterSubscriptionPricing and deletes it. Returns an error if one occurs.
func (c *FakeSecurityCenterSubscriptionPricings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(securitycentersubscriptionpricingsResource, name), &v1alpha1.SecurityCenterSubscriptionPricing{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecurityCenterSubscriptionPricings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(securitycentersubscriptionpricingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecurityCenterSubscriptionPricingList{})
	return err
}

// Patch applies the patch and returns the patched securityCenterSubscriptionPricing.
func (c *FakeSecurityCenterSubscriptionPricings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(securitycentersubscriptionpricingsResource, name, pt, data, subresources...), &v1alpha1.SecurityCenterSubscriptionPricing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityCenterSubscriptionPricing), err
}
