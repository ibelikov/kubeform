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

// AwsCognitoResourceServersGetter has a method to return a AwsCognitoResourceServerInterface.
// A group's client should implement this interface.
type AwsCognitoResourceServersGetter interface {
	AwsCognitoResourceServers() AwsCognitoResourceServerInterface
}

// AwsCognitoResourceServerInterface has methods to work with AwsCognitoResourceServer resources.
type AwsCognitoResourceServerInterface interface {
	Create(*v1alpha1.AwsCognitoResourceServer) (*v1alpha1.AwsCognitoResourceServer, error)
	Update(*v1alpha1.AwsCognitoResourceServer) (*v1alpha1.AwsCognitoResourceServer, error)
	UpdateStatus(*v1alpha1.AwsCognitoResourceServer) (*v1alpha1.AwsCognitoResourceServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCognitoResourceServer, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCognitoResourceServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoResourceServer, err error)
	AwsCognitoResourceServerExpansion
}

// awsCognitoResourceServers implements AwsCognitoResourceServerInterface
type awsCognitoResourceServers struct {
	client rest.Interface
}

// newAwsCognitoResourceServers returns a AwsCognitoResourceServers
func newAwsCognitoResourceServers(c *AwsV1alpha1Client) *awsCognitoResourceServers {
	return &awsCognitoResourceServers{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCognitoResourceServer, and returns the corresponding awsCognitoResourceServer object, and an error if there is any.
func (c *awsCognitoResourceServers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCognitoResourceServer, err error) {
	result = &v1alpha1.AwsCognitoResourceServer{}
	err = c.client.Get().
		Resource("awscognitoresourceservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCognitoResourceServers that match those selectors.
func (c *awsCognitoResourceServers) List(opts v1.ListOptions) (result *v1alpha1.AwsCognitoResourceServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCognitoResourceServerList{}
	err = c.client.Get().
		Resource("awscognitoresourceservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCognitoResourceServers.
func (c *awsCognitoResourceServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscognitoresourceservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCognitoResourceServer and creates it.  Returns the server's representation of the awsCognitoResourceServer, and an error, if there is any.
func (c *awsCognitoResourceServers) Create(awsCognitoResourceServer *v1alpha1.AwsCognitoResourceServer) (result *v1alpha1.AwsCognitoResourceServer, err error) {
	result = &v1alpha1.AwsCognitoResourceServer{}
	err = c.client.Post().
		Resource("awscognitoresourceservers").
		Body(awsCognitoResourceServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCognitoResourceServer and updates it. Returns the server's representation of the awsCognitoResourceServer, and an error, if there is any.
func (c *awsCognitoResourceServers) Update(awsCognitoResourceServer *v1alpha1.AwsCognitoResourceServer) (result *v1alpha1.AwsCognitoResourceServer, err error) {
	result = &v1alpha1.AwsCognitoResourceServer{}
	err = c.client.Put().
		Resource("awscognitoresourceservers").
		Name(awsCognitoResourceServer.Name).
		Body(awsCognitoResourceServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCognitoResourceServers) UpdateStatus(awsCognitoResourceServer *v1alpha1.AwsCognitoResourceServer) (result *v1alpha1.AwsCognitoResourceServer, err error) {
	result = &v1alpha1.AwsCognitoResourceServer{}
	err = c.client.Put().
		Resource("awscognitoresourceservers").
		Name(awsCognitoResourceServer.Name).
		SubResource("status").
		Body(awsCognitoResourceServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCognitoResourceServer and deletes it. Returns an error if one occurs.
func (c *awsCognitoResourceServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscognitoresourceservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCognitoResourceServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscognitoresourceservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCognitoResourceServer.
func (c *awsCognitoResourceServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCognitoResourceServer, err error) {
	result = &v1alpha1.AwsCognitoResourceServer{}
	err = c.client.Patch(pt).
		Resource("awscognitoresourceservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
