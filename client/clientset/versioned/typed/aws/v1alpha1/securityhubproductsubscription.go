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

// SecurityhubProductSubscriptionsGetter has a method to return a SecurityhubProductSubscriptionInterface.
// A group's client should implement this interface.
type SecurityhubProductSubscriptionsGetter interface {
	SecurityhubProductSubscriptions() SecurityhubProductSubscriptionInterface
}

// SecurityhubProductSubscriptionInterface has methods to work with SecurityhubProductSubscription resources.
type SecurityhubProductSubscriptionInterface interface {
	Create(*v1alpha1.SecurityhubProductSubscription) (*v1alpha1.SecurityhubProductSubscription, error)
	Update(*v1alpha1.SecurityhubProductSubscription) (*v1alpha1.SecurityhubProductSubscription, error)
	UpdateStatus(*v1alpha1.SecurityhubProductSubscription) (*v1alpha1.SecurityhubProductSubscription, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SecurityhubProductSubscription, error)
	List(opts v1.ListOptions) (*v1alpha1.SecurityhubProductSubscriptionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityhubProductSubscription, err error)
	SecurityhubProductSubscriptionExpansion
}

// securityhubProductSubscriptions implements SecurityhubProductSubscriptionInterface
type securityhubProductSubscriptions struct {
	client rest.Interface
}

// newSecurityhubProductSubscriptions returns a SecurityhubProductSubscriptions
func newSecurityhubProductSubscriptions(c *AwsV1alpha1Client) *securityhubProductSubscriptions {
	return &securityhubProductSubscriptions{
		client: c.RESTClient(),
	}
}

// Get takes name of the securityhubProductSubscription, and returns the corresponding securityhubProductSubscription object, and an error if there is any.
func (c *securityhubProductSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	result = &v1alpha1.SecurityhubProductSubscription{}
	err = c.client.Get().
		Resource("securityhubproductsubscriptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecurityhubProductSubscriptions that match those selectors.
func (c *securityhubProductSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.SecurityhubProductSubscriptionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecurityhubProductSubscriptionList{}
	err = c.client.Get().
		Resource("securityhubproductsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested securityhubProductSubscriptions.
func (c *securityhubProductSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("securityhubproductsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a securityhubProductSubscription and creates it.  Returns the server's representation of the securityhubProductSubscription, and an error, if there is any.
func (c *securityhubProductSubscriptions) Create(securityhubProductSubscription *v1alpha1.SecurityhubProductSubscription) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	result = &v1alpha1.SecurityhubProductSubscription{}
	err = c.client.Post().
		Resource("securityhubproductsubscriptions").
		Body(securityhubProductSubscription).
		Do().
		Into(result)
	return
}

// Update takes the representation of a securityhubProductSubscription and updates it. Returns the server's representation of the securityhubProductSubscription, and an error, if there is any.
func (c *securityhubProductSubscriptions) Update(securityhubProductSubscription *v1alpha1.SecurityhubProductSubscription) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	result = &v1alpha1.SecurityhubProductSubscription{}
	err = c.client.Put().
		Resource("securityhubproductsubscriptions").
		Name(securityhubProductSubscription.Name).
		Body(securityhubProductSubscription).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *securityhubProductSubscriptions) UpdateStatus(securityhubProductSubscription *v1alpha1.SecurityhubProductSubscription) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	result = &v1alpha1.SecurityhubProductSubscription{}
	err = c.client.Put().
		Resource("securityhubproductsubscriptions").
		Name(securityhubProductSubscription.Name).
		SubResource("status").
		Body(securityhubProductSubscription).
		Do().
		Into(result)
	return
}

// Delete takes name of the securityhubProductSubscription and deletes it. Returns an error if one occurs.
func (c *securityhubProductSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("securityhubproductsubscriptions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *securityhubProductSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("securityhubproductsubscriptions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched securityhubProductSubscription.
func (c *securityhubProductSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityhubProductSubscription, err error) {
	result = &v1alpha1.SecurityhubProductSubscription{}
	err = c.client.Patch(pt).
		Resource("securityhubproductsubscriptions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
