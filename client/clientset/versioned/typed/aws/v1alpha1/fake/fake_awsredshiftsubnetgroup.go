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

// FakeAwsRedshiftSubnetGroups implements AwsRedshiftSubnetGroupInterface
type FakeAwsRedshiftSubnetGroups struct {
	Fake *FakeAwsV1alpha1
}

var awsredshiftsubnetgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsredshiftsubnetgroups"}

var awsredshiftsubnetgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsRedshiftSubnetGroup"}

// Get takes name of the awsRedshiftSubnetGroup, and returns the corresponding awsRedshiftSubnetGroup object, and an error if there is any.
func (c *FakeAwsRedshiftSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsredshiftsubnetgroupsResource, name), &v1alpha1.AwsRedshiftSubnetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftSubnetGroup), err
}

// List takes label and field selectors, and returns the list of AwsRedshiftSubnetGroups that match those selectors.
func (c *FakeAwsRedshiftSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsRedshiftSubnetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsredshiftsubnetgroupsResource, awsredshiftsubnetgroupsKind, opts), &v1alpha1.AwsRedshiftSubnetGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsRedshiftSubnetGroupList{ListMeta: obj.(*v1alpha1.AwsRedshiftSubnetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsRedshiftSubnetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRedshiftSubnetGroups.
func (c *FakeAwsRedshiftSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsredshiftsubnetgroupsResource, opts))
}

// Create takes the representation of a awsRedshiftSubnetGroup and creates it.  Returns the server's representation of the awsRedshiftSubnetGroup, and an error, if there is any.
func (c *FakeAwsRedshiftSubnetGroups) Create(awsRedshiftSubnetGroup *v1alpha1.AwsRedshiftSubnetGroup) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsredshiftsubnetgroupsResource, awsRedshiftSubnetGroup), &v1alpha1.AwsRedshiftSubnetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftSubnetGroup), err
}

// Update takes the representation of a awsRedshiftSubnetGroup and updates it. Returns the server's representation of the awsRedshiftSubnetGroup, and an error, if there is any.
func (c *FakeAwsRedshiftSubnetGroups) Update(awsRedshiftSubnetGroup *v1alpha1.AwsRedshiftSubnetGroup) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsredshiftsubnetgroupsResource, awsRedshiftSubnetGroup), &v1alpha1.AwsRedshiftSubnetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftSubnetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsRedshiftSubnetGroups) UpdateStatus(awsRedshiftSubnetGroup *v1alpha1.AwsRedshiftSubnetGroup) (*v1alpha1.AwsRedshiftSubnetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsredshiftsubnetgroupsResource, "status", awsRedshiftSubnetGroup), &v1alpha1.AwsRedshiftSubnetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftSubnetGroup), err
}

// Delete takes name of the awsRedshiftSubnetGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsRedshiftSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsredshiftsubnetgroupsResource, name), &v1alpha1.AwsRedshiftSubnetGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRedshiftSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsredshiftsubnetgroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsRedshiftSubnetGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsRedshiftSubnetGroup.
func (c *FakeAwsRedshiftSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRedshiftSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsredshiftsubnetgroupsResource, name, pt, data, subresources...), &v1alpha1.AwsRedshiftSubnetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftSubnetGroup), err
}
