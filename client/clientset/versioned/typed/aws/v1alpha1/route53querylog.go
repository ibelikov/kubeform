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

// Route53QueryLogsGetter has a method to return a Route53QueryLogInterface.
// A group's client should implement this interface.
type Route53QueryLogsGetter interface {
	Route53QueryLogs() Route53QueryLogInterface
}

// Route53QueryLogInterface has methods to work with Route53QueryLog resources.
type Route53QueryLogInterface interface {
	Create(*v1alpha1.Route53QueryLog) (*v1alpha1.Route53QueryLog, error)
	Update(*v1alpha1.Route53QueryLog) (*v1alpha1.Route53QueryLog, error)
	UpdateStatus(*v1alpha1.Route53QueryLog) (*v1alpha1.Route53QueryLog, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Route53QueryLog, error)
	List(opts v1.ListOptions) (*v1alpha1.Route53QueryLogList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53QueryLog, err error)
	Route53QueryLogExpansion
}

// route53QueryLogs implements Route53QueryLogInterface
type route53QueryLogs struct {
	client rest.Interface
}

// newRoute53QueryLogs returns a Route53QueryLogs
func newRoute53QueryLogs(c *AwsV1alpha1Client) *route53QueryLogs {
	return &route53QueryLogs{
		client: c.RESTClient(),
	}
}

// Get takes name of the route53QueryLog, and returns the corresponding route53QueryLog object, and an error if there is any.
func (c *route53QueryLogs) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53QueryLog, err error) {
	result = &v1alpha1.Route53QueryLog{}
	err = c.client.Get().
		Resource("route53querylogs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Route53QueryLogs that match those selectors.
func (c *route53QueryLogs) List(opts v1.ListOptions) (result *v1alpha1.Route53QueryLogList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Route53QueryLogList{}
	err = c.client.Get().
		Resource("route53querylogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested route53QueryLogs.
func (c *route53QueryLogs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("route53querylogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a route53QueryLog and creates it.  Returns the server's representation of the route53QueryLog, and an error, if there is any.
func (c *route53QueryLogs) Create(route53QueryLog *v1alpha1.Route53QueryLog) (result *v1alpha1.Route53QueryLog, err error) {
	result = &v1alpha1.Route53QueryLog{}
	err = c.client.Post().
		Resource("route53querylogs").
		Body(route53QueryLog).
		Do().
		Into(result)
	return
}

// Update takes the representation of a route53QueryLog and updates it. Returns the server's representation of the route53QueryLog, and an error, if there is any.
func (c *route53QueryLogs) Update(route53QueryLog *v1alpha1.Route53QueryLog) (result *v1alpha1.Route53QueryLog, err error) {
	result = &v1alpha1.Route53QueryLog{}
	err = c.client.Put().
		Resource("route53querylogs").
		Name(route53QueryLog.Name).
		Body(route53QueryLog).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *route53QueryLogs) UpdateStatus(route53QueryLog *v1alpha1.Route53QueryLog) (result *v1alpha1.Route53QueryLog, err error) {
	result = &v1alpha1.Route53QueryLog{}
	err = c.client.Put().
		Resource("route53querylogs").
		Name(route53QueryLog.Name).
		SubResource("status").
		Body(route53QueryLog).
		Do().
		Into(result)
	return
}

// Delete takes name of the route53QueryLog and deletes it. Returns an error if one occurs.
func (c *route53QueryLogs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("route53querylogs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *route53QueryLogs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("route53querylogs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched route53QueryLog.
func (c *route53QueryLogs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53QueryLog, err error) {
	result = &v1alpha1.Route53QueryLog{}
	err = c.client.Patch(pt).
		Resource("route53querylogs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
