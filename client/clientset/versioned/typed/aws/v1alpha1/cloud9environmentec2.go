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

// Cloud9EnvironmentEc2sGetter has a method to return a Cloud9EnvironmentEc2Interface.
// A group's client should implement this interface.
type Cloud9EnvironmentEc2sGetter interface {
	Cloud9EnvironmentEc2s() Cloud9EnvironmentEc2Interface
}

// Cloud9EnvironmentEc2Interface has methods to work with Cloud9EnvironmentEc2 resources.
type Cloud9EnvironmentEc2Interface interface {
	Create(*v1alpha1.Cloud9EnvironmentEc2) (*v1alpha1.Cloud9EnvironmentEc2, error)
	Update(*v1alpha1.Cloud9EnvironmentEc2) (*v1alpha1.Cloud9EnvironmentEc2, error)
	UpdateStatus(*v1alpha1.Cloud9EnvironmentEc2) (*v1alpha1.Cloud9EnvironmentEc2, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Cloud9EnvironmentEc2, error)
	List(opts v1.ListOptions) (*v1alpha1.Cloud9EnvironmentEc2List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Cloud9EnvironmentEc2, err error)
	Cloud9EnvironmentEc2Expansion
}

// cloud9EnvironmentEc2s implements Cloud9EnvironmentEc2Interface
type cloud9EnvironmentEc2s struct {
	client rest.Interface
}

// newCloud9EnvironmentEc2s returns a Cloud9EnvironmentEc2s
func newCloud9EnvironmentEc2s(c *AwsV1alpha1Client) *cloud9EnvironmentEc2s {
	return &cloud9EnvironmentEc2s{
		client: c.RESTClient(),
	}
}

// Get takes name of the cloud9EnvironmentEc2, and returns the corresponding cloud9EnvironmentEc2 object, and an error if there is any.
func (c *cloud9EnvironmentEc2s) Get(name string, options v1.GetOptions) (result *v1alpha1.Cloud9EnvironmentEc2, err error) {
	result = &v1alpha1.Cloud9EnvironmentEc2{}
	err = c.client.Get().
		Resource("cloud9environmentec2s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Cloud9EnvironmentEc2s that match those selectors.
func (c *cloud9EnvironmentEc2s) List(opts v1.ListOptions) (result *v1alpha1.Cloud9EnvironmentEc2List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Cloud9EnvironmentEc2List{}
	err = c.client.Get().
		Resource("cloud9environmentec2s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloud9EnvironmentEc2s.
func (c *cloud9EnvironmentEc2s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cloud9environmentec2s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloud9EnvironmentEc2 and creates it.  Returns the server's representation of the cloud9EnvironmentEc2, and an error, if there is any.
func (c *cloud9EnvironmentEc2s) Create(cloud9EnvironmentEc2 *v1alpha1.Cloud9EnvironmentEc2) (result *v1alpha1.Cloud9EnvironmentEc2, err error) {
	result = &v1alpha1.Cloud9EnvironmentEc2{}
	err = c.client.Post().
		Resource("cloud9environmentec2s").
		Body(cloud9EnvironmentEc2).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloud9EnvironmentEc2 and updates it. Returns the server's representation of the cloud9EnvironmentEc2, and an error, if there is any.
func (c *cloud9EnvironmentEc2s) Update(cloud9EnvironmentEc2 *v1alpha1.Cloud9EnvironmentEc2) (result *v1alpha1.Cloud9EnvironmentEc2, err error) {
	result = &v1alpha1.Cloud9EnvironmentEc2{}
	err = c.client.Put().
		Resource("cloud9environmentec2s").
		Name(cloud9EnvironmentEc2.Name).
		Body(cloud9EnvironmentEc2).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloud9EnvironmentEc2s) UpdateStatus(cloud9EnvironmentEc2 *v1alpha1.Cloud9EnvironmentEc2) (result *v1alpha1.Cloud9EnvironmentEc2, err error) {
	result = &v1alpha1.Cloud9EnvironmentEc2{}
	err = c.client.Put().
		Resource("cloud9environmentec2s").
		Name(cloud9EnvironmentEc2.Name).
		SubResource("status").
		Body(cloud9EnvironmentEc2).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloud9EnvironmentEc2 and deletes it. Returns an error if one occurs.
func (c *cloud9EnvironmentEc2s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cloud9environmentec2s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloud9EnvironmentEc2s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cloud9environmentec2s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloud9EnvironmentEc2.
func (c *cloud9EnvironmentEc2s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Cloud9EnvironmentEc2, err error) {
	result = &v1alpha1.Cloud9EnvironmentEc2{}
	err = c.client.Patch(pt).
		Resource("cloud9environmentec2s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
