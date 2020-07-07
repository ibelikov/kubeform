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

// CognitoUserPoolDomainsGetter has a method to return a CognitoUserPoolDomainInterface.
// A group's client should implement this interface.
type CognitoUserPoolDomainsGetter interface {
	CognitoUserPoolDomains(namespace string) CognitoUserPoolDomainInterface
}

// CognitoUserPoolDomainInterface has methods to work with CognitoUserPoolDomain resources.
type CognitoUserPoolDomainInterface interface {
	Create(ctx context.Context, cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain, opts v1.CreateOptions) (*v1alpha1.CognitoUserPoolDomain, error)
	Update(ctx context.Context, cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain, opts v1.UpdateOptions) (*v1alpha1.CognitoUserPoolDomain, error)
	UpdateStatus(ctx context.Context, cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain, opts v1.UpdateOptions) (*v1alpha1.CognitoUserPoolDomain, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CognitoUserPoolDomain, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CognitoUserPoolDomainList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CognitoUserPoolDomain, err error)
	CognitoUserPoolDomainExpansion
}

// cognitoUserPoolDomains implements CognitoUserPoolDomainInterface
type cognitoUserPoolDomains struct {
	client rest.Interface
	ns     string
}

// newCognitoUserPoolDomains returns a CognitoUserPoolDomains
func newCognitoUserPoolDomains(c *AwsV1alpha1Client, namespace string) *cognitoUserPoolDomains {
	return &cognitoUserPoolDomains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cognitoUserPoolDomain, and returns the corresponding cognitoUserPoolDomain object, and an error if there is any.
func (c *cognitoUserPoolDomains) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	result = &v1alpha1.CognitoUserPoolDomain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CognitoUserPoolDomains that match those selectors.
func (c *cognitoUserPoolDomains) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CognitoUserPoolDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CognitoUserPoolDomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cognitoUserPoolDomains.
func (c *cognitoUserPoolDomains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cognitoUserPoolDomain and creates it.  Returns the server's representation of the cognitoUserPoolDomain, and an error, if there is any.
func (c *cognitoUserPoolDomains) Create(ctx context.Context, cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain, opts v1.CreateOptions) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	result = &v1alpha1.CognitoUserPoolDomain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cognitoUserPoolDomain).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cognitoUserPoolDomain and updates it. Returns the server's representation of the cognitoUserPoolDomain, and an error, if there is any.
func (c *cognitoUserPoolDomains) Update(ctx context.Context, cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain, opts v1.UpdateOptions) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	result = &v1alpha1.CognitoUserPoolDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		Name(cognitoUserPoolDomain.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cognitoUserPoolDomain).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cognitoUserPoolDomains) UpdateStatus(ctx context.Context, cognitoUserPoolDomain *v1alpha1.CognitoUserPoolDomain, opts v1.UpdateOptions) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	result = &v1alpha1.CognitoUserPoolDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		Name(cognitoUserPoolDomain.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cognitoUserPoolDomain).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cognitoUserPoolDomain and deletes it. Returns an error if one occurs.
func (c *cognitoUserPoolDomains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cognitoUserPoolDomains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cognitoUserPoolDomain.
func (c *cognitoUserPoolDomains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CognitoUserPoolDomain, err error) {
	result = &v1alpha1.CognitoUserPoolDomain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cognitouserpooldomains").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
