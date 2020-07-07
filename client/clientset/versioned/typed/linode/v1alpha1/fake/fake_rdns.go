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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeRdnses implements RdnsInterface
type FakeRdnses struct {
	Fake *FakeLinodeV1alpha1
	ns   string
}

var rdnsesResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "rdnses"}

var rdnsesKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "Rdns"}

// Get takes name of the rdns, and returns the corresponding rdns object, and an error if there is any.
func (c *FakeRdnses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rdnsesResource, c.ns, name), &v1alpha1.Rdns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// List takes label and field selectors, and returns the list of Rdnses that match those selectors.
func (c *FakeRdnses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RdnsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rdnsesResource, rdnsesKind, c.ns, opts), &v1alpha1.RdnsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RdnsList{ListMeta: obj.(*v1alpha1.RdnsList).ListMeta}
	for _, item := range obj.(*v1alpha1.RdnsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rdnses.
func (c *FakeRdnses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rdnsesResource, c.ns, opts))

}

// Create takes the representation of a rdns and creates it.  Returns the server's representation of the rdns, and an error, if there is any.
func (c *FakeRdnses) Create(ctx context.Context, rdns *v1alpha1.Rdns, opts v1.CreateOptions) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rdnsesResource, c.ns, rdns), &v1alpha1.Rdns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// Update takes the representation of a rdns and updates it. Returns the server's representation of the rdns, and an error, if there is any.
func (c *FakeRdnses) Update(ctx context.Context, rdns *v1alpha1.Rdns, opts v1.UpdateOptions) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rdnsesResource, c.ns, rdns), &v1alpha1.Rdns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRdnses) UpdateStatus(ctx context.Context, rdns *v1alpha1.Rdns, opts v1.UpdateOptions) (*v1alpha1.Rdns, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rdnsesResource, "status", c.ns, rdns), &v1alpha1.Rdns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// Delete takes name of the rdns and deletes it. Returns an error if one occurs.
func (c *FakeRdnses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rdnsesResource, c.ns, name), &v1alpha1.Rdns{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRdnses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rdnsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RdnsList{})
	return err
}

// Patch applies the patch and returns the patched rdns.
func (c *FakeRdnses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rdnsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Rdns{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}
