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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppmeshRoutes implements AppmeshRouteInterface
type FakeAppmeshRoutes struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var appmeshroutesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "appmeshroutes"}

var appmeshroutesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AppmeshRoute"}

// Get takes name of the appmeshRoute, and returns the corresponding appmeshRoute object, and an error if there is any.
func (c *FakeAppmeshRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.AppmeshRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appmeshroutesResource, c.ns, name), &v1alpha1.AppmeshRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshRoute), err
}

// List takes label and field selectors, and returns the list of AppmeshRoutes that match those selectors.
func (c *FakeAppmeshRoutes) List(opts v1.ListOptions) (result *v1alpha1.AppmeshRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appmeshroutesResource, appmeshroutesKind, c.ns, opts), &v1alpha1.AppmeshRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppmeshRouteList{ListMeta: obj.(*v1alpha1.AppmeshRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppmeshRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appmeshRoutes.
func (c *FakeAppmeshRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appmeshroutesResource, c.ns, opts))

}

// Create takes the representation of a appmeshRoute and creates it.  Returns the server's representation of the appmeshRoute, and an error, if there is any.
func (c *FakeAppmeshRoutes) Create(appmeshRoute *v1alpha1.AppmeshRoute) (result *v1alpha1.AppmeshRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appmeshroutesResource, c.ns, appmeshRoute), &v1alpha1.AppmeshRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshRoute), err
}

// Update takes the representation of a appmeshRoute and updates it. Returns the server's representation of the appmeshRoute, and an error, if there is any.
func (c *FakeAppmeshRoutes) Update(appmeshRoute *v1alpha1.AppmeshRoute) (result *v1alpha1.AppmeshRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appmeshroutesResource, c.ns, appmeshRoute), &v1alpha1.AppmeshRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppmeshRoutes) UpdateStatus(appmeshRoute *v1alpha1.AppmeshRoute) (*v1alpha1.AppmeshRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appmeshroutesResource, "status", c.ns, appmeshRoute), &v1alpha1.AppmeshRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshRoute), err
}

// Delete takes name of the appmeshRoute and deletes it. Returns an error if one occurs.
func (c *FakeAppmeshRoutes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appmeshroutesResource, c.ns, name), &v1alpha1.AppmeshRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppmeshRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appmeshroutesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppmeshRouteList{})
	return err
}

// Patch applies the patch and returns the patched appmeshRoute.
func (c *FakeAppmeshRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppmeshRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appmeshroutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppmeshRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshRoute), err
}
