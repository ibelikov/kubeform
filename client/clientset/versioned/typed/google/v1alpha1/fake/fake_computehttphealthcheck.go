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

// FakeComputeHttpHealthChecks implements ComputeHttpHealthCheckInterface
type FakeComputeHttpHealthChecks struct {
	Fake *FakeGoogleV1alpha1
}

var computehttphealthchecksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computehttphealthchecks"}

var computehttphealthchecksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeHttpHealthCheck"}

// Get takes name of the computeHttpHealthCheck, and returns the corresponding computeHttpHealthCheck object, and an error if there is any.
func (c *FakeComputeHttpHealthChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeHttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(computehttphealthchecksResource, name), &v1alpha1.ComputeHttpHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHttpHealthCheck), err
}

// List takes label and field selectors, and returns the list of ComputeHttpHealthChecks that match those selectors.
func (c *FakeComputeHttpHealthChecks) List(opts v1.ListOptions) (result *v1alpha1.ComputeHttpHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(computehttphealthchecksResource, computehttphealthchecksKind, opts), &v1alpha1.ComputeHttpHealthCheckList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeHttpHealthCheckList{ListMeta: obj.(*v1alpha1.ComputeHttpHealthCheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeHttpHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeHttpHealthChecks.
func (c *FakeComputeHttpHealthChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(computehttphealthchecksResource, opts))
}

// Create takes the representation of a computeHttpHealthCheck and creates it.  Returns the server's representation of the computeHttpHealthCheck, and an error, if there is any.
func (c *FakeComputeHttpHealthChecks) Create(computeHttpHealthCheck *v1alpha1.ComputeHttpHealthCheck) (result *v1alpha1.ComputeHttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(computehttphealthchecksResource, computeHttpHealthCheck), &v1alpha1.ComputeHttpHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHttpHealthCheck), err
}

// Update takes the representation of a computeHttpHealthCheck and updates it. Returns the server's representation of the computeHttpHealthCheck, and an error, if there is any.
func (c *FakeComputeHttpHealthChecks) Update(computeHttpHealthCheck *v1alpha1.ComputeHttpHealthCheck) (result *v1alpha1.ComputeHttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(computehttphealthchecksResource, computeHttpHealthCheck), &v1alpha1.ComputeHttpHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHttpHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeHttpHealthChecks) UpdateStatus(computeHttpHealthCheck *v1alpha1.ComputeHttpHealthCheck) (*v1alpha1.ComputeHttpHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(computehttphealthchecksResource, "status", computeHttpHealthCheck), &v1alpha1.ComputeHttpHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHttpHealthCheck), err
}

// Delete takes name of the computeHttpHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeComputeHttpHealthChecks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(computehttphealthchecksResource, name), &v1alpha1.ComputeHttpHealthCheck{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeHttpHealthChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(computehttphealthchecksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeHttpHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched computeHttpHealthCheck.
func (c *FakeComputeHttpHealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeHttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(computehttphealthchecksResource, name, pt, data, subresources...), &v1alpha1.ComputeHttpHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHttpHealthCheck), err
}