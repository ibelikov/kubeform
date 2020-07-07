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

// FakeDaxSubnetGroups implements DaxSubnetGroupInterface
type FakeDaxSubnetGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var daxsubnetgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "daxsubnetgroups"}

var daxsubnetgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DaxSubnetGroup"}

// Get takes name of the daxSubnetGroup, and returns the corresponding daxSubnetGroup object, and an error if there is any.
func (c *FakeDaxSubnetGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daxsubnetgroupsResource, c.ns, name), &v1alpha1.DaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxSubnetGroup), err
}

// List takes label and field selectors, and returns the list of DaxSubnetGroups that match those selectors.
func (c *FakeDaxSubnetGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DaxSubnetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daxsubnetgroupsResource, daxsubnetgroupsKind, c.ns, opts), &v1alpha1.DaxSubnetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DaxSubnetGroupList{ListMeta: obj.(*v1alpha1.DaxSubnetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DaxSubnetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daxSubnetGroups.
func (c *FakeDaxSubnetGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daxsubnetgroupsResource, c.ns, opts))

}

// Create takes the representation of a daxSubnetGroup and creates it.  Returns the server's representation of the daxSubnetGroup, and an error, if there is any.
func (c *FakeDaxSubnetGroups) Create(ctx context.Context, daxSubnetGroup *v1alpha1.DaxSubnetGroup, opts v1.CreateOptions) (result *v1alpha1.DaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daxsubnetgroupsResource, c.ns, daxSubnetGroup), &v1alpha1.DaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxSubnetGroup), err
}

// Update takes the representation of a daxSubnetGroup and updates it. Returns the server's representation of the daxSubnetGroup, and an error, if there is any.
func (c *FakeDaxSubnetGroups) Update(ctx context.Context, daxSubnetGroup *v1alpha1.DaxSubnetGroup, opts v1.UpdateOptions) (result *v1alpha1.DaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daxsubnetgroupsResource, c.ns, daxSubnetGroup), &v1alpha1.DaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxSubnetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDaxSubnetGroups) UpdateStatus(ctx context.Context, daxSubnetGroup *v1alpha1.DaxSubnetGroup, opts v1.UpdateOptions) (*v1alpha1.DaxSubnetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(daxsubnetgroupsResource, "status", c.ns, daxSubnetGroup), &v1alpha1.DaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxSubnetGroup), err
}

// Delete takes name of the daxSubnetGroup and deletes it. Returns an error if one occurs.
func (c *FakeDaxSubnetGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(daxsubnetgroupsResource, c.ns, name), &v1alpha1.DaxSubnetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaxSubnetGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daxsubnetgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DaxSubnetGroupList{})
	return err
}

// Patch applies the patch and returns the patched daxSubnetGroup.
func (c *FakeDaxSubnetGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DaxSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daxsubnetgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DaxSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DaxSubnetGroup), err
}
