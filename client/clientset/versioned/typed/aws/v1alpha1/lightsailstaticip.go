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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// LightsailStaticIpsGetter has a method to return a LightsailStaticIpInterface.
// A group's client should implement this interface.
type LightsailStaticIpsGetter interface {
	LightsailStaticIps() LightsailStaticIpInterface
}

// LightsailStaticIpInterface has methods to work with LightsailStaticIp resources.
type LightsailStaticIpInterface interface {
	Create(*v1alpha1.LightsailStaticIp) (*v1alpha1.LightsailStaticIp, error)
	Update(*v1alpha1.LightsailStaticIp) (*v1alpha1.LightsailStaticIp, error)
	UpdateStatus(*v1alpha1.LightsailStaticIp) (*v1alpha1.LightsailStaticIp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LightsailStaticIp, error)
	List(opts v1.ListOptions) (*v1alpha1.LightsailStaticIpList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailStaticIp, err error)
	LightsailStaticIpExpansion
}

// lightsailStaticIps implements LightsailStaticIpInterface
type lightsailStaticIps struct {
	client rest.Interface
}

// newLightsailStaticIps returns a LightsailStaticIps
func newLightsailStaticIps(c *AwsV1alpha1Client) *lightsailStaticIps {
	return &lightsailStaticIps{
		client: c.RESTClient(),
	}
}

// Get takes name of the lightsailStaticIp, and returns the corresponding lightsailStaticIp object, and an error if there is any.
func (c *lightsailStaticIps) Get(name string, options v1.GetOptions) (result *v1alpha1.LightsailStaticIp, err error) {
	result = &v1alpha1.LightsailStaticIp{}
	err = c.client.Get().
		Resource("lightsailstaticips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LightsailStaticIps that match those selectors.
func (c *lightsailStaticIps) List(opts v1.ListOptions) (result *v1alpha1.LightsailStaticIpList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LightsailStaticIpList{}
	err = c.client.Get().
		Resource("lightsailstaticips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lightsailStaticIps.
func (c *lightsailStaticIps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("lightsailstaticips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lightsailStaticIp and creates it.  Returns the server's representation of the lightsailStaticIp, and an error, if there is any.
func (c *lightsailStaticIps) Create(lightsailStaticIp *v1alpha1.LightsailStaticIp) (result *v1alpha1.LightsailStaticIp, err error) {
	result = &v1alpha1.LightsailStaticIp{}
	err = c.client.Post().
		Resource("lightsailstaticips").
		Body(lightsailStaticIp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lightsailStaticIp and updates it. Returns the server's representation of the lightsailStaticIp, and an error, if there is any.
func (c *lightsailStaticIps) Update(lightsailStaticIp *v1alpha1.LightsailStaticIp) (result *v1alpha1.LightsailStaticIp, err error) {
	result = &v1alpha1.LightsailStaticIp{}
	err = c.client.Put().
		Resource("lightsailstaticips").
		Name(lightsailStaticIp.Name).
		Body(lightsailStaticIp).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lightsailStaticIps) UpdateStatus(lightsailStaticIp *v1alpha1.LightsailStaticIp) (result *v1alpha1.LightsailStaticIp, err error) {
	result = &v1alpha1.LightsailStaticIp{}
	err = c.client.Put().
		Resource("lightsailstaticips").
		Name(lightsailStaticIp.Name).
		SubResource("status").
		Body(lightsailStaticIp).
		Do().
		Into(result)
	return
}

// Delete takes name of the lightsailStaticIp and deletes it. Returns an error if one occurs.
func (c *lightsailStaticIps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("lightsailstaticips").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lightsailStaticIps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("lightsailstaticips").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lightsailStaticIp.
func (c *lightsailStaticIps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailStaticIp, err error) {
	result = &v1alpha1.LightsailStaticIp{}
	err = c.client.Patch(pt).
		Resource("lightsailstaticips").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
