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

// FakeDxHostedPrivateVirtualInterfaces implements DxHostedPrivateVirtualInterfaceInterface
type FakeDxHostedPrivateVirtualInterfaces struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dxhostedprivatevirtualinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxhostedprivatevirtualinterfaces"}

var dxhostedprivatevirtualinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxHostedPrivateVirtualInterface"}

// Get takes name of the dxHostedPrivateVirtualInterface, and returns the corresponding dxHostedPrivateVirtualInterface object, and an error if there is any.
func (c *FakeDxHostedPrivateVirtualInterfaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dxhostedprivatevirtualinterfacesResource, c.ns, name), &v1alpha1.DxHostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterface), err
}

// List takes label and field selectors, and returns the list of DxHostedPrivateVirtualInterfaces that match those selectors.
func (c *FakeDxHostedPrivateVirtualInterfaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DxHostedPrivateVirtualInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dxhostedprivatevirtualinterfacesResource, dxhostedprivatevirtualinterfacesKind, c.ns, opts), &v1alpha1.DxHostedPrivateVirtualInterfaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxHostedPrivateVirtualInterfaceList{ListMeta: obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxHostedPrivateVirtualInterfaces.
func (c *FakeDxHostedPrivateVirtualInterfaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dxhostedprivatevirtualinterfacesResource, c.ns, opts))

}

// Create takes the representation of a dxHostedPrivateVirtualInterface and creates it.  Returns the server's representation of the dxHostedPrivateVirtualInterface, and an error, if there is any.
func (c *FakeDxHostedPrivateVirtualInterfaces) Create(ctx context.Context, dxHostedPrivateVirtualInterface *v1alpha1.DxHostedPrivateVirtualInterface, opts v1.CreateOptions) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dxhostedprivatevirtualinterfacesResource, c.ns, dxHostedPrivateVirtualInterface), &v1alpha1.DxHostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterface), err
}

// Update takes the representation of a dxHostedPrivateVirtualInterface and updates it. Returns the server's representation of the dxHostedPrivateVirtualInterface, and an error, if there is any.
func (c *FakeDxHostedPrivateVirtualInterfaces) Update(ctx context.Context, dxHostedPrivateVirtualInterface *v1alpha1.DxHostedPrivateVirtualInterface, opts v1.UpdateOptions) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dxhostedprivatevirtualinterfacesResource, c.ns, dxHostedPrivateVirtualInterface), &v1alpha1.DxHostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxHostedPrivateVirtualInterfaces) UpdateStatus(ctx context.Context, dxHostedPrivateVirtualInterface *v1alpha1.DxHostedPrivateVirtualInterface, opts v1.UpdateOptions) (*v1alpha1.DxHostedPrivateVirtualInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dxhostedprivatevirtualinterfacesResource, "status", c.ns, dxHostedPrivateVirtualInterface), &v1alpha1.DxHostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterface), err
}

// Delete takes name of the dxHostedPrivateVirtualInterface and deletes it. Returns an error if one occurs.
func (c *FakeDxHostedPrivateVirtualInterfaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dxhostedprivatevirtualinterfacesResource, c.ns, name), &v1alpha1.DxHostedPrivateVirtualInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxHostedPrivateVirtualInterfaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dxhostedprivatevirtualinterfacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxHostedPrivateVirtualInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched dxHostedPrivateVirtualInterface.
func (c *FakeDxHostedPrivateVirtualInterfaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dxhostedprivatevirtualinterfacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DxHostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterface), err
}
