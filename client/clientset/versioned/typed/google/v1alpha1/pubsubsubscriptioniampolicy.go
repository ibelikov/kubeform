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

// PubsubSubscriptionIamPoliciesGetter has a method to return a PubsubSubscriptionIamPolicyInterface.
// A group's client should implement this interface.
type PubsubSubscriptionIamPoliciesGetter interface {
	PubsubSubscriptionIamPolicies() PubsubSubscriptionIamPolicyInterface
}

// PubsubSubscriptionIamPolicyInterface has methods to work with PubsubSubscriptionIamPolicy resources.
type PubsubSubscriptionIamPolicyInterface interface {
	Create(*v1alpha1.PubsubSubscriptionIamPolicy) (*v1alpha1.PubsubSubscriptionIamPolicy, error)
	Update(*v1alpha1.PubsubSubscriptionIamPolicy) (*v1alpha1.PubsubSubscriptionIamPolicy, error)
	UpdateStatus(*v1alpha1.PubsubSubscriptionIamPolicy) (*v1alpha1.PubsubSubscriptionIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PubsubSubscriptionIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.PubsubSubscriptionIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubsubSubscriptionIamPolicy, err error)
	PubsubSubscriptionIamPolicyExpansion
}

// pubsubSubscriptionIamPolicies implements PubsubSubscriptionIamPolicyInterface
type pubsubSubscriptionIamPolicies struct {
	client rest.Interface
}

// newPubsubSubscriptionIamPolicies returns a PubsubSubscriptionIamPolicies
func newPubsubSubscriptionIamPolicies(c *GoogleV1alpha1Client) *pubsubSubscriptionIamPolicies {
	return &pubsubSubscriptionIamPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the pubsubSubscriptionIamPolicy, and returns the corresponding pubsubSubscriptionIamPolicy object, and an error if there is any.
func (c *pubsubSubscriptionIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.PubsubSubscriptionIamPolicy, err error) {
	result = &v1alpha1.PubsubSubscriptionIamPolicy{}
	err = c.client.Get().
		Resource("pubsubsubscriptioniampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PubsubSubscriptionIamPolicies that match those selectors.
func (c *pubsubSubscriptionIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.PubsubSubscriptionIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PubsubSubscriptionIamPolicyList{}
	err = c.client.Get().
		Resource("pubsubsubscriptioniampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pubsubSubscriptionIamPolicies.
func (c *pubsubSubscriptionIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("pubsubsubscriptioniampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pubsubSubscriptionIamPolicy and creates it.  Returns the server's representation of the pubsubSubscriptionIamPolicy, and an error, if there is any.
func (c *pubsubSubscriptionIamPolicies) Create(pubsubSubscriptionIamPolicy *v1alpha1.PubsubSubscriptionIamPolicy) (result *v1alpha1.PubsubSubscriptionIamPolicy, err error) {
	result = &v1alpha1.PubsubSubscriptionIamPolicy{}
	err = c.client.Post().
		Resource("pubsubsubscriptioniampolicies").
		Body(pubsubSubscriptionIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pubsubSubscriptionIamPolicy and updates it. Returns the server's representation of the pubsubSubscriptionIamPolicy, and an error, if there is any.
func (c *pubsubSubscriptionIamPolicies) Update(pubsubSubscriptionIamPolicy *v1alpha1.PubsubSubscriptionIamPolicy) (result *v1alpha1.PubsubSubscriptionIamPolicy, err error) {
	result = &v1alpha1.PubsubSubscriptionIamPolicy{}
	err = c.client.Put().
		Resource("pubsubsubscriptioniampolicies").
		Name(pubsubSubscriptionIamPolicy.Name).
		Body(pubsubSubscriptionIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pubsubSubscriptionIamPolicies) UpdateStatus(pubsubSubscriptionIamPolicy *v1alpha1.PubsubSubscriptionIamPolicy) (result *v1alpha1.PubsubSubscriptionIamPolicy, err error) {
	result = &v1alpha1.PubsubSubscriptionIamPolicy{}
	err = c.client.Put().
		Resource("pubsubsubscriptioniampolicies").
		Name(pubsubSubscriptionIamPolicy.Name).
		SubResource("status").
		Body(pubsubSubscriptionIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the pubsubSubscriptionIamPolicy and deletes it. Returns an error if one occurs.
func (c *pubsubSubscriptionIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pubsubsubscriptioniampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pubsubSubscriptionIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("pubsubsubscriptioniampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pubsubSubscriptionIamPolicy.
func (c *pubsubSubscriptionIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubsubSubscriptionIamPolicy, err error) {
	result = &v1alpha1.PubsubSubscriptionIamPolicy{}
	err = c.client.Patch(pt).
		Resource("pubsubsubscriptioniampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
