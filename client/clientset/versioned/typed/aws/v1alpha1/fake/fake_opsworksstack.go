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

// FakeOpsworksStacks implements OpsworksStackInterface
type FakeOpsworksStacks struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworksstacksResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksstacks"}

var opsworksstacksKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksStack"}

// Get takes name of the opsworksStack, and returns the corresponding opsworksStack object, and an error if there is any.
func (c *FakeOpsworksStacks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpsworksStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworksstacksResource, c.ns, name), &v1alpha1.OpsworksStack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksStack), err
}

// List takes label and field selectors, and returns the list of OpsworksStacks that match those selectors.
func (c *FakeOpsworksStacks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpsworksStackList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworksstacksResource, opsworksstacksKind, c.ns, opts), &v1alpha1.OpsworksStackList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksStackList{ListMeta: obj.(*v1alpha1.OpsworksStackList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksStackList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksStacks.
func (c *FakeOpsworksStacks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworksstacksResource, c.ns, opts))

}

// Create takes the representation of a opsworksStack and creates it.  Returns the server's representation of the opsworksStack, and an error, if there is any.
func (c *FakeOpsworksStacks) Create(ctx context.Context, opsworksStack *v1alpha1.OpsworksStack, opts v1.CreateOptions) (result *v1alpha1.OpsworksStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworksstacksResource, c.ns, opsworksStack), &v1alpha1.OpsworksStack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksStack), err
}

// Update takes the representation of a opsworksStack and updates it. Returns the server's representation of the opsworksStack, and an error, if there is any.
func (c *FakeOpsworksStacks) Update(ctx context.Context, opsworksStack *v1alpha1.OpsworksStack, opts v1.UpdateOptions) (result *v1alpha1.OpsworksStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworksstacksResource, c.ns, opsworksStack), &v1alpha1.OpsworksStack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksStack), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksStacks) UpdateStatus(ctx context.Context, opsworksStack *v1alpha1.OpsworksStack, opts v1.UpdateOptions) (*v1alpha1.OpsworksStack, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworksstacksResource, "status", c.ns, opsworksStack), &v1alpha1.OpsworksStack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksStack), err
}

// Delete takes name of the opsworksStack and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksStacks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworksstacksResource, c.ns, name), &v1alpha1.OpsworksStack{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksStacks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworksstacksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksStackList{})
	return err
}

// Patch applies the patch and returns the patched opsworksStack.
func (c *FakeOpsworksStacks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpsworksStack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworksstacksResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksStack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksStack), err
}
