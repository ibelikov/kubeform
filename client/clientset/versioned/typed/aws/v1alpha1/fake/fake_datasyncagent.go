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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeDatasyncAgents implements DatasyncAgentInterface
type FakeDatasyncAgents struct {
	Fake *FakeAwsV1alpha1
}

var datasyncagentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "datasyncagents"}

var datasyncagentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DatasyncAgent"}

// Get takes name of the datasyncAgent, and returns the corresponding datasyncAgent object, and an error if there is any.
func (c *FakeDatasyncAgents) Get(name string, options v1.GetOptions) (result *v1alpha1.DatasyncAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(datasyncagentsResource, name), &v1alpha1.DatasyncAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncAgent), err
}

// List takes label and field selectors, and returns the list of DatasyncAgents that match those selectors.
func (c *FakeDatasyncAgents) List(opts v1.ListOptions) (result *v1alpha1.DatasyncAgentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(datasyncagentsResource, datasyncagentsKind, opts), &v1alpha1.DatasyncAgentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatasyncAgentList{ListMeta: obj.(*v1alpha1.DatasyncAgentList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatasyncAgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasyncAgents.
func (c *FakeDatasyncAgents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(datasyncagentsResource, opts))
}

// Create takes the representation of a datasyncAgent and creates it.  Returns the server's representation of the datasyncAgent, and an error, if there is any.
func (c *FakeDatasyncAgents) Create(datasyncAgent *v1alpha1.DatasyncAgent) (result *v1alpha1.DatasyncAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(datasyncagentsResource, datasyncAgent), &v1alpha1.DatasyncAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncAgent), err
}

// Update takes the representation of a datasyncAgent and updates it. Returns the server's representation of the datasyncAgent, and an error, if there is any.
func (c *FakeDatasyncAgents) Update(datasyncAgent *v1alpha1.DatasyncAgent) (result *v1alpha1.DatasyncAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(datasyncagentsResource, datasyncAgent), &v1alpha1.DatasyncAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncAgent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatasyncAgents) UpdateStatus(datasyncAgent *v1alpha1.DatasyncAgent) (*v1alpha1.DatasyncAgent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(datasyncagentsResource, "status", datasyncAgent), &v1alpha1.DatasyncAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncAgent), err
}

// Delete takes name of the datasyncAgent and deletes it. Returns an error if one occurs.
func (c *FakeDatasyncAgents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(datasyncagentsResource, name), &v1alpha1.DatasyncAgent{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasyncAgents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(datasyncagentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatasyncAgentList{})
	return err
}

// Patch applies the patch and returns the patched datasyncAgent.
func (c *FakeDatasyncAgents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatasyncAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(datasyncagentsResource, name, pt, data, subresources...), &v1alpha1.DatasyncAgent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncAgent), err
}
