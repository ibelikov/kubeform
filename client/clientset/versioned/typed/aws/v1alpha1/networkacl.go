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

// NetworkACLsGetter has a method to return a NetworkACLInterface.
// A group's client should implement this interface.
type NetworkACLsGetter interface {
	NetworkACLs(namespace string) NetworkACLInterface
}

// NetworkACLInterface has methods to work with NetworkACL resources.
type NetworkACLInterface interface {
	Create(*v1alpha1.NetworkACL) (*v1alpha1.NetworkACL, error)
	Update(*v1alpha1.NetworkACL) (*v1alpha1.NetworkACL, error)
	UpdateStatus(*v1alpha1.NetworkACL) (*v1alpha1.NetworkACL, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkACL, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkACLList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkACL, err error)
	NetworkACLExpansion
}

// networkACLs implements NetworkACLInterface
type networkACLs struct {
	client rest.Interface
	ns     string
}

// newNetworkACLs returns a NetworkACLs
func newNetworkACLs(c *AwsV1alpha1Client, namespace string) *networkACLs {
	return &networkACLs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkACL, and returns the corresponding networkACL object, and an error if there is any.
func (c *networkACLs) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkACL, err error) {
	result = &v1alpha1.NetworkACL{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkACLs that match those selectors.
func (c *networkACLs) List(opts v1.ListOptions) (result *v1alpha1.NetworkACLList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkACLList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkACLs.
func (c *networkACLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkACL and creates it.  Returns the server's representation of the networkACL, and an error, if there is any.
func (c *networkACLs) Create(networkACL *v1alpha1.NetworkACL) (result *v1alpha1.NetworkACL, err error) {
	result = &v1alpha1.NetworkACL{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkacls").
		Body(networkACL).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkACL and updates it. Returns the server's representation of the networkACL, and an error, if there is any.
func (c *networkACLs) Update(networkACL *v1alpha1.NetworkACL) (result *v1alpha1.NetworkACL, err error) {
	result = &v1alpha1.NetworkACL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkacls").
		Name(networkACL.Name).
		Body(networkACL).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkACLs) UpdateStatus(networkACL *v1alpha1.NetworkACL) (result *v1alpha1.NetworkACL, err error) {
	result = &v1alpha1.NetworkACL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkacls").
		Name(networkACL.Name).
		SubResource("status").
		Body(networkACL).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkACL and deletes it. Returns an error if one occurs.
func (c *networkACLs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkACLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkACL.
func (c *networkACLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkACL, err error) {
	result = &v1alpha1.NetworkACL{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
