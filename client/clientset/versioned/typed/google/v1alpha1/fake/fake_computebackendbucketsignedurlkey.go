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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeBackendBucketSignedURLKeys implements ComputeBackendBucketSignedURLKeyInterface
type FakeComputeBackendBucketSignedURLKeys struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computebackendbucketsignedurlkeysResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computebackendbucketsignedurlkeys"}

var computebackendbucketsignedurlkeysKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeBackendBucketSignedURLKey"}

// Get takes name of the computeBackendBucketSignedURLKey, and returns the corresponding computeBackendBucketSignedURLKey object, and an error if there is any.
func (c *FakeComputeBackendBucketSignedURLKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeBackendBucketSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computebackendbucketsignedurlkeysResource, c.ns, name), &v1alpha1.ComputeBackendBucketSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucketSignedURLKey), err
}

// List takes label and field selectors, and returns the list of ComputeBackendBucketSignedURLKeys that match those selectors.
func (c *FakeComputeBackendBucketSignedURLKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeBackendBucketSignedURLKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computebackendbucketsignedurlkeysResource, computebackendbucketsignedurlkeysKind, c.ns, opts), &v1alpha1.ComputeBackendBucketSignedURLKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeBackendBucketSignedURLKeyList{ListMeta: obj.(*v1alpha1.ComputeBackendBucketSignedURLKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeBackendBucketSignedURLKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeBackendBucketSignedURLKeys.
func (c *FakeComputeBackendBucketSignedURLKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computebackendbucketsignedurlkeysResource, c.ns, opts))

}

// Create takes the representation of a computeBackendBucketSignedURLKey and creates it.  Returns the server's representation of the computeBackendBucketSignedURLKey, and an error, if there is any.
func (c *FakeComputeBackendBucketSignedURLKeys) Create(ctx context.Context, computeBackendBucketSignedURLKey *v1alpha1.ComputeBackendBucketSignedURLKey, opts v1.CreateOptions) (result *v1alpha1.ComputeBackendBucketSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computebackendbucketsignedurlkeysResource, c.ns, computeBackendBucketSignedURLKey), &v1alpha1.ComputeBackendBucketSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucketSignedURLKey), err
}

// Update takes the representation of a computeBackendBucketSignedURLKey and updates it. Returns the server's representation of the computeBackendBucketSignedURLKey, and an error, if there is any.
func (c *FakeComputeBackendBucketSignedURLKeys) Update(ctx context.Context, computeBackendBucketSignedURLKey *v1alpha1.ComputeBackendBucketSignedURLKey, opts v1.UpdateOptions) (result *v1alpha1.ComputeBackendBucketSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computebackendbucketsignedurlkeysResource, c.ns, computeBackendBucketSignedURLKey), &v1alpha1.ComputeBackendBucketSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucketSignedURLKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeBackendBucketSignedURLKeys) UpdateStatus(ctx context.Context, computeBackendBucketSignedURLKey *v1alpha1.ComputeBackendBucketSignedURLKey, opts v1.UpdateOptions) (*v1alpha1.ComputeBackendBucketSignedURLKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computebackendbucketsignedurlkeysResource, "status", c.ns, computeBackendBucketSignedURLKey), &v1alpha1.ComputeBackendBucketSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucketSignedURLKey), err
}

// Delete takes name of the computeBackendBucketSignedURLKey and deletes it. Returns an error if one occurs.
func (c *FakeComputeBackendBucketSignedURLKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computebackendbucketsignedurlkeysResource, c.ns, name), &v1alpha1.ComputeBackendBucketSignedURLKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeBackendBucketSignedURLKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computebackendbucketsignedurlkeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeBackendBucketSignedURLKeyList{})
	return err
}

// Patch applies the patch and returns the patched computeBackendBucketSignedURLKey.
func (c *FakeComputeBackendBucketSignedURLKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeBackendBucketSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computebackendbucketsignedurlkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeBackendBucketSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendBucketSignedURLKey), err
}