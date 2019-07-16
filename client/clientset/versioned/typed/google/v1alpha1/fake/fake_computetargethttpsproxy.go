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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeComputeTargetHttpsProxies implements ComputeTargetHttpsProxyInterface
type FakeComputeTargetHttpsProxies struct {
	Fake *FakeGoogleV1alpha1
}

var computetargethttpsproxiesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computetargethttpsproxies"}

var computetargethttpsproxiesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeTargetHttpsProxy"}

// Get takes name of the computeTargetHttpsProxy, and returns the corresponding computeTargetHttpsProxy object, and an error if there is any.
func (c *FakeComputeTargetHttpsProxies) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeTargetHttpsProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(computetargethttpsproxiesResource, name), &v1alpha1.ComputeTargetHttpsProxy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetHttpsProxy), err
}

// List takes label and field selectors, and returns the list of ComputeTargetHttpsProxies that match those selectors.
func (c *FakeComputeTargetHttpsProxies) List(opts v1.ListOptions) (result *v1alpha1.ComputeTargetHttpsProxyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(computetargethttpsproxiesResource, computetargethttpsproxiesKind, opts), &v1alpha1.ComputeTargetHttpsProxyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeTargetHttpsProxyList{ListMeta: obj.(*v1alpha1.ComputeTargetHttpsProxyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeTargetHttpsProxyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeTargetHttpsProxies.
func (c *FakeComputeTargetHttpsProxies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(computetargethttpsproxiesResource, opts))
}

// Create takes the representation of a computeTargetHttpsProxy and creates it.  Returns the server's representation of the computeTargetHttpsProxy, and an error, if there is any.
func (c *FakeComputeTargetHttpsProxies) Create(computeTargetHttpsProxy *v1alpha1.ComputeTargetHttpsProxy) (result *v1alpha1.ComputeTargetHttpsProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(computetargethttpsproxiesResource, computeTargetHttpsProxy), &v1alpha1.ComputeTargetHttpsProxy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetHttpsProxy), err
}

// Update takes the representation of a computeTargetHttpsProxy and updates it. Returns the server's representation of the computeTargetHttpsProxy, and an error, if there is any.
func (c *FakeComputeTargetHttpsProxies) Update(computeTargetHttpsProxy *v1alpha1.ComputeTargetHttpsProxy) (result *v1alpha1.ComputeTargetHttpsProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(computetargethttpsproxiesResource, computeTargetHttpsProxy), &v1alpha1.ComputeTargetHttpsProxy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetHttpsProxy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeTargetHttpsProxies) UpdateStatus(computeTargetHttpsProxy *v1alpha1.ComputeTargetHttpsProxy) (*v1alpha1.ComputeTargetHttpsProxy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(computetargethttpsproxiesResource, "status", computeTargetHttpsProxy), &v1alpha1.ComputeTargetHttpsProxy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetHttpsProxy), err
}

// Delete takes name of the computeTargetHttpsProxy and deletes it. Returns an error if one occurs.
func (c *FakeComputeTargetHttpsProxies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(computetargethttpsproxiesResource, name), &v1alpha1.ComputeTargetHttpsProxy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeTargetHttpsProxies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(computetargethttpsproxiesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeTargetHttpsProxyList{})
	return err
}

// Patch applies the patch and returns the patched computeTargetHttpsProxy.
func (c *FakeComputeTargetHttpsProxies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeTargetHttpsProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(computetargethttpsproxiesResource, name, pt, data, subresources...), &v1alpha1.ComputeTargetHttpsProxy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetHttpsProxy), err
}