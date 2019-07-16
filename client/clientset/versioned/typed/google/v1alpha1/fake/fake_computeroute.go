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

// FakeComputeRoutes implements ComputeRouteInterface
type FakeComputeRoutes struct {
	Fake *FakeGoogleV1alpha1
}

var computeroutesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeroutes"}

var computeroutesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeRoute"}

// Get takes name of the computeRoute, and returns the corresponding computeRoute object, and an error if there is any.
func (c *FakeComputeRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(computeroutesResource, name), &v1alpha1.ComputeRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRoute), err
}

// List takes label and field selectors, and returns the list of ComputeRoutes that match those selectors.
func (c *FakeComputeRoutes) List(opts v1.ListOptions) (result *v1alpha1.ComputeRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(computeroutesResource, computeroutesKind, opts), &v1alpha1.ComputeRouteList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeRouteList{ListMeta: obj.(*v1alpha1.ComputeRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeRoutes.
func (c *FakeComputeRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(computeroutesResource, opts))
}

// Create takes the representation of a computeRoute and creates it.  Returns the server's representation of the computeRoute, and an error, if there is any.
func (c *FakeComputeRoutes) Create(computeRoute *v1alpha1.ComputeRoute) (result *v1alpha1.ComputeRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(computeroutesResource, computeRoute), &v1alpha1.ComputeRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRoute), err
}

// Update takes the representation of a computeRoute and updates it. Returns the server's representation of the computeRoute, and an error, if there is any.
func (c *FakeComputeRoutes) Update(computeRoute *v1alpha1.ComputeRoute) (result *v1alpha1.ComputeRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(computeroutesResource, computeRoute), &v1alpha1.ComputeRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeRoutes) UpdateStatus(computeRoute *v1alpha1.ComputeRoute) (*v1alpha1.ComputeRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(computeroutesResource, "status", computeRoute), &v1alpha1.ComputeRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRoute), err
}

// Delete takes name of the computeRoute and deletes it. Returns an error if one occurs.
func (c *FakeComputeRoutes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(computeroutesResource, name), &v1alpha1.ComputeRoute{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(computeroutesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeRouteList{})
	return err
}

// Patch applies the patch and returns the patched computeRoute.
func (c *FakeComputeRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(computeroutesResource, name, pt, data, subresources...), &v1alpha1.ComputeRoute{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRoute), err
}
