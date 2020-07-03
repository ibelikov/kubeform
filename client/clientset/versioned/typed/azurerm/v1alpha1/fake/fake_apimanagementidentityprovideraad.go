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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiManagementIdentityProviderAads implements ApiManagementIdentityProviderAadInterface
type FakeApiManagementIdentityProviderAads struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementidentityprovideraadsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementidentityprovideraads"}

var apimanagementidentityprovideraadsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementIdentityProviderAad"}

// Get takes name of the apiManagementIdentityProviderAad, and returns the corresponding apiManagementIdentityProviderAad object, and an error if there is any.
func (c *FakeApiManagementIdentityProviderAads) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementIdentityProviderAad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementidentityprovideraadsResource, c.ns, name), &v1alpha1.ApiManagementIdentityProviderAad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementIdentityProviderAad), err
}

// List takes label and field selectors, and returns the list of ApiManagementIdentityProviderAads that match those selectors.
func (c *FakeApiManagementIdentityProviderAads) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementIdentityProviderAadList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementidentityprovideraadsResource, apimanagementidentityprovideraadsKind, c.ns, opts), &v1alpha1.ApiManagementIdentityProviderAadList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementIdentityProviderAadList{ListMeta: obj.(*v1alpha1.ApiManagementIdentityProviderAadList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementIdentityProviderAadList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementIdentityProviderAads.
func (c *FakeApiManagementIdentityProviderAads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementidentityprovideraadsResource, c.ns, opts))

}

// Create takes the representation of a apiManagementIdentityProviderAad and creates it.  Returns the server's representation of the apiManagementIdentityProviderAad, and an error, if there is any.
func (c *FakeApiManagementIdentityProviderAads) Create(ctx context.Context, apiManagementIdentityProviderAad *v1alpha1.ApiManagementIdentityProviderAad, opts v1.CreateOptions) (result *v1alpha1.ApiManagementIdentityProviderAad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementidentityprovideraadsResource, c.ns, apiManagementIdentityProviderAad), &v1alpha1.ApiManagementIdentityProviderAad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementIdentityProviderAad), err
}

// Update takes the representation of a apiManagementIdentityProviderAad and updates it. Returns the server's representation of the apiManagementIdentityProviderAad, and an error, if there is any.
func (c *FakeApiManagementIdentityProviderAads) Update(ctx context.Context, apiManagementIdentityProviderAad *v1alpha1.ApiManagementIdentityProviderAad, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementIdentityProviderAad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementidentityprovideraadsResource, c.ns, apiManagementIdentityProviderAad), &v1alpha1.ApiManagementIdentityProviderAad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementIdentityProviderAad), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementIdentityProviderAads) UpdateStatus(ctx context.Context, apiManagementIdentityProviderAad *v1alpha1.ApiManagementIdentityProviderAad, opts v1.UpdateOptions) (*v1alpha1.ApiManagementIdentityProviderAad, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementidentityprovideraadsResource, "status", c.ns, apiManagementIdentityProviderAad), &v1alpha1.ApiManagementIdentityProviderAad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementIdentityProviderAad), err
}

// Delete takes name of the apiManagementIdentityProviderAad and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementIdentityProviderAads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementidentityprovideraadsResource, c.ns, name), &v1alpha1.ApiManagementIdentityProviderAad{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementIdentityProviderAads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementidentityprovideraadsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementIdentityProviderAadList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementIdentityProviderAad.
func (c *FakeApiManagementIdentityProviderAads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementIdentityProviderAad, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementidentityprovideraadsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementIdentityProviderAad{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementIdentityProviderAad), err
}