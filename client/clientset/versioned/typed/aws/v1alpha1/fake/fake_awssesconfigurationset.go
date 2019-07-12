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

// FakeAwsSesConfigurationSets implements AwsSesConfigurationSetInterface
type FakeAwsSesConfigurationSets struct {
	Fake *FakeAwsV1alpha1
}

var awssesconfigurationsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awssesconfigurationsets"}

var awssesconfigurationsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSesConfigurationSet"}

// Get takes name of the awsSesConfigurationSet, and returns the corresponding awsSesConfigurationSet object, and an error if there is any.
func (c *FakeAwsSesConfigurationSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesConfigurationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awssesconfigurationsetsResource, name), &v1alpha1.AwsSesConfigurationSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesConfigurationSet), err
}

// List takes label and field selectors, and returns the list of AwsSesConfigurationSets that match those selectors.
func (c *FakeAwsSesConfigurationSets) List(opts v1.ListOptions) (result *v1alpha1.AwsSesConfigurationSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awssesconfigurationsetsResource, awssesconfigurationsetsKind, opts), &v1alpha1.AwsSesConfigurationSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSesConfigurationSetList{ListMeta: obj.(*v1alpha1.AwsSesConfigurationSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSesConfigurationSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSesConfigurationSets.
func (c *FakeAwsSesConfigurationSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awssesconfigurationsetsResource, opts))
}

// Create takes the representation of a awsSesConfigurationSet and creates it.  Returns the server's representation of the awsSesConfigurationSet, and an error, if there is any.
func (c *FakeAwsSesConfigurationSets) Create(awsSesConfigurationSet *v1alpha1.AwsSesConfigurationSet) (result *v1alpha1.AwsSesConfigurationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awssesconfigurationsetsResource, awsSesConfigurationSet), &v1alpha1.AwsSesConfigurationSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesConfigurationSet), err
}

// Update takes the representation of a awsSesConfigurationSet and updates it. Returns the server's representation of the awsSesConfigurationSet, and an error, if there is any.
func (c *FakeAwsSesConfigurationSets) Update(awsSesConfigurationSet *v1alpha1.AwsSesConfigurationSet) (result *v1alpha1.AwsSesConfigurationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awssesconfigurationsetsResource, awsSesConfigurationSet), &v1alpha1.AwsSesConfigurationSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesConfigurationSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSesConfigurationSets) UpdateStatus(awsSesConfigurationSet *v1alpha1.AwsSesConfigurationSet) (*v1alpha1.AwsSesConfigurationSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awssesconfigurationsetsResource, "status", awsSesConfigurationSet), &v1alpha1.AwsSesConfigurationSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesConfigurationSet), err
}

// Delete takes name of the awsSesConfigurationSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsSesConfigurationSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awssesconfigurationsetsResource, name), &v1alpha1.AwsSesConfigurationSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSesConfigurationSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awssesconfigurationsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSesConfigurationSetList{})
	return err
}

// Patch applies the patch and returns the patched awsSesConfigurationSet.
func (c *FakeAwsSesConfigurationSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesConfigurationSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awssesconfigurationsetsResource, name, pt, data, subresources...), &v1alpha1.AwsSesConfigurationSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesConfigurationSet), err
}
