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

// WafregionalRulesGetter has a method to return a WafregionalRuleInterface.
// A group's client should implement this interface.
type WafregionalRulesGetter interface {
	WafregionalRules(namespace string) WafregionalRuleInterface
}

// WafregionalRuleInterface has methods to work with WafregionalRule resources.
type WafregionalRuleInterface interface {
	Create(*v1alpha1.WafregionalRule) (*v1alpha1.WafregionalRule, error)
	Update(*v1alpha1.WafregionalRule) (*v1alpha1.WafregionalRule, error)
	UpdateStatus(*v1alpha1.WafregionalRule) (*v1alpha1.WafregionalRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafregionalRule, error)
	List(opts v1.ListOptions) (*v1alpha1.WafregionalRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalRule, err error)
	WafregionalRuleExpansion
}

// wafregionalRules implements WafregionalRuleInterface
type wafregionalRules struct {
	client rest.Interface
	ns     string
}

// newWafregionalRules returns a WafregionalRules
func newWafregionalRules(c *AwsV1alpha1Client, namespace string) *wafregionalRules {
	return &wafregionalRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafregionalRule, and returns the corresponding wafregionalRule object, and an error if there is any.
func (c *wafregionalRules) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalRule, err error) {
	result = &v1alpha1.WafregionalRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafregionalRules that match those selectors.
func (c *wafregionalRules) List(opts v1.ListOptions) (result *v1alpha1.WafregionalRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafregionalRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafregionalRules.
func (c *wafregionalRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafregionalRule and creates it.  Returns the server's representation of the wafregionalRule, and an error, if there is any.
func (c *wafregionalRules) Create(wafregionalRule *v1alpha1.WafregionalRule) (result *v1alpha1.WafregionalRule, err error) {
	result = &v1alpha1.WafregionalRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafregionalrules").
		Body(wafregionalRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafregionalRule and updates it. Returns the server's representation of the wafregionalRule, and an error, if there is any.
func (c *wafregionalRules) Update(wafregionalRule *v1alpha1.WafregionalRule) (result *v1alpha1.WafregionalRule, err error) {
	result = &v1alpha1.WafregionalRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalrules").
		Name(wafregionalRule.Name).
		Body(wafregionalRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafregionalRules) UpdateStatus(wafregionalRule *v1alpha1.WafregionalRule) (result *v1alpha1.WafregionalRule, err error) {
	result = &v1alpha1.WafregionalRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalrules").
		Name(wafregionalRule.Name).
		SubResource("status").
		Body(wafregionalRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafregionalRule and deletes it. Returns an error if one occurs.
func (c *wafregionalRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafregionalRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafregionalRule.
func (c *wafregionalRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalRule, err error) {
	result = &v1alpha1.WafregionalRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafregionalrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
