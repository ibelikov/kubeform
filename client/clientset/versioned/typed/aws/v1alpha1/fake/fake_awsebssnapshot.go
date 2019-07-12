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

// FakeAwsEbsSnapshots implements AwsEbsSnapshotInterface
type FakeAwsEbsSnapshots struct {
	Fake *FakeAwsV1alpha1
}

var awsebssnapshotsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsebssnapshots"}

var awsebssnapshotsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsEbsSnapshot"}

// Get takes name of the awsEbsSnapshot, and returns the corresponding awsEbsSnapshot object, and an error if there is any.
func (c *FakeAwsEbsSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEbsSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsebssnapshotsResource, name), &v1alpha1.AwsEbsSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsSnapshot), err
}

// List takes label and field selectors, and returns the list of AwsEbsSnapshots that match those selectors.
func (c *FakeAwsEbsSnapshots) List(opts v1.ListOptions) (result *v1alpha1.AwsEbsSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsebssnapshotsResource, awsebssnapshotsKind, opts), &v1alpha1.AwsEbsSnapshotList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEbsSnapshotList{ListMeta: obj.(*v1alpha1.AwsEbsSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsEbsSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEbsSnapshots.
func (c *FakeAwsEbsSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsebssnapshotsResource, opts))
}

// Create takes the representation of a awsEbsSnapshot and creates it.  Returns the server's representation of the awsEbsSnapshot, and an error, if there is any.
func (c *FakeAwsEbsSnapshots) Create(awsEbsSnapshot *v1alpha1.AwsEbsSnapshot) (result *v1alpha1.AwsEbsSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsebssnapshotsResource, awsEbsSnapshot), &v1alpha1.AwsEbsSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsSnapshot), err
}

// Update takes the representation of a awsEbsSnapshot and updates it. Returns the server's representation of the awsEbsSnapshot, and an error, if there is any.
func (c *FakeAwsEbsSnapshots) Update(awsEbsSnapshot *v1alpha1.AwsEbsSnapshot) (result *v1alpha1.AwsEbsSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsebssnapshotsResource, awsEbsSnapshot), &v1alpha1.AwsEbsSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsEbsSnapshots) UpdateStatus(awsEbsSnapshot *v1alpha1.AwsEbsSnapshot) (*v1alpha1.AwsEbsSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsebssnapshotsResource, "status", awsEbsSnapshot), &v1alpha1.AwsEbsSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsSnapshot), err
}

// Delete takes name of the awsEbsSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeAwsEbsSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsebssnapshotsResource, name), &v1alpha1.AwsEbsSnapshot{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEbsSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsebssnapshotsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEbsSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched awsEbsSnapshot.
func (c *FakeAwsEbsSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEbsSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsebssnapshotsResource, name, pt, data, subresources...), &v1alpha1.AwsEbsSnapshot{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEbsSnapshot), err
}
