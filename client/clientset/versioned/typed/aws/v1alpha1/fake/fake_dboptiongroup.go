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

// FakeDbOptionGroups implements DbOptionGroupInterface
type FakeDbOptionGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dboptiongroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dboptiongroups"}

var dboptiongroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DbOptionGroup"}

// Get takes name of the dbOptionGroup, and returns the corresponding dbOptionGroup object, and an error if there is any.
func (c *FakeDbOptionGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dboptiongroupsResource, c.ns, name), &v1alpha1.DbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbOptionGroup), err
}

// List takes label and field selectors, and returns the list of DbOptionGroups that match those selectors.
func (c *FakeDbOptionGroups) List(opts v1.ListOptions) (result *v1alpha1.DbOptionGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dboptiongroupsResource, dboptiongroupsKind, c.ns, opts), &v1alpha1.DbOptionGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DbOptionGroupList{ListMeta: obj.(*v1alpha1.DbOptionGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DbOptionGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dbOptionGroups.
func (c *FakeDbOptionGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dboptiongroupsResource, c.ns, opts))

}

// Create takes the representation of a dbOptionGroup and creates it.  Returns the server's representation of the dbOptionGroup, and an error, if there is any.
func (c *FakeDbOptionGroups) Create(dbOptionGroup *v1alpha1.DbOptionGroup) (result *v1alpha1.DbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dboptiongroupsResource, c.ns, dbOptionGroup), &v1alpha1.DbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbOptionGroup), err
}

// Update takes the representation of a dbOptionGroup and updates it. Returns the server's representation of the dbOptionGroup, and an error, if there is any.
func (c *FakeDbOptionGroups) Update(dbOptionGroup *v1alpha1.DbOptionGroup) (result *v1alpha1.DbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dboptiongroupsResource, c.ns, dbOptionGroup), &v1alpha1.DbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbOptionGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDbOptionGroups) UpdateStatus(dbOptionGroup *v1alpha1.DbOptionGroup) (*v1alpha1.DbOptionGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dboptiongroupsResource, "status", c.ns, dbOptionGroup), &v1alpha1.DbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbOptionGroup), err
}

// Delete takes name of the dbOptionGroup and deletes it. Returns an error if one occurs.
func (c *FakeDbOptionGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dboptiongroupsResource, c.ns, name), &v1alpha1.DbOptionGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDbOptionGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dboptiongroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DbOptionGroupList{})
	return err
}

// Patch applies the patch and returns the patched dbOptionGroup.
func (c *FakeDbOptionGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbOptionGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dboptiongroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DbOptionGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbOptionGroup), err
}
