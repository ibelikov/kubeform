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

// FakeAthenaWorkgroups implements AthenaWorkgroupInterface
type FakeAthenaWorkgroups struct {
	Fake *FakeAwsV1alpha1
}

var athenaworkgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "athenaworkgroups"}

var athenaworkgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AthenaWorkgroup"}

// Get takes name of the athenaWorkgroup, and returns the corresponding athenaWorkgroup object, and an error if there is any.
func (c *FakeAthenaWorkgroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AthenaWorkgroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(athenaworkgroupsResource, name), &v1alpha1.AthenaWorkgroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaWorkgroup), err
}

// List takes label and field selectors, and returns the list of AthenaWorkgroups that match those selectors.
func (c *FakeAthenaWorkgroups) List(opts v1.ListOptions) (result *v1alpha1.AthenaWorkgroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(athenaworkgroupsResource, athenaworkgroupsKind, opts), &v1alpha1.AthenaWorkgroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AthenaWorkgroupList{ListMeta: obj.(*v1alpha1.AthenaWorkgroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AthenaWorkgroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested athenaWorkgroups.
func (c *FakeAthenaWorkgroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(athenaworkgroupsResource, opts))
}

// Create takes the representation of a athenaWorkgroup and creates it.  Returns the server's representation of the athenaWorkgroup, and an error, if there is any.
func (c *FakeAthenaWorkgroups) Create(athenaWorkgroup *v1alpha1.AthenaWorkgroup) (result *v1alpha1.AthenaWorkgroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(athenaworkgroupsResource, athenaWorkgroup), &v1alpha1.AthenaWorkgroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaWorkgroup), err
}

// Update takes the representation of a athenaWorkgroup and updates it. Returns the server's representation of the athenaWorkgroup, and an error, if there is any.
func (c *FakeAthenaWorkgroups) Update(athenaWorkgroup *v1alpha1.AthenaWorkgroup) (result *v1alpha1.AthenaWorkgroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(athenaworkgroupsResource, athenaWorkgroup), &v1alpha1.AthenaWorkgroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaWorkgroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAthenaWorkgroups) UpdateStatus(athenaWorkgroup *v1alpha1.AthenaWorkgroup) (*v1alpha1.AthenaWorkgroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(athenaworkgroupsResource, "status", athenaWorkgroup), &v1alpha1.AthenaWorkgroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaWorkgroup), err
}

// Delete takes name of the athenaWorkgroup and deletes it. Returns an error if one occurs.
func (c *FakeAthenaWorkgroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(athenaworkgroupsResource, name), &v1alpha1.AthenaWorkgroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAthenaWorkgroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(athenaworkgroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AthenaWorkgroupList{})
	return err
}

// Patch applies the patch and returns the patched athenaWorkgroup.
func (c *FakeAthenaWorkgroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AthenaWorkgroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(athenaworkgroupsResource, name, pt, data, subresources...), &v1alpha1.AthenaWorkgroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaWorkgroup), err
}
