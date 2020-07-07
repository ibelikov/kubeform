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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeCodebuildWebhooks implements CodebuildWebhookInterface
type FakeCodebuildWebhooks struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var codebuildwebhooksResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "codebuildwebhooks"}

var codebuildwebhooksKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CodebuildWebhook"}

// Get takes name of the codebuildWebhook, and returns the corresponding codebuildWebhook object, and an error if there is any.
func (c *FakeCodebuildWebhooks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CodebuildWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(codebuildwebhooksResource, c.ns, name), &v1alpha1.CodebuildWebhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildWebhook), err
}

// List takes label and field selectors, and returns the list of CodebuildWebhooks that match those selectors.
func (c *FakeCodebuildWebhooks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CodebuildWebhookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(codebuildwebhooksResource, codebuildwebhooksKind, c.ns, opts), &v1alpha1.CodebuildWebhookList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CodebuildWebhookList{ListMeta: obj.(*v1alpha1.CodebuildWebhookList).ListMeta}
	for _, item := range obj.(*v1alpha1.CodebuildWebhookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested codebuildWebhooks.
func (c *FakeCodebuildWebhooks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(codebuildwebhooksResource, c.ns, opts))

}

// Create takes the representation of a codebuildWebhook and creates it.  Returns the server's representation of the codebuildWebhook, and an error, if there is any.
func (c *FakeCodebuildWebhooks) Create(ctx context.Context, codebuildWebhook *v1alpha1.CodebuildWebhook, opts v1.CreateOptions) (result *v1alpha1.CodebuildWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(codebuildwebhooksResource, c.ns, codebuildWebhook), &v1alpha1.CodebuildWebhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildWebhook), err
}

// Update takes the representation of a codebuildWebhook and updates it. Returns the server's representation of the codebuildWebhook, and an error, if there is any.
func (c *FakeCodebuildWebhooks) Update(ctx context.Context, codebuildWebhook *v1alpha1.CodebuildWebhook, opts v1.UpdateOptions) (result *v1alpha1.CodebuildWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(codebuildwebhooksResource, c.ns, codebuildWebhook), &v1alpha1.CodebuildWebhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildWebhook), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCodebuildWebhooks) UpdateStatus(ctx context.Context, codebuildWebhook *v1alpha1.CodebuildWebhook, opts v1.UpdateOptions) (*v1alpha1.CodebuildWebhook, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(codebuildwebhooksResource, "status", c.ns, codebuildWebhook), &v1alpha1.CodebuildWebhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildWebhook), err
}

// Delete takes name of the codebuildWebhook and deletes it. Returns an error if one occurs.
func (c *FakeCodebuildWebhooks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(codebuildwebhooksResource, c.ns, name), &v1alpha1.CodebuildWebhook{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCodebuildWebhooks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(codebuildwebhooksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CodebuildWebhookList{})
	return err
}

// Patch applies the patch and returns the patched codebuildWebhook.
func (c *FakeCodebuildWebhooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CodebuildWebhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(codebuildwebhooksResource, c.ns, name, pt, data, subresources...), &v1alpha1.CodebuildWebhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodebuildWebhook), err
}
