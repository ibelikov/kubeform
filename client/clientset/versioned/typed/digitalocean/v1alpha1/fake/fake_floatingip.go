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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeFloatingIps implements FloatingIpInterface
type FakeFloatingIps struct {
	Fake *FakeDigitaloceanV1alpha1
}

var floatingipsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "floatingips"}

var floatingipsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "FloatingIp"}

// Get takes name of the floatingIp, and returns the corresponding floatingIp object, and an error if there is any.
func (c *FakeFloatingIps) Get(name string, options v1.GetOptions) (result *v1alpha1.FloatingIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(floatingipsResource, name), &v1alpha1.FloatingIp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIp), err
}

// List takes label and field selectors, and returns the list of FloatingIps that match those selectors.
func (c *FakeFloatingIps) List(opts v1.ListOptions) (result *v1alpha1.FloatingIpList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(floatingipsResource, floatingipsKind, opts), &v1alpha1.FloatingIpList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FloatingIpList{ListMeta: obj.(*v1alpha1.FloatingIpList).ListMeta}
	for _, item := range obj.(*v1alpha1.FloatingIpList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested floatingIps.
func (c *FakeFloatingIps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(floatingipsResource, opts))
}

// Create takes the representation of a floatingIp and creates it.  Returns the server's representation of the floatingIp, and an error, if there is any.
func (c *FakeFloatingIps) Create(floatingIp *v1alpha1.FloatingIp) (result *v1alpha1.FloatingIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(floatingipsResource, floatingIp), &v1alpha1.FloatingIp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIp), err
}

// Update takes the representation of a floatingIp and updates it. Returns the server's representation of the floatingIp, and an error, if there is any.
func (c *FakeFloatingIps) Update(floatingIp *v1alpha1.FloatingIp) (result *v1alpha1.FloatingIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(floatingipsResource, floatingIp), &v1alpha1.FloatingIp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFloatingIps) UpdateStatus(floatingIp *v1alpha1.FloatingIp) (*v1alpha1.FloatingIp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(floatingipsResource, "status", floatingIp), &v1alpha1.FloatingIp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIp), err
}

// Delete takes name of the floatingIp and deletes it. Returns an error if one occurs.
func (c *FakeFloatingIps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(floatingipsResource, name), &v1alpha1.FloatingIp{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFloatingIps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(floatingipsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FloatingIpList{})
	return err
}

// Patch applies the patch and returns the patched floatingIp.
func (c *FakeFloatingIps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(floatingipsResource, name, pt, data, subresources...), &v1alpha1.FloatingIp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIp), err
}