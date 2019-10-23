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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkWatchersGetter has a method to return a NetworkWatcherInterface.
// A group's client should implement this interface.
type NetworkWatchersGetter interface {
	NetworkWatchers(namespace string) NetworkWatcherInterface
}

// NetworkWatcherInterface has methods to work with NetworkWatcher resources.
type NetworkWatcherInterface interface {
	Create(*v1alpha1.NetworkWatcher) (*v1alpha1.NetworkWatcher, error)
	Update(*v1alpha1.NetworkWatcher) (*v1alpha1.NetworkWatcher, error)
	UpdateStatus(*v1alpha1.NetworkWatcher) (*v1alpha1.NetworkWatcher, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkWatcher, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkWatcherList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkWatcher, err error)
	NetworkWatcherExpansion
}

// networkWatchers implements NetworkWatcherInterface
type networkWatchers struct {
	client rest.Interface
	ns     string
}

// newNetworkWatchers returns a NetworkWatchers
func newNetworkWatchers(c *AzurermV1alpha1Client, namespace string) *networkWatchers {
	return &networkWatchers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkWatcher, and returns the corresponding networkWatcher object, and an error if there is any.
func (c *networkWatchers) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkWatcher, err error) {
	result = &v1alpha1.NetworkWatcher{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkwatchers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkWatchers that match those selectors.
func (c *networkWatchers) List(opts v1.ListOptions) (result *v1alpha1.NetworkWatcherList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkWatcherList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkwatchers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkWatchers.
func (c *networkWatchers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkwatchers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkWatcher and creates it.  Returns the server's representation of the networkWatcher, and an error, if there is any.
func (c *networkWatchers) Create(networkWatcher *v1alpha1.NetworkWatcher) (result *v1alpha1.NetworkWatcher, err error) {
	result = &v1alpha1.NetworkWatcher{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkwatchers").
		Body(networkWatcher).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkWatcher and updates it. Returns the server's representation of the networkWatcher, and an error, if there is any.
func (c *networkWatchers) Update(networkWatcher *v1alpha1.NetworkWatcher) (result *v1alpha1.NetworkWatcher, err error) {
	result = &v1alpha1.NetworkWatcher{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkwatchers").
		Name(networkWatcher.Name).
		Body(networkWatcher).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkWatchers) UpdateStatus(networkWatcher *v1alpha1.NetworkWatcher) (result *v1alpha1.NetworkWatcher, err error) {
	result = &v1alpha1.NetworkWatcher{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkwatchers").
		Name(networkWatcher.Name).
		SubResource("status").
		Body(networkWatcher).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkWatcher and deletes it. Returns an error if one occurs.
func (c *networkWatchers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkwatchers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkWatchers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkwatchers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkWatcher.
func (c *networkWatchers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkWatcher, err error) {
	result = &v1alpha1.NetworkWatcher{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkwatchers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
