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

// ServicebusSubscriptionRulesGetter has a method to return a ServicebusSubscriptionRuleInterface.
// A group's client should implement this interface.
type ServicebusSubscriptionRulesGetter interface {
	ServicebusSubscriptionRules(namespace string) ServicebusSubscriptionRuleInterface
}

// ServicebusSubscriptionRuleInterface has methods to work with ServicebusSubscriptionRule resources.
type ServicebusSubscriptionRuleInterface interface {
	Create(*v1alpha1.ServicebusSubscriptionRule) (*v1alpha1.ServicebusSubscriptionRule, error)
	Update(*v1alpha1.ServicebusSubscriptionRule) (*v1alpha1.ServicebusSubscriptionRule, error)
	UpdateStatus(*v1alpha1.ServicebusSubscriptionRule) (*v1alpha1.ServicebusSubscriptionRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServicebusSubscriptionRule, error)
	List(opts v1.ListOptions) (*v1alpha1.ServicebusSubscriptionRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusSubscriptionRule, err error)
	ServicebusSubscriptionRuleExpansion
}

// servicebusSubscriptionRules implements ServicebusSubscriptionRuleInterface
type servicebusSubscriptionRules struct {
	client rest.Interface
	ns     string
}

// newServicebusSubscriptionRules returns a ServicebusSubscriptionRules
func newServicebusSubscriptionRules(c *AzurermV1alpha1Client, namespace string) *servicebusSubscriptionRules {
	return &servicebusSubscriptionRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the servicebusSubscriptionRule, and returns the corresponding servicebusSubscriptionRule object, and an error if there is any.
func (c *servicebusSubscriptionRules) Get(name string, options v1.GetOptions) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	result = &v1alpha1.ServicebusSubscriptionRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServicebusSubscriptionRules that match those selectors.
func (c *servicebusSubscriptionRules) List(opts v1.ListOptions) (result *v1alpha1.ServicebusSubscriptionRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServicebusSubscriptionRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested servicebusSubscriptionRules.
func (c *servicebusSubscriptionRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a servicebusSubscriptionRule and creates it.  Returns the server's representation of the servicebusSubscriptionRule, and an error, if there is any.
func (c *servicebusSubscriptionRules) Create(servicebusSubscriptionRule *v1alpha1.ServicebusSubscriptionRule) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	result = &v1alpha1.ServicebusSubscriptionRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		Body(servicebusSubscriptionRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a servicebusSubscriptionRule and updates it. Returns the server's representation of the servicebusSubscriptionRule, and an error, if there is any.
func (c *servicebusSubscriptionRules) Update(servicebusSubscriptionRule *v1alpha1.ServicebusSubscriptionRule) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	result = &v1alpha1.ServicebusSubscriptionRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		Name(servicebusSubscriptionRule.Name).
		Body(servicebusSubscriptionRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *servicebusSubscriptionRules) UpdateStatus(servicebusSubscriptionRule *v1alpha1.ServicebusSubscriptionRule) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	result = &v1alpha1.ServicebusSubscriptionRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		Name(servicebusSubscriptionRule.Name).
		SubResource("status").
		Body(servicebusSubscriptionRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the servicebusSubscriptionRule and deletes it. Returns an error if one occurs.
func (c *servicebusSubscriptionRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *servicebusSubscriptionRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched servicebusSubscriptionRule.
func (c *servicebusSubscriptionRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	result = &v1alpha1.ServicebusSubscriptionRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicebussubscriptionrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
