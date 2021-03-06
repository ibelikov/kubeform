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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeOpsworksUserProfiles implements OpsworksUserProfileInterface
type FakeOpsworksUserProfiles struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworksuserprofilesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksuserprofiles"}

var opsworksuserprofilesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksUserProfile"}

// Get takes name of the opsworksUserProfile, and returns the corresponding opsworksUserProfile object, and an error if there is any.
func (c *FakeOpsworksUserProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworksuserprofilesResource, c.ns, name), &v1alpha1.OpsworksUserProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksUserProfile), err
}

// List takes label and field selectors, and returns the list of OpsworksUserProfiles that match those selectors.
func (c *FakeOpsworksUserProfiles) List(opts v1.ListOptions) (result *v1alpha1.OpsworksUserProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworksuserprofilesResource, opsworksuserprofilesKind, c.ns, opts), &v1alpha1.OpsworksUserProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksUserProfileList{ListMeta: obj.(*v1alpha1.OpsworksUserProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksUserProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksUserProfiles.
func (c *FakeOpsworksUserProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworksuserprofilesResource, c.ns, opts))

}

// Create takes the representation of a opsworksUserProfile and creates it.  Returns the server's representation of the opsworksUserProfile, and an error, if there is any.
func (c *FakeOpsworksUserProfiles) Create(opsworksUserProfile *v1alpha1.OpsworksUserProfile) (result *v1alpha1.OpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworksuserprofilesResource, c.ns, opsworksUserProfile), &v1alpha1.OpsworksUserProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksUserProfile), err
}

// Update takes the representation of a opsworksUserProfile and updates it. Returns the server's representation of the opsworksUserProfile, and an error, if there is any.
func (c *FakeOpsworksUserProfiles) Update(opsworksUserProfile *v1alpha1.OpsworksUserProfile) (result *v1alpha1.OpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworksuserprofilesResource, c.ns, opsworksUserProfile), &v1alpha1.OpsworksUserProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksUserProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksUserProfiles) UpdateStatus(opsworksUserProfile *v1alpha1.OpsworksUserProfile) (*v1alpha1.OpsworksUserProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworksuserprofilesResource, "status", c.ns, opsworksUserProfile), &v1alpha1.OpsworksUserProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksUserProfile), err
}

// Delete takes name of the opsworksUserProfile and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksUserProfiles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworksuserprofilesResource, c.ns, name), &v1alpha1.OpsworksUserProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksUserProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworksuserprofilesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksUserProfileList{})
	return err
}

// Patch applies the patch and returns the patched opsworksUserProfile.
func (c *FakeOpsworksUserProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworksuserprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksUserProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksUserProfile), err
}
