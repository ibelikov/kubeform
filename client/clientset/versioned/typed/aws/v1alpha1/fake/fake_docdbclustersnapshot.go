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

// FakeDocdbClusterSnapshots implements DocdbClusterSnapshotInterface
type FakeDocdbClusterSnapshots struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var docdbclustersnapshotsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "docdbclustersnapshots"}

var docdbclustersnapshotsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DocdbClusterSnapshot"}

// Get takes name of the docdbClusterSnapshot, and returns the corresponding docdbClusterSnapshot object, and an error if there is any.
func (c *FakeDocdbClusterSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.DocdbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(docdbclustersnapshotsResource, c.ns, name), &v1alpha1.DocdbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterSnapshot), err
}

// List takes label and field selectors, and returns the list of DocdbClusterSnapshots that match those selectors.
func (c *FakeDocdbClusterSnapshots) List(opts v1.ListOptions) (result *v1alpha1.DocdbClusterSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(docdbclustersnapshotsResource, docdbclustersnapshotsKind, c.ns, opts), &v1alpha1.DocdbClusterSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DocdbClusterSnapshotList{ListMeta: obj.(*v1alpha1.DocdbClusterSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.DocdbClusterSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested docdbClusterSnapshots.
func (c *FakeDocdbClusterSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(docdbclustersnapshotsResource, c.ns, opts))

}

// Create takes the representation of a docdbClusterSnapshot and creates it.  Returns the server's representation of the docdbClusterSnapshot, and an error, if there is any.
func (c *FakeDocdbClusterSnapshots) Create(docdbClusterSnapshot *v1alpha1.DocdbClusterSnapshot) (result *v1alpha1.DocdbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(docdbclustersnapshotsResource, c.ns, docdbClusterSnapshot), &v1alpha1.DocdbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterSnapshot), err
}

// Update takes the representation of a docdbClusterSnapshot and updates it. Returns the server's representation of the docdbClusterSnapshot, and an error, if there is any.
func (c *FakeDocdbClusterSnapshots) Update(docdbClusterSnapshot *v1alpha1.DocdbClusterSnapshot) (result *v1alpha1.DocdbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(docdbclustersnapshotsResource, c.ns, docdbClusterSnapshot), &v1alpha1.DocdbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDocdbClusterSnapshots) UpdateStatus(docdbClusterSnapshot *v1alpha1.DocdbClusterSnapshot) (*v1alpha1.DocdbClusterSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(docdbclustersnapshotsResource, "status", c.ns, docdbClusterSnapshot), &v1alpha1.DocdbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterSnapshot), err
}

// Delete takes name of the docdbClusterSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeDocdbClusterSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(docdbclustersnapshotsResource, c.ns, name), &v1alpha1.DocdbClusterSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDocdbClusterSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(docdbclustersnapshotsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DocdbClusterSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched docdbClusterSnapshot.
func (c *FakeDocdbClusterSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DocdbClusterSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(docdbclustersnapshotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DocdbClusterSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterSnapshot), err
}
