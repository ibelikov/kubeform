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

// NetworkAclRulesGetter has a method to return a NetworkAclRuleInterface.
// A group's client should implement this interface.
type NetworkAclRulesGetter interface {
	NetworkAclRules() NetworkAclRuleInterface
}

// NetworkAclRuleInterface has methods to work with NetworkAclRule resources.
type NetworkAclRuleInterface interface {
	Create(*v1alpha1.NetworkAclRule) (*v1alpha1.NetworkAclRule, error)
	Update(*v1alpha1.NetworkAclRule) (*v1alpha1.NetworkAclRule, error)
	UpdateStatus(*v1alpha1.NetworkAclRule) (*v1alpha1.NetworkAclRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkAclRule, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkAclRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkAclRule, err error)
	NetworkAclRuleExpansion
}

// networkAclRules implements NetworkAclRuleInterface
type networkAclRules struct {
	client rest.Interface
}

// newNetworkAclRules returns a NetworkAclRules
func newNetworkAclRules(c *AwsV1alpha1Client) *networkAclRules {
	return &networkAclRules{
		client: c.RESTClient(),
	}
}

// Get takes name of the networkAclRule, and returns the corresponding networkAclRule object, and an error if there is any.
func (c *networkAclRules) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkAclRule, err error) {
	result = &v1alpha1.NetworkAclRule{}
	err = c.client.Get().
		Resource("networkaclrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkAclRules that match those selectors.
func (c *networkAclRules) List(opts v1.ListOptions) (result *v1alpha1.NetworkAclRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkAclRuleList{}
	err = c.client.Get().
		Resource("networkaclrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkAclRules.
func (c *networkAclRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("networkaclrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkAclRule and creates it.  Returns the server's representation of the networkAclRule, and an error, if there is any.
func (c *networkAclRules) Create(networkAclRule *v1alpha1.NetworkAclRule) (result *v1alpha1.NetworkAclRule, err error) {
	result = &v1alpha1.NetworkAclRule{}
	err = c.client.Post().
		Resource("networkaclrules").
		Body(networkAclRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkAclRule and updates it. Returns the server's representation of the networkAclRule, and an error, if there is any.
func (c *networkAclRules) Update(networkAclRule *v1alpha1.NetworkAclRule) (result *v1alpha1.NetworkAclRule, err error) {
	result = &v1alpha1.NetworkAclRule{}
	err = c.client.Put().
		Resource("networkaclrules").
		Name(networkAclRule.Name).
		Body(networkAclRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkAclRules) UpdateStatus(networkAclRule *v1alpha1.NetworkAclRule) (result *v1alpha1.NetworkAclRule, err error) {
	result = &v1alpha1.NetworkAclRule{}
	err = c.client.Put().
		Resource("networkaclrules").
		Name(networkAclRule.Name).
		SubResource("status").
		Body(networkAclRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkAclRule and deletes it. Returns an error if one occurs.
func (c *networkAclRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("networkaclrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkAclRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("networkaclrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkAclRule.
func (c *networkAclRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkAclRule, err error) {
	result = &v1alpha1.NetworkAclRule{}
	err = c.client.Patch(pt).
		Resource("networkaclrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
