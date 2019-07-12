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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermApiManagementAuthorizationServers implements AzurermApiManagementAuthorizationServerInterface
type FakeAzurermApiManagementAuthorizationServers struct {
	Fake *FakeAzurermV1alpha1
}

var azurermapimanagementauthorizationserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermapimanagementauthorizationservers"}

var azurermapimanagementauthorizationserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermApiManagementAuthorizationServer"}

// Get takes name of the azurermApiManagementAuthorizationServer, and returns the corresponding azurermApiManagementAuthorizationServer object, and an error if there is any.
func (c *FakeAzurermApiManagementAuthorizationServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermapimanagementauthorizationserversResource, name), &v1alpha1.AzurermApiManagementAuthorizationServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementAuthorizationServer), err
}

// List takes label and field selectors, and returns the list of AzurermApiManagementAuthorizationServers that match those selectors.
func (c *FakeAzurermApiManagementAuthorizationServers) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementAuthorizationServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermapimanagementauthorizationserversResource, azurermapimanagementauthorizationserversKind, opts), &v1alpha1.AzurermApiManagementAuthorizationServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermApiManagementAuthorizationServerList{ListMeta: obj.(*v1alpha1.AzurermApiManagementAuthorizationServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermApiManagementAuthorizationServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementAuthorizationServers.
func (c *FakeAzurermApiManagementAuthorizationServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermapimanagementauthorizationserversResource, opts))
}

// Create takes the representation of a azurermApiManagementAuthorizationServer and creates it.  Returns the server's representation of the azurermApiManagementAuthorizationServer, and an error, if there is any.
func (c *FakeAzurermApiManagementAuthorizationServers) Create(azurermApiManagementAuthorizationServer *v1alpha1.AzurermApiManagementAuthorizationServer) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermapimanagementauthorizationserversResource, azurermApiManagementAuthorizationServer), &v1alpha1.AzurermApiManagementAuthorizationServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementAuthorizationServer), err
}

// Update takes the representation of a azurermApiManagementAuthorizationServer and updates it. Returns the server's representation of the azurermApiManagementAuthorizationServer, and an error, if there is any.
func (c *FakeAzurermApiManagementAuthorizationServers) Update(azurermApiManagementAuthorizationServer *v1alpha1.AzurermApiManagementAuthorizationServer) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermapimanagementauthorizationserversResource, azurermApiManagementAuthorizationServer), &v1alpha1.AzurermApiManagementAuthorizationServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementAuthorizationServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermApiManagementAuthorizationServers) UpdateStatus(azurermApiManagementAuthorizationServer *v1alpha1.AzurermApiManagementAuthorizationServer) (*v1alpha1.AzurermApiManagementAuthorizationServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermapimanagementauthorizationserversResource, "status", azurermApiManagementAuthorizationServer), &v1alpha1.AzurermApiManagementAuthorizationServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementAuthorizationServer), err
}

// Delete takes name of the azurermApiManagementAuthorizationServer and deletes it. Returns an error if one occurs.
func (c *FakeAzurermApiManagementAuthorizationServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermapimanagementauthorizationserversResource, name), &v1alpha1.AzurermApiManagementAuthorizationServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermApiManagementAuthorizationServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermapimanagementauthorizationserversResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermApiManagementAuthorizationServerList{})
	return err
}

// Patch applies the patch and returns the patched azurermApiManagementAuthorizationServer.
func (c *FakeAzurermApiManagementAuthorizationServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementAuthorizationServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermapimanagementauthorizationserversResource, name, pt, data, subresources...), &v1alpha1.AzurermApiManagementAuthorizationServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementAuthorizationServer), err
}
