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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeSqlDatabaseInstances implements SqlDatabaseInstanceInterface
type FakeSqlDatabaseInstances struct {
	Fake *FakeGoogleV1alpha1
}

var sqldatabaseinstancesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "sqldatabaseinstances"}

var sqldatabaseinstancesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "SqlDatabaseInstance"}

// Get takes name of the sqlDatabaseInstance, and returns the corresponding sqlDatabaseInstance object, and an error if there is any.
func (c *FakeSqlDatabaseInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.SqlDatabaseInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(sqldatabaseinstancesResource, name), &v1alpha1.SqlDatabaseInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlDatabaseInstance), err
}

// List takes label and field selectors, and returns the list of SqlDatabaseInstances that match those selectors.
func (c *FakeSqlDatabaseInstances) List(opts v1.ListOptions) (result *v1alpha1.SqlDatabaseInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(sqldatabaseinstancesResource, sqldatabaseinstancesKind, opts), &v1alpha1.SqlDatabaseInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqlDatabaseInstanceList{ListMeta: obj.(*v1alpha1.SqlDatabaseInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqlDatabaseInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlDatabaseInstances.
func (c *FakeSqlDatabaseInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(sqldatabaseinstancesResource, opts))
}

// Create takes the representation of a sqlDatabaseInstance and creates it.  Returns the server's representation of the sqlDatabaseInstance, and an error, if there is any.
func (c *FakeSqlDatabaseInstances) Create(sqlDatabaseInstance *v1alpha1.SqlDatabaseInstance) (result *v1alpha1.SqlDatabaseInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(sqldatabaseinstancesResource, sqlDatabaseInstance), &v1alpha1.SqlDatabaseInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlDatabaseInstance), err
}

// Update takes the representation of a sqlDatabaseInstance and updates it. Returns the server's representation of the sqlDatabaseInstance, and an error, if there is any.
func (c *FakeSqlDatabaseInstances) Update(sqlDatabaseInstance *v1alpha1.SqlDatabaseInstance) (result *v1alpha1.SqlDatabaseInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(sqldatabaseinstancesResource, sqlDatabaseInstance), &v1alpha1.SqlDatabaseInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlDatabaseInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqlDatabaseInstances) UpdateStatus(sqlDatabaseInstance *v1alpha1.SqlDatabaseInstance) (*v1alpha1.SqlDatabaseInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(sqldatabaseinstancesResource, "status", sqlDatabaseInstance), &v1alpha1.SqlDatabaseInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlDatabaseInstance), err
}

// Delete takes name of the sqlDatabaseInstance and deletes it. Returns an error if one occurs.
func (c *FakeSqlDatabaseInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(sqldatabaseinstancesResource, name), &v1alpha1.SqlDatabaseInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlDatabaseInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(sqldatabaseinstancesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqlDatabaseInstanceList{})
	return err
}

// Patch applies the patch and returns the patched sqlDatabaseInstance.
func (c *FakeSqlDatabaseInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlDatabaseInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(sqldatabaseinstancesResource, name, pt, data, subresources...), &v1alpha1.SqlDatabaseInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlDatabaseInstance), err
}
