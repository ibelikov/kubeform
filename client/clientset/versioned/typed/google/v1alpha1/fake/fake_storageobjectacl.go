/*
Copyright The Kubeform Authors.

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

// FakeStorageObjectACLs implements StorageObjectACLInterface
type FakeStorageObjectACLs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var storageobjectaclsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storageobjectacls"}

var storageobjectaclsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageObjectACL"}

// Get takes name of the storageObjectACL, and returns the corresponding storageObjectACL object, and an error if there is any.
func (c *FakeStorageObjectACLs) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storageobjectaclsResource, c.ns, name), &v1alpha1.StorageObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectACL), err
}

// List takes label and field selectors, and returns the list of StorageObjectACLs that match those selectors.
func (c *FakeStorageObjectACLs) List(opts v1.ListOptions) (result *v1alpha1.StorageObjectACLList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storageobjectaclsResource, storageobjectaclsKind, c.ns, opts), &v1alpha1.StorageObjectACLList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageObjectACLList{ListMeta: obj.(*v1alpha1.StorageObjectACLList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageObjectACLList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageObjectACLs.
func (c *FakeStorageObjectACLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storageobjectaclsResource, c.ns, opts))

}

// Create takes the representation of a storageObjectACL and creates it.  Returns the server's representation of the storageObjectACL, and an error, if there is any.
func (c *FakeStorageObjectACLs) Create(storageObjectACL *v1alpha1.StorageObjectACL) (result *v1alpha1.StorageObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storageobjectaclsResource, c.ns, storageObjectACL), &v1alpha1.StorageObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectACL), err
}

// Update takes the representation of a storageObjectACL and updates it. Returns the server's representation of the storageObjectACL, and an error, if there is any.
func (c *FakeStorageObjectACLs) Update(storageObjectACL *v1alpha1.StorageObjectACL) (result *v1alpha1.StorageObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storageobjectaclsResource, c.ns, storageObjectACL), &v1alpha1.StorageObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectACL), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageObjectACLs) UpdateStatus(storageObjectACL *v1alpha1.StorageObjectACL) (*v1alpha1.StorageObjectACL, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storageobjectaclsResource, "status", c.ns, storageObjectACL), &v1alpha1.StorageObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectACL), err
}

// Delete takes name of the storageObjectACL and deletes it. Returns an error if one occurs.
func (c *FakeStorageObjectACLs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storageobjectaclsResource, c.ns, name), &v1alpha1.StorageObjectACL{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageObjectACLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storageobjectaclsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageObjectACLList{})
	return err
}

// Patch applies the patch and returns the patched storageObjectACL.
func (c *FakeStorageObjectACLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storageobjectaclsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectACL), err
}
