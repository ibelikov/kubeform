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

// FakeStorageDefaultObjectAccessControls implements StorageDefaultObjectAccessControlInterface
type FakeStorageDefaultObjectAccessControls struct {
	Fake *FakeGoogleV1alpha1
}

var storagedefaultobjectaccesscontrolsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagedefaultobjectaccesscontrols"}

var storagedefaultobjectaccesscontrolsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageDefaultObjectAccessControl"}

// Get takes name of the storageDefaultObjectAccessControl, and returns the corresponding storageDefaultObjectAccessControl object, and an error if there is any.
func (c *FakeStorageDefaultObjectAccessControls) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storagedefaultobjectaccesscontrolsResource, name), &v1alpha1.StorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAccessControl), err
}

// List takes label and field selectors, and returns the list of StorageDefaultObjectAccessControls that match those selectors.
func (c *FakeStorageDefaultObjectAccessControls) List(opts v1.ListOptions) (result *v1alpha1.StorageDefaultObjectAccessControlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storagedefaultobjectaccesscontrolsResource, storagedefaultobjectaccesscontrolsKind, opts), &v1alpha1.StorageDefaultObjectAccessControlList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageDefaultObjectAccessControlList{ListMeta: obj.(*v1alpha1.StorageDefaultObjectAccessControlList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageDefaultObjectAccessControlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageDefaultObjectAccessControls.
func (c *FakeStorageDefaultObjectAccessControls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storagedefaultobjectaccesscontrolsResource, opts))
}

// Create takes the representation of a storageDefaultObjectAccessControl and creates it.  Returns the server's representation of the storageDefaultObjectAccessControl, and an error, if there is any.
func (c *FakeStorageDefaultObjectAccessControls) Create(storageDefaultObjectAccessControl *v1alpha1.StorageDefaultObjectAccessControl) (result *v1alpha1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storagedefaultobjectaccesscontrolsResource, storageDefaultObjectAccessControl), &v1alpha1.StorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAccessControl), err
}

// Update takes the representation of a storageDefaultObjectAccessControl and updates it. Returns the server's representation of the storageDefaultObjectAccessControl, and an error, if there is any.
func (c *FakeStorageDefaultObjectAccessControls) Update(storageDefaultObjectAccessControl *v1alpha1.StorageDefaultObjectAccessControl) (result *v1alpha1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storagedefaultobjectaccesscontrolsResource, storageDefaultObjectAccessControl), &v1alpha1.StorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAccessControl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageDefaultObjectAccessControls) UpdateStatus(storageDefaultObjectAccessControl *v1alpha1.StorageDefaultObjectAccessControl) (*v1alpha1.StorageDefaultObjectAccessControl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(storagedefaultobjectaccesscontrolsResource, "status", storageDefaultObjectAccessControl), &v1alpha1.StorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAccessControl), err
}

// Delete takes name of the storageDefaultObjectAccessControl and deletes it. Returns an error if one occurs.
func (c *FakeStorageDefaultObjectAccessControls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storagedefaultobjectaccesscontrolsResource, name), &v1alpha1.StorageDefaultObjectAccessControl{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageDefaultObjectAccessControls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storagedefaultobjectaccesscontrolsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageDefaultObjectAccessControlList{})
	return err
}

// Patch applies the patch and returns the patched storageDefaultObjectAccessControl.
func (c *FakeStorageDefaultObjectAccessControls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageDefaultObjectAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storagedefaultobjectaccesscontrolsResource, name, pt, data, subresources...), &v1alpha1.StorageDefaultObjectAccessControl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAccessControl), err
}
