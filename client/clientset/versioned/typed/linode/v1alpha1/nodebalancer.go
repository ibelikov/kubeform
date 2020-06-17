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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// NodebalancersGetter has a method to return a NodebalancerInterface.
// A group's client should implement this interface.
type NodebalancersGetter interface {
	Nodebalancers(namespace string) NodebalancerInterface
}

// NodebalancerInterface has methods to work with Nodebalancer resources.
type NodebalancerInterface interface {
	Create(*v1alpha1.Nodebalancer) (*v1alpha1.Nodebalancer, error)
	Update(*v1alpha1.Nodebalancer) (*v1alpha1.Nodebalancer, error)
	UpdateStatus(*v1alpha1.Nodebalancer) (*v1alpha1.Nodebalancer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Nodebalancer, error)
	List(opts v1.ListOptions) (*v1alpha1.NodebalancerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Nodebalancer, err error)
	NodebalancerExpansion
}

// nodebalancers implements NodebalancerInterface
type nodebalancers struct {
	client rest.Interface
	ns     string
}

// newNodebalancers returns a Nodebalancers
func newNodebalancers(c *LinodeV1alpha1Client, namespace string) *nodebalancers {
	return &nodebalancers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodebalancer, and returns the corresponding nodebalancer object, and an error if there is any.
func (c *nodebalancers) Get(name string, options v1.GetOptions) (result *v1alpha1.Nodebalancer, err error) {
	result = &v1alpha1.Nodebalancer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodebalancers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Nodebalancers that match those selectors.
func (c *nodebalancers) List(opts v1.ListOptions) (result *v1alpha1.NodebalancerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodebalancerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodebalancers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodebalancers.
func (c *nodebalancers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodebalancers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nodebalancer and creates it.  Returns the server's representation of the nodebalancer, and an error, if there is any.
func (c *nodebalancers) Create(nodebalancer *v1alpha1.Nodebalancer) (result *v1alpha1.Nodebalancer, err error) {
	result = &v1alpha1.Nodebalancer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodebalancers").
		Body(nodebalancer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodebalancer and updates it. Returns the server's representation of the nodebalancer, and an error, if there is any.
func (c *nodebalancers) Update(nodebalancer *v1alpha1.Nodebalancer) (result *v1alpha1.Nodebalancer, err error) {
	result = &v1alpha1.Nodebalancer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodebalancers").
		Name(nodebalancer.Name).
		Body(nodebalancer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodebalancers) UpdateStatus(nodebalancer *v1alpha1.Nodebalancer) (result *v1alpha1.Nodebalancer, err error) {
	result = &v1alpha1.Nodebalancer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodebalancers").
		Name(nodebalancer.Name).
		SubResource("status").
		Body(nodebalancer).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodebalancer and deletes it. Returns an error if one occurs.
func (c *nodebalancers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodebalancers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodebalancers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodebalancers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodebalancer.
func (c *nodebalancers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Nodebalancer, err error) {
	result = &v1alpha1.Nodebalancer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodebalancers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
