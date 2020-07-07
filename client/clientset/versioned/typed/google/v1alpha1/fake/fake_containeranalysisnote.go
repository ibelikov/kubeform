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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeContainerAnalysisNotes implements ContainerAnalysisNoteInterface
type FakeContainerAnalysisNotes struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var containeranalysisnotesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "containeranalysisnotes"}

var containeranalysisnotesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ContainerAnalysisNote"}

// Get takes name of the containerAnalysisNote, and returns the corresponding containerAnalysisNote object, and an error if there is any.
func (c *FakeContainerAnalysisNotes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(containeranalysisnotesResource, c.ns, name), &v1alpha1.ContainerAnalysisNote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerAnalysisNote), err
}

// List takes label and field selectors, and returns the list of ContainerAnalysisNotes that match those selectors.
func (c *FakeContainerAnalysisNotes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ContainerAnalysisNoteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(containeranalysisnotesResource, containeranalysisnotesKind, c.ns, opts), &v1alpha1.ContainerAnalysisNoteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ContainerAnalysisNoteList{ListMeta: obj.(*v1alpha1.ContainerAnalysisNoteList).ListMeta}
	for _, item := range obj.(*v1alpha1.ContainerAnalysisNoteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerAnalysisNotes.
func (c *FakeContainerAnalysisNotes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(containeranalysisnotesResource, c.ns, opts))

}

// Create takes the representation of a containerAnalysisNote and creates it.  Returns the server's representation of the containerAnalysisNote, and an error, if there is any.
func (c *FakeContainerAnalysisNotes) Create(ctx context.Context, containerAnalysisNote *v1alpha1.ContainerAnalysisNote, opts v1.CreateOptions) (result *v1alpha1.ContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(containeranalysisnotesResource, c.ns, containerAnalysisNote), &v1alpha1.ContainerAnalysisNote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerAnalysisNote), err
}

// Update takes the representation of a containerAnalysisNote and updates it. Returns the server's representation of the containerAnalysisNote, and an error, if there is any.
func (c *FakeContainerAnalysisNotes) Update(ctx context.Context, containerAnalysisNote *v1alpha1.ContainerAnalysisNote, opts v1.UpdateOptions) (result *v1alpha1.ContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(containeranalysisnotesResource, c.ns, containerAnalysisNote), &v1alpha1.ContainerAnalysisNote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerAnalysisNote), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerAnalysisNotes) UpdateStatus(ctx context.Context, containerAnalysisNote *v1alpha1.ContainerAnalysisNote, opts v1.UpdateOptions) (*v1alpha1.ContainerAnalysisNote, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(containeranalysisnotesResource, "status", c.ns, containerAnalysisNote), &v1alpha1.ContainerAnalysisNote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerAnalysisNote), err
}

// Delete takes name of the containerAnalysisNote and deletes it. Returns an error if one occurs.
func (c *FakeContainerAnalysisNotes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(containeranalysisnotesResource, c.ns, name), &v1alpha1.ContainerAnalysisNote{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerAnalysisNotes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(containeranalysisnotesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ContainerAnalysisNoteList{})
	return err
}

// Patch applies the patch and returns the patched containerAnalysisNote.
func (c *FakeContainerAnalysisNotes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ContainerAnalysisNote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(containeranalysisnotesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ContainerAnalysisNote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerAnalysisNote), err
}
