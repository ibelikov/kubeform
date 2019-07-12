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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsIamGroupPolicyAttachmentsGetter has a method to return a AwsIamGroupPolicyAttachmentInterface.
// A group's client should implement this interface.
type AwsIamGroupPolicyAttachmentsGetter interface {
	AwsIamGroupPolicyAttachments() AwsIamGroupPolicyAttachmentInterface
}

// AwsIamGroupPolicyAttachmentInterface has methods to work with AwsIamGroupPolicyAttachment resources.
type AwsIamGroupPolicyAttachmentInterface interface {
	Create(*v1alpha1.AwsIamGroupPolicyAttachment) (*v1alpha1.AwsIamGroupPolicyAttachment, error)
	Update(*v1alpha1.AwsIamGroupPolicyAttachment) (*v1alpha1.AwsIamGroupPolicyAttachment, error)
	UpdateStatus(*v1alpha1.AwsIamGroupPolicyAttachment) (*v1alpha1.AwsIamGroupPolicyAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIamGroupPolicyAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIamGroupPolicyAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error)
	AwsIamGroupPolicyAttachmentExpansion
}

// awsIamGroupPolicyAttachments implements AwsIamGroupPolicyAttachmentInterface
type awsIamGroupPolicyAttachments struct {
	client rest.Interface
}

// newAwsIamGroupPolicyAttachments returns a AwsIamGroupPolicyAttachments
func newAwsIamGroupPolicyAttachments(c *AwsV1alpha1Client) *awsIamGroupPolicyAttachments {
	return &awsIamGroupPolicyAttachments{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsIamGroupPolicyAttachment, and returns the corresponding awsIamGroupPolicyAttachment object, and an error if there is any.
func (c *awsIamGroupPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamGroupPolicyAttachment{}
	err = c.client.Get().
		Resource("awsiamgrouppolicyattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamGroupPolicyAttachments that match those selectors.
func (c *awsIamGroupPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsIamGroupPolicyAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsIamGroupPolicyAttachmentList{}
	err = c.client.Get().
		Resource("awsiamgrouppolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamGroupPolicyAttachments.
func (c *awsIamGroupPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsiamgrouppolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsIamGroupPolicyAttachment and creates it.  Returns the server's representation of the awsIamGroupPolicyAttachment, and an error, if there is any.
func (c *awsIamGroupPolicyAttachments) Create(awsIamGroupPolicyAttachment *v1alpha1.AwsIamGroupPolicyAttachment) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamGroupPolicyAttachment{}
	err = c.client.Post().
		Resource("awsiamgrouppolicyattachments").
		Body(awsIamGroupPolicyAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamGroupPolicyAttachment and updates it. Returns the server's representation of the awsIamGroupPolicyAttachment, and an error, if there is any.
func (c *awsIamGroupPolicyAttachments) Update(awsIamGroupPolicyAttachment *v1alpha1.AwsIamGroupPolicyAttachment) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamGroupPolicyAttachment{}
	err = c.client.Put().
		Resource("awsiamgrouppolicyattachments").
		Name(awsIamGroupPolicyAttachment.Name).
		Body(awsIamGroupPolicyAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsIamGroupPolicyAttachments) UpdateStatus(awsIamGroupPolicyAttachment *v1alpha1.AwsIamGroupPolicyAttachment) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamGroupPolicyAttachment{}
	err = c.client.Put().
		Resource("awsiamgrouppolicyattachments").
		Name(awsIamGroupPolicyAttachment.Name).
		SubResource("status").
		Body(awsIamGroupPolicyAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamGroupPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *awsIamGroupPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsiamgrouppolicyattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamGroupPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsiamgrouppolicyattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamGroupPolicyAttachment.
func (c *awsIamGroupPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroupPolicyAttachment, err error) {
	result = &v1alpha1.AwsIamGroupPolicyAttachment{}
	err = c.client.Patch(pt).
		Resource("awsiamgrouppolicyattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
