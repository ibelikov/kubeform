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

// AzurermRecoveryServicesProtectedVmsGetter has a method to return a AzurermRecoveryServicesProtectedVmInterface.
// A group's client should implement this interface.
type AzurermRecoveryServicesProtectedVmsGetter interface {
	AzurermRecoveryServicesProtectedVms() AzurermRecoveryServicesProtectedVmInterface
}

// AzurermRecoveryServicesProtectedVmInterface has methods to work with AzurermRecoveryServicesProtectedVm resources.
type AzurermRecoveryServicesProtectedVmInterface interface {
	Create(*v1alpha1.AzurermRecoveryServicesProtectedVm) (*v1alpha1.AzurermRecoveryServicesProtectedVm, error)
	Update(*v1alpha1.AzurermRecoveryServicesProtectedVm) (*v1alpha1.AzurermRecoveryServicesProtectedVm, error)
	UpdateStatus(*v1alpha1.AzurermRecoveryServicesProtectedVm) (*v1alpha1.AzurermRecoveryServicesProtectedVm, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermRecoveryServicesProtectedVm, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermRecoveryServicesProtectedVmList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermRecoveryServicesProtectedVm, err error)
	AzurermRecoveryServicesProtectedVmExpansion
}

// azurermRecoveryServicesProtectedVms implements AzurermRecoveryServicesProtectedVmInterface
type azurermRecoveryServicesProtectedVms struct {
	client rest.Interface
}

// newAzurermRecoveryServicesProtectedVms returns a AzurermRecoveryServicesProtectedVms
func newAzurermRecoveryServicesProtectedVms(c *AzurermV1alpha1Client) *azurermRecoveryServicesProtectedVms {
	return &azurermRecoveryServicesProtectedVms{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermRecoveryServicesProtectedVm, and returns the corresponding azurermRecoveryServicesProtectedVm object, and an error if there is any.
func (c *azurermRecoveryServicesProtectedVms) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermRecoveryServicesProtectedVm, err error) {
	result = &v1alpha1.AzurermRecoveryServicesProtectedVm{}
	err = c.client.Get().
		Resource("azurermrecoveryservicesprotectedvms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermRecoveryServicesProtectedVms that match those selectors.
func (c *azurermRecoveryServicesProtectedVms) List(opts v1.ListOptions) (result *v1alpha1.AzurermRecoveryServicesProtectedVmList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermRecoveryServicesProtectedVmList{}
	err = c.client.Get().
		Resource("azurermrecoveryservicesprotectedvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermRecoveryServicesProtectedVms.
func (c *azurermRecoveryServicesProtectedVms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermrecoveryservicesprotectedvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermRecoveryServicesProtectedVm and creates it.  Returns the server's representation of the azurermRecoveryServicesProtectedVm, and an error, if there is any.
func (c *azurermRecoveryServicesProtectedVms) Create(azurermRecoveryServicesProtectedVm *v1alpha1.AzurermRecoveryServicesProtectedVm) (result *v1alpha1.AzurermRecoveryServicesProtectedVm, err error) {
	result = &v1alpha1.AzurermRecoveryServicesProtectedVm{}
	err = c.client.Post().
		Resource("azurermrecoveryservicesprotectedvms").
		Body(azurermRecoveryServicesProtectedVm).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermRecoveryServicesProtectedVm and updates it. Returns the server's representation of the azurermRecoveryServicesProtectedVm, and an error, if there is any.
func (c *azurermRecoveryServicesProtectedVms) Update(azurermRecoveryServicesProtectedVm *v1alpha1.AzurermRecoveryServicesProtectedVm) (result *v1alpha1.AzurermRecoveryServicesProtectedVm, err error) {
	result = &v1alpha1.AzurermRecoveryServicesProtectedVm{}
	err = c.client.Put().
		Resource("azurermrecoveryservicesprotectedvms").
		Name(azurermRecoveryServicesProtectedVm.Name).
		Body(azurermRecoveryServicesProtectedVm).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermRecoveryServicesProtectedVms) UpdateStatus(azurermRecoveryServicesProtectedVm *v1alpha1.AzurermRecoveryServicesProtectedVm) (result *v1alpha1.AzurermRecoveryServicesProtectedVm, err error) {
	result = &v1alpha1.AzurermRecoveryServicesProtectedVm{}
	err = c.client.Put().
		Resource("azurermrecoveryservicesprotectedvms").
		Name(azurermRecoveryServicesProtectedVm.Name).
		SubResource("status").
		Body(azurermRecoveryServicesProtectedVm).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermRecoveryServicesProtectedVm and deletes it. Returns an error if one occurs.
func (c *azurermRecoveryServicesProtectedVms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermrecoveryservicesprotectedvms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermRecoveryServicesProtectedVms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermrecoveryservicesprotectedvms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermRecoveryServicesProtectedVm.
func (c *azurermRecoveryServicesProtectedVms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermRecoveryServicesProtectedVm, err error) {
	result = &v1alpha1.AzurermRecoveryServicesProtectedVm{}
	err = c.client.Patch(pt).
		Resource("azurermrecoveryservicesprotectedvms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
