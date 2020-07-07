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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeStorageBucketAccessControls implements StorageBucketAccessControlInterface
type FakeStorageBucketAccessControls struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var storagebucketaccesscontrolsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagebucketaccesscontrols"}

var storagebucketaccesscontrolsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageBucketAccessControl"}

// Get takes name of the storageBucketAccessControl, and returns the corresponding storageBucketAccessControl object, and an error if there is any.
func (c *FakeStorageBucketAccessControls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageBucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagebucketaccesscontrolsResource, c.ns, name), &v1alpha1.StorageBucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketAccessControl), err
}

// List takes label and field selectors, and returns the list of StorageBucketAccessControls that match those selectors.
func (c *FakeStorageBucketAccessControls) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageBucketAccessControlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagebucketaccesscontrolsResource, storagebucketaccesscontrolsKind, c.ns, opts), &v1alpha1.StorageBucketAccessControlList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageBucketAccessControlList{ListMeta: obj.(*v1alpha1.StorageBucketAccessControlList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageBucketAccessControlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageBucketAccessControls.
func (c *FakeStorageBucketAccessControls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagebucketaccesscontrolsResource, c.ns, opts))

}

// Create takes the representation of a storageBucketAccessControl and creates it.  Returns the server's representation of the storageBucketAccessControl, and an error, if there is any.
func (c *FakeStorageBucketAccessControls) Create(ctx context.Context, storageBucketAccessControl *v1alpha1.StorageBucketAccessControl, opts v1.CreateOptions) (result *v1alpha1.StorageBucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagebucketaccesscontrolsResource, c.ns, storageBucketAccessControl), &v1alpha1.StorageBucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketAccessControl), err
}

// Update takes the representation of a storageBucketAccessControl and updates it. Returns the server's representation of the storageBucketAccessControl, and an error, if there is any.
func (c *FakeStorageBucketAccessControls) Update(ctx context.Context, storageBucketAccessControl *v1alpha1.StorageBucketAccessControl, opts v1.UpdateOptions) (result *v1alpha1.StorageBucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagebucketaccesscontrolsResource, c.ns, storageBucketAccessControl), &v1alpha1.StorageBucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketAccessControl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageBucketAccessControls) UpdateStatus(ctx context.Context, storageBucketAccessControl *v1alpha1.StorageBucketAccessControl, opts v1.UpdateOptions) (*v1alpha1.StorageBucketAccessControl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagebucketaccesscontrolsResource, "status", c.ns, storageBucketAccessControl), &v1alpha1.StorageBucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketAccessControl), err
}

// Delete takes name of the storageBucketAccessControl and deletes it. Returns an error if one occurs.
func (c *FakeStorageBucketAccessControls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagebucketaccesscontrolsResource, c.ns, name), &v1alpha1.StorageBucketAccessControl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageBucketAccessControls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagebucketaccesscontrolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageBucketAccessControlList{})
	return err
}

// Patch applies the patch and returns the patched storageBucketAccessControl.
func (c *FakeStorageBucketAccessControls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageBucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagebucketaccesscontrolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageBucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketAccessControl), err
}
