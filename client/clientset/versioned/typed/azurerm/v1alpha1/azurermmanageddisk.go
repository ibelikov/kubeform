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

// AzurermManagedDisksGetter has a method to return a AzurermManagedDiskInterface.
// A group's client should implement this interface.
type AzurermManagedDisksGetter interface {
	AzurermManagedDisks() AzurermManagedDiskInterface
}

// AzurermManagedDiskInterface has methods to work with AzurermManagedDisk resources.
type AzurermManagedDiskInterface interface {
	Create(*v1alpha1.AzurermManagedDisk) (*v1alpha1.AzurermManagedDisk, error)
	Update(*v1alpha1.AzurermManagedDisk) (*v1alpha1.AzurermManagedDisk, error)
	UpdateStatus(*v1alpha1.AzurermManagedDisk) (*v1alpha1.AzurermManagedDisk, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermManagedDisk, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermManagedDiskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermManagedDisk, err error)
	AzurermManagedDiskExpansion
}

// azurermManagedDisks implements AzurermManagedDiskInterface
type azurermManagedDisks struct {
	client rest.Interface
}

// newAzurermManagedDisks returns a AzurermManagedDisks
func newAzurermManagedDisks(c *AzurermV1alpha1Client) *azurermManagedDisks {
	return &azurermManagedDisks{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermManagedDisk, and returns the corresponding azurermManagedDisk object, and an error if there is any.
func (c *azurermManagedDisks) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermManagedDisk, err error) {
	result = &v1alpha1.AzurermManagedDisk{}
	err = c.client.Get().
		Resource("azurermmanageddisks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermManagedDisks that match those selectors.
func (c *azurermManagedDisks) List(opts v1.ListOptions) (result *v1alpha1.AzurermManagedDiskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermManagedDiskList{}
	err = c.client.Get().
		Resource("azurermmanageddisks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermManagedDisks.
func (c *azurermManagedDisks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermmanageddisks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermManagedDisk and creates it.  Returns the server's representation of the azurermManagedDisk, and an error, if there is any.
func (c *azurermManagedDisks) Create(azurermManagedDisk *v1alpha1.AzurermManagedDisk) (result *v1alpha1.AzurermManagedDisk, err error) {
	result = &v1alpha1.AzurermManagedDisk{}
	err = c.client.Post().
		Resource("azurermmanageddisks").
		Body(azurermManagedDisk).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermManagedDisk and updates it. Returns the server's representation of the azurermManagedDisk, and an error, if there is any.
func (c *azurermManagedDisks) Update(azurermManagedDisk *v1alpha1.AzurermManagedDisk) (result *v1alpha1.AzurermManagedDisk, err error) {
	result = &v1alpha1.AzurermManagedDisk{}
	err = c.client.Put().
		Resource("azurermmanageddisks").
		Name(azurermManagedDisk.Name).
		Body(azurermManagedDisk).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermManagedDisks) UpdateStatus(azurermManagedDisk *v1alpha1.AzurermManagedDisk) (result *v1alpha1.AzurermManagedDisk, err error) {
	result = &v1alpha1.AzurermManagedDisk{}
	err = c.client.Put().
		Resource("azurermmanageddisks").
		Name(azurermManagedDisk.Name).
		SubResource("status").
		Body(azurermManagedDisk).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermManagedDisk and deletes it. Returns an error if one occurs.
func (c *azurermManagedDisks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermmanageddisks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermManagedDisks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermmanageddisks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermManagedDisk.
func (c *azurermManagedDisks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermManagedDisk, err error) {
	result = &v1alpha1.AzurermManagedDisk{}
	err = c.client.Patch(pt).
		Resource("azurermmanageddisks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
