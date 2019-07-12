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

// FakeAzurermHdinsightKafkaClusters implements AzurermHdinsightKafkaClusterInterface
type FakeAzurermHdinsightKafkaClusters struct {
	Fake *FakeAzurermV1alpha1
}

var azurermhdinsightkafkaclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermhdinsightkafkaclusters"}

var azurermhdinsightkafkaclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermHdinsightKafkaCluster"}

// Get takes name of the azurermHdinsightKafkaCluster, and returns the corresponding azurermHdinsightKafkaCluster object, and an error if there is any.
func (c *FakeAzurermHdinsightKafkaClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermHdinsightKafkaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermhdinsightkafkaclustersResource, name), &v1alpha1.AzurermHdinsightKafkaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightKafkaCluster), err
}

// List takes label and field selectors, and returns the list of AzurermHdinsightKafkaClusters that match those selectors.
func (c *FakeAzurermHdinsightKafkaClusters) List(opts v1.ListOptions) (result *v1alpha1.AzurermHdinsightKafkaClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermhdinsightkafkaclustersResource, azurermhdinsightkafkaclustersKind, opts), &v1alpha1.AzurermHdinsightKafkaClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermHdinsightKafkaClusterList{ListMeta: obj.(*v1alpha1.AzurermHdinsightKafkaClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermHdinsightKafkaClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermHdinsightKafkaClusters.
func (c *FakeAzurermHdinsightKafkaClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermhdinsightkafkaclustersResource, opts))
}

// Create takes the representation of a azurermHdinsightKafkaCluster and creates it.  Returns the server's representation of the azurermHdinsightKafkaCluster, and an error, if there is any.
func (c *FakeAzurermHdinsightKafkaClusters) Create(azurermHdinsightKafkaCluster *v1alpha1.AzurermHdinsightKafkaCluster) (result *v1alpha1.AzurermHdinsightKafkaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermhdinsightkafkaclustersResource, azurermHdinsightKafkaCluster), &v1alpha1.AzurermHdinsightKafkaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightKafkaCluster), err
}

// Update takes the representation of a azurermHdinsightKafkaCluster and updates it. Returns the server's representation of the azurermHdinsightKafkaCluster, and an error, if there is any.
func (c *FakeAzurermHdinsightKafkaClusters) Update(azurermHdinsightKafkaCluster *v1alpha1.AzurermHdinsightKafkaCluster) (result *v1alpha1.AzurermHdinsightKafkaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermhdinsightkafkaclustersResource, azurermHdinsightKafkaCluster), &v1alpha1.AzurermHdinsightKafkaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightKafkaCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermHdinsightKafkaClusters) UpdateStatus(azurermHdinsightKafkaCluster *v1alpha1.AzurermHdinsightKafkaCluster) (*v1alpha1.AzurermHdinsightKafkaCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermhdinsightkafkaclustersResource, "status", azurermHdinsightKafkaCluster), &v1alpha1.AzurermHdinsightKafkaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightKafkaCluster), err
}

// Delete takes name of the azurermHdinsightKafkaCluster and deletes it. Returns an error if one occurs.
func (c *FakeAzurermHdinsightKafkaClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermhdinsightkafkaclustersResource, name), &v1alpha1.AzurermHdinsightKafkaCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermHdinsightKafkaClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermhdinsightkafkaclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermHdinsightKafkaClusterList{})
	return err
}

// Patch applies the patch and returns the patched azurermHdinsightKafkaCluster.
func (c *FakeAzurermHdinsightKafkaClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermHdinsightKafkaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermhdinsightkafkaclustersResource, name, pt, data, subresources...), &v1alpha1.AzurermHdinsightKafkaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightKafkaCluster), err
}
