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

// FakeEc2TransitGateways implements Ec2TransitGatewayInterface
type FakeEc2TransitGateways struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ec2transitgatewaysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ec2transitgateways"}

var ec2transitgatewaysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Ec2TransitGateway"}

// Get takes name of the ec2TransitGateway, and returns the corresponding ec2TransitGateway object, and an error if there is any.
func (c *FakeEc2TransitGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ec2TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ec2transitgatewaysResource, c.ns, name), &v1alpha1.Ec2TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGateway), err
}

// List takes label and field selectors, and returns the list of Ec2TransitGateways that match those selectors.
func (c *FakeEc2TransitGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Ec2TransitGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ec2transitgatewaysResource, ec2transitgatewaysKind, c.ns, opts), &v1alpha1.Ec2TransitGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Ec2TransitGatewayList{ListMeta: obj.(*v1alpha1.Ec2TransitGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.Ec2TransitGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ec2TransitGateways.
func (c *FakeEc2TransitGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ec2transitgatewaysResource, c.ns, opts))

}

// Create takes the representation of a ec2TransitGateway and creates it.  Returns the server's representation of the ec2TransitGateway, and an error, if there is any.
func (c *FakeEc2TransitGateways) Create(ctx context.Context, ec2TransitGateway *v1alpha1.Ec2TransitGateway, opts v1.CreateOptions) (result *v1alpha1.Ec2TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ec2transitgatewaysResource, c.ns, ec2TransitGateway), &v1alpha1.Ec2TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGateway), err
}

// Update takes the representation of a ec2TransitGateway and updates it. Returns the server's representation of the ec2TransitGateway, and an error, if there is any.
func (c *FakeEc2TransitGateways) Update(ctx context.Context, ec2TransitGateway *v1alpha1.Ec2TransitGateway, opts v1.UpdateOptions) (result *v1alpha1.Ec2TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ec2transitgatewaysResource, c.ns, ec2TransitGateway), &v1alpha1.Ec2TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEc2TransitGateways) UpdateStatus(ctx context.Context, ec2TransitGateway *v1alpha1.Ec2TransitGateway, opts v1.UpdateOptions) (*v1alpha1.Ec2TransitGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ec2transitgatewaysResource, "status", c.ns, ec2TransitGateway), &v1alpha1.Ec2TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGateway), err
}

// Delete takes name of the ec2TransitGateway and deletes it. Returns an error if one occurs.
func (c *FakeEc2TransitGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ec2transitgatewaysResource, c.ns, name), &v1alpha1.Ec2TransitGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEc2TransitGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ec2transitgatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.Ec2TransitGatewayList{})
	return err
}

// Patch applies the patch and returns the patched ec2TransitGateway.
func (c *FakeEc2TransitGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ec2TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ec2transitgatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ec2TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGateway), err
}
