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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleContainerAnalysisNotes implements GoogleContainerAnalysisNoteInterface
type FakeGoogleContainerAnalysisNotes struct {
	Fake *FakeGoogleV1alpha1
}

var googlecontaineranalysisnotesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecontaineranalysisnotes"}

var googlecontaineranalysisnotesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleContainerAnalysisNote"}

// Get takes name of the googleContainerAnalysisNote, and returns the corresponding googleContainerAnalysisNote object, and an error if there is any.
func (c *FakeGoogleContainerAnalysisNotes) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecontaineranalysisnotesResource, name), &v1alpha1.GoogleContainerAnalysisNote{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerAnalysisNote), err
}

// List takes label and field selectors, and returns the list of GoogleContainerAnalysisNotes that match those selectors.
func (c *FakeGoogleContainerAnalysisNotes) List(opts v1.ListOptions) (result *v1alpha1.GoogleContainerAnalysisNoteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecontaineranalysisnotesResource, googlecontaineranalysisnotesKind, opts), &v1alpha1.GoogleContainerAnalysisNoteList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleContainerAnalysisNoteList{ListMeta: obj.(*v1alpha1.GoogleContainerAnalysisNoteList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleContainerAnalysisNoteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleContainerAnalysisNotes.
func (c *FakeGoogleContainerAnalysisNotes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecontaineranalysisnotesResource, opts))
}

// Create takes the representation of a googleContainerAnalysisNote and creates it.  Returns the server's representation of the googleContainerAnalysisNote, and an error, if there is any.
func (c *FakeGoogleContainerAnalysisNotes) Create(googleContainerAnalysisNote *v1alpha1.GoogleContainerAnalysisNote) (result *v1alpha1.GoogleContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecontaineranalysisnotesResource, googleContainerAnalysisNote), &v1alpha1.GoogleContainerAnalysisNote{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerAnalysisNote), err
}

// Update takes the representation of a googleContainerAnalysisNote and updates it. Returns the server's representation of the googleContainerAnalysisNote, and an error, if there is any.
func (c *FakeGoogleContainerAnalysisNotes) Update(googleContainerAnalysisNote *v1alpha1.GoogleContainerAnalysisNote) (result *v1alpha1.GoogleContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecontaineranalysisnotesResource, googleContainerAnalysisNote), &v1alpha1.GoogleContainerAnalysisNote{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerAnalysisNote), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleContainerAnalysisNotes) UpdateStatus(googleContainerAnalysisNote *v1alpha1.GoogleContainerAnalysisNote) (*v1alpha1.GoogleContainerAnalysisNote, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecontaineranalysisnotesResource, "status", googleContainerAnalysisNote), &v1alpha1.GoogleContainerAnalysisNote{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerAnalysisNote), err
}

// Delete takes name of the googleContainerAnalysisNote and deletes it. Returns an error if one occurs.
func (c *FakeGoogleContainerAnalysisNotes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecontaineranalysisnotesResource, name), &v1alpha1.GoogleContainerAnalysisNote{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleContainerAnalysisNotes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecontaineranalysisnotesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleContainerAnalysisNoteList{})
	return err
}

// Patch applies the patch and returns the patched googleContainerAnalysisNote.
func (c *FakeGoogleContainerAnalysisNotes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecontaineranalysisnotesResource, name, pt, data, subresources...), &v1alpha1.GoogleContainerAnalysisNote{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleContainerAnalysisNote), err
}
