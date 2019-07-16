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

// SpannerDatabaseIamMembersGetter has a method to return a SpannerDatabaseIamMemberInterface.
// A group's client should implement this interface.
type SpannerDatabaseIamMembersGetter interface {
	SpannerDatabaseIamMembers() SpannerDatabaseIamMemberInterface
}

// SpannerDatabaseIamMemberInterface has methods to work with SpannerDatabaseIamMember resources.
type SpannerDatabaseIamMemberInterface interface {
	Create(*v1alpha1.SpannerDatabaseIamMember) (*v1alpha1.SpannerDatabaseIamMember, error)
	Update(*v1alpha1.SpannerDatabaseIamMember) (*v1alpha1.SpannerDatabaseIamMember, error)
	UpdateStatus(*v1alpha1.SpannerDatabaseIamMember) (*v1alpha1.SpannerDatabaseIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SpannerDatabaseIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.SpannerDatabaseIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerDatabaseIamMember, err error)
	SpannerDatabaseIamMemberExpansion
}

// spannerDatabaseIamMembers implements SpannerDatabaseIamMemberInterface
type spannerDatabaseIamMembers struct {
	client rest.Interface
}

// newSpannerDatabaseIamMembers returns a SpannerDatabaseIamMembers
func newSpannerDatabaseIamMembers(c *GoogleV1alpha1Client) *spannerDatabaseIamMembers {
	return &spannerDatabaseIamMembers{
		client: c.RESTClient(),
	}
}

// Get takes name of the spannerDatabaseIamMember, and returns the corresponding spannerDatabaseIamMember object, and an error if there is any.
func (c *spannerDatabaseIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.SpannerDatabaseIamMember, err error) {
	result = &v1alpha1.SpannerDatabaseIamMember{}
	err = c.client.Get().
		Resource("spannerdatabaseiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpannerDatabaseIamMembers that match those selectors.
func (c *spannerDatabaseIamMembers) List(opts v1.ListOptions) (result *v1alpha1.SpannerDatabaseIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpannerDatabaseIamMemberList{}
	err = c.client.Get().
		Resource("spannerdatabaseiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spannerDatabaseIamMembers.
func (c *spannerDatabaseIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("spannerdatabaseiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a spannerDatabaseIamMember and creates it.  Returns the server's representation of the spannerDatabaseIamMember, and an error, if there is any.
func (c *spannerDatabaseIamMembers) Create(spannerDatabaseIamMember *v1alpha1.SpannerDatabaseIamMember) (result *v1alpha1.SpannerDatabaseIamMember, err error) {
	result = &v1alpha1.SpannerDatabaseIamMember{}
	err = c.client.Post().
		Resource("spannerdatabaseiammembers").
		Body(spannerDatabaseIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a spannerDatabaseIamMember and updates it. Returns the server's representation of the spannerDatabaseIamMember, and an error, if there is any.
func (c *spannerDatabaseIamMembers) Update(spannerDatabaseIamMember *v1alpha1.SpannerDatabaseIamMember) (result *v1alpha1.SpannerDatabaseIamMember, err error) {
	result = &v1alpha1.SpannerDatabaseIamMember{}
	err = c.client.Put().
		Resource("spannerdatabaseiammembers").
		Name(spannerDatabaseIamMember.Name).
		Body(spannerDatabaseIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *spannerDatabaseIamMembers) UpdateStatus(spannerDatabaseIamMember *v1alpha1.SpannerDatabaseIamMember) (result *v1alpha1.SpannerDatabaseIamMember, err error) {
	result = &v1alpha1.SpannerDatabaseIamMember{}
	err = c.client.Put().
		Resource("spannerdatabaseiammembers").
		Name(spannerDatabaseIamMember.Name).
		SubResource("status").
		Body(spannerDatabaseIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the spannerDatabaseIamMember and deletes it. Returns an error if one occurs.
func (c *spannerDatabaseIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("spannerdatabaseiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spannerDatabaseIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("spannerdatabaseiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched spannerDatabaseIamMember.
func (c *spannerDatabaseIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerDatabaseIamMember, err error) {
	result = &v1alpha1.SpannerDatabaseIamMember{}
	err = c.client.Patch(pt).
		Resource("spannerdatabaseiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
