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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVpcEndpoints implements VpcEndpointInterface
type FakeVpcEndpoints struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpcendpointsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpcendpoints"}

var vpcendpointsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpcEndpoint"}

// Get takes name of the vpcEndpoint, and returns the corresponding vpcEndpoint object, and an error if there is any.
func (c *FakeVpcEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.VpcEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcendpointsResource, c.ns, name), &v1alpha1.VpcEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpoint), err
}

// List takes label and field selectors, and returns the list of VpcEndpoints that match those selectors.
func (c *FakeVpcEndpoints) List(opts v1.ListOptions) (result *v1alpha1.VpcEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcendpointsResource, vpcendpointsKind, c.ns, opts), &v1alpha1.VpcEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcEndpointList{ListMeta: obj.(*v1alpha1.VpcEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcEndpoints.
func (c *FakeVpcEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcendpointsResource, c.ns, opts))

}

// Create takes the representation of a vpcEndpoint and creates it.  Returns the server's representation of the vpcEndpoint, and an error, if there is any.
func (c *FakeVpcEndpoints) Create(vpcEndpoint *v1alpha1.VpcEndpoint) (result *v1alpha1.VpcEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcendpointsResource, c.ns, vpcEndpoint), &v1alpha1.VpcEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpoint), err
}

// Update takes the representation of a vpcEndpoint and updates it. Returns the server's representation of the vpcEndpoint, and an error, if there is any.
func (c *FakeVpcEndpoints) Update(vpcEndpoint *v1alpha1.VpcEndpoint) (result *v1alpha1.VpcEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcendpointsResource, c.ns, vpcEndpoint), &v1alpha1.VpcEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcEndpoints) UpdateStatus(vpcEndpoint *v1alpha1.VpcEndpoint) (*v1alpha1.VpcEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcendpointsResource, "status", c.ns, vpcEndpoint), &v1alpha1.VpcEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpoint), err
}

// Delete takes name of the vpcEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeVpcEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcendpointsResource, c.ns, name), &v1alpha1.VpcEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcendpointsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcEndpointList{})
	return err
}

// Patch applies the patch and returns the patched vpcEndpoint.
func (c *FakeVpcEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpoint), err
}
