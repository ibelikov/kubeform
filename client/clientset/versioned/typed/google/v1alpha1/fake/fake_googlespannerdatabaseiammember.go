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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleSpannerDatabaseIamMembers implements GoogleSpannerDatabaseIamMemberInterface
type FakeGoogleSpannerDatabaseIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var googlespannerdatabaseiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlespannerdatabaseiammembers"}

var googlespannerdatabaseiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleSpannerDatabaseIamMember"}

// Get takes name of the googleSpannerDatabaseIamMember, and returns the corresponding googleSpannerDatabaseIamMember object, and an error if there is any.
func (c *FakeGoogleSpannerDatabaseIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlespannerdatabaseiammembersResource, name), &v1alpha1.GoogleSpannerDatabaseIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamMember), err
}

// List takes label and field selectors, and returns the list of GoogleSpannerDatabaseIamMembers that match those selectors.
func (c *FakeGoogleSpannerDatabaseIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleSpannerDatabaseIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlespannerdatabaseiammembersResource, googlespannerdatabaseiammembersKind, opts), &v1alpha1.GoogleSpannerDatabaseIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleSpannerDatabaseIamMemberList{ListMeta: obj.(*v1alpha1.GoogleSpannerDatabaseIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleSpannerDatabaseIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleSpannerDatabaseIamMembers.
func (c *FakeGoogleSpannerDatabaseIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlespannerdatabaseiammembersResource, opts))
}

// Create takes the representation of a googleSpannerDatabaseIamMember and creates it.  Returns the server's representation of the googleSpannerDatabaseIamMember, and an error, if there is any.
func (c *FakeGoogleSpannerDatabaseIamMembers) Create(googleSpannerDatabaseIamMember *v1alpha1.GoogleSpannerDatabaseIamMember) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlespannerdatabaseiammembersResource, googleSpannerDatabaseIamMember), &v1alpha1.GoogleSpannerDatabaseIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamMember), err
}

// Update takes the representation of a googleSpannerDatabaseIamMember and updates it. Returns the server's representation of the googleSpannerDatabaseIamMember, and an error, if there is any.
func (c *FakeGoogleSpannerDatabaseIamMembers) Update(googleSpannerDatabaseIamMember *v1alpha1.GoogleSpannerDatabaseIamMember) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlespannerdatabaseiammembersResource, googleSpannerDatabaseIamMember), &v1alpha1.GoogleSpannerDatabaseIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleSpannerDatabaseIamMembers) UpdateStatus(googleSpannerDatabaseIamMember *v1alpha1.GoogleSpannerDatabaseIamMember) (*v1alpha1.GoogleSpannerDatabaseIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlespannerdatabaseiammembersResource, "status", googleSpannerDatabaseIamMember), &v1alpha1.GoogleSpannerDatabaseIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamMember), err
}

// Delete takes name of the googleSpannerDatabaseIamMember and deletes it. Returns an error if one occurs.
func (c *FakeGoogleSpannerDatabaseIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlespannerdatabaseiammembersResource, name), &v1alpha1.GoogleSpannerDatabaseIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleSpannerDatabaseIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlespannerdatabaseiammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleSpannerDatabaseIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched googleSpannerDatabaseIamMember.
func (c *FakeGoogleSpannerDatabaseIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlespannerdatabaseiammembersResource, name, pt, data, subresources...), &v1alpha1.GoogleSpannerDatabaseIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamMember), err
}
