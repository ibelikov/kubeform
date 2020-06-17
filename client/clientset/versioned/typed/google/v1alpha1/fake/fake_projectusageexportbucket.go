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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeProjectUsageExportBuckets implements ProjectUsageExportBucketInterface
type FakeProjectUsageExportBuckets struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var projectusageexportbucketsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "projectusageexportbuckets"}

var projectusageexportbucketsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ProjectUsageExportBucket"}

// Get takes name of the projectUsageExportBucket, and returns the corresponding projectUsageExportBucket object, and an error if there is any.
func (c *FakeProjectUsageExportBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectusageexportbucketsResource, c.ns, name), &v1alpha1.ProjectUsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectUsageExportBucket), err
}

// List takes label and field selectors, and returns the list of ProjectUsageExportBuckets that match those selectors.
func (c *FakeProjectUsageExportBuckets) List(opts v1.ListOptions) (result *v1alpha1.ProjectUsageExportBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectusageexportbucketsResource, projectusageexportbucketsKind, c.ns, opts), &v1alpha1.ProjectUsageExportBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProjectUsageExportBucketList{ListMeta: obj.(*v1alpha1.ProjectUsageExportBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProjectUsageExportBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectUsageExportBuckets.
func (c *FakeProjectUsageExportBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectusageexportbucketsResource, c.ns, opts))

}

// Create takes the representation of a projectUsageExportBucket and creates it.  Returns the server's representation of the projectUsageExportBucket, and an error, if there is any.
func (c *FakeProjectUsageExportBuckets) Create(projectUsageExportBucket *v1alpha1.ProjectUsageExportBucket) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectusageexportbucketsResource, c.ns, projectUsageExportBucket), &v1alpha1.ProjectUsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectUsageExportBucket), err
}

// Update takes the representation of a projectUsageExportBucket and updates it. Returns the server's representation of the projectUsageExportBucket, and an error, if there is any.
func (c *FakeProjectUsageExportBuckets) Update(projectUsageExportBucket *v1alpha1.ProjectUsageExportBucket) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectusageexportbucketsResource, c.ns, projectUsageExportBucket), &v1alpha1.ProjectUsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectUsageExportBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectUsageExportBuckets) UpdateStatus(projectUsageExportBucket *v1alpha1.ProjectUsageExportBucket) (*v1alpha1.ProjectUsageExportBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(projectusageexportbucketsResource, "status", c.ns, projectUsageExportBucket), &v1alpha1.ProjectUsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectUsageExportBucket), err
}

// Delete takes name of the projectUsageExportBucket and deletes it. Returns an error if one occurs.
func (c *FakeProjectUsageExportBuckets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(projectusageexportbucketsResource, c.ns, name), &v1alpha1.ProjectUsageExportBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectUsageExportBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectusageexportbucketsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProjectUsageExportBucketList{})
	return err
}

// Patch applies the patch and returns the patched projectUsageExportBucket.
func (c *FakeProjectUsageExportBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectusageexportbucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProjectUsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectUsageExportBucket), err
}
