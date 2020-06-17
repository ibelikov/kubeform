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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeNetworkInterfaceApplicationSecurityGroupAssociations implements NetworkInterfaceApplicationSecurityGroupAssociationInterface
type FakeNetworkInterfaceApplicationSecurityGroupAssociations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networkinterfaceapplicationsecuritygroupassociationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkinterfaceapplicationsecuritygroupassociations"}

var networkinterfaceapplicationsecuritygroupassociationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkInterfaceApplicationSecurityGroupAssociation"}

// Get takes name of the networkInterfaceApplicationSecurityGroupAssociation, and returns the corresponding networkInterfaceApplicationSecurityGroupAssociation object, and an error if there is any.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkinterfaceapplicationsecuritygroupassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaceApplicationSecurityGroupAssociations that match those selectors.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) List(opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkinterfaceapplicationsecuritygroupassociationsResource, networkinterfaceapplicationsecuritygroupassociationsKind, c.ns, opts), &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociationList{ListMeta: obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaceApplicationSecurityGroupAssociations.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkinterfaceapplicationsecuritygroupassociationsResource, c.ns, opts))

}

// Create takes the representation of a networkInterfaceApplicationSecurityGroupAssociation and creates it.  Returns the server's representation of the networkInterfaceApplicationSecurityGroupAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) Create(networkInterfaceApplicationSecurityGroupAssociation *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation) (result *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkinterfaceapplicationsecuritygroupassociationsResource, c.ns, networkInterfaceApplicationSecurityGroupAssociation), &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation), err
}

// Update takes the representation of a networkInterfaceApplicationSecurityGroupAssociation and updates it. Returns the server's representation of the networkInterfaceApplicationSecurityGroupAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) Update(networkInterfaceApplicationSecurityGroupAssociation *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation) (result *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkinterfaceapplicationsecuritygroupassociationsResource, c.ns, networkInterfaceApplicationSecurityGroupAssociation), &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) UpdateStatus(networkInterfaceApplicationSecurityGroupAssociation *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation) (*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkinterfaceapplicationsecuritygroupassociationsResource, "status", c.ns, networkInterfaceApplicationSecurityGroupAssociation), &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation), err
}

// Delete takes name of the networkInterfaceApplicationSecurityGroupAssociation and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkinterfaceapplicationsecuritygroupassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkinterfaceapplicationsecuritygroupassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociationList{})
	return err
}

// Patch applies the patch and returns the patched networkInterfaceApplicationSecurityGroupAssociation.
func (c *FakeNetworkInterfaceApplicationSecurityGroupAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkinterfaceapplicationsecuritygroupassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation), err
}
