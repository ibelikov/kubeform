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

// FakeCodepipelineWebhooks implements CodepipelineWebhookInterface
type FakeCodepipelineWebhooks struct {
	Fake *FakeAwsV1alpha1
}

var codepipelinewebhooksResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "codepipelinewebhooks"}

var codepipelinewebhooksKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CodepipelineWebhook"}

// Get takes name of the codepipelineWebhook, and returns the corresponding codepipelineWebhook object, and an error if there is any.
func (c *FakeCodepipelineWebhooks) Get(name string, options v1.GetOptions) (result *v1alpha1.CodepipelineWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(codepipelinewebhooksResource, name), &v1alpha1.CodepipelineWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodepipelineWebhook), err
}

// List takes label and field selectors, and returns the list of CodepipelineWebhooks that match those selectors.
func (c *FakeCodepipelineWebhooks) List(opts v1.ListOptions) (result *v1alpha1.CodepipelineWebhookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(codepipelinewebhooksResource, codepipelinewebhooksKind, opts), &v1alpha1.CodepipelineWebhookList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CodepipelineWebhookList{ListMeta: obj.(*v1alpha1.CodepipelineWebhookList).ListMeta}
	for _, item := range obj.(*v1alpha1.CodepipelineWebhookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested codepipelineWebhooks.
func (c *FakeCodepipelineWebhooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(codepipelinewebhooksResource, opts))
}

// Create takes the representation of a codepipelineWebhook and creates it.  Returns the server's representation of the codepipelineWebhook, and an error, if there is any.
func (c *FakeCodepipelineWebhooks) Create(codepipelineWebhook *v1alpha1.CodepipelineWebhook) (result *v1alpha1.CodepipelineWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(codepipelinewebhooksResource, codepipelineWebhook), &v1alpha1.CodepipelineWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodepipelineWebhook), err
}

// Update takes the representation of a codepipelineWebhook and updates it. Returns the server's representation of the codepipelineWebhook, and an error, if there is any.
func (c *FakeCodepipelineWebhooks) Update(codepipelineWebhook *v1alpha1.CodepipelineWebhook) (result *v1alpha1.CodepipelineWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(codepipelinewebhooksResource, codepipelineWebhook), &v1alpha1.CodepipelineWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodepipelineWebhook), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCodepipelineWebhooks) UpdateStatus(codepipelineWebhook *v1alpha1.CodepipelineWebhook) (*v1alpha1.CodepipelineWebhook, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(codepipelinewebhooksResource, "status", codepipelineWebhook), &v1alpha1.CodepipelineWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodepipelineWebhook), err
}

// Delete takes name of the codepipelineWebhook and deletes it. Returns an error if one occurs.
func (c *FakeCodepipelineWebhooks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(codepipelinewebhooksResource, name), &v1alpha1.CodepipelineWebhook{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCodepipelineWebhooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(codepipelinewebhooksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CodepipelineWebhookList{})
	return err
}

// Patch applies the patch and returns the patched codepipelineWebhook.
func (c *FakeCodepipelineWebhooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodepipelineWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(codepipelinewebhooksResource, name, pt, data, subresources...), &v1alpha1.CodepipelineWebhook{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodepipelineWebhook), err
}
