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

// FakeDatasyncTasks implements DatasyncTaskInterface
type FakeDatasyncTasks struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var datasynctasksResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "datasynctasks"}

var datasynctasksKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DatasyncTask"}

// Get takes name of the datasyncTask, and returns the corresponding datasyncTask object, and an error if there is any.
func (c *FakeDatasyncTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.DatasyncTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasynctasksResource, c.ns, name), &v1alpha1.DatasyncTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncTask), err
}

// List takes label and field selectors, and returns the list of DatasyncTasks that match those selectors.
func (c *FakeDatasyncTasks) List(opts v1.ListOptions) (result *v1alpha1.DatasyncTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasynctasksResource, datasynctasksKind, c.ns, opts), &v1alpha1.DatasyncTaskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatasyncTaskList{ListMeta: obj.(*v1alpha1.DatasyncTaskList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatasyncTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasyncTasks.
func (c *FakeDatasyncTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasynctasksResource, c.ns, opts))

}

// Create takes the representation of a datasyncTask and creates it.  Returns the server's representation of the datasyncTask, and an error, if there is any.
func (c *FakeDatasyncTasks) Create(datasyncTask *v1alpha1.DatasyncTask) (result *v1alpha1.DatasyncTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasynctasksResource, c.ns, datasyncTask), &v1alpha1.DatasyncTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncTask), err
}

// Update takes the representation of a datasyncTask and updates it. Returns the server's representation of the datasyncTask, and an error, if there is any.
func (c *FakeDatasyncTasks) Update(datasyncTask *v1alpha1.DatasyncTask) (result *v1alpha1.DatasyncTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasynctasksResource, c.ns, datasyncTask), &v1alpha1.DatasyncTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncTask), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatasyncTasks) UpdateStatus(datasyncTask *v1alpha1.DatasyncTask) (*v1alpha1.DatasyncTask, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datasynctasksResource, "status", c.ns, datasyncTask), &v1alpha1.DatasyncTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncTask), err
}

// Delete takes name of the datasyncTask and deletes it. Returns an error if one occurs.
func (c *FakeDatasyncTasks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasynctasksResource, c.ns, name), &v1alpha1.DatasyncTask{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasyncTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasynctasksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatasyncTaskList{})
	return err
}

// Patch applies the patch and returns the patched datasyncTask.
func (c *FakeDatasyncTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatasyncTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasynctasksResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatasyncTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncTask), err
}
