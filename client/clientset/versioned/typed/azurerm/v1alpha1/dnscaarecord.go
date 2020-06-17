/*
Copyright The Kubeform Authors.

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

// DnsCaaRecordsGetter has a method to return a DnsCaaRecordInterface.
// A group's client should implement this interface.
type DnsCaaRecordsGetter interface {
	DnsCaaRecords(namespace string) DnsCaaRecordInterface
}

// DnsCaaRecordInterface has methods to work with DnsCaaRecord resources.
type DnsCaaRecordInterface interface {
	Create(*v1alpha1.DnsCaaRecord) (*v1alpha1.DnsCaaRecord, error)
	Update(*v1alpha1.DnsCaaRecord) (*v1alpha1.DnsCaaRecord, error)
	UpdateStatus(*v1alpha1.DnsCaaRecord) (*v1alpha1.DnsCaaRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DnsCaaRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.DnsCaaRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsCaaRecord, err error)
	DnsCaaRecordExpansion
}

// dnsCaaRecords implements DnsCaaRecordInterface
type dnsCaaRecords struct {
	client rest.Interface
	ns     string
}

// newDnsCaaRecords returns a DnsCaaRecords
func newDnsCaaRecords(c *AzurermV1alpha1Client, namespace string) *dnsCaaRecords {
	return &dnsCaaRecords{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dnsCaaRecord, and returns the corresponding dnsCaaRecord object, and an error if there is any.
func (c *dnsCaaRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsCaaRecord, err error) {
	result = &v1alpha1.DnsCaaRecord{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnscaarecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DnsCaaRecords that match those selectors.
func (c *dnsCaaRecords) List(opts v1.ListOptions) (result *v1alpha1.DnsCaaRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DnsCaaRecordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnscaarecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dnsCaaRecords.
func (c *dnsCaaRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnscaarecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dnsCaaRecord and creates it.  Returns the server's representation of the dnsCaaRecord, and an error, if there is any.
func (c *dnsCaaRecords) Create(dnsCaaRecord *v1alpha1.DnsCaaRecord) (result *v1alpha1.DnsCaaRecord, err error) {
	result = &v1alpha1.DnsCaaRecord{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnscaarecords").
		Body(dnsCaaRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dnsCaaRecord and updates it. Returns the server's representation of the dnsCaaRecord, and an error, if there is any.
func (c *dnsCaaRecords) Update(dnsCaaRecord *v1alpha1.DnsCaaRecord) (result *v1alpha1.DnsCaaRecord, err error) {
	result = &v1alpha1.DnsCaaRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnscaarecords").
		Name(dnsCaaRecord.Name).
		Body(dnsCaaRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dnsCaaRecords) UpdateStatus(dnsCaaRecord *v1alpha1.DnsCaaRecord) (result *v1alpha1.DnsCaaRecord, err error) {
	result = &v1alpha1.DnsCaaRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnscaarecords").
		Name(dnsCaaRecord.Name).
		SubResource("status").
		Body(dnsCaaRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the dnsCaaRecord and deletes it. Returns an error if one occurs.
func (c *dnsCaaRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnscaarecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dnsCaaRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnscaarecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dnsCaaRecord.
func (c *dnsCaaRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsCaaRecord, err error) {
	result = &v1alpha1.DnsCaaRecord{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnscaarecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
