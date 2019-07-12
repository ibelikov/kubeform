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

// AwsDynamodbTablesGetter has a method to return a AwsDynamodbTableInterface.
// A group's client should implement this interface.
type AwsDynamodbTablesGetter interface {
	AwsDynamodbTables() AwsDynamodbTableInterface
}

// AwsDynamodbTableInterface has methods to work with AwsDynamodbTable resources.
type AwsDynamodbTableInterface interface {
	Create(*v1alpha1.AwsDynamodbTable) (*v1alpha1.AwsDynamodbTable, error)
	Update(*v1alpha1.AwsDynamodbTable) (*v1alpha1.AwsDynamodbTable, error)
	UpdateStatus(*v1alpha1.AwsDynamodbTable) (*v1alpha1.AwsDynamodbTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDynamodbTable, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDynamodbTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDynamodbTable, err error)
	AwsDynamodbTableExpansion
}

// awsDynamodbTables implements AwsDynamodbTableInterface
type awsDynamodbTables struct {
	client rest.Interface
}

// newAwsDynamodbTables returns a AwsDynamodbTables
func newAwsDynamodbTables(c *AwsV1alpha1Client) *awsDynamodbTables {
	return &awsDynamodbTables{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDynamodbTable, and returns the corresponding awsDynamodbTable object, and an error if there is any.
func (c *awsDynamodbTables) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDynamodbTable, err error) {
	result = &v1alpha1.AwsDynamodbTable{}
	err = c.client.Get().
		Resource("awsdynamodbtables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDynamodbTables that match those selectors.
func (c *awsDynamodbTables) List(opts v1.ListOptions) (result *v1alpha1.AwsDynamodbTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDynamodbTableList{}
	err = c.client.Get().
		Resource("awsdynamodbtables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDynamodbTables.
func (c *awsDynamodbTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdynamodbtables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDynamodbTable and creates it.  Returns the server's representation of the awsDynamodbTable, and an error, if there is any.
func (c *awsDynamodbTables) Create(awsDynamodbTable *v1alpha1.AwsDynamodbTable) (result *v1alpha1.AwsDynamodbTable, err error) {
	result = &v1alpha1.AwsDynamodbTable{}
	err = c.client.Post().
		Resource("awsdynamodbtables").
		Body(awsDynamodbTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDynamodbTable and updates it. Returns the server's representation of the awsDynamodbTable, and an error, if there is any.
func (c *awsDynamodbTables) Update(awsDynamodbTable *v1alpha1.AwsDynamodbTable) (result *v1alpha1.AwsDynamodbTable, err error) {
	result = &v1alpha1.AwsDynamodbTable{}
	err = c.client.Put().
		Resource("awsdynamodbtables").
		Name(awsDynamodbTable.Name).
		Body(awsDynamodbTable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDynamodbTables) UpdateStatus(awsDynamodbTable *v1alpha1.AwsDynamodbTable) (result *v1alpha1.AwsDynamodbTable, err error) {
	result = &v1alpha1.AwsDynamodbTable{}
	err = c.client.Put().
		Resource("awsdynamodbtables").
		Name(awsDynamodbTable.Name).
		SubResource("status").
		Body(awsDynamodbTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDynamodbTable and deletes it. Returns an error if one occurs.
func (c *awsDynamodbTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdynamodbtables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDynamodbTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdynamodbtables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDynamodbTable.
func (c *awsDynamodbTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDynamodbTable, err error) {
	result = &v1alpha1.AwsDynamodbTable{}
	err = c.client.Patch(pt).
		Resource("awsdynamodbtables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
