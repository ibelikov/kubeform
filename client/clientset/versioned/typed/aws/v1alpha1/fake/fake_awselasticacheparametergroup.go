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

// FakeAwsElasticacheParameterGroups implements AwsElasticacheParameterGroupInterface
type FakeAwsElasticacheParameterGroups struct {
	Fake *FakeAwsV1alpha1
}

var awselasticacheparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awselasticacheparametergroups"}

var awselasticacheparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsElasticacheParameterGroup"}

// Get takes name of the awsElasticacheParameterGroup, and returns the corresponding awsElasticacheParameterGroup object, and an error if there is any.
func (c *FakeAwsElasticacheParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awselasticacheparametergroupsResource, name), &v1alpha1.AwsElasticacheParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheParameterGroup), err
}

// List takes label and field selectors, and returns the list of AwsElasticacheParameterGroups that match those selectors.
func (c *FakeAwsElasticacheParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticacheParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awselasticacheparametergroupsResource, awselasticacheparametergroupsKind, opts), &v1alpha1.AwsElasticacheParameterGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElasticacheParameterGroupList{ListMeta: obj.(*v1alpha1.AwsElasticacheParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsElasticacheParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticacheParameterGroups.
func (c *FakeAwsElasticacheParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awselasticacheparametergroupsResource, opts))
}

// Create takes the representation of a awsElasticacheParameterGroup and creates it.  Returns the server's representation of the awsElasticacheParameterGroup, and an error, if there is any.
func (c *FakeAwsElasticacheParameterGroups) Create(awsElasticacheParameterGroup *v1alpha1.AwsElasticacheParameterGroup) (result *v1alpha1.AwsElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awselasticacheparametergroupsResource, awsElasticacheParameterGroup), &v1alpha1.AwsElasticacheParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheParameterGroup), err
}

// Update takes the representation of a awsElasticacheParameterGroup and updates it. Returns the server's representation of the awsElasticacheParameterGroup, and an error, if there is any.
func (c *FakeAwsElasticacheParameterGroups) Update(awsElasticacheParameterGroup *v1alpha1.AwsElasticacheParameterGroup) (result *v1alpha1.AwsElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awselasticacheparametergroupsResource, awsElasticacheParameterGroup), &v1alpha1.AwsElasticacheParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsElasticacheParameterGroups) UpdateStatus(awsElasticacheParameterGroup *v1alpha1.AwsElasticacheParameterGroup) (*v1alpha1.AwsElasticacheParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awselasticacheparametergroupsResource, "status", awsElasticacheParameterGroup), &v1alpha1.AwsElasticacheParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheParameterGroup), err
}

// Delete takes name of the awsElasticacheParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticacheParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awselasticacheparametergroupsResource, name), &v1alpha1.AwsElasticacheParameterGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticacheParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awselasticacheparametergroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElasticacheParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticacheParameterGroup.
func (c *FakeAwsElasticacheParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awselasticacheparametergroupsResource, name, pt, data, subresources...), &v1alpha1.AwsElasticacheParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticacheParameterGroup), err
}
