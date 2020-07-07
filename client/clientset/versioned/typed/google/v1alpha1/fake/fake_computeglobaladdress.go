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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeComputeGlobalAddresses implements ComputeGlobalAddressInterface
type FakeComputeGlobalAddresses struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeglobaladdressesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeglobaladdresses"}

var computeglobaladdressesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeGlobalAddress"}

// Get takes name of the computeGlobalAddress, and returns the corresponding computeGlobalAddress object, and an error if there is any.
func (c *FakeComputeGlobalAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeglobaladdressesResource, c.ns, name), &v1alpha1.ComputeGlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalAddress), err
}

// List takes label and field selectors, and returns the list of ComputeGlobalAddresses that match those selectors.
func (c *FakeComputeGlobalAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeGlobalAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeglobaladdressesResource, computeglobaladdressesKind, c.ns, opts), &v1alpha1.ComputeGlobalAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeGlobalAddressList{ListMeta: obj.(*v1alpha1.ComputeGlobalAddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeGlobalAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeGlobalAddresses.
func (c *FakeComputeGlobalAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeglobaladdressesResource, c.ns, opts))

}

// Create takes the representation of a computeGlobalAddress and creates it.  Returns the server's representation of the computeGlobalAddress, and an error, if there is any.
func (c *FakeComputeGlobalAddresses) Create(ctx context.Context, computeGlobalAddress *v1alpha1.ComputeGlobalAddress, opts v1.CreateOptions) (result *v1alpha1.ComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeglobaladdressesResource, c.ns, computeGlobalAddress), &v1alpha1.ComputeGlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalAddress), err
}

// Update takes the representation of a computeGlobalAddress and updates it. Returns the server's representation of the computeGlobalAddress, and an error, if there is any.
func (c *FakeComputeGlobalAddresses) Update(ctx context.Context, computeGlobalAddress *v1alpha1.ComputeGlobalAddress, opts v1.UpdateOptions) (result *v1alpha1.ComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeglobaladdressesResource, c.ns, computeGlobalAddress), &v1alpha1.ComputeGlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalAddress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeGlobalAddresses) UpdateStatus(ctx context.Context, computeGlobalAddress *v1alpha1.ComputeGlobalAddress, opts v1.UpdateOptions) (*v1alpha1.ComputeGlobalAddress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeglobaladdressesResource, "status", c.ns, computeGlobalAddress), &v1alpha1.ComputeGlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalAddress), err
}

// Delete takes name of the computeGlobalAddress and deletes it. Returns an error if one occurs.
func (c *FakeComputeGlobalAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeglobaladdressesResource, c.ns, name), &v1alpha1.ComputeGlobalAddress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeGlobalAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeglobaladdressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeGlobalAddressList{})
	return err
}

// Patch applies the patch and returns the patched computeGlobalAddress.
func (c *FakeComputeGlobalAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeglobaladdressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeGlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalAddress), err
}
