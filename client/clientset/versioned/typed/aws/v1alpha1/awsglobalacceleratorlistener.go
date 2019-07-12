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

// AwsGlobalacceleratorListenersGetter has a method to return a AwsGlobalacceleratorListenerInterface.
// A group's client should implement this interface.
type AwsGlobalacceleratorListenersGetter interface {
	AwsGlobalacceleratorListeners() AwsGlobalacceleratorListenerInterface
}

// AwsGlobalacceleratorListenerInterface has methods to work with AwsGlobalacceleratorListener resources.
type AwsGlobalacceleratorListenerInterface interface {
	Create(*v1alpha1.AwsGlobalacceleratorListener) (*v1alpha1.AwsGlobalacceleratorListener, error)
	Update(*v1alpha1.AwsGlobalacceleratorListener) (*v1alpha1.AwsGlobalacceleratorListener, error)
	UpdateStatus(*v1alpha1.AwsGlobalacceleratorListener) (*v1alpha1.AwsGlobalacceleratorListener, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsGlobalacceleratorListener, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsGlobalacceleratorListenerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlobalacceleratorListener, err error)
	AwsGlobalacceleratorListenerExpansion
}

// awsGlobalacceleratorListeners implements AwsGlobalacceleratorListenerInterface
type awsGlobalacceleratorListeners struct {
	client rest.Interface
}

// newAwsGlobalacceleratorListeners returns a AwsGlobalacceleratorListeners
func newAwsGlobalacceleratorListeners(c *AwsV1alpha1Client) *awsGlobalacceleratorListeners {
	return &awsGlobalacceleratorListeners{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsGlobalacceleratorListener, and returns the corresponding awsGlobalacceleratorListener object, and an error if there is any.
func (c *awsGlobalacceleratorListeners) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	result = &v1alpha1.AwsGlobalacceleratorListener{}
	err = c.client.Get().
		Resource("awsglobalacceleratorlisteners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGlobalacceleratorListeners that match those selectors.
func (c *awsGlobalacceleratorListeners) List(opts v1.ListOptions) (result *v1alpha1.AwsGlobalacceleratorListenerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsGlobalacceleratorListenerList{}
	err = c.client.Get().
		Resource("awsglobalacceleratorlisteners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGlobalacceleratorListeners.
func (c *awsGlobalacceleratorListeners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsglobalacceleratorlisteners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsGlobalacceleratorListener and creates it.  Returns the server's representation of the awsGlobalacceleratorListener, and an error, if there is any.
func (c *awsGlobalacceleratorListeners) Create(awsGlobalacceleratorListener *v1alpha1.AwsGlobalacceleratorListener) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	result = &v1alpha1.AwsGlobalacceleratorListener{}
	err = c.client.Post().
		Resource("awsglobalacceleratorlisteners").
		Body(awsGlobalacceleratorListener).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGlobalacceleratorListener and updates it. Returns the server's representation of the awsGlobalacceleratorListener, and an error, if there is any.
func (c *awsGlobalacceleratorListeners) Update(awsGlobalacceleratorListener *v1alpha1.AwsGlobalacceleratorListener) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	result = &v1alpha1.AwsGlobalacceleratorListener{}
	err = c.client.Put().
		Resource("awsglobalacceleratorlisteners").
		Name(awsGlobalacceleratorListener.Name).
		Body(awsGlobalacceleratorListener).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsGlobalacceleratorListeners) UpdateStatus(awsGlobalacceleratorListener *v1alpha1.AwsGlobalacceleratorListener) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	result = &v1alpha1.AwsGlobalacceleratorListener{}
	err = c.client.Put().
		Resource("awsglobalacceleratorlisteners").
		Name(awsGlobalacceleratorListener.Name).
		SubResource("status").
		Body(awsGlobalacceleratorListener).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGlobalacceleratorListener and deletes it. Returns an error if one occurs.
func (c *awsGlobalacceleratorListeners) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsglobalacceleratorlisteners").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGlobalacceleratorListeners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsglobalacceleratorlisteners").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGlobalacceleratorListener.
func (c *awsGlobalacceleratorListeners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlobalacceleratorListener, err error) {
	result = &v1alpha1.AwsGlobalacceleratorListener{}
	err = c.client.Patch(pt).
		Resource("awsglobalacceleratorlisteners").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
