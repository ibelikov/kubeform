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

// PacketCapturesGetter has a method to return a PacketCaptureInterface.
// A group's client should implement this interface.
type PacketCapturesGetter interface {
	PacketCaptures() PacketCaptureInterface
}

// PacketCaptureInterface has methods to work with PacketCapture resources.
type PacketCaptureInterface interface {
	Create(*v1alpha1.PacketCapture) (*v1alpha1.PacketCapture, error)
	Update(*v1alpha1.PacketCapture) (*v1alpha1.PacketCapture, error)
	UpdateStatus(*v1alpha1.PacketCapture) (*v1alpha1.PacketCapture, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PacketCapture, error)
	List(opts v1.ListOptions) (*v1alpha1.PacketCaptureList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PacketCapture, err error)
	PacketCaptureExpansion
}

// packetCaptures implements PacketCaptureInterface
type packetCaptures struct {
	client rest.Interface
}

// newPacketCaptures returns a PacketCaptures
func newPacketCaptures(c *AzurermV1alpha1Client) *packetCaptures {
	return &packetCaptures{
		client: c.RESTClient(),
	}
}

// Get takes name of the packetCapture, and returns the corresponding packetCapture object, and an error if there is any.
func (c *packetCaptures) Get(name string, options v1.GetOptions) (result *v1alpha1.PacketCapture, err error) {
	result = &v1alpha1.PacketCapture{}
	err = c.client.Get().
		Resource("packetcaptures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PacketCaptures that match those selectors.
func (c *packetCaptures) List(opts v1.ListOptions) (result *v1alpha1.PacketCaptureList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PacketCaptureList{}
	err = c.client.Get().
		Resource("packetcaptures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested packetCaptures.
func (c *packetCaptures) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("packetcaptures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a packetCapture and creates it.  Returns the server's representation of the packetCapture, and an error, if there is any.
func (c *packetCaptures) Create(packetCapture *v1alpha1.PacketCapture) (result *v1alpha1.PacketCapture, err error) {
	result = &v1alpha1.PacketCapture{}
	err = c.client.Post().
		Resource("packetcaptures").
		Body(packetCapture).
		Do().
		Into(result)
	return
}

// Update takes the representation of a packetCapture and updates it. Returns the server's representation of the packetCapture, and an error, if there is any.
func (c *packetCaptures) Update(packetCapture *v1alpha1.PacketCapture) (result *v1alpha1.PacketCapture, err error) {
	result = &v1alpha1.PacketCapture{}
	err = c.client.Put().
		Resource("packetcaptures").
		Name(packetCapture.Name).
		Body(packetCapture).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *packetCaptures) UpdateStatus(packetCapture *v1alpha1.PacketCapture) (result *v1alpha1.PacketCapture, err error) {
	result = &v1alpha1.PacketCapture{}
	err = c.client.Put().
		Resource("packetcaptures").
		Name(packetCapture.Name).
		SubResource("status").
		Body(packetCapture).
		Do().
		Into(result)
	return
}

// Delete takes name of the packetCapture and deletes it. Returns an error if one occurs.
func (c *packetCaptures) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("packetcaptures").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *packetCaptures) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("packetcaptures").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched packetCapture.
func (c *packetCaptures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PacketCapture, err error) {
	result = &v1alpha1.PacketCapture{}
	err = c.client.Patch(pt).
		Resource("packetcaptures").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
