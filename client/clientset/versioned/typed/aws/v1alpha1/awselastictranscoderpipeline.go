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

// AwsElastictranscoderPipelinesGetter has a method to return a AwsElastictranscoderPipelineInterface.
// A group's client should implement this interface.
type AwsElastictranscoderPipelinesGetter interface {
	AwsElastictranscoderPipelines() AwsElastictranscoderPipelineInterface
}

// AwsElastictranscoderPipelineInterface has methods to work with AwsElastictranscoderPipeline resources.
type AwsElastictranscoderPipelineInterface interface {
	Create(*v1alpha1.AwsElastictranscoderPipeline) (*v1alpha1.AwsElastictranscoderPipeline, error)
	Update(*v1alpha1.AwsElastictranscoderPipeline) (*v1alpha1.AwsElastictranscoderPipeline, error)
	UpdateStatus(*v1alpha1.AwsElastictranscoderPipeline) (*v1alpha1.AwsElastictranscoderPipeline, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsElastictranscoderPipeline, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsElastictranscoderPipelineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElastictranscoderPipeline, err error)
	AwsElastictranscoderPipelineExpansion
}

// awsElastictranscoderPipelines implements AwsElastictranscoderPipelineInterface
type awsElastictranscoderPipelines struct {
	client rest.Interface
}

// newAwsElastictranscoderPipelines returns a AwsElastictranscoderPipelines
func newAwsElastictranscoderPipelines(c *AwsV1alpha1Client) *awsElastictranscoderPipelines {
	return &awsElastictranscoderPipelines{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsElastictranscoderPipeline, and returns the corresponding awsElastictranscoderPipeline object, and an error if there is any.
func (c *awsElastictranscoderPipelines) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	result = &v1alpha1.AwsElastictranscoderPipeline{}
	err = c.client.Get().
		Resource("awselastictranscoderpipelines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsElastictranscoderPipelines that match those selectors.
func (c *awsElastictranscoderPipelines) List(opts v1.ListOptions) (result *v1alpha1.AwsElastictranscoderPipelineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsElastictranscoderPipelineList{}
	err = c.client.Get().
		Resource("awselastictranscoderpipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsElastictranscoderPipelines.
func (c *awsElastictranscoderPipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awselastictranscoderpipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsElastictranscoderPipeline and creates it.  Returns the server's representation of the awsElastictranscoderPipeline, and an error, if there is any.
func (c *awsElastictranscoderPipelines) Create(awsElastictranscoderPipeline *v1alpha1.AwsElastictranscoderPipeline) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	result = &v1alpha1.AwsElastictranscoderPipeline{}
	err = c.client.Post().
		Resource("awselastictranscoderpipelines").
		Body(awsElastictranscoderPipeline).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsElastictranscoderPipeline and updates it. Returns the server's representation of the awsElastictranscoderPipeline, and an error, if there is any.
func (c *awsElastictranscoderPipelines) Update(awsElastictranscoderPipeline *v1alpha1.AwsElastictranscoderPipeline) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	result = &v1alpha1.AwsElastictranscoderPipeline{}
	err = c.client.Put().
		Resource("awselastictranscoderpipelines").
		Name(awsElastictranscoderPipeline.Name).
		Body(awsElastictranscoderPipeline).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsElastictranscoderPipelines) UpdateStatus(awsElastictranscoderPipeline *v1alpha1.AwsElastictranscoderPipeline) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	result = &v1alpha1.AwsElastictranscoderPipeline{}
	err = c.client.Put().
		Resource("awselastictranscoderpipelines").
		Name(awsElastictranscoderPipeline.Name).
		SubResource("status").
		Body(awsElastictranscoderPipeline).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsElastictranscoderPipeline and deletes it. Returns an error if one occurs.
func (c *awsElastictranscoderPipelines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awselastictranscoderpipelines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsElastictranscoderPipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awselastictranscoderpipelines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsElastictranscoderPipeline.
func (c *awsElastictranscoderPipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElastictranscoderPipeline, err error) {
	result = &v1alpha1.AwsElastictranscoderPipeline{}
	err = c.client.Patch(pt).
		Resource("awselastictranscoderpipelines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
