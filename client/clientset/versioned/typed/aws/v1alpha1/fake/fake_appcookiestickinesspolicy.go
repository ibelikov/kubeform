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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppCookieStickinessPolicies implements AppCookieStickinessPolicyInterface
type FakeAppCookieStickinessPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var appcookiestickinesspoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "appcookiestickinesspolicies"}

var appcookiestickinesspoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AppCookieStickinessPolicy"}

// Get takes name of the appCookieStickinessPolicy, and returns the corresponding appCookieStickinessPolicy object, and an error if there is any.
func (c *FakeAppCookieStickinessPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appcookiestickinesspoliciesResource, c.ns, name), &v1alpha1.AppCookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppCookieStickinessPolicy), err
}

// List takes label and field selectors, and returns the list of AppCookieStickinessPolicies that match those selectors.
func (c *FakeAppCookieStickinessPolicies) List(opts v1.ListOptions) (result *v1alpha1.AppCookieStickinessPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appcookiestickinesspoliciesResource, appcookiestickinesspoliciesKind, c.ns, opts), &v1alpha1.AppCookieStickinessPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppCookieStickinessPolicyList{ListMeta: obj.(*v1alpha1.AppCookieStickinessPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppCookieStickinessPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appCookieStickinessPolicies.
func (c *FakeAppCookieStickinessPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appcookiestickinesspoliciesResource, c.ns, opts))

}

// Create takes the representation of a appCookieStickinessPolicy and creates it.  Returns the server's representation of the appCookieStickinessPolicy, and an error, if there is any.
func (c *FakeAppCookieStickinessPolicies) Create(appCookieStickinessPolicy *v1alpha1.AppCookieStickinessPolicy) (result *v1alpha1.AppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appcookiestickinesspoliciesResource, c.ns, appCookieStickinessPolicy), &v1alpha1.AppCookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppCookieStickinessPolicy), err
}

// Update takes the representation of a appCookieStickinessPolicy and updates it. Returns the server's representation of the appCookieStickinessPolicy, and an error, if there is any.
func (c *FakeAppCookieStickinessPolicies) Update(appCookieStickinessPolicy *v1alpha1.AppCookieStickinessPolicy) (result *v1alpha1.AppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appcookiestickinesspoliciesResource, c.ns, appCookieStickinessPolicy), &v1alpha1.AppCookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppCookieStickinessPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppCookieStickinessPolicies) UpdateStatus(appCookieStickinessPolicy *v1alpha1.AppCookieStickinessPolicy) (*v1alpha1.AppCookieStickinessPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appcookiestickinesspoliciesResource, "status", c.ns, appCookieStickinessPolicy), &v1alpha1.AppCookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppCookieStickinessPolicy), err
}

// Delete takes name of the appCookieStickinessPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAppCookieStickinessPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appcookiestickinesspoliciesResource, c.ns, name), &v1alpha1.AppCookieStickinessPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppCookieStickinessPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appcookiestickinesspoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppCookieStickinessPolicyList{})
	return err
}

// Patch applies the patch and returns the patched appCookieStickinessPolicy.
func (c *FakeAppCookieStickinessPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppCookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appcookiestickinesspoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppCookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppCookieStickinessPolicy), err
}
