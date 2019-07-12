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

// AzurermDnsCnameRecordsGetter has a method to return a AzurermDnsCnameRecordInterface.
// A group's client should implement this interface.
type AzurermDnsCnameRecordsGetter interface {
	AzurermDnsCnameRecords() AzurermDnsCnameRecordInterface
}

// AzurermDnsCnameRecordInterface has methods to work with AzurermDnsCnameRecord resources.
type AzurermDnsCnameRecordInterface interface {
	Create(*v1alpha1.AzurermDnsCnameRecord) (*v1alpha1.AzurermDnsCnameRecord, error)
	Update(*v1alpha1.AzurermDnsCnameRecord) (*v1alpha1.AzurermDnsCnameRecord, error)
	UpdateStatus(*v1alpha1.AzurermDnsCnameRecord) (*v1alpha1.AzurermDnsCnameRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDnsCnameRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDnsCnameRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsCnameRecord, err error)
	AzurermDnsCnameRecordExpansion
}

// azurermDnsCnameRecords implements AzurermDnsCnameRecordInterface
type azurermDnsCnameRecords struct {
	client rest.Interface
}

// newAzurermDnsCnameRecords returns a AzurermDnsCnameRecords
func newAzurermDnsCnameRecords(c *AzurermV1alpha1Client) *azurermDnsCnameRecords {
	return &azurermDnsCnameRecords{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDnsCnameRecord, and returns the corresponding azurermDnsCnameRecord object, and an error if there is any.
func (c *azurermDnsCnameRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsCnameRecord, err error) {
	result = &v1alpha1.AzurermDnsCnameRecord{}
	err = c.client.Get().
		Resource("azurermdnscnamerecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDnsCnameRecords that match those selectors.
func (c *azurermDnsCnameRecords) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsCnameRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDnsCnameRecordList{}
	err = c.client.Get().
		Resource("azurermdnscnamerecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDnsCnameRecords.
func (c *azurermDnsCnameRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdnscnamerecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDnsCnameRecord and creates it.  Returns the server's representation of the azurermDnsCnameRecord, and an error, if there is any.
func (c *azurermDnsCnameRecords) Create(azurermDnsCnameRecord *v1alpha1.AzurermDnsCnameRecord) (result *v1alpha1.AzurermDnsCnameRecord, err error) {
	result = &v1alpha1.AzurermDnsCnameRecord{}
	err = c.client.Post().
		Resource("azurermdnscnamerecords").
		Body(azurermDnsCnameRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDnsCnameRecord and updates it. Returns the server's representation of the azurermDnsCnameRecord, and an error, if there is any.
func (c *azurermDnsCnameRecords) Update(azurermDnsCnameRecord *v1alpha1.AzurermDnsCnameRecord) (result *v1alpha1.AzurermDnsCnameRecord, err error) {
	result = &v1alpha1.AzurermDnsCnameRecord{}
	err = c.client.Put().
		Resource("azurermdnscnamerecords").
		Name(azurermDnsCnameRecord.Name).
		Body(azurermDnsCnameRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDnsCnameRecords) UpdateStatus(azurermDnsCnameRecord *v1alpha1.AzurermDnsCnameRecord) (result *v1alpha1.AzurermDnsCnameRecord, err error) {
	result = &v1alpha1.AzurermDnsCnameRecord{}
	err = c.client.Put().
		Resource("azurermdnscnamerecords").
		Name(azurermDnsCnameRecord.Name).
		SubResource("status").
		Body(azurermDnsCnameRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDnsCnameRecord and deletes it. Returns an error if one occurs.
func (c *azurermDnsCnameRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdnscnamerecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDnsCnameRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdnscnamerecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDnsCnameRecord.
func (c *azurermDnsCnameRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsCnameRecord, err error) {
	result = &v1alpha1.AzurermDnsCnameRecord{}
	err = c.client.Patch(pt).
		Resource("azurermdnscnamerecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
