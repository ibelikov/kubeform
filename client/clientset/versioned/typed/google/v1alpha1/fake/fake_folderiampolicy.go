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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFolderIamPolicies implements FolderIamPolicyInterface
type FakeFolderIamPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var folderiampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "folderiampolicies"}

var folderiampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "FolderIamPolicy"}

// Get takes name of the folderIamPolicy, and returns the corresponding folderIamPolicy object, and an error if there is any.
func (c *FakeFolderIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.FolderIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(folderiampoliciesResource, c.ns, name), &v1alpha1.FolderIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamPolicy), err
}

// List takes label and field selectors, and returns the list of FolderIamPolicies that match those selectors.
func (c *FakeFolderIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.FolderIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(folderiampoliciesResource, folderiampoliciesKind, c.ns, opts), &v1alpha1.FolderIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FolderIamPolicyList{ListMeta: obj.(*v1alpha1.FolderIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.FolderIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested folderIamPolicies.
func (c *FakeFolderIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(folderiampoliciesResource, c.ns, opts))

}

// Create takes the representation of a folderIamPolicy and creates it.  Returns the server's representation of the folderIamPolicy, and an error, if there is any.
func (c *FakeFolderIamPolicies) Create(folderIamPolicy *v1alpha1.FolderIamPolicy) (result *v1alpha1.FolderIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(folderiampoliciesResource, c.ns, folderIamPolicy), &v1alpha1.FolderIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamPolicy), err
}

// Update takes the representation of a folderIamPolicy and updates it. Returns the server's representation of the folderIamPolicy, and an error, if there is any.
func (c *FakeFolderIamPolicies) Update(folderIamPolicy *v1alpha1.FolderIamPolicy) (result *v1alpha1.FolderIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(folderiampoliciesResource, c.ns, folderIamPolicy), &v1alpha1.FolderIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFolderIamPolicies) UpdateStatus(folderIamPolicy *v1alpha1.FolderIamPolicy) (*v1alpha1.FolderIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(folderiampoliciesResource, "status", c.ns, folderIamPolicy), &v1alpha1.FolderIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamPolicy), err
}

// Delete takes name of the folderIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeFolderIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(folderiampoliciesResource, c.ns, name), &v1alpha1.FolderIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFolderIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(folderiampoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FolderIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched folderIamPolicy.
func (c *FakeFolderIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FolderIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(folderiampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FolderIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderIamPolicy), err
}
