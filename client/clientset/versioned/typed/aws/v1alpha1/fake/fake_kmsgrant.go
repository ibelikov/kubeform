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

// FakeKmsGrants implements KmsGrantInterface
type FakeKmsGrants struct {
	Fake *FakeAwsV1alpha1
}

var kmsgrantsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "kmsgrants"}

var kmsgrantsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "KmsGrant"}

// Get takes name of the kmsGrant, and returns the corresponding kmsGrant object, and an error if there is any.
func (c *FakeKmsGrants) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kmsgrantsResource, name), &v1alpha1.KmsGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsGrant), err
}

// List takes label and field selectors, and returns the list of KmsGrants that match those selectors.
func (c *FakeKmsGrants) List(opts v1.ListOptions) (result *v1alpha1.KmsGrantList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kmsgrantsResource, kmsgrantsKind, opts), &v1alpha1.KmsGrantList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KmsGrantList{ListMeta: obj.(*v1alpha1.KmsGrantList).ListMeta}
	for _, item := range obj.(*v1alpha1.KmsGrantList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kmsGrants.
func (c *FakeKmsGrants) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kmsgrantsResource, opts))
}

// Create takes the representation of a kmsGrant and creates it.  Returns the server's representation of the kmsGrant, and an error, if there is any.
func (c *FakeKmsGrants) Create(kmsGrant *v1alpha1.KmsGrant) (result *v1alpha1.KmsGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kmsgrantsResource, kmsGrant), &v1alpha1.KmsGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsGrant), err
}

// Update takes the representation of a kmsGrant and updates it. Returns the server's representation of the kmsGrant, and an error, if there is any.
func (c *FakeKmsGrants) Update(kmsGrant *v1alpha1.KmsGrant) (result *v1alpha1.KmsGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kmsgrantsResource, kmsGrant), &v1alpha1.KmsGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsGrant), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKmsGrants) UpdateStatus(kmsGrant *v1alpha1.KmsGrant) (*v1alpha1.KmsGrant, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kmsgrantsResource, "status", kmsGrant), &v1alpha1.KmsGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsGrant), err
}

// Delete takes name of the kmsGrant and deletes it. Returns an error if one occurs.
func (c *FakeKmsGrants) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kmsgrantsResource, name), &v1alpha1.KmsGrant{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKmsGrants) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kmsgrantsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KmsGrantList{})
	return err
}

// Patch applies the patch and returns the patched kmsGrant.
func (c *FakeKmsGrants) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kmsgrantsResource, name, pt, data, subresources...), &v1alpha1.KmsGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsGrant), err
}