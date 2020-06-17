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

// Ec2TransitGatewaysGetter has a method to return a Ec2TransitGatewayInterface.
// A group's client should implement this interface.
type Ec2TransitGatewaysGetter interface {
	Ec2TransitGateways(namespace string) Ec2TransitGatewayInterface
}

// Ec2TransitGatewayInterface has methods to work with Ec2TransitGateway resources.
type Ec2TransitGatewayInterface interface {
	Create(*v1alpha1.Ec2TransitGateway) (*v1alpha1.Ec2TransitGateway, error)
	Update(*v1alpha1.Ec2TransitGateway) (*v1alpha1.Ec2TransitGateway, error)
	UpdateStatus(*v1alpha1.Ec2TransitGateway) (*v1alpha1.Ec2TransitGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Ec2TransitGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.Ec2TransitGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2TransitGateway, err error)
	Ec2TransitGatewayExpansion
}

// ec2TransitGateways implements Ec2TransitGatewayInterface
type ec2TransitGateways struct {
	client rest.Interface
	ns     string
}

// newEc2TransitGateways returns a Ec2TransitGateways
func newEc2TransitGateways(c *AwsV1alpha1Client, namespace string) *ec2TransitGateways {
	return &ec2TransitGateways{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ec2TransitGateway, and returns the corresponding ec2TransitGateway object, and an error if there is any.
func (c *ec2TransitGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.Ec2TransitGateway, err error) {
	result = &v1alpha1.Ec2TransitGateway{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ec2TransitGateways that match those selectors.
func (c *ec2TransitGateways) List(opts v1.ListOptions) (result *v1alpha1.Ec2TransitGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Ec2TransitGatewayList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ec2TransitGateways.
func (c *ec2TransitGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ec2TransitGateway and creates it.  Returns the server's representation of the ec2TransitGateway, and an error, if there is any.
func (c *ec2TransitGateways) Create(ec2TransitGateway *v1alpha1.Ec2TransitGateway) (result *v1alpha1.Ec2TransitGateway, err error) {
	result = &v1alpha1.Ec2TransitGateway{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		Body(ec2TransitGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ec2TransitGateway and updates it. Returns the server's representation of the ec2TransitGateway, and an error, if there is any.
func (c *ec2TransitGateways) Update(ec2TransitGateway *v1alpha1.Ec2TransitGateway) (result *v1alpha1.Ec2TransitGateway, err error) {
	result = &v1alpha1.Ec2TransitGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		Name(ec2TransitGateway.Name).
		Body(ec2TransitGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ec2TransitGateways) UpdateStatus(ec2TransitGateway *v1alpha1.Ec2TransitGateway) (result *v1alpha1.Ec2TransitGateway, err error) {
	result = &v1alpha1.Ec2TransitGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		Name(ec2TransitGateway.Name).
		SubResource("status").
		Body(ec2TransitGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the ec2TransitGateway and deletes it. Returns an error if one occurs.
func (c *ec2TransitGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ec2TransitGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ec2transitgateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ec2TransitGateway.
func (c *ec2TransitGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2TransitGateway, err error) {
	result = &v1alpha1.Ec2TransitGateway{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ec2transitgateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
