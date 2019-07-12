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

// FakeAwsCloudwatchEventPermissions implements AwsCloudwatchEventPermissionInterface
type FakeAwsCloudwatchEventPermissions struct {
	Fake *FakeAwsV1alpha1
}

var awscloudwatcheventpermissionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awscloudwatcheventpermissions"}

var awscloudwatcheventpermissionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsCloudwatchEventPermission"}

// Get takes name of the awsCloudwatchEventPermission, and returns the corresponding awsCloudwatchEventPermission object, and an error if there is any.
func (c *FakeAwsCloudwatchEventPermissions) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awscloudwatcheventpermissionsResource, name), &v1alpha1.AwsCloudwatchEventPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventPermission), err
}

// List takes label and field selectors, and returns the list of AwsCloudwatchEventPermissions that match those selectors.
func (c *FakeAwsCloudwatchEventPermissions) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchEventPermissionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awscloudwatcheventpermissionsResource, awscloudwatcheventpermissionsKind, opts), &v1alpha1.AwsCloudwatchEventPermissionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCloudwatchEventPermissionList{ListMeta: obj.(*v1alpha1.AwsCloudwatchEventPermissionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsCloudwatchEventPermissionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchEventPermissions.
func (c *FakeAwsCloudwatchEventPermissions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awscloudwatcheventpermissionsResource, opts))
}

// Create takes the representation of a awsCloudwatchEventPermission and creates it.  Returns the server's representation of the awsCloudwatchEventPermission, and an error, if there is any.
func (c *FakeAwsCloudwatchEventPermissions) Create(awsCloudwatchEventPermission *v1alpha1.AwsCloudwatchEventPermission) (result *v1alpha1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awscloudwatcheventpermissionsResource, awsCloudwatchEventPermission), &v1alpha1.AwsCloudwatchEventPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventPermission), err
}

// Update takes the representation of a awsCloudwatchEventPermission and updates it. Returns the server's representation of the awsCloudwatchEventPermission, and an error, if there is any.
func (c *FakeAwsCloudwatchEventPermissions) Update(awsCloudwatchEventPermission *v1alpha1.AwsCloudwatchEventPermission) (result *v1alpha1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awscloudwatcheventpermissionsResource, awsCloudwatchEventPermission), &v1alpha1.AwsCloudwatchEventPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventPermission), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsCloudwatchEventPermissions) UpdateStatus(awsCloudwatchEventPermission *v1alpha1.AwsCloudwatchEventPermission) (*v1alpha1.AwsCloudwatchEventPermission, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awscloudwatcheventpermissionsResource, "status", awsCloudwatchEventPermission), &v1alpha1.AwsCloudwatchEventPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventPermission), err
}

// Delete takes name of the awsCloudwatchEventPermission and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudwatchEventPermissions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awscloudwatcheventpermissionsResource, name), &v1alpha1.AwsCloudwatchEventPermission{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudwatchEventPermissions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awscloudwatcheventpermissionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCloudwatchEventPermissionList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudwatchEventPermission.
func (c *FakeAwsCloudwatchEventPermissions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchEventPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awscloudwatcheventpermissionsResource, name, pt, data, subresources...), &v1alpha1.AwsCloudwatchEventPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCloudwatchEventPermission), err
}
