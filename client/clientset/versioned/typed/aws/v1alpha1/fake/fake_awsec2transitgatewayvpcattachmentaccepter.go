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

// FakeAwsEc2TransitGatewayVpcAttachmentAccepters implements AwsEc2TransitGatewayVpcAttachmentAccepterInterface
type FakeAwsEc2TransitGatewayVpcAttachmentAccepters struct {
	Fake *FakeAwsV1alpha1
}

var awsec2transitgatewayvpcattachmentacceptersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsec2transitgatewayvpcattachmentaccepters"}

var awsec2transitgatewayvpcattachmentacceptersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsEc2TransitGatewayVpcAttachmentAccepter"}

// Get takes name of the awsEc2TransitGatewayVpcAttachmentAccepter, and returns the corresponding awsEc2TransitGatewayVpcAttachmentAccepter object, and an error if there is any.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsec2transitgatewayvpcattachmentacceptersResource, name), &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter), err
}

// List takes label and field selectors, and returns the list of AwsEc2TransitGatewayVpcAttachmentAccepters that match those selectors.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) List(opts v1.ListOptions) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsec2transitgatewayvpcattachmentacceptersResource, awsec2transitgatewayvpcattachmentacceptersKind, opts), &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepterList{ListMeta: obj.(*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEc2TransitGatewayVpcAttachmentAccepters.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsec2transitgatewayvpcattachmentacceptersResource, opts))
}

// Create takes the representation of a awsEc2TransitGatewayVpcAttachmentAccepter and creates it.  Returns the server's representation of the awsEc2TransitGatewayVpcAttachmentAccepter, and an error, if there is any.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) Create(awsEc2TransitGatewayVpcAttachmentAccepter *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsec2transitgatewayvpcattachmentacceptersResource, awsEc2TransitGatewayVpcAttachmentAccepter), &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter), err
}

// Update takes the representation of a awsEc2TransitGatewayVpcAttachmentAccepter and updates it. Returns the server's representation of the awsEc2TransitGatewayVpcAttachmentAccepter, and an error, if there is any.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) Update(awsEc2TransitGatewayVpcAttachmentAccepter *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsec2transitgatewayvpcattachmentacceptersResource, awsEc2TransitGatewayVpcAttachmentAccepter), &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) UpdateStatus(awsEc2TransitGatewayVpcAttachmentAccepter *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter) (*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsec2transitgatewayvpcattachmentacceptersResource, "status", awsEc2TransitGatewayVpcAttachmentAccepter), &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter), err
}

// Delete takes name of the awsEc2TransitGatewayVpcAttachmentAccepter and deletes it. Returns an error if one occurs.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsec2transitgatewayvpcattachmentacceptersResource, name), &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsec2transitgatewayvpcattachmentacceptersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepterList{})
	return err
}

// Patch applies the patch and returns the patched awsEc2TransitGatewayVpcAttachmentAccepter.
func (c *FakeAwsEc2TransitGatewayVpcAttachmentAccepters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsec2transitgatewayvpcattachmentacceptersResource, name, pt, data, subresources...), &v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEc2TransitGatewayVpcAttachmentAccepter), err
}