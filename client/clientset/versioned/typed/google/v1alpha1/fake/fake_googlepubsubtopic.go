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

// FakeGooglePubsubTopics implements GooglePubsubTopicInterface
type FakeGooglePubsubTopics struct {
	Fake *FakeGoogleV1alpha1
}

var googlepubsubtopicsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlepubsubtopics"}

var googlepubsubtopicsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GooglePubsubTopic"}

// Get takes name of the googlePubsubTopic, and returns the corresponding googlePubsubTopic object, and an error if there is any.
func (c *FakeGooglePubsubTopics) Get(name string, options v1.GetOptions) (result *v1alpha1.GooglePubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlepubsubtopicsResource, name), &v1alpha1.GooglePubsubTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubTopic), err
}

// List takes label and field selectors, and returns the list of GooglePubsubTopics that match those selectors.
func (c *FakeGooglePubsubTopics) List(opts v1.ListOptions) (result *v1alpha1.GooglePubsubTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlepubsubtopicsResource, googlepubsubtopicsKind, opts), &v1alpha1.GooglePubsubTopicList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GooglePubsubTopicList{ListMeta: obj.(*v1alpha1.GooglePubsubTopicList).ListMeta}
	for _, item := range obj.(*v1alpha1.GooglePubsubTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googlePubsubTopics.
func (c *FakeGooglePubsubTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlepubsubtopicsResource, opts))
}

// Create takes the representation of a googlePubsubTopic and creates it.  Returns the server's representation of the googlePubsubTopic, and an error, if there is any.
func (c *FakeGooglePubsubTopics) Create(googlePubsubTopic *v1alpha1.GooglePubsubTopic) (result *v1alpha1.GooglePubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlepubsubtopicsResource, googlePubsubTopic), &v1alpha1.GooglePubsubTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubTopic), err
}

// Update takes the representation of a googlePubsubTopic and updates it. Returns the server's representation of the googlePubsubTopic, and an error, if there is any.
func (c *FakeGooglePubsubTopics) Update(googlePubsubTopic *v1alpha1.GooglePubsubTopic) (result *v1alpha1.GooglePubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlepubsubtopicsResource, googlePubsubTopic), &v1alpha1.GooglePubsubTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGooglePubsubTopics) UpdateStatus(googlePubsubTopic *v1alpha1.GooglePubsubTopic) (*v1alpha1.GooglePubsubTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlepubsubtopicsResource, "status", googlePubsubTopic), &v1alpha1.GooglePubsubTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubTopic), err
}

// Delete takes name of the googlePubsubTopic and deletes it. Returns an error if one occurs.
func (c *FakeGooglePubsubTopics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlepubsubtopicsResource, name), &v1alpha1.GooglePubsubTopic{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGooglePubsubTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlepubsubtopicsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GooglePubsubTopicList{})
	return err
}

// Patch applies the patch and returns the patched googlePubsubTopic.
func (c *FakeGooglePubsubTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GooglePubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlepubsubtopicsResource, name, pt, data, subresources...), &v1alpha1.GooglePubsubTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubTopic), err
}
