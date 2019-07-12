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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermCosmosdbCassandraKeyspacesGetter has a method to return a AzurermCosmosdbCassandraKeyspaceInterface.
// A group's client should implement this interface.
type AzurermCosmosdbCassandraKeyspacesGetter interface {
	AzurermCosmosdbCassandraKeyspaces() AzurermCosmosdbCassandraKeyspaceInterface
}

// AzurermCosmosdbCassandraKeyspaceInterface has methods to work with AzurermCosmosdbCassandraKeyspace resources.
type AzurermCosmosdbCassandraKeyspaceInterface interface {
	Create(*v1alpha1.AzurermCosmosdbCassandraKeyspace) (*v1alpha1.AzurermCosmosdbCassandraKeyspace, error)
	Update(*v1alpha1.AzurermCosmosdbCassandraKeyspace) (*v1alpha1.AzurermCosmosdbCassandraKeyspace, error)
	UpdateStatus(*v1alpha1.AzurermCosmosdbCassandraKeyspace) (*v1alpha1.AzurermCosmosdbCassandraKeyspace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermCosmosdbCassandraKeyspace, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermCosmosdbCassandraKeyspaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCosmosdbCassandraKeyspace, err error)
	AzurermCosmosdbCassandraKeyspaceExpansion
}

// azurermCosmosdbCassandraKeyspaces implements AzurermCosmosdbCassandraKeyspaceInterface
type azurermCosmosdbCassandraKeyspaces struct {
	client rest.Interface
}

// newAzurermCosmosdbCassandraKeyspaces returns a AzurermCosmosdbCassandraKeyspaces
func newAzurermCosmosdbCassandraKeyspaces(c *AzurermV1alpha1Client) *azurermCosmosdbCassandraKeyspaces {
	return &azurermCosmosdbCassandraKeyspaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermCosmosdbCassandraKeyspace, and returns the corresponding azurermCosmosdbCassandraKeyspace object, and an error if there is any.
func (c *azurermCosmosdbCassandraKeyspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermCosmosdbCassandraKeyspace, err error) {
	result = &v1alpha1.AzurermCosmosdbCassandraKeyspace{}
	err = c.client.Get().
		Resource("azurermcosmosdbcassandrakeyspaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermCosmosdbCassandraKeyspaces that match those selectors.
func (c *azurermCosmosdbCassandraKeyspaces) List(opts v1.ListOptions) (result *v1alpha1.AzurermCosmosdbCassandraKeyspaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermCosmosdbCassandraKeyspaceList{}
	err = c.client.Get().
		Resource("azurermcosmosdbcassandrakeyspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermCosmosdbCassandraKeyspaces.
func (c *azurermCosmosdbCassandraKeyspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermcosmosdbcassandrakeyspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermCosmosdbCassandraKeyspace and creates it.  Returns the server's representation of the azurermCosmosdbCassandraKeyspace, and an error, if there is any.
func (c *azurermCosmosdbCassandraKeyspaces) Create(azurermCosmosdbCassandraKeyspace *v1alpha1.AzurermCosmosdbCassandraKeyspace) (result *v1alpha1.AzurermCosmosdbCassandraKeyspace, err error) {
	result = &v1alpha1.AzurermCosmosdbCassandraKeyspace{}
	err = c.client.Post().
		Resource("azurermcosmosdbcassandrakeyspaces").
		Body(azurermCosmosdbCassandraKeyspace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermCosmosdbCassandraKeyspace and updates it. Returns the server's representation of the azurermCosmosdbCassandraKeyspace, and an error, if there is any.
func (c *azurermCosmosdbCassandraKeyspaces) Update(azurermCosmosdbCassandraKeyspace *v1alpha1.AzurermCosmosdbCassandraKeyspace) (result *v1alpha1.AzurermCosmosdbCassandraKeyspace, err error) {
	result = &v1alpha1.AzurermCosmosdbCassandraKeyspace{}
	err = c.client.Put().
		Resource("azurermcosmosdbcassandrakeyspaces").
		Name(azurermCosmosdbCassandraKeyspace.Name).
		Body(azurermCosmosdbCassandraKeyspace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermCosmosdbCassandraKeyspaces) UpdateStatus(azurermCosmosdbCassandraKeyspace *v1alpha1.AzurermCosmosdbCassandraKeyspace) (result *v1alpha1.AzurermCosmosdbCassandraKeyspace, err error) {
	result = &v1alpha1.AzurermCosmosdbCassandraKeyspace{}
	err = c.client.Put().
		Resource("azurermcosmosdbcassandrakeyspaces").
		Name(azurermCosmosdbCassandraKeyspace.Name).
		SubResource("status").
		Body(azurermCosmosdbCassandraKeyspace).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermCosmosdbCassandraKeyspace and deletes it. Returns an error if one occurs.
func (c *azurermCosmosdbCassandraKeyspaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermcosmosdbcassandrakeyspaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermCosmosdbCassandraKeyspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermcosmosdbcassandrakeyspaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermCosmosdbCassandraKeyspace.
func (c *azurermCosmosdbCassandraKeyspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermCosmosdbCassandraKeyspace, err error) {
	result = &v1alpha1.AzurermCosmosdbCassandraKeyspace{}
	err = c.client.Patch(pt).
		Resource("azurermcosmosdbcassandrakeyspaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
