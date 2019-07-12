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

// AwsDefaultRouteTablesGetter has a method to return a AwsDefaultRouteTableInterface.
// A group's client should implement this interface.
type AwsDefaultRouteTablesGetter interface {
	AwsDefaultRouteTables() AwsDefaultRouteTableInterface
}

// AwsDefaultRouteTableInterface has methods to work with AwsDefaultRouteTable resources.
type AwsDefaultRouteTableInterface interface {
	Create(*v1alpha1.AwsDefaultRouteTable) (*v1alpha1.AwsDefaultRouteTable, error)
	Update(*v1alpha1.AwsDefaultRouteTable) (*v1alpha1.AwsDefaultRouteTable, error)
	UpdateStatus(*v1alpha1.AwsDefaultRouteTable) (*v1alpha1.AwsDefaultRouteTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDefaultRouteTable, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDefaultRouteTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultRouteTable, err error)
	AwsDefaultRouteTableExpansion
}

// awsDefaultRouteTables implements AwsDefaultRouteTableInterface
type awsDefaultRouteTables struct {
	client rest.Interface
}

// newAwsDefaultRouteTables returns a AwsDefaultRouteTables
func newAwsDefaultRouteTables(c *AwsV1alpha1Client) *awsDefaultRouteTables {
	return &awsDefaultRouteTables{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDefaultRouteTable, and returns the corresponding awsDefaultRouteTable object, and an error if there is any.
func (c *awsDefaultRouteTables) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDefaultRouteTable, err error) {
	result = &v1alpha1.AwsDefaultRouteTable{}
	err = c.client.Get().
		Resource("awsdefaultroutetables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDefaultRouteTables that match those selectors.
func (c *awsDefaultRouteTables) List(opts v1.ListOptions) (result *v1alpha1.AwsDefaultRouteTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDefaultRouteTableList{}
	err = c.client.Get().
		Resource("awsdefaultroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDefaultRouteTables.
func (c *awsDefaultRouteTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdefaultroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDefaultRouteTable and creates it.  Returns the server's representation of the awsDefaultRouteTable, and an error, if there is any.
func (c *awsDefaultRouteTables) Create(awsDefaultRouteTable *v1alpha1.AwsDefaultRouteTable) (result *v1alpha1.AwsDefaultRouteTable, err error) {
	result = &v1alpha1.AwsDefaultRouteTable{}
	err = c.client.Post().
		Resource("awsdefaultroutetables").
		Body(awsDefaultRouteTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDefaultRouteTable and updates it. Returns the server's representation of the awsDefaultRouteTable, and an error, if there is any.
func (c *awsDefaultRouteTables) Update(awsDefaultRouteTable *v1alpha1.AwsDefaultRouteTable) (result *v1alpha1.AwsDefaultRouteTable, err error) {
	result = &v1alpha1.AwsDefaultRouteTable{}
	err = c.client.Put().
		Resource("awsdefaultroutetables").
		Name(awsDefaultRouteTable.Name).
		Body(awsDefaultRouteTable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDefaultRouteTables) UpdateStatus(awsDefaultRouteTable *v1alpha1.AwsDefaultRouteTable) (result *v1alpha1.AwsDefaultRouteTable, err error) {
	result = &v1alpha1.AwsDefaultRouteTable{}
	err = c.client.Put().
		Resource("awsdefaultroutetables").
		Name(awsDefaultRouteTable.Name).
		SubResource("status").
		Body(awsDefaultRouteTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDefaultRouteTable and deletes it. Returns an error if one occurs.
func (c *awsDefaultRouteTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdefaultroutetables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDefaultRouteTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdefaultroutetables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDefaultRouteTable.
func (c *awsDefaultRouteTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultRouteTable, err error) {
	result = &v1alpha1.AwsDefaultRouteTable{}
	err = c.client.Patch(pt).
		Resource("awsdefaultroutetables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
