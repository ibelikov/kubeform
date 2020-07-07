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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeDatabaseClusters implements DatabaseClusterInterface
type FakeDatabaseClusters struct {
	Fake *FakeDigitaloceanV1alpha1
	ns   string
}

var databaseclustersResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "databaseclusters"}

var databaseclustersKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DatabaseCluster"}

// Get takes name of the databaseCluster, and returns the corresponding databaseCluster object, and an error if there is any.
func (c *FakeDatabaseClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatabaseCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(databaseclustersResource, c.ns, name), &v1alpha1.DatabaseCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseCluster), err
}

// List takes label and field selectors, and returns the list of DatabaseClusters that match those selectors.
func (c *FakeDatabaseClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatabaseClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(databaseclustersResource, databaseclustersKind, c.ns, opts), &v1alpha1.DatabaseClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatabaseClusterList{ListMeta: obj.(*v1alpha1.DatabaseClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatabaseClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested databaseClusters.
func (c *FakeDatabaseClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(databaseclustersResource, c.ns, opts))

}

// Create takes the representation of a databaseCluster and creates it.  Returns the server's representation of the databaseCluster, and an error, if there is any.
func (c *FakeDatabaseClusters) Create(ctx context.Context, databaseCluster *v1alpha1.DatabaseCluster, opts v1.CreateOptions) (result *v1alpha1.DatabaseCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(databaseclustersResource, c.ns, databaseCluster), &v1alpha1.DatabaseCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseCluster), err
}

// Update takes the representation of a databaseCluster and updates it. Returns the server's representation of the databaseCluster, and an error, if there is any.
func (c *FakeDatabaseClusters) Update(ctx context.Context, databaseCluster *v1alpha1.DatabaseCluster, opts v1.UpdateOptions) (result *v1alpha1.DatabaseCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(databaseclustersResource, c.ns, databaseCluster), &v1alpha1.DatabaseCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatabaseClusters) UpdateStatus(ctx context.Context, databaseCluster *v1alpha1.DatabaseCluster, opts v1.UpdateOptions) (*v1alpha1.DatabaseCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(databaseclustersResource, "status", c.ns, databaseCluster), &v1alpha1.DatabaseCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseCluster), err
}

// Delete takes name of the databaseCluster and deletes it. Returns an error if one occurs.
func (c *FakeDatabaseClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(databaseclustersResource, c.ns, name), &v1alpha1.DatabaseCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatabaseClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(databaseclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatabaseClusterList{})
	return err
}

// Patch applies the patch and returns the patched databaseCluster.
func (c *FakeDatabaseClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatabaseCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(databaseclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatabaseCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseCluster), err
}
