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

// FakeCognitoIdentityPoolRolesAttachments implements CognitoIdentityPoolRolesAttachmentInterface
type FakeCognitoIdentityPoolRolesAttachments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cognitoidentitypoolrolesattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cognitoidentitypoolrolesattachments"}

var cognitoidentitypoolrolesattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CognitoIdentityPoolRolesAttachment"}

// Get takes name of the cognitoIdentityPoolRolesAttachment, and returns the corresponding cognitoIdentityPoolRolesAttachment object, and an error if there is any.
func (c *FakeCognitoIdentityPoolRolesAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CognitoIdentityPoolRolesAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cognitoidentitypoolrolesattachmentsResource, c.ns, name), &v1alpha1.CognitoIdentityPoolRolesAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityPoolRolesAttachment), err
}

// List takes label and field selectors, and returns the list of CognitoIdentityPoolRolesAttachments that match those selectors.
func (c *FakeCognitoIdentityPoolRolesAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CognitoIdentityPoolRolesAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cognitoidentitypoolrolesattachmentsResource, cognitoidentitypoolrolesattachmentsKind, c.ns, opts), &v1alpha1.CognitoIdentityPoolRolesAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CognitoIdentityPoolRolesAttachmentList{ListMeta: obj.(*v1alpha1.CognitoIdentityPoolRolesAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.CognitoIdentityPoolRolesAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cognitoIdentityPoolRolesAttachments.
func (c *FakeCognitoIdentityPoolRolesAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cognitoidentitypoolrolesattachmentsResource, c.ns, opts))

}

// Create takes the representation of a cognitoIdentityPoolRolesAttachment and creates it.  Returns the server's representation of the cognitoIdentityPoolRolesAttachment, and an error, if there is any.
func (c *FakeCognitoIdentityPoolRolesAttachments) Create(ctx context.Context, cognitoIdentityPoolRolesAttachment *v1alpha1.CognitoIdentityPoolRolesAttachment, opts v1.CreateOptions) (result *v1alpha1.CognitoIdentityPoolRolesAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cognitoidentitypoolrolesattachmentsResource, c.ns, cognitoIdentityPoolRolesAttachment), &v1alpha1.CognitoIdentityPoolRolesAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityPoolRolesAttachment), err
}

// Update takes the representation of a cognitoIdentityPoolRolesAttachment and updates it. Returns the server's representation of the cognitoIdentityPoolRolesAttachment, and an error, if there is any.
func (c *FakeCognitoIdentityPoolRolesAttachments) Update(ctx context.Context, cognitoIdentityPoolRolesAttachment *v1alpha1.CognitoIdentityPoolRolesAttachment, opts v1.UpdateOptions) (result *v1alpha1.CognitoIdentityPoolRolesAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cognitoidentitypoolrolesattachmentsResource, c.ns, cognitoIdentityPoolRolesAttachment), &v1alpha1.CognitoIdentityPoolRolesAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityPoolRolesAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCognitoIdentityPoolRolesAttachments) UpdateStatus(ctx context.Context, cognitoIdentityPoolRolesAttachment *v1alpha1.CognitoIdentityPoolRolesAttachment, opts v1.UpdateOptions) (*v1alpha1.CognitoIdentityPoolRolesAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cognitoidentitypoolrolesattachmentsResource, "status", c.ns, cognitoIdentityPoolRolesAttachment), &v1alpha1.CognitoIdentityPoolRolesAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityPoolRolesAttachment), err
}

// Delete takes name of the cognitoIdentityPoolRolesAttachment and deletes it. Returns an error if one occurs.
func (c *FakeCognitoIdentityPoolRolesAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cognitoidentitypoolrolesattachmentsResource, c.ns, name), &v1alpha1.CognitoIdentityPoolRolesAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCognitoIdentityPoolRolesAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cognitoidentitypoolrolesattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CognitoIdentityPoolRolesAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched cognitoIdentityPoolRolesAttachment.
func (c *FakeCognitoIdentityPoolRolesAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CognitoIdentityPoolRolesAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cognitoidentitypoolrolesattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CognitoIdentityPoolRolesAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitoIdentityPoolRolesAttachment), err
}
