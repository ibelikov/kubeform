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

// Route53ResolverRuleAssociationsGetter has a method to return a Route53ResolverRuleAssociationInterface.
// A group's client should implement this interface.
type Route53ResolverRuleAssociationsGetter interface {
	Route53ResolverRuleAssociations() Route53ResolverRuleAssociationInterface
}

// Route53ResolverRuleAssociationInterface has methods to work with Route53ResolverRuleAssociation resources.
type Route53ResolverRuleAssociationInterface interface {
	Create(*v1alpha1.Route53ResolverRuleAssociation) (*v1alpha1.Route53ResolverRuleAssociation, error)
	Update(*v1alpha1.Route53ResolverRuleAssociation) (*v1alpha1.Route53ResolverRuleAssociation, error)
	UpdateStatus(*v1alpha1.Route53ResolverRuleAssociation) (*v1alpha1.Route53ResolverRuleAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Route53ResolverRuleAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.Route53ResolverRuleAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ResolverRuleAssociation, err error)
	Route53ResolverRuleAssociationExpansion
}

// route53ResolverRuleAssociations implements Route53ResolverRuleAssociationInterface
type route53ResolverRuleAssociations struct {
	client rest.Interface
}

// newRoute53ResolverRuleAssociations returns a Route53ResolverRuleAssociations
func newRoute53ResolverRuleAssociations(c *AwsV1alpha1Client) *route53ResolverRuleAssociations {
	return &route53ResolverRuleAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the route53ResolverRuleAssociation, and returns the corresponding route53ResolverRuleAssociation object, and an error if there is any.
func (c *route53ResolverRuleAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	result = &v1alpha1.Route53ResolverRuleAssociation{}
	err = c.client.Get().
		Resource("route53resolverruleassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Route53ResolverRuleAssociations that match those selectors.
func (c *route53ResolverRuleAssociations) List(opts v1.ListOptions) (result *v1alpha1.Route53ResolverRuleAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Route53ResolverRuleAssociationList{}
	err = c.client.Get().
		Resource("route53resolverruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested route53ResolverRuleAssociations.
func (c *route53ResolverRuleAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("route53resolverruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a route53ResolverRuleAssociation and creates it.  Returns the server's representation of the route53ResolverRuleAssociation, and an error, if there is any.
func (c *route53ResolverRuleAssociations) Create(route53ResolverRuleAssociation *v1alpha1.Route53ResolverRuleAssociation) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	result = &v1alpha1.Route53ResolverRuleAssociation{}
	err = c.client.Post().
		Resource("route53resolverruleassociations").
		Body(route53ResolverRuleAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a route53ResolverRuleAssociation and updates it. Returns the server's representation of the route53ResolverRuleAssociation, and an error, if there is any.
func (c *route53ResolverRuleAssociations) Update(route53ResolverRuleAssociation *v1alpha1.Route53ResolverRuleAssociation) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	result = &v1alpha1.Route53ResolverRuleAssociation{}
	err = c.client.Put().
		Resource("route53resolverruleassociations").
		Name(route53ResolverRuleAssociation.Name).
		Body(route53ResolverRuleAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *route53ResolverRuleAssociations) UpdateStatus(route53ResolverRuleAssociation *v1alpha1.Route53ResolverRuleAssociation) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	result = &v1alpha1.Route53ResolverRuleAssociation{}
	err = c.client.Put().
		Resource("route53resolverruleassociations").
		Name(route53ResolverRuleAssociation.Name).
		SubResource("status").
		Body(route53ResolverRuleAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the route53ResolverRuleAssociation and deletes it. Returns an error if one occurs.
func (c *route53ResolverRuleAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("route53resolverruleassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *route53ResolverRuleAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("route53resolverruleassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched route53ResolverRuleAssociation.
func (c *route53ResolverRuleAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ResolverRuleAssociation, err error) {
	result = &v1alpha1.Route53ResolverRuleAssociation{}
	err = c.client.Patch(pt).
		Resource("route53resolverruleassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
