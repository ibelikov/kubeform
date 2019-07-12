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

// FakeAzurermHdinsightRserverClusters implements AzurermHdinsightRserverClusterInterface
type FakeAzurermHdinsightRserverClusters struct {
	Fake *FakeAzurermV1alpha1
}

var azurermhdinsightrserverclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermhdinsightrserverclusters"}

var azurermhdinsightrserverclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermHdinsightRserverCluster"}

// Get takes name of the azurermHdinsightRserverCluster, and returns the corresponding azurermHdinsightRserverCluster object, and an error if there is any.
func (c *FakeAzurermHdinsightRserverClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermHdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermhdinsightrserverclustersResource, name), &v1alpha1.AzurermHdinsightRserverCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightRserverCluster), err
}

// List takes label and field selectors, and returns the list of AzurermHdinsightRserverClusters that match those selectors.
func (c *FakeAzurermHdinsightRserverClusters) List(opts v1.ListOptions) (result *v1alpha1.AzurermHdinsightRserverClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermhdinsightrserverclustersResource, azurermhdinsightrserverclustersKind, opts), &v1alpha1.AzurermHdinsightRserverClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermHdinsightRserverClusterList{ListMeta: obj.(*v1alpha1.AzurermHdinsightRserverClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermHdinsightRserverClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermHdinsightRserverClusters.
func (c *FakeAzurermHdinsightRserverClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermhdinsightrserverclustersResource, opts))
}

// Create takes the representation of a azurermHdinsightRserverCluster and creates it.  Returns the server's representation of the azurermHdinsightRserverCluster, and an error, if there is any.
func (c *FakeAzurermHdinsightRserverClusters) Create(azurermHdinsightRserverCluster *v1alpha1.AzurermHdinsightRserverCluster) (result *v1alpha1.AzurermHdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermhdinsightrserverclustersResource, azurermHdinsightRserverCluster), &v1alpha1.AzurermHdinsightRserverCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightRserverCluster), err
}

// Update takes the representation of a azurermHdinsightRserverCluster and updates it. Returns the server's representation of the azurermHdinsightRserverCluster, and an error, if there is any.
func (c *FakeAzurermHdinsightRserverClusters) Update(azurermHdinsightRserverCluster *v1alpha1.AzurermHdinsightRserverCluster) (result *v1alpha1.AzurermHdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermhdinsightrserverclustersResource, azurermHdinsightRserverCluster), &v1alpha1.AzurermHdinsightRserverCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightRserverCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermHdinsightRserverClusters) UpdateStatus(azurermHdinsightRserverCluster *v1alpha1.AzurermHdinsightRserverCluster) (*v1alpha1.AzurermHdinsightRserverCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermhdinsightrserverclustersResource, "status", azurermHdinsightRserverCluster), &v1alpha1.AzurermHdinsightRserverCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightRserverCluster), err
}

// Delete takes name of the azurermHdinsightRserverCluster and deletes it. Returns an error if one occurs.
func (c *FakeAzurermHdinsightRserverClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermhdinsightrserverclustersResource, name), &v1alpha1.AzurermHdinsightRserverCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermHdinsightRserverClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermhdinsightrserverclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermHdinsightRserverClusterList{})
	return err
}

// Patch applies the patch and returns the patched azurermHdinsightRserverCluster.
func (c *FakeAzurermHdinsightRserverClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermHdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermhdinsightrserverclustersResource, name, pt, data, subresources...), &v1alpha1.AzurermHdinsightRserverCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightRserverCluster), err
}
