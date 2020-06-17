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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// SubnetsGetter has a method to return a SubnetInterface.
// A group's client should implement this interface.
type SubnetsGetter interface {
	Subnets(namespace string) SubnetInterface
}

// SubnetInterface has methods to work with Subnet resources.
type SubnetInterface interface {
	Create(*v1alpha1.Subnet) (*v1alpha1.Subnet, error)
	Update(*v1alpha1.Subnet) (*v1alpha1.Subnet, error)
	UpdateStatus(*v1alpha1.Subnet) (*v1alpha1.Subnet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Subnet, error)
	List(opts v1.ListOptions) (*v1alpha1.SubnetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Subnet, err error)
	SubnetExpansion
}

// subnets implements SubnetInterface
type subnets struct {
	client rest.Interface
	ns     string
}

// newSubnets returns a Subnets
func newSubnets(c *AwsV1alpha1Client, namespace string) *subnets {
	return &subnets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the subnet, and returns the corresponding subnet object, and an error if there is any.
func (c *subnets) Get(name string, options v1.GetOptions) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Subnets that match those selectors.
func (c *subnets) List(opts v1.ListOptions) (result *v1alpha1.SubnetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SubnetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested subnets.
func (c *subnets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("subnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a subnet and creates it.  Returns the server's representation of the subnet, and an error, if there is any.
func (c *subnets) Create(subnet *v1alpha1.Subnet) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("subnets").
		Body(subnet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a subnet and updates it. Returns the server's representation of the subnet, and an error, if there is any.
func (c *subnets) Update(subnet *v1alpha1.Subnet) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnets").
		Name(subnet.Name).
		Body(subnet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *subnets) UpdateStatus(subnet *v1alpha1.Subnet) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnets").
		Name(subnet.Name).
		SubResource("status").
		Body(subnet).
		Do().
		Into(result)
	return
}

// Delete takes name of the subnet and deletes it. Returns an error if one occurs.
func (c *subnets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *subnets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched subnet.
func (c *subnets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("subnets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
