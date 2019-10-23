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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DynamodbTablesGetter has a method to return a DynamodbTableInterface.
// A group's client should implement this interface.
type DynamodbTablesGetter interface {
	DynamodbTables(namespace string) DynamodbTableInterface
}

// DynamodbTableInterface has methods to work with DynamodbTable resources.
type DynamodbTableInterface interface {
	Create(*v1alpha1.DynamodbTable) (*v1alpha1.DynamodbTable, error)
	Update(*v1alpha1.DynamodbTable) (*v1alpha1.DynamodbTable, error)
	UpdateStatus(*v1alpha1.DynamodbTable) (*v1alpha1.DynamodbTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DynamodbTable, error)
	List(opts v1.ListOptions) (*v1alpha1.DynamodbTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DynamodbTable, err error)
	DynamodbTableExpansion
}

// dynamodbTables implements DynamodbTableInterface
type dynamodbTables struct {
	client rest.Interface
	ns     string
}

// newDynamodbTables returns a DynamodbTables
func newDynamodbTables(c *AwsV1alpha1Client, namespace string) *dynamodbTables {
	return &dynamodbTables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dynamodbTable, and returns the corresponding dynamodbTable object, and an error if there is any.
func (c *dynamodbTables) Get(name string, options v1.GetOptions) (result *v1alpha1.DynamodbTable, err error) {
	result = &v1alpha1.DynamodbTable{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dynamodbtables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DynamodbTables that match those selectors.
func (c *dynamodbTables) List(opts v1.ListOptions) (result *v1alpha1.DynamodbTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DynamodbTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dynamodbtables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dynamodbTables.
func (c *dynamodbTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dynamodbtables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dynamodbTable and creates it.  Returns the server's representation of the dynamodbTable, and an error, if there is any.
func (c *dynamodbTables) Create(dynamodbTable *v1alpha1.DynamodbTable) (result *v1alpha1.DynamodbTable, err error) {
	result = &v1alpha1.DynamodbTable{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dynamodbtables").
		Body(dynamodbTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dynamodbTable and updates it. Returns the server's representation of the dynamodbTable, and an error, if there is any.
func (c *dynamodbTables) Update(dynamodbTable *v1alpha1.DynamodbTable) (result *v1alpha1.DynamodbTable, err error) {
	result = &v1alpha1.DynamodbTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dynamodbtables").
		Name(dynamodbTable.Name).
		Body(dynamodbTable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dynamodbTables) UpdateStatus(dynamodbTable *v1alpha1.DynamodbTable) (result *v1alpha1.DynamodbTable, err error) {
	result = &v1alpha1.DynamodbTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dynamodbtables").
		Name(dynamodbTable.Name).
		SubResource("status").
		Body(dynamodbTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the dynamodbTable and deletes it. Returns an error if one occurs.
func (c *dynamodbTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dynamodbtables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dynamodbTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dynamodbtables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dynamodbTable.
func (c *dynamodbTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DynamodbTable, err error) {
	result = &v1alpha1.DynamodbTable{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dynamodbtables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
