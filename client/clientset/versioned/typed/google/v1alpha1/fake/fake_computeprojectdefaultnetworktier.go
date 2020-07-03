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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeProjectDefaultNetworkTiers implements ComputeProjectDefaultNetworkTierInterface
type FakeComputeProjectDefaultNetworkTiers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeprojectdefaultnetworktiersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeprojectdefaultnetworktiers"}

var computeprojectdefaultnetworktiersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeProjectDefaultNetworkTier"}

// Get takes name of the computeProjectDefaultNetworkTier, and returns the corresponding computeProjectDefaultNetworkTier object, and an error if there is any.
func (c *FakeComputeProjectDefaultNetworkTiers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeProjectDefaultNetworkTier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeprojectdefaultnetworktiersResource, c.ns, name), &v1alpha1.ComputeProjectDefaultNetworkTier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectDefaultNetworkTier), err
}

// List takes label and field selectors, and returns the list of ComputeProjectDefaultNetworkTiers that match those selectors.
func (c *FakeComputeProjectDefaultNetworkTiers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeProjectDefaultNetworkTierList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeprojectdefaultnetworktiersResource, computeprojectdefaultnetworktiersKind, c.ns, opts), &v1alpha1.ComputeProjectDefaultNetworkTierList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeProjectDefaultNetworkTierList{ListMeta: obj.(*v1alpha1.ComputeProjectDefaultNetworkTierList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeProjectDefaultNetworkTierList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeProjectDefaultNetworkTiers.
func (c *FakeComputeProjectDefaultNetworkTiers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeprojectdefaultnetworktiersResource, c.ns, opts))

}

// Create takes the representation of a computeProjectDefaultNetworkTier and creates it.  Returns the server's representation of the computeProjectDefaultNetworkTier, and an error, if there is any.
func (c *FakeComputeProjectDefaultNetworkTiers) Create(ctx context.Context, computeProjectDefaultNetworkTier *v1alpha1.ComputeProjectDefaultNetworkTier, opts v1.CreateOptions) (result *v1alpha1.ComputeProjectDefaultNetworkTier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeprojectdefaultnetworktiersResource, c.ns, computeProjectDefaultNetworkTier), &v1alpha1.ComputeProjectDefaultNetworkTier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectDefaultNetworkTier), err
}

// Update takes the representation of a computeProjectDefaultNetworkTier and updates it. Returns the server's representation of the computeProjectDefaultNetworkTier, and an error, if there is any.
func (c *FakeComputeProjectDefaultNetworkTiers) Update(ctx context.Context, computeProjectDefaultNetworkTier *v1alpha1.ComputeProjectDefaultNetworkTier, opts v1.UpdateOptions) (result *v1alpha1.ComputeProjectDefaultNetworkTier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeprojectdefaultnetworktiersResource, c.ns, computeProjectDefaultNetworkTier), &v1alpha1.ComputeProjectDefaultNetworkTier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectDefaultNetworkTier), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeProjectDefaultNetworkTiers) UpdateStatus(ctx context.Context, computeProjectDefaultNetworkTier *v1alpha1.ComputeProjectDefaultNetworkTier, opts v1.UpdateOptions) (*v1alpha1.ComputeProjectDefaultNetworkTier, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeprojectdefaultnetworktiersResource, "status", c.ns, computeProjectDefaultNetworkTier), &v1alpha1.ComputeProjectDefaultNetworkTier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectDefaultNetworkTier), err
}

// Delete takes name of the computeProjectDefaultNetworkTier and deletes it. Returns an error if one occurs.
func (c *FakeComputeProjectDefaultNetworkTiers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeprojectdefaultnetworktiersResource, c.ns, name), &v1alpha1.ComputeProjectDefaultNetworkTier{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeProjectDefaultNetworkTiers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeprojectdefaultnetworktiersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeProjectDefaultNetworkTierList{})
	return err
}

// Patch applies the patch and returns the patched computeProjectDefaultNetworkTier.
func (c *FakeComputeProjectDefaultNetworkTiers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeProjectDefaultNetworkTier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeprojectdefaultnetworktiersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeProjectDefaultNetworkTier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectDefaultNetworkTier), err
}