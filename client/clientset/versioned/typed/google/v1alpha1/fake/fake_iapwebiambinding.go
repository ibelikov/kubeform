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

// FakeIapWebIamBindings implements IapWebIamBindingInterface
type FakeIapWebIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var iapwebiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "iapwebiambindings"}

var iapwebiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "IapWebIamBinding"}

// Get takes name of the iapWebIamBinding, and returns the corresponding iapWebIamBinding object, and an error if there is any.
func (c *FakeIapWebIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IapWebIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iapwebiambindingsResource, c.ns, name), &v1alpha1.IapWebIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebIamBinding), err
}

// List takes label and field selectors, and returns the list of IapWebIamBindings that match those selectors.
func (c *FakeIapWebIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IapWebIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iapwebiambindingsResource, iapwebiambindingsKind, c.ns, opts), &v1alpha1.IapWebIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IapWebIamBindingList{ListMeta: obj.(*v1alpha1.IapWebIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.IapWebIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iapWebIamBindings.
func (c *FakeIapWebIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iapwebiambindingsResource, c.ns, opts))

}

// Create takes the representation of a iapWebIamBinding and creates it.  Returns the server's representation of the iapWebIamBinding, and an error, if there is any.
func (c *FakeIapWebIamBindings) Create(ctx context.Context, iapWebIamBinding *v1alpha1.IapWebIamBinding, opts v1.CreateOptions) (result *v1alpha1.IapWebIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iapwebiambindingsResource, c.ns, iapWebIamBinding), &v1alpha1.IapWebIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebIamBinding), err
}

// Update takes the representation of a iapWebIamBinding and updates it. Returns the server's representation of the iapWebIamBinding, and an error, if there is any.
func (c *FakeIapWebIamBindings) Update(ctx context.Context, iapWebIamBinding *v1alpha1.IapWebIamBinding, opts v1.UpdateOptions) (result *v1alpha1.IapWebIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iapwebiambindingsResource, c.ns, iapWebIamBinding), &v1alpha1.IapWebIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIapWebIamBindings) UpdateStatus(ctx context.Context, iapWebIamBinding *v1alpha1.IapWebIamBinding, opts v1.UpdateOptions) (*v1alpha1.IapWebIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iapwebiambindingsResource, "status", c.ns, iapWebIamBinding), &v1alpha1.IapWebIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebIamBinding), err
}

// Delete takes name of the iapWebIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeIapWebIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iapwebiambindingsResource, c.ns, name), &v1alpha1.IapWebIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIapWebIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iapwebiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IapWebIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched iapWebIamBinding.
func (c *FakeIapWebIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IapWebIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iapwebiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IapWebIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebIamBinding), err
}