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

// ElbAttachmentsGetter has a method to return a ElbAttachmentInterface.
// A group's client should implement this interface.
type ElbAttachmentsGetter interface {
	ElbAttachments() ElbAttachmentInterface
}

// ElbAttachmentInterface has methods to work with ElbAttachment resources.
type ElbAttachmentInterface interface {
	Create(*v1alpha1.ElbAttachment) (*v1alpha1.ElbAttachment, error)
	Update(*v1alpha1.ElbAttachment) (*v1alpha1.ElbAttachment, error)
	UpdateStatus(*v1alpha1.ElbAttachment) (*v1alpha1.ElbAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ElbAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.ElbAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElbAttachment, err error)
	ElbAttachmentExpansion
}

// elbAttachments implements ElbAttachmentInterface
type elbAttachments struct {
	client rest.Interface
}

// newElbAttachments returns a ElbAttachments
func newElbAttachments(c *AwsV1alpha1Client) *elbAttachments {
	return &elbAttachments{
		client: c.RESTClient(),
	}
}

// Get takes name of the elbAttachment, and returns the corresponding elbAttachment object, and an error if there is any.
func (c *elbAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.ElbAttachment, err error) {
	result = &v1alpha1.ElbAttachment{}
	err = c.client.Get().
		Resource("elbattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElbAttachments that match those selectors.
func (c *elbAttachments) List(opts v1.ListOptions) (result *v1alpha1.ElbAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElbAttachmentList{}
	err = c.client.Get().
		Resource("elbattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elbAttachments.
func (c *elbAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("elbattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a elbAttachment and creates it.  Returns the server's representation of the elbAttachment, and an error, if there is any.
func (c *elbAttachments) Create(elbAttachment *v1alpha1.ElbAttachment) (result *v1alpha1.ElbAttachment, err error) {
	result = &v1alpha1.ElbAttachment{}
	err = c.client.Post().
		Resource("elbattachments").
		Body(elbAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a elbAttachment and updates it. Returns the server's representation of the elbAttachment, and an error, if there is any.
func (c *elbAttachments) Update(elbAttachment *v1alpha1.ElbAttachment) (result *v1alpha1.ElbAttachment, err error) {
	result = &v1alpha1.ElbAttachment{}
	err = c.client.Put().
		Resource("elbattachments").
		Name(elbAttachment.Name).
		Body(elbAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *elbAttachments) UpdateStatus(elbAttachment *v1alpha1.ElbAttachment) (result *v1alpha1.ElbAttachment, err error) {
	result = &v1alpha1.ElbAttachment{}
	err = c.client.Put().
		Resource("elbattachments").
		Name(elbAttachment.Name).
		SubResource("status").
		Body(elbAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the elbAttachment and deletes it. Returns an error if one occurs.
func (c *elbAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("elbattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elbAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("elbattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched elbAttachment.
func (c *elbAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElbAttachment, err error) {
	result = &v1alpha1.ElbAttachment{}
	err = c.client.Patch(pt).
		Resource("elbattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
