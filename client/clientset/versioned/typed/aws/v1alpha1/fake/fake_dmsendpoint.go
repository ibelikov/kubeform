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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeDmsEndpoints implements DmsEndpointInterface
type FakeDmsEndpoints struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dmsendpointsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dmsendpoints"}

var dmsendpointsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DmsEndpoint"}

// Get takes name of the dmsEndpoint, and returns the corresponding dmsEndpoint object, and an error if there is any.
func (c *FakeDmsEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DmsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dmsendpointsResource, c.ns, name), &v1alpha1.DmsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsEndpoint), err
}

// List takes label and field selectors, and returns the list of DmsEndpoints that match those selectors.
func (c *FakeDmsEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DmsEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dmsendpointsResource, dmsendpointsKind, c.ns, opts), &v1alpha1.DmsEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DmsEndpointList{ListMeta: obj.(*v1alpha1.DmsEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.DmsEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dmsEndpoints.
func (c *FakeDmsEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dmsendpointsResource, c.ns, opts))

}

// Create takes the representation of a dmsEndpoint and creates it.  Returns the server's representation of the dmsEndpoint, and an error, if there is any.
func (c *FakeDmsEndpoints) Create(ctx context.Context, dmsEndpoint *v1alpha1.DmsEndpoint, opts v1.CreateOptions) (result *v1alpha1.DmsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dmsendpointsResource, c.ns, dmsEndpoint), &v1alpha1.DmsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsEndpoint), err
}

// Update takes the representation of a dmsEndpoint and updates it. Returns the server's representation of the dmsEndpoint, and an error, if there is any.
func (c *FakeDmsEndpoints) Update(ctx context.Context, dmsEndpoint *v1alpha1.DmsEndpoint, opts v1.UpdateOptions) (result *v1alpha1.DmsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dmsendpointsResource, c.ns, dmsEndpoint), &v1alpha1.DmsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDmsEndpoints) UpdateStatus(ctx context.Context, dmsEndpoint *v1alpha1.DmsEndpoint, opts v1.UpdateOptions) (*v1alpha1.DmsEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dmsendpointsResource, "status", c.ns, dmsEndpoint), &v1alpha1.DmsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsEndpoint), err
}

// Delete takes name of the dmsEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeDmsEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dmsendpointsResource, c.ns, name), &v1alpha1.DmsEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDmsEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dmsendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DmsEndpointList{})
	return err
}

// Patch applies the patch and returns the patched dmsEndpoint.
func (c *FakeDmsEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DmsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dmsendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DmsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsEndpoint), err
}
