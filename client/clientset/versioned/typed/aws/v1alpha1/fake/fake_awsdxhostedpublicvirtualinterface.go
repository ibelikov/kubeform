/*
Copyright 2019 The Kubeform Authors.

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

// FakeAwsDxHostedPublicVirtualInterfaces implements AwsDxHostedPublicVirtualInterfaceInterface
type FakeAwsDxHostedPublicVirtualInterfaces struct {
	Fake *FakeAwsV1alpha1
}

var awsdxhostedpublicvirtualinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdxhostedpublicvirtualinterfaces"}

var awsdxhostedpublicvirtualinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDxHostedPublicVirtualInterface"}

// Get takes name of the awsDxHostedPublicVirtualInterface, and returns the corresponding awsDxHostedPublicVirtualInterface object, and an error if there is any.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdxhostedpublicvirtualinterfacesResource, name), &v1alpha1.AwsDxHostedPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxHostedPublicVirtualInterface), err
}

// List takes label and field selectors, and returns the list of AwsDxHostedPublicVirtualInterfaces that match those selectors.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) List(opts v1.ListOptions) (result *v1alpha1.AwsDxHostedPublicVirtualInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdxhostedpublicvirtualinterfacesResource, awsdxhostedpublicvirtualinterfacesKind, opts), &v1alpha1.AwsDxHostedPublicVirtualInterfaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDxHostedPublicVirtualInterfaceList{ListMeta: obj.(*v1alpha1.AwsDxHostedPublicVirtualInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDxHostedPublicVirtualInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDxHostedPublicVirtualInterfaces.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdxhostedpublicvirtualinterfacesResource, opts))
}

// Create takes the representation of a awsDxHostedPublicVirtualInterface and creates it.  Returns the server's representation of the awsDxHostedPublicVirtualInterface, and an error, if there is any.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) Create(awsDxHostedPublicVirtualInterface *v1alpha1.AwsDxHostedPublicVirtualInterface) (result *v1alpha1.AwsDxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdxhostedpublicvirtualinterfacesResource, awsDxHostedPublicVirtualInterface), &v1alpha1.AwsDxHostedPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxHostedPublicVirtualInterface), err
}

// Update takes the representation of a awsDxHostedPublicVirtualInterface and updates it. Returns the server's representation of the awsDxHostedPublicVirtualInterface, and an error, if there is any.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) Update(awsDxHostedPublicVirtualInterface *v1alpha1.AwsDxHostedPublicVirtualInterface) (result *v1alpha1.AwsDxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdxhostedpublicvirtualinterfacesResource, awsDxHostedPublicVirtualInterface), &v1alpha1.AwsDxHostedPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxHostedPublicVirtualInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDxHostedPublicVirtualInterfaces) UpdateStatus(awsDxHostedPublicVirtualInterface *v1alpha1.AwsDxHostedPublicVirtualInterface) (*v1alpha1.AwsDxHostedPublicVirtualInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdxhostedpublicvirtualinterfacesResource, "status", awsDxHostedPublicVirtualInterface), &v1alpha1.AwsDxHostedPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxHostedPublicVirtualInterface), err
}

// Delete takes name of the awsDxHostedPublicVirtualInterface and deletes it. Returns an error if one occurs.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdxhostedpublicvirtualinterfacesResource, name), &v1alpha1.AwsDxHostedPublicVirtualInterface{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdxhostedpublicvirtualinterfacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDxHostedPublicVirtualInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched awsDxHostedPublicVirtualInterface.
func (c *FakeAwsDxHostedPublicVirtualInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDxHostedPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdxhostedpublicvirtualinterfacesResource, name, pt, data, subresources...), &v1alpha1.AwsDxHostedPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxHostedPublicVirtualInterface), err
}
