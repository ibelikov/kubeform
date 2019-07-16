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

// AthenaDatabasesGetter has a method to return a AthenaDatabaseInterface.
// A group's client should implement this interface.
type AthenaDatabasesGetter interface {
	AthenaDatabases() AthenaDatabaseInterface
}

// AthenaDatabaseInterface has methods to work with AthenaDatabase resources.
type AthenaDatabaseInterface interface {
	Create(*v1alpha1.AthenaDatabase) (*v1alpha1.AthenaDatabase, error)
	Update(*v1alpha1.AthenaDatabase) (*v1alpha1.AthenaDatabase, error)
	UpdateStatus(*v1alpha1.AthenaDatabase) (*v1alpha1.AthenaDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AthenaDatabase, error)
	List(opts v1.ListOptions) (*v1alpha1.AthenaDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AthenaDatabase, err error)
	AthenaDatabaseExpansion
}

// athenaDatabases implements AthenaDatabaseInterface
type athenaDatabases struct {
	client rest.Interface
}

// newAthenaDatabases returns a AthenaDatabases
func newAthenaDatabases(c *AwsV1alpha1Client) *athenaDatabases {
	return &athenaDatabases{
		client: c.RESTClient(),
	}
}

// Get takes name of the athenaDatabase, and returns the corresponding athenaDatabase object, and an error if there is any.
func (c *athenaDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.AthenaDatabase, err error) {
	result = &v1alpha1.AthenaDatabase{}
	err = c.client.Get().
		Resource("athenadatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AthenaDatabases that match those selectors.
func (c *athenaDatabases) List(opts v1.ListOptions) (result *v1alpha1.AthenaDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AthenaDatabaseList{}
	err = c.client.Get().
		Resource("athenadatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested athenaDatabases.
func (c *athenaDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("athenadatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a athenaDatabase and creates it.  Returns the server's representation of the athenaDatabase, and an error, if there is any.
func (c *athenaDatabases) Create(athenaDatabase *v1alpha1.AthenaDatabase) (result *v1alpha1.AthenaDatabase, err error) {
	result = &v1alpha1.AthenaDatabase{}
	err = c.client.Post().
		Resource("athenadatabases").
		Body(athenaDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a athenaDatabase and updates it. Returns the server's representation of the athenaDatabase, and an error, if there is any.
func (c *athenaDatabases) Update(athenaDatabase *v1alpha1.AthenaDatabase) (result *v1alpha1.AthenaDatabase, err error) {
	result = &v1alpha1.AthenaDatabase{}
	err = c.client.Put().
		Resource("athenadatabases").
		Name(athenaDatabase.Name).
		Body(athenaDatabase).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *athenaDatabases) UpdateStatus(athenaDatabase *v1alpha1.AthenaDatabase) (result *v1alpha1.AthenaDatabase, err error) {
	result = &v1alpha1.AthenaDatabase{}
	err = c.client.Put().
		Resource("athenadatabases").
		Name(athenaDatabase.Name).
		SubResource("status").
		Body(athenaDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the athenaDatabase and deletes it. Returns an error if one occurs.
func (c *athenaDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("athenadatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *athenaDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("athenadatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched athenaDatabase.
func (c *athenaDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AthenaDatabase, err error) {
	result = &v1alpha1.AthenaDatabase{}
	err = c.client.Patch(pt).
		Resource("athenadatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
