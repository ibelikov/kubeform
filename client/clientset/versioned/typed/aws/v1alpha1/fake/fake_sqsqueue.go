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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeSqsQueues implements SqsQueueInterface
type FakeSqsQueues struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var sqsqueuesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "sqsqueues"}

var sqsqueuesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SqsQueue"}

// Get takes name of the sqsQueue, and returns the corresponding sqsQueue object, and an error if there is any.
func (c *FakeSqsQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.SqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqsqueuesResource, c.ns, name), &v1alpha1.SqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqsQueue), err
}

// List takes label and field selectors, and returns the list of SqsQueues that match those selectors.
func (c *FakeSqsQueues) List(opts v1.ListOptions) (result *v1alpha1.SqsQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqsqueuesResource, sqsqueuesKind, c.ns, opts), &v1alpha1.SqsQueueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqsQueueList{ListMeta: obj.(*v1alpha1.SqsQueueList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqsQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqsQueues.
func (c *FakeSqsQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqsqueuesResource, c.ns, opts))

}

// Create takes the representation of a sqsQueue and creates it.  Returns the server's representation of the sqsQueue, and an error, if there is any.
func (c *FakeSqsQueues) Create(sqsQueue *v1alpha1.SqsQueue) (result *v1alpha1.SqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqsqueuesResource, c.ns, sqsQueue), &v1alpha1.SqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqsQueue), err
}

// Update takes the representation of a sqsQueue and updates it. Returns the server's representation of the sqsQueue, and an error, if there is any.
func (c *FakeSqsQueues) Update(sqsQueue *v1alpha1.SqsQueue) (result *v1alpha1.SqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqsqueuesResource, c.ns, sqsQueue), &v1alpha1.SqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqsQueue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqsQueues) UpdateStatus(sqsQueue *v1alpha1.SqsQueue) (*v1alpha1.SqsQueue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqsqueuesResource, "status", c.ns, sqsQueue), &v1alpha1.SqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqsQueue), err
}

// Delete takes name of the sqsQueue and deletes it. Returns an error if one occurs.
func (c *FakeSqsQueues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqsqueuesResource, c.ns, name), &v1alpha1.SqsQueue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqsQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqsqueuesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqsQueueList{})
	return err
}

// Patch applies the patch and returns the patched sqsQueue.
func (c *FakeSqsQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqsQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqsqueuesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SqsQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqsQueue), err
}
