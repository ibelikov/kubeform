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

// FakeCdns implements CdnInterface
type FakeCdns struct {
	Fake *FakeDigitaloceanV1alpha1
	ns   string
}

var cdnsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "cdns"}

var cdnsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "Cdn"}

// Get takes name of the cdn, and returns the corresponding cdn object, and an error if there is any.
func (c *FakeCdns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Cdn, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cdnsResource, c.ns, name), &v1alpha1.Cdn{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cdn), err
}

// List takes label and field selectors, and returns the list of Cdns that match those selectors.
func (c *FakeCdns) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CdnList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cdnsResource, cdnsKind, c.ns, opts), &v1alpha1.CdnList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CdnList{ListMeta: obj.(*v1alpha1.CdnList).ListMeta}
	for _, item := range obj.(*v1alpha1.CdnList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cdns.
func (c *FakeCdns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cdnsResource, c.ns, opts))

}

// Create takes the representation of a cdn and creates it.  Returns the server's representation of the cdn, and an error, if there is any.
func (c *FakeCdns) Create(ctx context.Context, cdn *v1alpha1.Cdn, opts v1.CreateOptions) (result *v1alpha1.Cdn, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cdnsResource, c.ns, cdn), &v1alpha1.Cdn{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cdn), err
}

// Update takes the representation of a cdn and updates it. Returns the server's representation of the cdn, and an error, if there is any.
func (c *FakeCdns) Update(ctx context.Context, cdn *v1alpha1.Cdn, opts v1.UpdateOptions) (result *v1alpha1.Cdn, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cdnsResource, c.ns, cdn), &v1alpha1.Cdn{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cdn), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCdns) UpdateStatus(ctx context.Context, cdn *v1alpha1.Cdn, opts v1.UpdateOptions) (*v1alpha1.Cdn, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cdnsResource, "status", c.ns, cdn), &v1alpha1.Cdn{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cdn), err
}

// Delete takes name of the cdn and deletes it. Returns an error if one occurs.
func (c *FakeCdns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cdnsResource, c.ns, name), &v1alpha1.Cdn{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCdns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cdnsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CdnList{})
	return err
}

// Patch applies the patch and returns the patched cdn.
func (c *FakeCdns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Cdn, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cdnsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Cdn{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cdn), err
}
