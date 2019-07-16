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

// CodebuildWebhooksGetter has a method to return a CodebuildWebhookInterface.
// A group's client should implement this interface.
type CodebuildWebhooksGetter interface {
	CodebuildWebhooks() CodebuildWebhookInterface
}

// CodebuildWebhookInterface has methods to work with CodebuildWebhook resources.
type CodebuildWebhookInterface interface {
	Create(*v1alpha1.CodebuildWebhook) (*v1alpha1.CodebuildWebhook, error)
	Update(*v1alpha1.CodebuildWebhook) (*v1alpha1.CodebuildWebhook, error)
	UpdateStatus(*v1alpha1.CodebuildWebhook) (*v1alpha1.CodebuildWebhook, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CodebuildWebhook, error)
	List(opts v1.ListOptions) (*v1alpha1.CodebuildWebhookList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodebuildWebhook, err error)
	CodebuildWebhookExpansion
}

// codebuildWebhooks implements CodebuildWebhookInterface
type codebuildWebhooks struct {
	client rest.Interface
}

// newCodebuildWebhooks returns a CodebuildWebhooks
func newCodebuildWebhooks(c *AwsV1alpha1Client) *codebuildWebhooks {
	return &codebuildWebhooks{
		client: c.RESTClient(),
	}
}

// Get takes name of the codebuildWebhook, and returns the corresponding codebuildWebhook object, and an error if there is any.
func (c *codebuildWebhooks) Get(name string, options v1.GetOptions) (result *v1alpha1.CodebuildWebhook, err error) {
	result = &v1alpha1.CodebuildWebhook{}
	err = c.client.Get().
		Resource("codebuildwebhooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CodebuildWebhooks that match those selectors.
func (c *codebuildWebhooks) List(opts v1.ListOptions) (result *v1alpha1.CodebuildWebhookList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CodebuildWebhookList{}
	err = c.client.Get().
		Resource("codebuildwebhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested codebuildWebhooks.
func (c *codebuildWebhooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("codebuildwebhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a codebuildWebhook and creates it.  Returns the server's representation of the codebuildWebhook, and an error, if there is any.
func (c *codebuildWebhooks) Create(codebuildWebhook *v1alpha1.CodebuildWebhook) (result *v1alpha1.CodebuildWebhook, err error) {
	result = &v1alpha1.CodebuildWebhook{}
	err = c.client.Post().
		Resource("codebuildwebhooks").
		Body(codebuildWebhook).
		Do().
		Into(result)
	return
}

// Update takes the representation of a codebuildWebhook and updates it. Returns the server's representation of the codebuildWebhook, and an error, if there is any.
func (c *codebuildWebhooks) Update(codebuildWebhook *v1alpha1.CodebuildWebhook) (result *v1alpha1.CodebuildWebhook, err error) {
	result = &v1alpha1.CodebuildWebhook{}
	err = c.client.Put().
		Resource("codebuildwebhooks").
		Name(codebuildWebhook.Name).
		Body(codebuildWebhook).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *codebuildWebhooks) UpdateStatus(codebuildWebhook *v1alpha1.CodebuildWebhook) (result *v1alpha1.CodebuildWebhook, err error) {
	result = &v1alpha1.CodebuildWebhook{}
	err = c.client.Put().
		Resource("codebuildwebhooks").
		Name(codebuildWebhook.Name).
		SubResource("status").
		Body(codebuildWebhook).
		Do().
		Into(result)
	return
}

// Delete takes name of the codebuildWebhook and deletes it. Returns an error if one occurs.
func (c *codebuildWebhooks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("codebuildwebhooks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *codebuildWebhooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("codebuildwebhooks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched codebuildWebhook.
func (c *codebuildWebhooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodebuildWebhook, err error) {
	result = &v1alpha1.CodebuildWebhook{}
	err = c.client.Patch(pt).
		Resource("codebuildwebhooks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
