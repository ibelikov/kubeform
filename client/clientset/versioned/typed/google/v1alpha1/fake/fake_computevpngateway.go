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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeVPNGateways implements ComputeVPNGatewayInterface
type FakeComputeVPNGateways struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computevpngatewaysResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computevpngateways"}

var computevpngatewaysKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeVPNGateway"}

// Get takes name of the computeVPNGateway, and returns the corresponding computeVPNGateway object, and an error if there is any.
func (c *FakeComputeVPNGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computevpngatewaysResource, c.ns, name), &v1alpha1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeVPNGateway), err
}

// List takes label and field selectors, and returns the list of ComputeVPNGateways that match those selectors.
func (c *FakeComputeVPNGateways) List(opts v1.ListOptions) (result *v1alpha1.ComputeVPNGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computevpngatewaysResource, computevpngatewaysKind, c.ns, opts), &v1alpha1.ComputeVPNGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeVPNGatewayList{ListMeta: obj.(*v1alpha1.ComputeVPNGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeVPNGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeVPNGateways.
func (c *FakeComputeVPNGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computevpngatewaysResource, c.ns, opts))

}

// Create takes the representation of a computeVPNGateway and creates it.  Returns the server's representation of the computeVPNGateway, and an error, if there is any.
func (c *FakeComputeVPNGateways) Create(computeVPNGateway *v1alpha1.ComputeVPNGateway) (result *v1alpha1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computevpngatewaysResource, c.ns, computeVPNGateway), &v1alpha1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeVPNGateway), err
}

// Update takes the representation of a computeVPNGateway and updates it. Returns the server's representation of the computeVPNGateway, and an error, if there is any.
func (c *FakeComputeVPNGateways) Update(computeVPNGateway *v1alpha1.ComputeVPNGateway) (result *v1alpha1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computevpngatewaysResource, c.ns, computeVPNGateway), &v1alpha1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeVPNGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeVPNGateways) UpdateStatus(computeVPNGateway *v1alpha1.ComputeVPNGateway) (*v1alpha1.ComputeVPNGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computevpngatewaysResource, "status", c.ns, computeVPNGateway), &v1alpha1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeVPNGateway), err
}

// Delete takes name of the computeVPNGateway and deletes it. Returns an error if one occurs.
func (c *FakeComputeVPNGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computevpngatewaysResource, c.ns, name), &v1alpha1.ComputeVPNGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeVPNGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computevpngatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeVPNGatewayList{})
	return err
}

// Patch applies the patch and returns the patched computeVPNGateway.
func (c *FakeComputeVPNGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computevpngatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeVPNGateway), err
}
