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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// NetworkACLRulesGetter has a method to return a NetworkACLRuleInterface.
// A group's client should implement this interface.
type NetworkACLRulesGetter interface {
	NetworkACLRules(namespace string) NetworkACLRuleInterface
}

// NetworkACLRuleInterface has methods to work with NetworkACLRule resources.
type NetworkACLRuleInterface interface {
	Create(ctx context.Context, networkACLRule *v1alpha1.NetworkACLRule, opts v1.CreateOptions) (*v1alpha1.NetworkACLRule, error)
	Update(ctx context.Context, networkACLRule *v1alpha1.NetworkACLRule, opts v1.UpdateOptions) (*v1alpha1.NetworkACLRule, error)
	UpdateStatus(ctx context.Context, networkACLRule *v1alpha1.NetworkACLRule, opts v1.UpdateOptions) (*v1alpha1.NetworkACLRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NetworkACLRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NetworkACLRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkACLRule, err error)
	NetworkACLRuleExpansion
}

// networkACLRules implements NetworkACLRuleInterface
type networkACLRules struct {
	client rest.Interface
	ns     string
}

// newNetworkACLRules returns a NetworkACLRules
func newNetworkACLRules(c *AwsV1alpha1Client, namespace string) *networkACLRules {
	return &networkACLRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkACLRule, and returns the corresponding networkACLRule object, and an error if there is any.
func (c *networkACLRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkACLRule, err error) {
	result = &v1alpha1.NetworkACLRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkaclrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkACLRules that match those selectors.
func (c *networkACLRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkACLRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkACLRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkaclrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkACLRules.
func (c *networkACLRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkaclrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkACLRule and creates it.  Returns the server's representation of the networkACLRule, and an error, if there is any.
func (c *networkACLRules) Create(ctx context.Context, networkACLRule *v1alpha1.NetworkACLRule, opts v1.CreateOptions) (result *v1alpha1.NetworkACLRule, err error) {
	result = &v1alpha1.NetworkACLRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkaclrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkACLRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkACLRule and updates it. Returns the server's representation of the networkACLRule, and an error, if there is any.
func (c *networkACLRules) Update(ctx context.Context, networkACLRule *v1alpha1.NetworkACLRule, opts v1.UpdateOptions) (result *v1alpha1.NetworkACLRule, err error) {
	result = &v1alpha1.NetworkACLRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkaclrules").
		Name(networkACLRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkACLRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkACLRules) UpdateStatus(ctx context.Context, networkACLRule *v1alpha1.NetworkACLRule, opts v1.UpdateOptions) (result *v1alpha1.NetworkACLRule, err error) {
	result = &v1alpha1.NetworkACLRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkaclrules").
		Name(networkACLRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkACLRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkACLRule and deletes it. Returns an error if one occurs.
func (c *networkACLRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkaclrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkACLRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkaclrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkACLRule.
func (c *networkACLRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkACLRule, err error) {
	result = &v1alpha1.NetworkACLRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkaclrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
