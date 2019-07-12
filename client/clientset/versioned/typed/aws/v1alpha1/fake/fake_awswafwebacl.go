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

// FakeAwsWafWebAcls implements AwsWafWebAclInterface
type FakeAwsWafWebAcls struct {
	Fake *FakeAwsV1alpha1
}

var awswafwebaclsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awswafwebacls"}

var awswafwebaclsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsWafWebAcl"}

// Get takes name of the awsWafWebAcl, and returns the corresponding awsWafWebAcl object, and an error if there is any.
func (c *FakeAwsWafWebAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awswafwebaclsResource, name), &v1alpha1.AwsWafWebAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafWebAcl), err
}

// List takes label and field selectors, and returns the list of AwsWafWebAcls that match those selectors.
func (c *FakeAwsWafWebAcls) List(opts v1.ListOptions) (result *v1alpha1.AwsWafWebAclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awswafwebaclsResource, awswafwebaclsKind, opts), &v1alpha1.AwsWafWebAclList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsWafWebAclList{ListMeta: obj.(*v1alpha1.AwsWafWebAclList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsWafWebAclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafWebAcls.
func (c *FakeAwsWafWebAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awswafwebaclsResource, opts))
}

// Create takes the representation of a awsWafWebAcl and creates it.  Returns the server's representation of the awsWafWebAcl, and an error, if there is any.
func (c *FakeAwsWafWebAcls) Create(awsWafWebAcl *v1alpha1.AwsWafWebAcl) (result *v1alpha1.AwsWafWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awswafwebaclsResource, awsWafWebAcl), &v1alpha1.AwsWafWebAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafWebAcl), err
}

// Update takes the representation of a awsWafWebAcl and updates it. Returns the server's representation of the awsWafWebAcl, and an error, if there is any.
func (c *FakeAwsWafWebAcls) Update(awsWafWebAcl *v1alpha1.AwsWafWebAcl) (result *v1alpha1.AwsWafWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awswafwebaclsResource, awsWafWebAcl), &v1alpha1.AwsWafWebAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafWebAcl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsWafWebAcls) UpdateStatus(awsWafWebAcl *v1alpha1.AwsWafWebAcl) (*v1alpha1.AwsWafWebAcl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awswafwebaclsResource, "status", awsWafWebAcl), &v1alpha1.AwsWafWebAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafWebAcl), err
}

// Delete takes name of the awsWafWebAcl and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafWebAcls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awswafwebaclsResource, name), &v1alpha1.AwsWafWebAcl{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafWebAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awswafwebaclsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsWafWebAclList{})
	return err
}

// Patch applies the patch and returns the patched awsWafWebAcl.
func (c *FakeAwsWafWebAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafWebAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awswafwebaclsResource, name, pt, data, subresources...), &v1alpha1.AwsWafWebAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsWafWebAcl), err
}
