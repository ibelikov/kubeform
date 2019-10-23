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

// PostgresqlFirewallRulesGetter has a method to return a PostgresqlFirewallRuleInterface.
// A group's client should implement this interface.
type PostgresqlFirewallRulesGetter interface {
	PostgresqlFirewallRules(namespace string) PostgresqlFirewallRuleInterface
}

// PostgresqlFirewallRuleInterface has methods to work with PostgresqlFirewallRule resources.
type PostgresqlFirewallRuleInterface interface {
	Create(*v1alpha1.PostgresqlFirewallRule) (*v1alpha1.PostgresqlFirewallRule, error)
	Update(*v1alpha1.PostgresqlFirewallRule) (*v1alpha1.PostgresqlFirewallRule, error)
	UpdateStatus(*v1alpha1.PostgresqlFirewallRule) (*v1alpha1.PostgresqlFirewallRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PostgresqlFirewallRule, error)
	List(opts v1.ListOptions) (*v1alpha1.PostgresqlFirewallRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresqlFirewallRule, err error)
	PostgresqlFirewallRuleExpansion
}

// postgresqlFirewallRules implements PostgresqlFirewallRuleInterface
type postgresqlFirewallRules struct {
	client rest.Interface
	ns     string
}

// newPostgresqlFirewallRules returns a PostgresqlFirewallRules
func newPostgresqlFirewallRules(c *AzurermV1alpha1Client, namespace string) *postgresqlFirewallRules {
	return &postgresqlFirewallRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the postgresqlFirewallRule, and returns the corresponding postgresqlFirewallRule object, and an error if there is any.
func (c *postgresqlFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.PostgresqlFirewallRule, err error) {
	result = &v1alpha1.PostgresqlFirewallRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PostgresqlFirewallRules that match those selectors.
func (c *postgresqlFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.PostgresqlFirewallRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PostgresqlFirewallRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested postgresqlFirewallRules.
func (c *postgresqlFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a postgresqlFirewallRule and creates it.  Returns the server's representation of the postgresqlFirewallRule, and an error, if there is any.
func (c *postgresqlFirewallRules) Create(postgresqlFirewallRule *v1alpha1.PostgresqlFirewallRule) (result *v1alpha1.PostgresqlFirewallRule, err error) {
	result = &v1alpha1.PostgresqlFirewallRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		Body(postgresqlFirewallRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a postgresqlFirewallRule and updates it. Returns the server's representation of the postgresqlFirewallRule, and an error, if there is any.
func (c *postgresqlFirewallRules) Update(postgresqlFirewallRule *v1alpha1.PostgresqlFirewallRule) (result *v1alpha1.PostgresqlFirewallRule, err error) {
	result = &v1alpha1.PostgresqlFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		Name(postgresqlFirewallRule.Name).
		Body(postgresqlFirewallRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *postgresqlFirewallRules) UpdateStatus(postgresqlFirewallRule *v1alpha1.PostgresqlFirewallRule) (result *v1alpha1.PostgresqlFirewallRule, err error) {
	result = &v1alpha1.PostgresqlFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		Name(postgresqlFirewallRule.Name).
		SubResource("status").
		Body(postgresqlFirewallRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the postgresqlFirewallRule and deletes it. Returns an error if one occurs.
func (c *postgresqlFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *postgresqlFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched postgresqlFirewallRule.
func (c *postgresqlFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresqlFirewallRule, err error) {
	result = &v1alpha1.PostgresqlFirewallRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("postgresqlfirewallrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
