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

// FakeGoogleComputeInstanceGroupManagers implements GoogleComputeInstanceGroupManagerInterface
type FakeGoogleComputeInstanceGroupManagers struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputeinstancegroupmanagersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputeinstancegroupmanagers"}

var googlecomputeinstancegroupmanagersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeInstanceGroupManager"}

// Get takes name of the googleComputeInstanceGroupManager, and returns the corresponding googleComputeInstanceGroupManager object, and an error if there is any.
func (c *FakeGoogleComputeInstanceGroupManagers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputeinstancegroupmanagersResource, name), &v1alpha1.GoogleComputeInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeInstanceGroupManager), err
}

// List takes label and field selectors, and returns the list of GoogleComputeInstanceGroupManagers that match those selectors.
func (c *FakeGoogleComputeInstanceGroupManagers) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeInstanceGroupManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputeinstancegroupmanagersResource, googlecomputeinstancegroupmanagersKind, opts), &v1alpha1.GoogleComputeInstanceGroupManagerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeInstanceGroupManagerList{ListMeta: obj.(*v1alpha1.GoogleComputeInstanceGroupManagerList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeInstanceGroupManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeInstanceGroupManagers.
func (c *FakeGoogleComputeInstanceGroupManagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputeinstancegroupmanagersResource, opts))
}

// Create takes the representation of a googleComputeInstanceGroupManager and creates it.  Returns the server's representation of the googleComputeInstanceGroupManager, and an error, if there is any.
func (c *FakeGoogleComputeInstanceGroupManagers) Create(googleComputeInstanceGroupManager *v1alpha1.GoogleComputeInstanceGroupManager) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputeinstancegroupmanagersResource, googleComputeInstanceGroupManager), &v1alpha1.GoogleComputeInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeInstanceGroupManager), err
}

// Update takes the representation of a googleComputeInstanceGroupManager and updates it. Returns the server's representation of the googleComputeInstanceGroupManager, and an error, if there is any.
func (c *FakeGoogleComputeInstanceGroupManagers) Update(googleComputeInstanceGroupManager *v1alpha1.GoogleComputeInstanceGroupManager) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputeinstancegroupmanagersResource, googleComputeInstanceGroupManager), &v1alpha1.GoogleComputeInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeInstanceGroupManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeInstanceGroupManagers) UpdateStatus(googleComputeInstanceGroupManager *v1alpha1.GoogleComputeInstanceGroupManager) (*v1alpha1.GoogleComputeInstanceGroupManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputeinstancegroupmanagersResource, "status", googleComputeInstanceGroupManager), &v1alpha1.GoogleComputeInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeInstanceGroupManager), err
}

// Delete takes name of the googleComputeInstanceGroupManager and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeInstanceGroupManagers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputeinstancegroupmanagersResource, name), &v1alpha1.GoogleComputeInstanceGroupManager{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeInstanceGroupManagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputeinstancegroupmanagersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeInstanceGroupManagerList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeInstanceGroupManager.
func (c *FakeGoogleComputeInstanceGroupManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputeinstancegroupmanagersResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeInstanceGroupManager), err
}
