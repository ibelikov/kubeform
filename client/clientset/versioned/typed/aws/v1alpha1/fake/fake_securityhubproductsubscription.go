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

// FakeSecurityhubProductSubscriptions implements SecurityhubProductSubscriptionInterface
type FakeSecurityhubProductSubscriptions struct {
	Fake *FakeAwsV1alpha1
}

var securityhubproductsubscriptionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "securityhubproductsubscriptions"}

var securityhubproductsubscriptionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SecurityhubProductSubscription"}

// Get takes name of the securityhubProductSubscription, and returns the corresponding securityhubProductSubscription object, and an error if there is any.
func (c *FakeSecurityhubProductSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(securityhubproductsubscriptionsResource, name), &v1alpha1.SecurityhubProductSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubProductSubscription), err
}

// List takes label and field selectors, and returns the list of SecurityhubProductSubscriptions that match those selectors.
func (c *FakeSecurityhubProductSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.SecurityhubProductSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(securityhubproductsubscriptionsResource, securityhubproductsubscriptionsKind, opts), &v1alpha1.SecurityhubProductSubscriptionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecurityhubProductSubscriptionList{ListMeta: obj.(*v1alpha1.SecurityhubProductSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecurityhubProductSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested securityhubProductSubscriptions.
func (c *FakeSecurityhubProductSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(securityhubproductsubscriptionsResource, opts))
}

// Create takes the representation of a securityhubProductSubscription and creates it.  Returns the server's representation of the securityhubProductSubscription, and an error, if there is any.
func (c *FakeSecurityhubProductSubscriptions) Create(securityhubProductSubscription *v1alpha1.SecurityhubProductSubscription) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(securityhubproductsubscriptionsResource, securityhubProductSubscription), &v1alpha1.SecurityhubProductSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubProductSubscription), err
}

// Update takes the representation of a securityhubProductSubscription and updates it. Returns the server's representation of the securityhubProductSubscription, and an error, if there is any.
func (c *FakeSecurityhubProductSubscriptions) Update(securityhubProductSubscription *v1alpha1.SecurityhubProductSubscription) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(securityhubproductsubscriptionsResource, securityhubProductSubscription), &v1alpha1.SecurityhubProductSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubProductSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecurityhubProductSubscriptions) UpdateStatus(securityhubProductSubscription *v1alpha1.SecurityhubProductSubscription) (*v1alpha1.SecurityhubProductSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(securityhubproductsubscriptionsResource, "status", securityhubProductSubscription), &v1alpha1.SecurityhubProductSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubProductSubscription), err
}

// Delete takes name of the securityhubProductSubscription and deletes it. Returns an error if one occurs.
func (c *FakeSecurityhubProductSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(securityhubproductsubscriptionsResource, name), &v1alpha1.SecurityhubProductSubscription{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecurityhubProductSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(securityhubproductsubscriptionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecurityhubProductSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched securityhubProductSubscription.
func (c *FakeSecurityhubProductSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(securityhubproductsubscriptionsResource, name, pt, data, subresources...), &v1alpha1.SecurityhubProductSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecurityhubProductSubscription), err
}
