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

// FakeEips implements EipInterface
type FakeEips struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var eipsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "eips"}

var eipsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Eip"}

// Get takes name of the eip, and returns the corresponding eip object, and an error if there is any.
func (c *FakeEips) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Eip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eipsResource, c.ns, name), &v1alpha1.Eip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eip), err
}

// List takes label and field selectors, and returns the list of Eips that match those selectors.
func (c *FakeEips) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eipsResource, eipsKind, c.ns, opts), &v1alpha1.EipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EipList{ListMeta: obj.(*v1alpha1.EipList).ListMeta}
	for _, item := range obj.(*v1alpha1.EipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eips.
func (c *FakeEips) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eipsResource, c.ns, opts))

}

// Create takes the representation of a eip and creates it.  Returns the server's representation of the eip, and an error, if there is any.
func (c *FakeEips) Create(ctx context.Context, eip *v1alpha1.Eip, opts v1.CreateOptions) (result *v1alpha1.Eip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eipsResource, c.ns, eip), &v1alpha1.Eip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eip), err
}

// Update takes the representation of a eip and updates it. Returns the server's representation of the eip, and an error, if there is any.
func (c *FakeEips) Update(ctx context.Context, eip *v1alpha1.Eip, opts v1.UpdateOptions) (result *v1alpha1.Eip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eipsResource, c.ns, eip), &v1alpha1.Eip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eip), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEips) UpdateStatus(ctx context.Context, eip *v1alpha1.Eip, opts v1.UpdateOptions) (*v1alpha1.Eip, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eipsResource, "status", c.ns, eip), &v1alpha1.Eip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eip), err
}

// Delete takes name of the eip and deletes it. Returns an error if one occurs.
func (c *FakeEips) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eipsResource, c.ns, name), &v1alpha1.Eip{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEips) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EipList{})
	return err
}

// Patch applies the patch and returns the patched eip.
func (c *FakeEips) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Eip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Eip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eip), err
}
