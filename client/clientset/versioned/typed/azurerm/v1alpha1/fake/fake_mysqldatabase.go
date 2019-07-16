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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeMysqlDatabases implements MysqlDatabaseInterface
type FakeMysqlDatabases struct {
	Fake *FakeAzurermV1alpha1
}

var mysqldatabasesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mysqldatabases"}

var mysqldatabasesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MysqlDatabase"}

// Get takes name of the mysqlDatabase, and returns the corresponding mysqlDatabase object, and an error if there is any.
func (c *FakeMysqlDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(mysqldatabasesResource, name), &v1alpha1.MysqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

// List takes label and field selectors, and returns the list of MysqlDatabases that match those selectors.
func (c *FakeMysqlDatabases) List(opts v1.ListOptions) (result *v1alpha1.MysqlDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(mysqldatabasesResource, mysqldatabasesKind, opts), &v1alpha1.MysqlDatabaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlDatabaseList{ListMeta: obj.(*v1alpha1.MysqlDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mysqlDatabases.
func (c *FakeMysqlDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(mysqldatabasesResource, opts))
}

// Create takes the representation of a mysqlDatabase and creates it.  Returns the server's representation of the mysqlDatabase, and an error, if there is any.
func (c *FakeMysqlDatabases) Create(mysqlDatabase *v1alpha1.MysqlDatabase) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(mysqldatabasesResource, mysqlDatabase), &v1alpha1.MysqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

// Update takes the representation of a mysqlDatabase and updates it. Returns the server's representation of the mysqlDatabase, and an error, if there is any.
func (c *FakeMysqlDatabases) Update(mysqlDatabase *v1alpha1.MysqlDatabase) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(mysqldatabasesResource, mysqlDatabase), &v1alpha1.MysqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMysqlDatabases) UpdateStatus(mysqlDatabase *v1alpha1.MysqlDatabase) (*v1alpha1.MysqlDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(mysqldatabasesResource, "status", mysqlDatabase), &v1alpha1.MysqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

// Delete takes name of the mysqlDatabase and deletes it. Returns an error if one occurs.
func (c *FakeMysqlDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(mysqldatabasesResource, name), &v1alpha1.MysqlDatabase{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMysqlDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(mysqldatabasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched mysqlDatabase.
func (c *FakeMysqlDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(mysqldatabasesResource, name, pt, data, subresources...), &v1alpha1.MysqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}
