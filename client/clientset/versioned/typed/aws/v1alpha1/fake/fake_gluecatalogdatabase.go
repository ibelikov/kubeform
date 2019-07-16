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

// FakeGlueCatalogDatabases implements GlueCatalogDatabaseInterface
type FakeGlueCatalogDatabases struct {
	Fake *FakeAwsV1alpha1
}

var gluecatalogdatabasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "gluecatalogdatabases"}

var gluecatalogdatabasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GlueCatalogDatabase"}

// Get takes name of the glueCatalogDatabase, and returns the corresponding glueCatalogDatabase object, and an error if there is any.
func (c *FakeGlueCatalogDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.GlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(gluecatalogdatabasesResource, name), &v1alpha1.GlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogDatabase), err
}

// List takes label and field selectors, and returns the list of GlueCatalogDatabases that match those selectors.
func (c *FakeGlueCatalogDatabases) List(opts v1.ListOptions) (result *v1alpha1.GlueCatalogDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(gluecatalogdatabasesResource, gluecatalogdatabasesKind, opts), &v1alpha1.GlueCatalogDatabaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlueCatalogDatabaseList{ListMeta: obj.(*v1alpha1.GlueCatalogDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlueCatalogDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested glueCatalogDatabases.
func (c *FakeGlueCatalogDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(gluecatalogdatabasesResource, opts))
}

// Create takes the representation of a glueCatalogDatabase and creates it.  Returns the server's representation of the glueCatalogDatabase, and an error, if there is any.
func (c *FakeGlueCatalogDatabases) Create(glueCatalogDatabase *v1alpha1.GlueCatalogDatabase) (result *v1alpha1.GlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(gluecatalogdatabasesResource, glueCatalogDatabase), &v1alpha1.GlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogDatabase), err
}

// Update takes the representation of a glueCatalogDatabase and updates it. Returns the server's representation of the glueCatalogDatabase, and an error, if there is any.
func (c *FakeGlueCatalogDatabases) Update(glueCatalogDatabase *v1alpha1.GlueCatalogDatabase) (result *v1alpha1.GlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(gluecatalogdatabasesResource, glueCatalogDatabase), &v1alpha1.GlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlueCatalogDatabases) UpdateStatus(glueCatalogDatabase *v1alpha1.GlueCatalogDatabase) (*v1alpha1.GlueCatalogDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(gluecatalogdatabasesResource, "status", glueCatalogDatabase), &v1alpha1.GlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogDatabase), err
}

// Delete takes name of the glueCatalogDatabase and deletes it. Returns an error if one occurs.
func (c *FakeGlueCatalogDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(gluecatalogdatabasesResource, name), &v1alpha1.GlueCatalogDatabase{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlueCatalogDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(gluecatalogdatabasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlueCatalogDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched glueCatalogDatabase.
func (c *FakeGlueCatalogDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueCatalogDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(gluecatalogdatabasesResource, name, pt, data, subresources...), &v1alpha1.GlueCatalogDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCatalogDatabase), err
}
