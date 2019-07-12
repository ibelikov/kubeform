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

// FakeGoogleContainerNodePools implements GoogleContainerNodePoolInterface
type FakeGoogleContainerNodePools struct {
	Fake *FakeGoogleV1alpha1
}

var googlecontainernodepoolsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecontainernodepools"}

var googlecontainernodepoolsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleContainerNodePool"}

// Get takes name of the googleContainerNodePool, and returns the corresponding googleContainerNodePool object, and an error if there is any.
func (c *FakeGoogleContainerNodePools) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecontainernodepoolsResource, name), &v1alpha1.GoogleContainerNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerNodePool), err
}

// List takes label and field selectors, and returns the list of GoogleContainerNodePools that match those selectors.
func (c *FakeGoogleContainerNodePools) List(opts v1.ListOptions) (result *v1alpha1.GoogleContainerNodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecontainernodepoolsResource, googlecontainernodepoolsKind, opts), &v1alpha1.GoogleContainerNodePoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleContainerNodePoolList{ListMeta: obj.(*v1alpha1.GoogleContainerNodePoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleContainerNodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleContainerNodePools.
func (c *FakeGoogleContainerNodePools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecontainernodepoolsResource, opts))
}

// Create takes the representation of a googleContainerNodePool and creates it.  Returns the server's representation of the googleContainerNodePool, and an error, if there is any.
func (c *FakeGoogleContainerNodePools) Create(googleContainerNodePool *v1alpha1.GoogleContainerNodePool) (result *v1alpha1.GoogleContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecontainernodepoolsResource, googleContainerNodePool), &v1alpha1.GoogleContainerNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerNodePool), err
}

// Update takes the representation of a googleContainerNodePool and updates it. Returns the server's representation of the googleContainerNodePool, and an error, if there is any.
func (c *FakeGoogleContainerNodePools) Update(googleContainerNodePool *v1alpha1.GoogleContainerNodePool) (result *v1alpha1.GoogleContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecontainernodepoolsResource, googleContainerNodePool), &v1alpha1.GoogleContainerNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerNodePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleContainerNodePools) UpdateStatus(googleContainerNodePool *v1alpha1.GoogleContainerNodePool) (*v1alpha1.GoogleContainerNodePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecontainernodepoolsResource, "status", googleContainerNodePool), &v1alpha1.GoogleContainerNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerNodePool), err
}

// Delete takes name of the googleContainerNodePool and deletes it. Returns an error if one occurs.
func (c *FakeGoogleContainerNodePools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecontainernodepoolsResource, name), &v1alpha1.GoogleContainerNodePool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleContainerNodePools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecontainernodepoolsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleContainerNodePoolList{})
	return err
}

// Patch applies the patch and returns the patched googleContainerNodePool.
func (c *FakeGoogleContainerNodePools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecontainernodepoolsResource, name, pt, data, subresources...), &v1alpha1.GoogleContainerNodePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerNodePool), err
}
