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

// FakeProxyProtocolPolicies implements ProxyProtocolPolicyInterface
type FakeProxyProtocolPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var proxyprotocolpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "proxyprotocolpolicies"}

var proxyprotocolpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ProxyProtocolPolicy"}

// Get takes name of the proxyProtocolPolicy, and returns the corresponding proxyProtocolPolicy object, and an error if there is any.
func (c *FakeProxyProtocolPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ProxyProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(proxyprotocolpoliciesResource, c.ns, name), &v1alpha1.ProxyProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyProtocolPolicy), err
}

// List takes label and field selectors, and returns the list of ProxyProtocolPolicies that match those selectors.
func (c *FakeProxyProtocolPolicies) List(opts v1.ListOptions) (result *v1alpha1.ProxyProtocolPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(proxyprotocolpoliciesResource, proxyprotocolpoliciesKind, c.ns, opts), &v1alpha1.ProxyProtocolPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProxyProtocolPolicyList{ListMeta: obj.(*v1alpha1.ProxyProtocolPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProxyProtocolPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested proxyProtocolPolicies.
func (c *FakeProxyProtocolPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(proxyprotocolpoliciesResource, c.ns, opts))

}

// Create takes the representation of a proxyProtocolPolicy and creates it.  Returns the server's representation of the proxyProtocolPolicy, and an error, if there is any.
func (c *FakeProxyProtocolPolicies) Create(proxyProtocolPolicy *v1alpha1.ProxyProtocolPolicy) (result *v1alpha1.ProxyProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(proxyprotocolpoliciesResource, c.ns, proxyProtocolPolicy), &v1alpha1.ProxyProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyProtocolPolicy), err
}

// Update takes the representation of a proxyProtocolPolicy and updates it. Returns the server's representation of the proxyProtocolPolicy, and an error, if there is any.
func (c *FakeProxyProtocolPolicies) Update(proxyProtocolPolicy *v1alpha1.ProxyProtocolPolicy) (result *v1alpha1.ProxyProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(proxyprotocolpoliciesResource, c.ns, proxyProtocolPolicy), &v1alpha1.ProxyProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyProtocolPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProxyProtocolPolicies) UpdateStatus(proxyProtocolPolicy *v1alpha1.ProxyProtocolPolicy) (*v1alpha1.ProxyProtocolPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(proxyprotocolpoliciesResource, "status", c.ns, proxyProtocolPolicy), &v1alpha1.ProxyProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyProtocolPolicy), err
}

// Delete takes name of the proxyProtocolPolicy and deletes it. Returns an error if one occurs.
func (c *FakeProxyProtocolPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(proxyprotocolpoliciesResource, c.ns, name), &v1alpha1.ProxyProtocolPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProxyProtocolPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(proxyprotocolpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProxyProtocolPolicyList{})
	return err
}

// Patch applies the patch and returns the patched proxyProtocolPolicy.
func (c *FakeProxyProtocolPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProxyProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(proxyprotocolpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProxyProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProxyProtocolPolicy), err
}
