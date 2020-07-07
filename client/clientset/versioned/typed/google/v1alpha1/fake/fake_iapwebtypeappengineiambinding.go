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

// FakeIapWebTypeAppEngineIamBindings implements IapWebTypeAppEngineIamBindingInterface
type FakeIapWebTypeAppEngineIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var iapwebtypeappengineiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "iapwebtypeappengineiambindings"}

var iapwebtypeappengineiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "IapWebTypeAppEngineIamBinding"}

// Get takes name of the iapWebTypeAppEngineIamBinding, and returns the corresponding iapWebTypeAppEngineIamBinding object, and an error if there is any.
func (c *FakeIapWebTypeAppEngineIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IapWebTypeAppEngineIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iapwebtypeappengineiambindingsResource, c.ns, name), &v1alpha1.IapWebTypeAppEngineIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebTypeAppEngineIamBinding), err
}

// List takes label and field selectors, and returns the list of IapWebTypeAppEngineIamBindings that match those selectors.
func (c *FakeIapWebTypeAppEngineIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IapWebTypeAppEngineIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iapwebtypeappengineiambindingsResource, iapwebtypeappengineiambindingsKind, c.ns, opts), &v1alpha1.IapWebTypeAppEngineIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IapWebTypeAppEngineIamBindingList{ListMeta: obj.(*v1alpha1.IapWebTypeAppEngineIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.IapWebTypeAppEngineIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iapWebTypeAppEngineIamBindings.
func (c *FakeIapWebTypeAppEngineIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iapwebtypeappengineiambindingsResource, c.ns, opts))

}

// Create takes the representation of a iapWebTypeAppEngineIamBinding and creates it.  Returns the server's representation of the iapWebTypeAppEngineIamBinding, and an error, if there is any.
func (c *FakeIapWebTypeAppEngineIamBindings) Create(ctx context.Context, iapWebTypeAppEngineIamBinding *v1alpha1.IapWebTypeAppEngineIamBinding, opts v1.CreateOptions) (result *v1alpha1.IapWebTypeAppEngineIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iapwebtypeappengineiambindingsResource, c.ns, iapWebTypeAppEngineIamBinding), &v1alpha1.IapWebTypeAppEngineIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebTypeAppEngineIamBinding), err
}

// Update takes the representation of a iapWebTypeAppEngineIamBinding and updates it. Returns the server's representation of the iapWebTypeAppEngineIamBinding, and an error, if there is any.
func (c *FakeIapWebTypeAppEngineIamBindings) Update(ctx context.Context, iapWebTypeAppEngineIamBinding *v1alpha1.IapWebTypeAppEngineIamBinding, opts v1.UpdateOptions) (result *v1alpha1.IapWebTypeAppEngineIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iapwebtypeappengineiambindingsResource, c.ns, iapWebTypeAppEngineIamBinding), &v1alpha1.IapWebTypeAppEngineIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebTypeAppEngineIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIapWebTypeAppEngineIamBindings) UpdateStatus(ctx context.Context, iapWebTypeAppEngineIamBinding *v1alpha1.IapWebTypeAppEngineIamBinding, opts v1.UpdateOptions) (*v1alpha1.IapWebTypeAppEngineIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iapwebtypeappengineiambindingsResource, "status", c.ns, iapWebTypeAppEngineIamBinding), &v1alpha1.IapWebTypeAppEngineIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebTypeAppEngineIamBinding), err
}

// Delete takes name of the iapWebTypeAppEngineIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeIapWebTypeAppEngineIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iapwebtypeappengineiambindingsResource, c.ns, name), &v1alpha1.IapWebTypeAppEngineIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIapWebTypeAppEngineIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iapwebtypeappengineiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IapWebTypeAppEngineIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched iapWebTypeAppEngineIamBinding.
func (c *FakeIapWebTypeAppEngineIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IapWebTypeAppEngineIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iapwebtypeappengineiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IapWebTypeAppEngineIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapWebTypeAppEngineIamBinding), err
}
