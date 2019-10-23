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

// DmsEndpointsGetter has a method to return a DmsEndpointInterface.
// A group's client should implement this interface.
type DmsEndpointsGetter interface {
	DmsEndpoints(namespace string) DmsEndpointInterface
}

// DmsEndpointInterface has methods to work with DmsEndpoint resources.
type DmsEndpointInterface interface {
	Create(*v1alpha1.DmsEndpoint) (*v1alpha1.DmsEndpoint, error)
	Update(*v1alpha1.DmsEndpoint) (*v1alpha1.DmsEndpoint, error)
	UpdateStatus(*v1alpha1.DmsEndpoint) (*v1alpha1.DmsEndpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DmsEndpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.DmsEndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DmsEndpoint, err error)
	DmsEndpointExpansion
}

// dmsEndpoints implements DmsEndpointInterface
type dmsEndpoints struct {
	client rest.Interface
	ns     string
}

// newDmsEndpoints returns a DmsEndpoints
func newDmsEndpoints(c *AwsV1alpha1Client, namespace string) *dmsEndpoints {
	return &dmsEndpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dmsEndpoint, and returns the corresponding dmsEndpoint object, and an error if there is any.
func (c *dmsEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.DmsEndpoint, err error) {
	result = &v1alpha1.DmsEndpoint{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dmsendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DmsEndpoints that match those selectors.
func (c *dmsEndpoints) List(opts v1.ListOptions) (result *v1alpha1.DmsEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DmsEndpointList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dmsendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dmsEndpoints.
func (c *dmsEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dmsendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dmsEndpoint and creates it.  Returns the server's representation of the dmsEndpoint, and an error, if there is any.
func (c *dmsEndpoints) Create(dmsEndpoint *v1alpha1.DmsEndpoint) (result *v1alpha1.DmsEndpoint, err error) {
	result = &v1alpha1.DmsEndpoint{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dmsendpoints").
		Body(dmsEndpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dmsEndpoint and updates it. Returns the server's representation of the dmsEndpoint, and an error, if there is any.
func (c *dmsEndpoints) Update(dmsEndpoint *v1alpha1.DmsEndpoint) (result *v1alpha1.DmsEndpoint, err error) {
	result = &v1alpha1.DmsEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dmsendpoints").
		Name(dmsEndpoint.Name).
		Body(dmsEndpoint).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dmsEndpoints) UpdateStatus(dmsEndpoint *v1alpha1.DmsEndpoint) (result *v1alpha1.DmsEndpoint, err error) {
	result = &v1alpha1.DmsEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dmsendpoints").
		Name(dmsEndpoint.Name).
		SubResource("status").
		Body(dmsEndpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the dmsEndpoint and deletes it. Returns an error if one occurs.
func (c *dmsEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dmsendpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dmsEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dmsendpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dmsEndpoint.
func (c *dmsEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DmsEndpoint, err error) {
	result = &v1alpha1.DmsEndpoint{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dmsendpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
