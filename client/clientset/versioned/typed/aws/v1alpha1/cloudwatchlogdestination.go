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

// CloudwatchLogDestinationsGetter has a method to return a CloudwatchLogDestinationInterface.
// A group's client should implement this interface.
type CloudwatchLogDestinationsGetter interface {
	CloudwatchLogDestinations() CloudwatchLogDestinationInterface
}

// CloudwatchLogDestinationInterface has methods to work with CloudwatchLogDestination resources.
type CloudwatchLogDestinationInterface interface {
	Create(*v1alpha1.CloudwatchLogDestination) (*v1alpha1.CloudwatchLogDestination, error)
	Update(*v1alpha1.CloudwatchLogDestination) (*v1alpha1.CloudwatchLogDestination, error)
	UpdateStatus(*v1alpha1.CloudwatchLogDestination) (*v1alpha1.CloudwatchLogDestination, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudwatchLogDestination, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudwatchLogDestinationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogDestination, err error)
	CloudwatchLogDestinationExpansion
}

// cloudwatchLogDestinations implements CloudwatchLogDestinationInterface
type cloudwatchLogDestinations struct {
	client rest.Interface
}

// newCloudwatchLogDestinations returns a CloudwatchLogDestinations
func newCloudwatchLogDestinations(c *AwsV1alpha1Client) *cloudwatchLogDestinations {
	return &cloudwatchLogDestinations{
		client: c.RESTClient(),
	}
}

// Get takes name of the cloudwatchLogDestination, and returns the corresponding cloudwatchLogDestination object, and an error if there is any.
func (c *cloudwatchLogDestinations) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogDestination, err error) {
	result = &v1alpha1.CloudwatchLogDestination{}
	err = c.client.Get().
		Resource("cloudwatchlogdestinations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudwatchLogDestinations that match those selectors.
func (c *cloudwatchLogDestinations) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogDestinationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudwatchLogDestinationList{}
	err = c.client.Get().
		Resource("cloudwatchlogdestinations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogDestinations.
func (c *cloudwatchLogDestinations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cloudwatchlogdestinations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudwatchLogDestination and creates it.  Returns the server's representation of the cloudwatchLogDestination, and an error, if there is any.
func (c *cloudwatchLogDestinations) Create(cloudwatchLogDestination *v1alpha1.CloudwatchLogDestination) (result *v1alpha1.CloudwatchLogDestination, err error) {
	result = &v1alpha1.CloudwatchLogDestination{}
	err = c.client.Post().
		Resource("cloudwatchlogdestinations").
		Body(cloudwatchLogDestination).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudwatchLogDestination and updates it. Returns the server's representation of the cloudwatchLogDestination, and an error, if there is any.
func (c *cloudwatchLogDestinations) Update(cloudwatchLogDestination *v1alpha1.CloudwatchLogDestination) (result *v1alpha1.CloudwatchLogDestination, err error) {
	result = &v1alpha1.CloudwatchLogDestination{}
	err = c.client.Put().
		Resource("cloudwatchlogdestinations").
		Name(cloudwatchLogDestination.Name).
		Body(cloudwatchLogDestination).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudwatchLogDestinations) UpdateStatus(cloudwatchLogDestination *v1alpha1.CloudwatchLogDestination) (result *v1alpha1.CloudwatchLogDestination, err error) {
	result = &v1alpha1.CloudwatchLogDestination{}
	err = c.client.Put().
		Resource("cloudwatchlogdestinations").
		Name(cloudwatchLogDestination.Name).
		SubResource("status").
		Body(cloudwatchLogDestination).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudwatchLogDestination and deletes it. Returns an error if one occurs.
func (c *cloudwatchLogDestinations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cloudwatchlogdestinations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudwatchLogDestinations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cloudwatchlogdestinations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudwatchLogDestination.
func (c *cloudwatchLogDestinations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogDestination, err error) {
	result = &v1alpha1.CloudwatchLogDestination{}
	err = c.client.Patch(pt).
		Resource("cloudwatchlogdestinations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
