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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeLogAnalyticsSolutions implements LogAnalyticsSolutionInterface
type FakeLogAnalyticsSolutions struct {
	Fake *FakeAzurermV1alpha1
}

var loganalyticssolutionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "loganalyticssolutions"}

var loganalyticssolutionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogAnalyticsSolution"}

// Get takes name of the logAnalyticsSolution, and returns the corresponding logAnalyticsSolution object, and an error if there is any.
func (c *FakeLogAnalyticsSolutions) Get(name string, options v1.GetOptions) (result *v1alpha1.LogAnalyticsSolution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(loganalyticssolutionsResource, name), &v1alpha1.LogAnalyticsSolution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsSolution), err
}

// List takes label and field selectors, and returns the list of LogAnalyticsSolutions that match those selectors.
func (c *FakeLogAnalyticsSolutions) List(opts v1.ListOptions) (result *v1alpha1.LogAnalyticsSolutionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(loganalyticssolutionsResource, loganalyticssolutionsKind, opts), &v1alpha1.LogAnalyticsSolutionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogAnalyticsSolutionList{ListMeta: obj.(*v1alpha1.LogAnalyticsSolutionList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogAnalyticsSolutionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logAnalyticsSolutions.
func (c *FakeLogAnalyticsSolutions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(loganalyticssolutionsResource, opts))
}

// Create takes the representation of a logAnalyticsSolution and creates it.  Returns the server's representation of the logAnalyticsSolution, and an error, if there is any.
func (c *FakeLogAnalyticsSolutions) Create(logAnalyticsSolution *v1alpha1.LogAnalyticsSolution) (result *v1alpha1.LogAnalyticsSolution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(loganalyticssolutionsResource, logAnalyticsSolution), &v1alpha1.LogAnalyticsSolution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsSolution), err
}

// Update takes the representation of a logAnalyticsSolution and updates it. Returns the server's representation of the logAnalyticsSolution, and an error, if there is any.
func (c *FakeLogAnalyticsSolutions) Update(logAnalyticsSolution *v1alpha1.LogAnalyticsSolution) (result *v1alpha1.LogAnalyticsSolution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(loganalyticssolutionsResource, logAnalyticsSolution), &v1alpha1.LogAnalyticsSolution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsSolution), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogAnalyticsSolutions) UpdateStatus(logAnalyticsSolution *v1alpha1.LogAnalyticsSolution) (*v1alpha1.LogAnalyticsSolution, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(loganalyticssolutionsResource, "status", logAnalyticsSolution), &v1alpha1.LogAnalyticsSolution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsSolution), err
}

// Delete takes name of the logAnalyticsSolution and deletes it. Returns an error if one occurs.
func (c *FakeLogAnalyticsSolutions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(loganalyticssolutionsResource, name), &v1alpha1.LogAnalyticsSolution{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogAnalyticsSolutions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(loganalyticssolutionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogAnalyticsSolutionList{})
	return err
}

// Patch applies the patch and returns the patched logAnalyticsSolution.
func (c *FakeLogAnalyticsSolutions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogAnalyticsSolution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(loganalyticssolutionsResource, name, pt, data, subresources...), &v1alpha1.LogAnalyticsSolution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogAnalyticsSolution), err
}
