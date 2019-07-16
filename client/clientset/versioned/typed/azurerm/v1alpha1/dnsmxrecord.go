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

// DnsMxRecordsGetter has a method to return a DnsMxRecordInterface.
// A group's client should implement this interface.
type DnsMxRecordsGetter interface {
	DnsMxRecords() DnsMxRecordInterface
}

// DnsMxRecordInterface has methods to work with DnsMxRecord resources.
type DnsMxRecordInterface interface {
	Create(*v1alpha1.DnsMxRecord) (*v1alpha1.DnsMxRecord, error)
	Update(*v1alpha1.DnsMxRecord) (*v1alpha1.DnsMxRecord, error)
	UpdateStatus(*v1alpha1.DnsMxRecord) (*v1alpha1.DnsMxRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DnsMxRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.DnsMxRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsMxRecord, err error)
	DnsMxRecordExpansion
}

// dnsMxRecords implements DnsMxRecordInterface
type dnsMxRecords struct {
	client rest.Interface
}

// newDnsMxRecords returns a DnsMxRecords
func newDnsMxRecords(c *AzurermV1alpha1Client) *dnsMxRecords {
	return &dnsMxRecords{
		client: c.RESTClient(),
	}
}

// Get takes name of the dnsMxRecord, and returns the corresponding dnsMxRecord object, and an error if there is any.
func (c *dnsMxRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsMxRecord, err error) {
	result = &v1alpha1.DnsMxRecord{}
	err = c.client.Get().
		Resource("dnsmxrecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DnsMxRecords that match those selectors.
func (c *dnsMxRecords) List(opts v1.ListOptions) (result *v1alpha1.DnsMxRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DnsMxRecordList{}
	err = c.client.Get().
		Resource("dnsmxrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dnsMxRecords.
func (c *dnsMxRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("dnsmxrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dnsMxRecord and creates it.  Returns the server's representation of the dnsMxRecord, and an error, if there is any.
func (c *dnsMxRecords) Create(dnsMxRecord *v1alpha1.DnsMxRecord) (result *v1alpha1.DnsMxRecord, err error) {
	result = &v1alpha1.DnsMxRecord{}
	err = c.client.Post().
		Resource("dnsmxrecords").
		Body(dnsMxRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dnsMxRecord and updates it. Returns the server's representation of the dnsMxRecord, and an error, if there is any.
func (c *dnsMxRecords) Update(dnsMxRecord *v1alpha1.DnsMxRecord) (result *v1alpha1.DnsMxRecord, err error) {
	result = &v1alpha1.DnsMxRecord{}
	err = c.client.Put().
		Resource("dnsmxrecords").
		Name(dnsMxRecord.Name).
		Body(dnsMxRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dnsMxRecords) UpdateStatus(dnsMxRecord *v1alpha1.DnsMxRecord) (result *v1alpha1.DnsMxRecord, err error) {
	result = &v1alpha1.DnsMxRecord{}
	err = c.client.Put().
		Resource("dnsmxrecords").
		Name(dnsMxRecord.Name).
		SubResource("status").
		Body(dnsMxRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the dnsMxRecord and deletes it. Returns an error if one occurs.
func (c *dnsMxRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("dnsmxrecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dnsMxRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("dnsmxrecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dnsMxRecord.
func (c *dnsMxRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsMxRecord, err error) {
	result = &v1alpha1.DnsMxRecord{}
	err = c.client.Patch(pt).
		Resource("dnsmxrecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
