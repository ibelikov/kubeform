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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeDxHostedPrivateVirtualInterfaceAccepters implements DxHostedPrivateVirtualInterfaceAccepterInterface
type FakeDxHostedPrivateVirtualInterfaceAccepters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dxhostedprivatevirtualinterfaceacceptersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxhostedprivatevirtualinterfaceaccepters"}

var dxhostedprivatevirtualinterfaceacceptersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxHostedPrivateVirtualInterfaceAccepter"}

// Get takes name of the dxHostedPrivateVirtualInterfaceAccepter, and returns the corresponding dxHostedPrivateVirtualInterfaceAccepter object, and an error if there is any.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) Get(name string, options v1.GetOptions) (result *v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dxhostedprivatevirtualinterfaceacceptersResource, c.ns, name), &v1alpha1.DxHostedPrivateVirtualInterfaceAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter), err
}

// List takes label and field selectors, and returns the list of DxHostedPrivateVirtualInterfaceAccepters that match those selectors.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) List(opts v1.ListOptions) (result *v1alpha1.DxHostedPrivateVirtualInterfaceAccepterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dxhostedprivatevirtualinterfaceacceptersResource, dxhostedprivatevirtualinterfaceacceptersKind, c.ns, opts), &v1alpha1.DxHostedPrivateVirtualInterfaceAccepterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxHostedPrivateVirtualInterfaceAccepterList{ListMeta: obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepterList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxHostedPrivateVirtualInterfaceAccepters.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dxhostedprivatevirtualinterfaceacceptersResource, c.ns, opts))

}

// Create takes the representation of a dxHostedPrivateVirtualInterfaceAccepter and creates it.  Returns the server's representation of the dxHostedPrivateVirtualInterfaceAccepter, and an error, if there is any.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) Create(dxHostedPrivateVirtualInterfaceAccepter *v1alpha1.DxHostedPrivateVirtualInterfaceAccepter) (result *v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dxhostedprivatevirtualinterfaceacceptersResource, c.ns, dxHostedPrivateVirtualInterfaceAccepter), &v1alpha1.DxHostedPrivateVirtualInterfaceAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter), err
}

// Update takes the representation of a dxHostedPrivateVirtualInterfaceAccepter and updates it. Returns the server's representation of the dxHostedPrivateVirtualInterfaceAccepter, and an error, if there is any.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) Update(dxHostedPrivateVirtualInterfaceAccepter *v1alpha1.DxHostedPrivateVirtualInterfaceAccepter) (result *v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dxhostedprivatevirtualinterfaceacceptersResource, c.ns, dxHostedPrivateVirtualInterfaceAccepter), &v1alpha1.DxHostedPrivateVirtualInterfaceAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) UpdateStatus(dxHostedPrivateVirtualInterfaceAccepter *v1alpha1.DxHostedPrivateVirtualInterfaceAccepter) (*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dxhostedprivatevirtualinterfaceacceptersResource, "status", c.ns, dxHostedPrivateVirtualInterfaceAccepter), &v1alpha1.DxHostedPrivateVirtualInterfaceAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter), err
}

// Delete takes name of the dxHostedPrivateVirtualInterfaceAccepter and deletes it. Returns an error if one occurs.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dxhostedprivatevirtualinterfaceacceptersResource, c.ns, name), &v1alpha1.DxHostedPrivateVirtualInterfaceAccepter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dxhostedprivatevirtualinterfaceacceptersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxHostedPrivateVirtualInterfaceAccepterList{})
	return err
}

// Patch applies the patch and returns the patched dxHostedPrivateVirtualInterfaceAccepter.
func (c *FakeDxHostedPrivateVirtualInterfaceAccepters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxHostedPrivateVirtualInterfaceAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dxhostedprivatevirtualinterfaceacceptersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DxHostedPrivateVirtualInterfaceAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxHostedPrivateVirtualInterfaceAccepter), err
}
