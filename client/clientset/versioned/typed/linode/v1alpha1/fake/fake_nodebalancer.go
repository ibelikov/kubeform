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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodebalancers implements NodebalancerInterface
type FakeNodebalancers struct {
	Fake *FakeLinodeV1alpha1
	ns   string
}

var nodebalancersResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "nodebalancers"}

var nodebalancersKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "Nodebalancer"}

// Get takes name of the nodebalancer, and returns the corresponding nodebalancer object, and an error if there is any.
func (c *FakeNodebalancers) Get(name string, options v1.GetOptions) (result *v1alpha1.Nodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodebalancersResource, c.ns, name), &v1alpha1.Nodebalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nodebalancer), err
}

// List takes label and field selectors, and returns the list of Nodebalancers that match those selectors.
func (c *FakeNodebalancers) List(opts v1.ListOptions) (result *v1alpha1.NodebalancerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodebalancersResource, nodebalancersKind, c.ns, opts), &v1alpha1.NodebalancerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodebalancerList{ListMeta: obj.(*v1alpha1.NodebalancerList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodebalancerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodebalancers.
func (c *FakeNodebalancers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodebalancersResource, c.ns, opts))

}

// Create takes the representation of a nodebalancer and creates it.  Returns the server's representation of the nodebalancer, and an error, if there is any.
func (c *FakeNodebalancers) Create(nodebalancer *v1alpha1.Nodebalancer) (result *v1alpha1.Nodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodebalancersResource, c.ns, nodebalancer), &v1alpha1.Nodebalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nodebalancer), err
}

// Update takes the representation of a nodebalancer and updates it. Returns the server's representation of the nodebalancer, and an error, if there is any.
func (c *FakeNodebalancers) Update(nodebalancer *v1alpha1.Nodebalancer) (result *v1alpha1.Nodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodebalancersResource, c.ns, nodebalancer), &v1alpha1.Nodebalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nodebalancer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodebalancers) UpdateStatus(nodebalancer *v1alpha1.Nodebalancer) (*v1alpha1.Nodebalancer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodebalancersResource, "status", c.ns, nodebalancer), &v1alpha1.Nodebalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nodebalancer), err
}

// Delete takes name of the nodebalancer and deletes it. Returns an error if one occurs.
func (c *FakeNodebalancers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodebalancersResource, c.ns, name), &v1alpha1.Nodebalancer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodebalancers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodebalancersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodebalancerList{})
	return err
}

// Patch applies the patch and returns the patched nodebalancer.
func (c *FakeNodebalancers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Nodebalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodebalancersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Nodebalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Nodebalancer), err
}
