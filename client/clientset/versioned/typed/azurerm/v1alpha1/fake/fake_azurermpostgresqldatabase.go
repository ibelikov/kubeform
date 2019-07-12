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

// FakeAzurermPostgresqlDatabases implements AzurermPostgresqlDatabaseInterface
type FakeAzurermPostgresqlDatabases struct {
	Fake *FakeAzurermV1alpha1
}

var azurermpostgresqldatabasesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermpostgresqldatabases"}

var azurermpostgresqldatabasesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermPostgresqlDatabase"}

// Get takes name of the azurermPostgresqlDatabase, and returns the corresponding azurermPostgresqlDatabase object, and an error if there is any.
func (c *FakeAzurermPostgresqlDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermPostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermpostgresqldatabasesResource, name), &v1alpha1.AzurermPostgresqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermPostgresqlDatabase), err
}

// List takes label and field selectors, and returns the list of AzurermPostgresqlDatabases that match those selectors.
func (c *FakeAzurermPostgresqlDatabases) List(opts v1.ListOptions) (result *v1alpha1.AzurermPostgresqlDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermpostgresqldatabasesResource, azurermpostgresqldatabasesKind, opts), &v1alpha1.AzurermPostgresqlDatabaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermPostgresqlDatabaseList{ListMeta: obj.(*v1alpha1.AzurermPostgresqlDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermPostgresqlDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermPostgresqlDatabases.
func (c *FakeAzurermPostgresqlDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermpostgresqldatabasesResource, opts))
}

// Create takes the representation of a azurermPostgresqlDatabase and creates it.  Returns the server's representation of the azurermPostgresqlDatabase, and an error, if there is any.
func (c *FakeAzurermPostgresqlDatabases) Create(azurermPostgresqlDatabase *v1alpha1.AzurermPostgresqlDatabase) (result *v1alpha1.AzurermPostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermpostgresqldatabasesResource, azurermPostgresqlDatabase), &v1alpha1.AzurermPostgresqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermPostgresqlDatabase), err
}

// Update takes the representation of a azurermPostgresqlDatabase and updates it. Returns the server's representation of the azurermPostgresqlDatabase, and an error, if there is any.
func (c *FakeAzurermPostgresqlDatabases) Update(azurermPostgresqlDatabase *v1alpha1.AzurermPostgresqlDatabase) (result *v1alpha1.AzurermPostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermpostgresqldatabasesResource, azurermPostgresqlDatabase), &v1alpha1.AzurermPostgresqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermPostgresqlDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermPostgresqlDatabases) UpdateStatus(azurermPostgresqlDatabase *v1alpha1.AzurermPostgresqlDatabase) (*v1alpha1.AzurermPostgresqlDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermpostgresqldatabasesResource, "status", azurermPostgresqlDatabase), &v1alpha1.AzurermPostgresqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermPostgresqlDatabase), err
}

// Delete takes name of the azurermPostgresqlDatabase and deletes it. Returns an error if one occurs.
func (c *FakeAzurermPostgresqlDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermpostgresqldatabasesResource, name), &v1alpha1.AzurermPostgresqlDatabase{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermPostgresqlDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermpostgresqldatabasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermPostgresqlDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched azurermPostgresqlDatabase.
func (c *FakeAzurermPostgresqlDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermPostgresqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermpostgresqldatabasesResource, name, pt, data, subresources...), &v1alpha1.AzurermPostgresqlDatabase{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermPostgresqlDatabase), err
}
