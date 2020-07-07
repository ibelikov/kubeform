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

// FakeDropletSnapshots implements DropletSnapshotInterface
type FakeDropletSnapshots struct {
	Fake *FakeDigitaloceanV1alpha1
	ns   string
}

var dropletsnapshotsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "dropletsnapshots"}

var dropletsnapshotsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DropletSnapshot"}

// Get takes name of the dropletSnapshot, and returns the corresponding dropletSnapshot object, and an error if there is any.
func (c *FakeDropletSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DropletSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dropletsnapshotsResource, c.ns, name), &v1alpha1.DropletSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DropletSnapshot), err
}

// List takes label and field selectors, and returns the list of DropletSnapshots that match those selectors.
func (c *FakeDropletSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DropletSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dropletsnapshotsResource, dropletsnapshotsKind, c.ns, opts), &v1alpha1.DropletSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DropletSnapshotList{ListMeta: obj.(*v1alpha1.DropletSnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.DropletSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dropletSnapshots.
func (c *FakeDropletSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dropletsnapshotsResource, c.ns, opts))

}

// Create takes the representation of a dropletSnapshot and creates it.  Returns the server's representation of the dropletSnapshot, and an error, if there is any.
func (c *FakeDropletSnapshots) Create(ctx context.Context, dropletSnapshot *v1alpha1.DropletSnapshot, opts v1.CreateOptions) (result *v1alpha1.DropletSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dropletsnapshotsResource, c.ns, dropletSnapshot), &v1alpha1.DropletSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DropletSnapshot), err
}

// Update takes the representation of a dropletSnapshot and updates it. Returns the server's representation of the dropletSnapshot, and an error, if there is any.
func (c *FakeDropletSnapshots) Update(ctx context.Context, dropletSnapshot *v1alpha1.DropletSnapshot, opts v1.UpdateOptions) (result *v1alpha1.DropletSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dropletsnapshotsResource, c.ns, dropletSnapshot), &v1alpha1.DropletSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DropletSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDropletSnapshots) UpdateStatus(ctx context.Context, dropletSnapshot *v1alpha1.DropletSnapshot, opts v1.UpdateOptions) (*v1alpha1.DropletSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dropletsnapshotsResource, "status", c.ns, dropletSnapshot), &v1alpha1.DropletSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DropletSnapshot), err
}

// Delete takes name of the dropletSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeDropletSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dropletsnapshotsResource, c.ns, name), &v1alpha1.DropletSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDropletSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dropletsnapshotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DropletSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched dropletSnapshot.
func (c *FakeDropletSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DropletSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dropletsnapshotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DropletSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DropletSnapshot), err
}
