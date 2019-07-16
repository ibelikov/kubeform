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

// FakeNetworkInterfaces implements NetworkInterfaceInterface
type FakeNetworkInterfaces struct {
	Fake *FakeAwsV1alpha1
}

var networkinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "networkinterfaces"}

var networkinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "NetworkInterface"}

// Get takes name of the networkInterface, and returns the corresponding networkInterface object, and an error if there is any.
func (c *FakeNetworkInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(networkinterfacesResource, name), &v1alpha1.NetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaces that match those selectors.
func (c *FakeNetworkInterfaces) List(opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(networkinterfacesResource, networkinterfacesKind, opts), &v1alpha1.NetworkInterfaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceList{ListMeta: obj.(*v1alpha1.NetworkInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaces.
func (c *FakeNetworkInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(networkinterfacesResource, opts))
}

// Create takes the representation of a networkInterface and creates it.  Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *FakeNetworkInterfaces) Create(networkInterface *v1alpha1.NetworkInterface) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(networkinterfacesResource, networkInterface), &v1alpha1.NetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// Update takes the representation of a networkInterface and updates it. Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *FakeNetworkInterfaces) Update(networkInterface *v1alpha1.NetworkInterface) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(networkinterfacesResource, networkInterface), &v1alpha1.NetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkInterfaces) UpdateStatus(networkInterface *v1alpha1.NetworkInterface) (*v1alpha1.NetworkInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(networkinterfacesResource, "status", networkInterface), &v1alpha1.NetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// Delete takes name of the networkInterface and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(networkinterfacesResource, name), &v1alpha1.NetworkInterface{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(networkinterfacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched networkInterface.
func (c *FakeNetworkInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(networkinterfacesResource, name, pt, data, subresources...), &v1alpha1.NetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}
