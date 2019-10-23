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

// FakeBatchComputeEnvironments implements BatchComputeEnvironmentInterface
type FakeBatchComputeEnvironments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var batchcomputeenvironmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "batchcomputeenvironments"}

var batchcomputeenvironmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "BatchComputeEnvironment"}

// Get takes name of the batchComputeEnvironment, and returns the corresponding batchComputeEnvironment object, and an error if there is any.
func (c *FakeBatchComputeEnvironments) Get(name string, options v1.GetOptions) (result *v1alpha1.BatchComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(batchcomputeenvironmentsResource, c.ns, name), &v1alpha1.BatchComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchComputeEnvironment), err
}

// List takes label and field selectors, and returns the list of BatchComputeEnvironments that match those selectors.
func (c *FakeBatchComputeEnvironments) List(opts v1.ListOptions) (result *v1alpha1.BatchComputeEnvironmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(batchcomputeenvironmentsResource, batchcomputeenvironmentsKind, c.ns, opts), &v1alpha1.BatchComputeEnvironmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BatchComputeEnvironmentList{ListMeta: obj.(*v1alpha1.BatchComputeEnvironmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.BatchComputeEnvironmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested batchComputeEnvironments.
func (c *FakeBatchComputeEnvironments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(batchcomputeenvironmentsResource, c.ns, opts))

}

// Create takes the representation of a batchComputeEnvironment and creates it.  Returns the server's representation of the batchComputeEnvironment, and an error, if there is any.
func (c *FakeBatchComputeEnvironments) Create(batchComputeEnvironment *v1alpha1.BatchComputeEnvironment) (result *v1alpha1.BatchComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(batchcomputeenvironmentsResource, c.ns, batchComputeEnvironment), &v1alpha1.BatchComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchComputeEnvironment), err
}

// Update takes the representation of a batchComputeEnvironment and updates it. Returns the server's representation of the batchComputeEnvironment, and an error, if there is any.
func (c *FakeBatchComputeEnvironments) Update(batchComputeEnvironment *v1alpha1.BatchComputeEnvironment) (result *v1alpha1.BatchComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(batchcomputeenvironmentsResource, c.ns, batchComputeEnvironment), &v1alpha1.BatchComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchComputeEnvironment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBatchComputeEnvironments) UpdateStatus(batchComputeEnvironment *v1alpha1.BatchComputeEnvironment) (*v1alpha1.BatchComputeEnvironment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(batchcomputeenvironmentsResource, "status", c.ns, batchComputeEnvironment), &v1alpha1.BatchComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchComputeEnvironment), err
}

// Delete takes name of the batchComputeEnvironment and deletes it. Returns an error if one occurs.
func (c *FakeBatchComputeEnvironments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(batchcomputeenvironmentsResource, c.ns, name), &v1alpha1.BatchComputeEnvironment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBatchComputeEnvironments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(batchcomputeenvironmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BatchComputeEnvironmentList{})
	return err
}

// Patch applies the patch and returns the patched batchComputeEnvironment.
func (c *FakeBatchComputeEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(batchcomputeenvironmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BatchComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchComputeEnvironment), err
}
