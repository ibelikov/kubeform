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

// FakeHdinsightMlServicesClusters implements HdinsightMlServicesClusterInterface
type FakeHdinsightMlServicesClusters struct {
	Fake *FakeAzurermV1alpha1
}

var hdinsightmlservicesclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "hdinsightmlservicesclusters"}

var hdinsightmlservicesclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "HdinsightMlServicesCluster"}

// Get takes name of the hdinsightMlServicesCluster, and returns the corresponding hdinsightMlServicesCluster object, and an error if there is any.
func (c *FakeHdinsightMlServicesClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(hdinsightmlservicesclustersResource, name), &v1alpha1.HdinsightMlServicesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightMlServicesCluster), err
}

// List takes label and field selectors, and returns the list of HdinsightMlServicesClusters that match those selectors.
func (c *FakeHdinsightMlServicesClusters) List(opts v1.ListOptions) (result *v1alpha1.HdinsightMlServicesClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(hdinsightmlservicesclustersResource, hdinsightmlservicesclustersKind, opts), &v1alpha1.HdinsightMlServicesClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HdinsightMlServicesClusterList{ListMeta: obj.(*v1alpha1.HdinsightMlServicesClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.HdinsightMlServicesClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hdinsightMlServicesClusters.
func (c *FakeHdinsightMlServicesClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(hdinsightmlservicesclustersResource, opts))
}

// Create takes the representation of a hdinsightMlServicesCluster and creates it.  Returns the server's representation of the hdinsightMlServicesCluster, and an error, if there is any.
func (c *FakeHdinsightMlServicesClusters) Create(hdinsightMlServicesCluster *v1alpha1.HdinsightMlServicesCluster) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(hdinsightmlservicesclustersResource, hdinsightMlServicesCluster), &v1alpha1.HdinsightMlServicesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightMlServicesCluster), err
}

// Update takes the representation of a hdinsightMlServicesCluster and updates it. Returns the server's representation of the hdinsightMlServicesCluster, and an error, if there is any.
func (c *FakeHdinsightMlServicesClusters) Update(hdinsightMlServicesCluster *v1alpha1.HdinsightMlServicesCluster) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(hdinsightmlservicesclustersResource, hdinsightMlServicesCluster), &v1alpha1.HdinsightMlServicesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightMlServicesCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHdinsightMlServicesClusters) UpdateStatus(hdinsightMlServicesCluster *v1alpha1.HdinsightMlServicesCluster) (*v1alpha1.HdinsightMlServicesCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(hdinsightmlservicesclustersResource, "status", hdinsightMlServicesCluster), &v1alpha1.HdinsightMlServicesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightMlServicesCluster), err
}

// Delete takes name of the hdinsightMlServicesCluster and deletes it. Returns an error if one occurs.
func (c *FakeHdinsightMlServicesClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(hdinsightmlservicesclustersResource, name), &v1alpha1.HdinsightMlServicesCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHdinsightMlServicesClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(hdinsightmlservicesclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HdinsightMlServicesClusterList{})
	return err
}

// Patch applies the patch and returns the patched hdinsightMlServicesCluster.
func (c *FakeHdinsightMlServicesClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(hdinsightmlservicesclustersResource, name, pt, data, subresources...), &v1alpha1.HdinsightMlServicesCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightMlServicesCluster), err
}
