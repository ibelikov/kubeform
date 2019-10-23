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

// SagemakerNotebookInstanceLifecycleConfigurationsGetter has a method to return a SagemakerNotebookInstanceLifecycleConfigurationInterface.
// A group's client should implement this interface.
type SagemakerNotebookInstanceLifecycleConfigurationsGetter interface {
	SagemakerNotebookInstanceLifecycleConfigurations(namespace string) SagemakerNotebookInstanceLifecycleConfigurationInterface
}

// SagemakerNotebookInstanceLifecycleConfigurationInterface has methods to work with SagemakerNotebookInstanceLifecycleConfiguration resources.
type SagemakerNotebookInstanceLifecycleConfigurationInterface interface {
	Create(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, error)
	Update(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, error)
	UpdateStatus(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error)
	SagemakerNotebookInstanceLifecycleConfigurationExpansion
}

// sagemakerNotebookInstanceLifecycleConfigurations implements SagemakerNotebookInstanceLifecycleConfigurationInterface
type sagemakerNotebookInstanceLifecycleConfigurations struct {
	client rest.Interface
	ns     string
}

// newSagemakerNotebookInstanceLifecycleConfigurations returns a SagemakerNotebookInstanceLifecycleConfigurations
func newSagemakerNotebookInstanceLifecycleConfigurations(c *AwsV1alpha1Client, namespace string) *sagemakerNotebookInstanceLifecycleConfigurations {
	return &sagemakerNotebookInstanceLifecycleConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sagemakerNotebookInstanceLifecycleConfiguration, and returns the corresponding sagemakerNotebookInstanceLifecycleConfiguration object, and an error if there is any.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	result = &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SagemakerNotebookInstanceLifecycleConfigurations that match those selectors.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) List(opts v1.ListOptions) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sagemakerNotebookInstanceLifecycleConfigurations.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sagemakerNotebookInstanceLifecycleConfiguration and creates it.  Returns the server's representation of the sagemakerNotebookInstanceLifecycleConfiguration, and an error, if there is any.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) Create(sagemakerNotebookInstanceLifecycleConfiguration *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	result = &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		Body(sagemakerNotebookInstanceLifecycleConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sagemakerNotebookInstanceLifecycleConfiguration and updates it. Returns the server's representation of the sagemakerNotebookInstanceLifecycleConfiguration, and an error, if there is any.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) Update(sagemakerNotebookInstanceLifecycleConfiguration *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	result = &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		Name(sagemakerNotebookInstanceLifecycleConfiguration.Name).
		Body(sagemakerNotebookInstanceLifecycleConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sagemakerNotebookInstanceLifecycleConfigurations) UpdateStatus(sagemakerNotebookInstanceLifecycleConfiguration *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	result = &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		Name(sagemakerNotebookInstanceLifecycleConfiguration.Name).
		SubResource("status").
		Body(sagemakerNotebookInstanceLifecycleConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the sagemakerNotebookInstanceLifecycleConfiguration and deletes it. Returns an error if one occurs.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sagemakerNotebookInstanceLifecycleConfiguration.
func (c *sagemakerNotebookInstanceLifecycleConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	result = &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sagemakernotebookinstancelifecycleconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
