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

// FakeDirectoryServiceDirectories implements DirectoryServiceDirectoryInterface
type FakeDirectoryServiceDirectories struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var directoryservicedirectoriesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "directoryservicedirectories"}

var directoryservicedirectoriesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DirectoryServiceDirectory"}

// Get takes name of the directoryServiceDirectory, and returns the corresponding directoryServiceDirectory object, and an error if there is any.
func (c *FakeDirectoryServiceDirectories) Get(name string, options v1.GetOptions) (result *v1alpha1.DirectoryServiceDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(directoryservicedirectoriesResource, c.ns, name), &v1alpha1.DirectoryServiceDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceDirectory), err
}

// List takes label and field selectors, and returns the list of DirectoryServiceDirectories that match those selectors.
func (c *FakeDirectoryServiceDirectories) List(opts v1.ListOptions) (result *v1alpha1.DirectoryServiceDirectoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(directoryservicedirectoriesResource, directoryservicedirectoriesKind, c.ns, opts), &v1alpha1.DirectoryServiceDirectoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DirectoryServiceDirectoryList{ListMeta: obj.(*v1alpha1.DirectoryServiceDirectoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.DirectoryServiceDirectoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directoryServiceDirectories.
func (c *FakeDirectoryServiceDirectories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(directoryservicedirectoriesResource, c.ns, opts))

}

// Create takes the representation of a directoryServiceDirectory and creates it.  Returns the server's representation of the directoryServiceDirectory, and an error, if there is any.
func (c *FakeDirectoryServiceDirectories) Create(directoryServiceDirectory *v1alpha1.DirectoryServiceDirectory) (result *v1alpha1.DirectoryServiceDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(directoryservicedirectoriesResource, c.ns, directoryServiceDirectory), &v1alpha1.DirectoryServiceDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceDirectory), err
}

// Update takes the representation of a directoryServiceDirectory and updates it. Returns the server's representation of the directoryServiceDirectory, and an error, if there is any.
func (c *FakeDirectoryServiceDirectories) Update(directoryServiceDirectory *v1alpha1.DirectoryServiceDirectory) (result *v1alpha1.DirectoryServiceDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(directoryservicedirectoriesResource, c.ns, directoryServiceDirectory), &v1alpha1.DirectoryServiceDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceDirectory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectoryServiceDirectories) UpdateStatus(directoryServiceDirectory *v1alpha1.DirectoryServiceDirectory) (*v1alpha1.DirectoryServiceDirectory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(directoryservicedirectoriesResource, "status", c.ns, directoryServiceDirectory), &v1alpha1.DirectoryServiceDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceDirectory), err
}

// Delete takes name of the directoryServiceDirectory and deletes it. Returns an error if one occurs.
func (c *FakeDirectoryServiceDirectories) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(directoryservicedirectoriesResource, c.ns, name), &v1alpha1.DirectoryServiceDirectory{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectoryServiceDirectories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(directoryservicedirectoriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DirectoryServiceDirectoryList{})
	return err
}

// Patch applies the patch and returns the patched directoryServiceDirectory.
func (c *FakeDirectoryServiceDirectories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DirectoryServiceDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(directoryservicedirectoriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DirectoryServiceDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryServiceDirectory), err
}
