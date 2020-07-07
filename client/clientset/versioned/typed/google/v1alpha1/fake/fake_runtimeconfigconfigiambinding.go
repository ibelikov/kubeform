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

// FakeRuntimeconfigConfigIamBindings implements RuntimeconfigConfigIamBindingInterface
type FakeRuntimeconfigConfigIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var runtimeconfigconfigiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "runtimeconfigconfigiambindings"}

var runtimeconfigconfigiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "RuntimeconfigConfigIamBinding"}

// Get takes name of the runtimeconfigConfigIamBinding, and returns the corresponding runtimeconfigConfigIamBinding object, and an error if there is any.
func (c *FakeRuntimeconfigConfigIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RuntimeconfigConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(runtimeconfigconfigiambindingsResource, c.ns, name), &v1alpha1.RuntimeconfigConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamBinding), err
}

// List takes label and field selectors, and returns the list of RuntimeconfigConfigIamBindings that match those selectors.
func (c *FakeRuntimeconfigConfigIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RuntimeconfigConfigIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(runtimeconfigconfigiambindingsResource, runtimeconfigconfigiambindingsKind, c.ns, opts), &v1alpha1.RuntimeconfigConfigIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RuntimeconfigConfigIamBindingList{ListMeta: obj.(*v1alpha1.RuntimeconfigConfigIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.RuntimeconfigConfigIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runtimeconfigConfigIamBindings.
func (c *FakeRuntimeconfigConfigIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(runtimeconfigconfigiambindingsResource, c.ns, opts))

}

// Create takes the representation of a runtimeconfigConfigIamBinding and creates it.  Returns the server's representation of the runtimeconfigConfigIamBinding, and an error, if there is any.
func (c *FakeRuntimeconfigConfigIamBindings) Create(ctx context.Context, runtimeconfigConfigIamBinding *v1alpha1.RuntimeconfigConfigIamBinding, opts v1.CreateOptions) (result *v1alpha1.RuntimeconfigConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(runtimeconfigconfigiambindingsResource, c.ns, runtimeconfigConfigIamBinding), &v1alpha1.RuntimeconfigConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamBinding), err
}

// Update takes the representation of a runtimeconfigConfigIamBinding and updates it. Returns the server's representation of the runtimeconfigConfigIamBinding, and an error, if there is any.
func (c *FakeRuntimeconfigConfigIamBindings) Update(ctx context.Context, runtimeconfigConfigIamBinding *v1alpha1.RuntimeconfigConfigIamBinding, opts v1.UpdateOptions) (result *v1alpha1.RuntimeconfigConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(runtimeconfigconfigiambindingsResource, c.ns, runtimeconfigConfigIamBinding), &v1alpha1.RuntimeconfigConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRuntimeconfigConfigIamBindings) UpdateStatus(ctx context.Context, runtimeconfigConfigIamBinding *v1alpha1.RuntimeconfigConfigIamBinding, opts v1.UpdateOptions) (*v1alpha1.RuntimeconfigConfigIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(runtimeconfigconfigiambindingsResource, "status", c.ns, runtimeconfigConfigIamBinding), &v1alpha1.RuntimeconfigConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamBinding), err
}

// Delete takes name of the runtimeconfigConfigIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeRuntimeconfigConfigIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(runtimeconfigconfigiambindingsResource, c.ns, name), &v1alpha1.RuntimeconfigConfigIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuntimeconfigConfigIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(runtimeconfigconfigiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RuntimeconfigConfigIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched runtimeconfigConfigIamBinding.
func (c *FakeRuntimeconfigConfigIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RuntimeconfigConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(runtimeconfigconfigiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RuntimeconfigConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamBinding), err
}
