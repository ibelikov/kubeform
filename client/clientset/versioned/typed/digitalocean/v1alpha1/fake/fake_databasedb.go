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

// FakeDatabaseDbs implements DatabaseDbInterface
type FakeDatabaseDbs struct {
	Fake *FakeDigitaloceanV1alpha1
	ns   string
}

var databasedbsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "databasedbs"}

var databasedbsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DatabaseDb"}

// Get takes name of the databaseDb, and returns the corresponding databaseDb object, and an error if there is any.
func (c *FakeDatabaseDbs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatabaseDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(databasedbsResource, c.ns, name), &v1alpha1.DatabaseDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseDb), err
}

// List takes label and field selectors, and returns the list of DatabaseDbs that match those selectors.
func (c *FakeDatabaseDbs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatabaseDbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(databasedbsResource, databasedbsKind, c.ns, opts), &v1alpha1.DatabaseDbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatabaseDbList{ListMeta: obj.(*v1alpha1.DatabaseDbList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatabaseDbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested databaseDbs.
func (c *FakeDatabaseDbs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(databasedbsResource, c.ns, opts))

}

// Create takes the representation of a databaseDb and creates it.  Returns the server's representation of the databaseDb, and an error, if there is any.
func (c *FakeDatabaseDbs) Create(ctx context.Context, databaseDb *v1alpha1.DatabaseDb, opts v1.CreateOptions) (result *v1alpha1.DatabaseDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(databasedbsResource, c.ns, databaseDb), &v1alpha1.DatabaseDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseDb), err
}

// Update takes the representation of a databaseDb and updates it. Returns the server's representation of the databaseDb, and an error, if there is any.
func (c *FakeDatabaseDbs) Update(ctx context.Context, databaseDb *v1alpha1.DatabaseDb, opts v1.UpdateOptions) (result *v1alpha1.DatabaseDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(databasedbsResource, c.ns, databaseDb), &v1alpha1.DatabaseDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseDb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatabaseDbs) UpdateStatus(ctx context.Context, databaseDb *v1alpha1.DatabaseDb, opts v1.UpdateOptions) (*v1alpha1.DatabaseDb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(databasedbsResource, "status", c.ns, databaseDb), &v1alpha1.DatabaseDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseDb), err
}

// Delete takes name of the databaseDb and deletes it. Returns an error if one occurs.
func (c *FakeDatabaseDbs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(databasedbsResource, c.ns, name), &v1alpha1.DatabaseDb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatabaseDbs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(databasedbsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatabaseDbList{})
	return err
}

// Patch applies the patch and returns the patched databaseDb.
func (c *FakeDatabaseDbs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatabaseDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(databasedbsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatabaseDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseDb), err
}
