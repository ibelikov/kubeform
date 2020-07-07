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

// FakeIapAppEngineServiceIamBindings implements IapAppEngineServiceIamBindingInterface
type FakeIapAppEngineServiceIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var iapappengineserviceiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "iapappengineserviceiambindings"}

var iapappengineserviceiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "IapAppEngineServiceIamBinding"}

// Get takes name of the iapAppEngineServiceIamBinding, and returns the corresponding iapAppEngineServiceIamBinding object, and an error if there is any.
func (c *FakeIapAppEngineServiceIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IapAppEngineServiceIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iapappengineserviceiambindingsResource, c.ns, name), &v1alpha1.IapAppEngineServiceIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamBinding), err
}

// List takes label and field selectors, and returns the list of IapAppEngineServiceIamBindings that match those selectors.
func (c *FakeIapAppEngineServiceIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IapAppEngineServiceIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iapappengineserviceiambindingsResource, iapappengineserviceiambindingsKind, c.ns, opts), &v1alpha1.IapAppEngineServiceIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IapAppEngineServiceIamBindingList{ListMeta: obj.(*v1alpha1.IapAppEngineServiceIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.IapAppEngineServiceIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iapAppEngineServiceIamBindings.
func (c *FakeIapAppEngineServiceIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iapappengineserviceiambindingsResource, c.ns, opts))

}

// Create takes the representation of a iapAppEngineServiceIamBinding and creates it.  Returns the server's representation of the iapAppEngineServiceIamBinding, and an error, if there is any.
func (c *FakeIapAppEngineServiceIamBindings) Create(ctx context.Context, iapAppEngineServiceIamBinding *v1alpha1.IapAppEngineServiceIamBinding, opts v1.CreateOptions) (result *v1alpha1.IapAppEngineServiceIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iapappengineserviceiambindingsResource, c.ns, iapAppEngineServiceIamBinding), &v1alpha1.IapAppEngineServiceIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamBinding), err
}

// Update takes the representation of a iapAppEngineServiceIamBinding and updates it. Returns the server's representation of the iapAppEngineServiceIamBinding, and an error, if there is any.
func (c *FakeIapAppEngineServiceIamBindings) Update(ctx context.Context, iapAppEngineServiceIamBinding *v1alpha1.IapAppEngineServiceIamBinding, opts v1.UpdateOptions) (result *v1alpha1.IapAppEngineServiceIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iapappengineserviceiambindingsResource, c.ns, iapAppEngineServiceIamBinding), &v1alpha1.IapAppEngineServiceIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIapAppEngineServiceIamBindings) UpdateStatus(ctx context.Context, iapAppEngineServiceIamBinding *v1alpha1.IapAppEngineServiceIamBinding, opts v1.UpdateOptions) (*v1alpha1.IapAppEngineServiceIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iapappengineserviceiambindingsResource, "status", c.ns, iapAppEngineServiceIamBinding), &v1alpha1.IapAppEngineServiceIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamBinding), err
}

// Delete takes name of the iapAppEngineServiceIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeIapAppEngineServiceIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iapappengineserviceiambindingsResource, c.ns, name), &v1alpha1.IapAppEngineServiceIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIapAppEngineServiceIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iapappengineserviceiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IapAppEngineServiceIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched iapAppEngineServiceIamBinding.
func (c *FakeIapAppEngineServiceIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IapAppEngineServiceIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iapappengineserviceiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IapAppEngineServiceIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamBinding), err
}
