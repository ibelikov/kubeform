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

// FakeEc2Fleets implements Ec2FleetInterface
type FakeEc2Fleets struct {
	Fake *FakeAwsV1alpha1
}

var ec2fleetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ec2fleets"}

var ec2fleetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Ec2Fleet"}

// Get takes name of the ec2Fleet, and returns the corresponding ec2Fleet object, and an error if there is any.
func (c *FakeEc2Fleets) Get(name string, options v1.GetOptions) (result *v1alpha1.Ec2Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ec2fleetsResource, name), &v1alpha1.Ec2Fleet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2Fleet), err
}

// List takes label and field selectors, and returns the list of Ec2Fleets that match those selectors.
func (c *FakeEc2Fleets) List(opts v1.ListOptions) (result *v1alpha1.Ec2FleetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ec2fleetsResource, ec2fleetsKind, opts), &v1alpha1.Ec2FleetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Ec2FleetList{ListMeta: obj.(*v1alpha1.Ec2FleetList).ListMeta}
	for _, item := range obj.(*v1alpha1.Ec2FleetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ec2Fleets.
func (c *FakeEc2Fleets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ec2fleetsResource, opts))
}

// Create takes the representation of a ec2Fleet and creates it.  Returns the server's representation of the ec2Fleet, and an error, if there is any.
func (c *FakeEc2Fleets) Create(ec2Fleet *v1alpha1.Ec2Fleet) (result *v1alpha1.Ec2Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ec2fleetsResource, ec2Fleet), &v1alpha1.Ec2Fleet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2Fleet), err
}

// Update takes the representation of a ec2Fleet and updates it. Returns the server's representation of the ec2Fleet, and an error, if there is any.
func (c *FakeEc2Fleets) Update(ec2Fleet *v1alpha1.Ec2Fleet) (result *v1alpha1.Ec2Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ec2fleetsResource, ec2Fleet), &v1alpha1.Ec2Fleet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2Fleet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEc2Fleets) UpdateStatus(ec2Fleet *v1alpha1.Ec2Fleet) (*v1alpha1.Ec2Fleet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ec2fleetsResource, "status", ec2Fleet), &v1alpha1.Ec2Fleet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2Fleet), err
}

// Delete takes name of the ec2Fleet and deletes it. Returns an error if one occurs.
func (c *FakeEc2Fleets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ec2fleetsResource, name), &v1alpha1.Ec2Fleet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEc2Fleets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ec2fleetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.Ec2FleetList{})
	return err
}

// Patch applies the patch and returns the patched ec2Fleet.
func (c *FakeEc2Fleets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ec2fleetsResource, name, pt, data, subresources...), &v1alpha1.Ec2Fleet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2Fleet), err
}
