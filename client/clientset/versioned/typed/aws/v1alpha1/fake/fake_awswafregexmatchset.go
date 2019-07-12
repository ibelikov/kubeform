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

// FakeAwsWafRegexMatchSets implements AwsWafRegexMatchSetInterface
type FakeAwsWafRegexMatchSets struct {
	Fake *FakeAwsV1alpha1
}

var awswafregexmatchsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awswafregexmatchsets"}

var awswafregexmatchsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsWafRegexMatchSet"}

// Get takes name of the awsWafRegexMatchSet, and returns the corresponding awsWafRegexMatchSet object, and an error if there is any.
func (c *FakeAwsWafRegexMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awswafregexmatchsetsResource, name), &v1alpha1.AwsWafRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafRegexMatchSet), err
}

// List takes label and field selectors, and returns the list of AwsWafRegexMatchSets that match those selectors.
func (c *FakeAwsWafRegexMatchSets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafRegexMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awswafregexmatchsetsResource, awswafregexmatchsetsKind, opts), &v1alpha1.AwsWafRegexMatchSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafRegexMatchSetList{ListMeta: obj.(*v1alpha1.AwsWafRegexMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsWafRegexMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafRegexMatchSets.
func (c *FakeAwsWafRegexMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awswafregexmatchsetsResource, opts))
}

// Create takes the representation of a awsWafRegexMatchSet and creates it.  Returns the server's representation of the awsWafRegexMatchSet, and an error, if there is any.
func (c *FakeAwsWafRegexMatchSets) Create(awsWafRegexMatchSet *v1alpha1.AwsWafRegexMatchSet) (result *v1alpha1.AwsWafRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awswafregexmatchsetsResource, awsWafRegexMatchSet), &v1alpha1.AwsWafRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafRegexMatchSet), err
}

// Update takes the representation of a awsWafRegexMatchSet and updates it. Returns the server's representation of the awsWafRegexMatchSet, and an error, if there is any.
func (c *FakeAwsWafRegexMatchSets) Update(awsWafRegexMatchSet *v1alpha1.AwsWafRegexMatchSet) (result *v1alpha1.AwsWafRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awswafregexmatchsetsResource, awsWafRegexMatchSet), &v1alpha1.AwsWafRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafRegexMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsWafRegexMatchSets) UpdateStatus(awsWafRegexMatchSet *v1alpha1.AwsWafRegexMatchSet) (*v1alpha1.AwsWafRegexMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awswafregexmatchsetsResource, "status", awsWafRegexMatchSet), &v1alpha1.AwsWafRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafRegexMatchSet), err
}

// Delete takes name of the awsWafRegexMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafRegexMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awswafregexmatchsetsResource, name), &v1alpha1.AwsWafRegexMatchSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafRegexMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awswafregexmatchsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafRegexMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafRegexMatchSet.
func (c *FakeAwsWafRegexMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafRegexMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awswafregexmatchsetsResource, name, pt, data, subresources...), &v1alpha1.AwsWafRegexMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafRegexMatchSet), err
}
