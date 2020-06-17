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

// ConfigDeliveryChannelsGetter has a method to return a ConfigDeliveryChannelInterface.
// A group's client should implement this interface.
type ConfigDeliveryChannelsGetter interface {
	ConfigDeliveryChannels(namespace string) ConfigDeliveryChannelInterface
}

// ConfigDeliveryChannelInterface has methods to work with ConfigDeliveryChannel resources.
type ConfigDeliveryChannelInterface interface {
	Create(*v1alpha1.ConfigDeliveryChannel) (*v1alpha1.ConfigDeliveryChannel, error)
	Update(*v1alpha1.ConfigDeliveryChannel) (*v1alpha1.ConfigDeliveryChannel, error)
	UpdateStatus(*v1alpha1.ConfigDeliveryChannel) (*v1alpha1.ConfigDeliveryChannel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ConfigDeliveryChannel, error)
	List(opts v1.ListOptions) (*v1alpha1.ConfigDeliveryChannelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigDeliveryChannel, err error)
	ConfigDeliveryChannelExpansion
}

// configDeliveryChannels implements ConfigDeliveryChannelInterface
type configDeliveryChannels struct {
	client rest.Interface
	ns     string
}

// newConfigDeliveryChannels returns a ConfigDeliveryChannels
func newConfigDeliveryChannels(c *AwsV1alpha1Client, namespace string) *configDeliveryChannels {
	return &configDeliveryChannels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configDeliveryChannel, and returns the corresponding configDeliveryChannel object, and an error if there is any.
func (c *configDeliveryChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigDeliveryChannel, err error) {
	result = &v1alpha1.ConfigDeliveryChannel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigDeliveryChannels that match those selectors.
func (c *configDeliveryChannels) List(opts v1.ListOptions) (result *v1alpha1.ConfigDeliveryChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigDeliveryChannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configDeliveryChannels.
func (c *configDeliveryChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a configDeliveryChannel and creates it.  Returns the server's representation of the configDeliveryChannel, and an error, if there is any.
func (c *configDeliveryChannels) Create(configDeliveryChannel *v1alpha1.ConfigDeliveryChannel) (result *v1alpha1.ConfigDeliveryChannel, err error) {
	result = &v1alpha1.ConfigDeliveryChannel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		Body(configDeliveryChannel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a configDeliveryChannel and updates it. Returns the server's representation of the configDeliveryChannel, and an error, if there is any.
func (c *configDeliveryChannels) Update(configDeliveryChannel *v1alpha1.ConfigDeliveryChannel) (result *v1alpha1.ConfigDeliveryChannel, err error) {
	result = &v1alpha1.ConfigDeliveryChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		Name(configDeliveryChannel.Name).
		Body(configDeliveryChannel).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *configDeliveryChannels) UpdateStatus(configDeliveryChannel *v1alpha1.ConfigDeliveryChannel) (result *v1alpha1.ConfigDeliveryChannel, err error) {
	result = &v1alpha1.ConfigDeliveryChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		Name(configDeliveryChannel.Name).
		SubResource("status").
		Body(configDeliveryChannel).
		Do().
		Into(result)
	return
}

// Delete takes name of the configDeliveryChannel and deletes it. Returns an error if one occurs.
func (c *configDeliveryChannels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configDeliveryChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configdeliverychannels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched configDeliveryChannel.
func (c *configDeliveryChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigDeliveryChannel, err error) {
	result = &v1alpha1.ConfigDeliveryChannel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configdeliverychannels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
