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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeEbsVolumes implements EbsVolumeInterface
type FakeEbsVolumes struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ebsvolumesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ebsvolumes"}

var ebsvolumesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EbsVolume"}

// Get takes name of the ebsVolume, and returns the corresponding ebsVolume object, and an error if there is any.
func (c *FakeEbsVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ebsvolumesResource, c.ns, name), &v1alpha1.EbsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsVolume), err
}

// List takes label and field selectors, and returns the list of EbsVolumes that match those selectors.
func (c *FakeEbsVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EbsVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ebsvolumesResource, ebsvolumesKind, c.ns, opts), &v1alpha1.EbsVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EbsVolumeList{ListMeta: obj.(*v1alpha1.EbsVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.EbsVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ebsVolumes.
func (c *FakeEbsVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ebsvolumesResource, c.ns, opts))

}

// Create takes the representation of a ebsVolume and creates it.  Returns the server's representation of the ebsVolume, and an error, if there is any.
func (c *FakeEbsVolumes) Create(ctx context.Context, ebsVolume *v1alpha1.EbsVolume, opts v1.CreateOptions) (result *v1alpha1.EbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ebsvolumesResource, c.ns, ebsVolume), &v1alpha1.EbsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsVolume), err
}

// Update takes the representation of a ebsVolume and updates it. Returns the server's representation of the ebsVolume, and an error, if there is any.
func (c *FakeEbsVolumes) Update(ctx context.Context, ebsVolume *v1alpha1.EbsVolume, opts v1.UpdateOptions) (result *v1alpha1.EbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ebsvolumesResource, c.ns, ebsVolume), &v1alpha1.EbsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEbsVolumes) UpdateStatus(ctx context.Context, ebsVolume *v1alpha1.EbsVolume, opts v1.UpdateOptions) (*v1alpha1.EbsVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ebsvolumesResource, "status", c.ns, ebsVolume), &v1alpha1.EbsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsVolume), err
}

// Delete takes name of the ebsVolume and deletes it. Returns an error if one occurs.
func (c *FakeEbsVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ebsvolumesResource, c.ns, name), &v1alpha1.EbsVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEbsVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ebsvolumesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EbsVolumeList{})
	return err
}

// Patch applies the patch and returns the patched ebsVolume.
func (c *FakeEbsVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EbsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ebsvolumesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EbsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EbsVolume), err
}
