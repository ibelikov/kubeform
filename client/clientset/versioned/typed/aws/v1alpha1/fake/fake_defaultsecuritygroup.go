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

// FakeDefaultSecurityGroups implements DefaultSecurityGroupInterface
type FakeDefaultSecurityGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var defaultsecuritygroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "defaultsecuritygroups"}

var defaultsecuritygroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DefaultSecurityGroup"}

// Get takes name of the defaultSecurityGroup, and returns the corresponding defaultSecurityGroup object, and an error if there is any.
func (c *FakeDefaultSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DefaultSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(defaultsecuritygroupsResource, c.ns, name), &v1alpha1.DefaultSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultSecurityGroup), err
}

// List takes label and field selectors, and returns the list of DefaultSecurityGroups that match those selectors.
func (c *FakeDefaultSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.DefaultSecurityGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(defaultsecuritygroupsResource, defaultsecuritygroupsKind, c.ns, opts), &v1alpha1.DefaultSecurityGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DefaultSecurityGroupList{ListMeta: obj.(*v1alpha1.DefaultSecurityGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DefaultSecurityGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested defaultSecurityGroups.
func (c *FakeDefaultSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(defaultsecuritygroupsResource, c.ns, opts))

}

// Create takes the representation of a defaultSecurityGroup and creates it.  Returns the server's representation of the defaultSecurityGroup, and an error, if there is any.
func (c *FakeDefaultSecurityGroups) Create(defaultSecurityGroup *v1alpha1.DefaultSecurityGroup) (result *v1alpha1.DefaultSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(defaultsecuritygroupsResource, c.ns, defaultSecurityGroup), &v1alpha1.DefaultSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultSecurityGroup), err
}

// Update takes the representation of a defaultSecurityGroup and updates it. Returns the server's representation of the defaultSecurityGroup, and an error, if there is any.
func (c *FakeDefaultSecurityGroups) Update(defaultSecurityGroup *v1alpha1.DefaultSecurityGroup) (result *v1alpha1.DefaultSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(defaultsecuritygroupsResource, c.ns, defaultSecurityGroup), &v1alpha1.DefaultSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultSecurityGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDefaultSecurityGroups) UpdateStatus(defaultSecurityGroup *v1alpha1.DefaultSecurityGroup) (*v1alpha1.DefaultSecurityGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(defaultsecuritygroupsResource, "status", c.ns, defaultSecurityGroup), &v1alpha1.DefaultSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultSecurityGroup), err
}

// Delete takes name of the defaultSecurityGroup and deletes it. Returns an error if one occurs.
func (c *FakeDefaultSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(defaultsecuritygroupsResource, c.ns, name), &v1alpha1.DefaultSecurityGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDefaultSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(defaultsecuritygroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DefaultSecurityGroupList{})
	return err
}

// Patch applies the patch and returns the patched defaultSecurityGroup.
func (c *FakeDefaultSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(defaultsecuritygroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DefaultSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultSecurityGroup), err
}
