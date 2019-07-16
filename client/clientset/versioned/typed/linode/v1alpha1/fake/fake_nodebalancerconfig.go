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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeNodebalancerConfigs implements NodebalancerConfigInterface
type FakeNodebalancerConfigs struct {
	Fake *FakeLinodeV1alpha1
}

var nodebalancerconfigsResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "nodebalancerconfigs"}

var nodebalancerconfigsKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "NodebalancerConfig"}

// Get takes name of the nodebalancerConfig, and returns the corresponding nodebalancerConfig object, and an error if there is any.
func (c *FakeNodebalancerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.NodebalancerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodebalancerconfigsResource, name), &v1alpha1.NodebalancerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerConfig), err
}

// List takes label and field selectors, and returns the list of NodebalancerConfigs that match those selectors.
func (c *FakeNodebalancerConfigs) List(opts v1.ListOptions) (result *v1alpha1.NodebalancerConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodebalancerconfigsResource, nodebalancerconfigsKind, opts), &v1alpha1.NodebalancerConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodebalancerConfigList{ListMeta: obj.(*v1alpha1.NodebalancerConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodebalancerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodebalancerConfigs.
func (c *FakeNodebalancerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodebalancerconfigsResource, opts))
}

// Create takes the representation of a nodebalancerConfig and creates it.  Returns the server's representation of the nodebalancerConfig, and an error, if there is any.
func (c *FakeNodebalancerConfigs) Create(nodebalancerConfig *v1alpha1.NodebalancerConfig) (result *v1alpha1.NodebalancerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodebalancerconfigsResource, nodebalancerConfig), &v1alpha1.NodebalancerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerConfig), err
}

// Update takes the representation of a nodebalancerConfig and updates it. Returns the server's representation of the nodebalancerConfig, and an error, if there is any.
func (c *FakeNodebalancerConfigs) Update(nodebalancerConfig *v1alpha1.NodebalancerConfig) (result *v1alpha1.NodebalancerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodebalancerconfigsResource, nodebalancerConfig), &v1alpha1.NodebalancerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodebalancerConfigs) UpdateStatus(nodebalancerConfig *v1alpha1.NodebalancerConfig) (*v1alpha1.NodebalancerConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodebalancerconfigsResource, "status", nodebalancerConfig), &v1alpha1.NodebalancerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerConfig), err
}

// Delete takes name of the nodebalancerConfig and deletes it. Returns an error if one occurs.
func (c *FakeNodebalancerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodebalancerconfigsResource, name), &v1alpha1.NodebalancerConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodebalancerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodebalancerconfigsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodebalancerConfigList{})
	return err
}

// Patch applies the patch and returns the patched nodebalancerConfig.
func (c *FakeNodebalancerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodebalancerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodebalancerconfigsResource, name, pt, data, subresources...), &v1alpha1.NodebalancerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodebalancerConfig), err
}