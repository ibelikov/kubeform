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

// FakeStreamAnalyticsFunctionJavascriptUdves implements StreamAnalyticsFunctionJavascriptUdfInterface
type FakeStreamAnalyticsFunctionJavascriptUdves struct {
	Fake *FakeAzurermV1alpha1
}

var streamanalyticsfunctionjavascriptudvesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "streamanalyticsfunctionjavascriptudves"}

var streamanalyticsfunctionjavascriptudvesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StreamAnalyticsFunctionJavascriptUdf"}

// Get takes name of the streamAnalyticsFunctionJavascriptUdf, and returns the corresponding streamAnalyticsFunctionJavascriptUdf object, and an error if there is any.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) Get(name string, options v1.GetOptions) (result *v1alpha1.StreamAnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(streamanalyticsfunctionjavascriptudvesResource, name), &v1alpha1.StreamAnalyticsFunctionJavascriptUdf{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsFunctionJavascriptUdf), err
}

// List takes label and field selectors, and returns the list of StreamAnalyticsFunctionJavascriptUdves that match those selectors.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) List(opts v1.ListOptions) (result *v1alpha1.StreamAnalyticsFunctionJavascriptUdfList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(streamanalyticsfunctionjavascriptudvesResource, streamanalyticsfunctionjavascriptudvesKind, opts), &v1alpha1.StreamAnalyticsFunctionJavascriptUdfList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StreamAnalyticsFunctionJavascriptUdfList{ListMeta: obj.(*v1alpha1.StreamAnalyticsFunctionJavascriptUdfList).ListMeta}
	for _, item := range obj.(*v1alpha1.StreamAnalyticsFunctionJavascriptUdfList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested streamAnalyticsFunctionJavascriptUdves.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(streamanalyticsfunctionjavascriptudvesResource, opts))
}

// Create takes the representation of a streamAnalyticsFunctionJavascriptUdf and creates it.  Returns the server's representation of the streamAnalyticsFunctionJavascriptUdf, and an error, if there is any.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) Create(streamAnalyticsFunctionJavascriptUdf *v1alpha1.StreamAnalyticsFunctionJavascriptUdf) (result *v1alpha1.StreamAnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(streamanalyticsfunctionjavascriptudvesResource, streamAnalyticsFunctionJavascriptUdf), &v1alpha1.StreamAnalyticsFunctionJavascriptUdf{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsFunctionJavascriptUdf), err
}

// Update takes the representation of a streamAnalyticsFunctionJavascriptUdf and updates it. Returns the server's representation of the streamAnalyticsFunctionJavascriptUdf, and an error, if there is any.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) Update(streamAnalyticsFunctionJavascriptUdf *v1alpha1.StreamAnalyticsFunctionJavascriptUdf) (result *v1alpha1.StreamAnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(streamanalyticsfunctionjavascriptudvesResource, streamAnalyticsFunctionJavascriptUdf), &v1alpha1.StreamAnalyticsFunctionJavascriptUdf{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsFunctionJavascriptUdf), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) UpdateStatus(streamAnalyticsFunctionJavascriptUdf *v1alpha1.StreamAnalyticsFunctionJavascriptUdf) (*v1alpha1.StreamAnalyticsFunctionJavascriptUdf, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(streamanalyticsfunctionjavascriptudvesResource, "status", streamAnalyticsFunctionJavascriptUdf), &v1alpha1.StreamAnalyticsFunctionJavascriptUdf{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsFunctionJavascriptUdf), err
}

// Delete takes name of the streamAnalyticsFunctionJavascriptUdf and deletes it. Returns an error if one occurs.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(streamanalyticsfunctionjavascriptudvesResource, name), &v1alpha1.StreamAnalyticsFunctionJavascriptUdf{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(streamanalyticsfunctionjavascriptudvesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StreamAnalyticsFunctionJavascriptUdfList{})
	return err
}

// Patch applies the patch and returns the patched streamAnalyticsFunctionJavascriptUdf.
func (c *FakeStreamAnalyticsFunctionJavascriptUdves) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StreamAnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(streamanalyticsfunctionjavascriptudvesResource, name, pt, data, subresources...), &v1alpha1.StreamAnalyticsFunctionJavascriptUdf{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsFunctionJavascriptUdf), err
}
