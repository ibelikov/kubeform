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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeRouterInterfacesGetter has a method to return a ComputeRouterInterfaceInterface.
// A group's client should implement this interface.
type ComputeRouterInterfacesGetter interface {
	ComputeRouterInterfaces() ComputeRouterInterfaceInterface
}

// ComputeRouterInterfaceInterface has methods to work with ComputeRouterInterface resources.
type ComputeRouterInterfaceInterface interface {
	Create(*v1alpha1.ComputeRouterInterface) (*v1alpha1.ComputeRouterInterface, error)
	Update(*v1alpha1.ComputeRouterInterface) (*v1alpha1.ComputeRouterInterface, error)
	UpdateStatus(*v1alpha1.ComputeRouterInterface) (*v1alpha1.ComputeRouterInterface, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeRouterInterface, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeRouterInterfaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouterInterface, err error)
	ComputeRouterInterfaceExpansion
}

// computeRouterInterfaces implements ComputeRouterInterfaceInterface
type computeRouterInterfaces struct {
	client rest.Interface
}

// newComputeRouterInterfaces returns a ComputeRouterInterfaces
func newComputeRouterInterfaces(c *GoogleV1alpha1Client) *computeRouterInterfaces {
	return &computeRouterInterfaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the computeRouterInterface, and returns the corresponding computeRouterInterface object, and an error if there is any.
func (c *computeRouterInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRouterInterface, err error) {
	result = &v1alpha1.ComputeRouterInterface{}
	err = c.client.Get().
		Resource("computerouterinterfaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeRouterInterfaces that match those selectors.
func (c *computeRouterInterfaces) List(opts v1.ListOptions) (result *v1alpha1.ComputeRouterInterfaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeRouterInterfaceList{}
	err = c.client.Get().
		Resource("computerouterinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeRouterInterfaces.
func (c *computeRouterInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("computerouterinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeRouterInterface and creates it.  Returns the server's representation of the computeRouterInterface, and an error, if there is any.
func (c *computeRouterInterfaces) Create(computeRouterInterface *v1alpha1.ComputeRouterInterface) (result *v1alpha1.ComputeRouterInterface, err error) {
	result = &v1alpha1.ComputeRouterInterface{}
	err = c.client.Post().
		Resource("computerouterinterfaces").
		Body(computeRouterInterface).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeRouterInterface and updates it. Returns the server's representation of the computeRouterInterface, and an error, if there is any.
func (c *computeRouterInterfaces) Update(computeRouterInterface *v1alpha1.ComputeRouterInterface) (result *v1alpha1.ComputeRouterInterface, err error) {
	result = &v1alpha1.ComputeRouterInterface{}
	err = c.client.Put().
		Resource("computerouterinterfaces").
		Name(computeRouterInterface.Name).
		Body(computeRouterInterface).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeRouterInterfaces) UpdateStatus(computeRouterInterface *v1alpha1.ComputeRouterInterface) (result *v1alpha1.ComputeRouterInterface, err error) {
	result = &v1alpha1.ComputeRouterInterface{}
	err = c.client.Put().
		Resource("computerouterinterfaces").
		Name(computeRouterInterface.Name).
		SubResource("status").
		Body(computeRouterInterface).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeRouterInterface and deletes it. Returns an error if one occurs.
func (c *computeRouterInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("computerouterinterfaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeRouterInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("computerouterinterfaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeRouterInterface.
func (c *computeRouterInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouterInterface, err error) {
	result = &v1alpha1.ComputeRouterInterface{}
	err = c.client.Patch(pt).
		Resource("computerouterinterfaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
