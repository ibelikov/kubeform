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

// GoogleSpannerDatabaseIamMembersGetter has a method to return a GoogleSpannerDatabaseIamMemberInterface.
// A group's client should implement this interface.
type GoogleSpannerDatabaseIamMembersGetter interface {
	GoogleSpannerDatabaseIamMembers() GoogleSpannerDatabaseIamMemberInterface
}

// GoogleSpannerDatabaseIamMemberInterface has methods to work with GoogleSpannerDatabaseIamMember resources.
type GoogleSpannerDatabaseIamMemberInterface interface {
	Create(*v1alpha1.GoogleSpannerDatabaseIamMember) (*v1alpha1.GoogleSpannerDatabaseIamMember, error)
	Update(*v1alpha1.GoogleSpannerDatabaseIamMember) (*v1alpha1.GoogleSpannerDatabaseIamMember, error)
	UpdateStatus(*v1alpha1.GoogleSpannerDatabaseIamMember) (*v1alpha1.GoogleSpannerDatabaseIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleSpannerDatabaseIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleSpannerDatabaseIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error)
	GoogleSpannerDatabaseIamMemberExpansion
}

// googleSpannerDatabaseIamMembers implements GoogleSpannerDatabaseIamMemberInterface
type googleSpannerDatabaseIamMembers struct {
	client rest.Interface
}

// newGoogleSpannerDatabaseIamMembers returns a GoogleSpannerDatabaseIamMembers
func newGoogleSpannerDatabaseIamMembers(c *GoogleV1alpha1Client) *googleSpannerDatabaseIamMembers {
	return &googleSpannerDatabaseIamMembers{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleSpannerDatabaseIamMember, and returns the corresponding googleSpannerDatabaseIamMember object, and an error if there is any.
func (c *googleSpannerDatabaseIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	result = &v1alpha1.GoogleSpannerDatabaseIamMember{}
	err = c.client.Get().
		Resource("googlespannerdatabaseiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleSpannerDatabaseIamMembers that match those selectors.
func (c *googleSpannerDatabaseIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleSpannerDatabaseIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleSpannerDatabaseIamMemberList{}
	err = c.client.Get().
		Resource("googlespannerdatabaseiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleSpannerDatabaseIamMembers.
func (c *googleSpannerDatabaseIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlespannerdatabaseiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleSpannerDatabaseIamMember and creates it.  Returns the server's representation of the googleSpannerDatabaseIamMember, and an error, if there is any.
func (c *googleSpannerDatabaseIamMembers) Create(googleSpannerDatabaseIamMember *v1alpha1.GoogleSpannerDatabaseIamMember) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	result = &v1alpha1.GoogleSpannerDatabaseIamMember{}
	err = c.client.Post().
		Resource("googlespannerdatabaseiammembers").
		Body(googleSpannerDatabaseIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleSpannerDatabaseIamMember and updates it. Returns the server's representation of the googleSpannerDatabaseIamMember, and an error, if there is any.
func (c *googleSpannerDatabaseIamMembers) Update(googleSpannerDatabaseIamMember *v1alpha1.GoogleSpannerDatabaseIamMember) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	result = &v1alpha1.GoogleSpannerDatabaseIamMember{}
	err = c.client.Put().
		Resource("googlespannerdatabaseiammembers").
		Name(googleSpannerDatabaseIamMember.Name).
		Body(googleSpannerDatabaseIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleSpannerDatabaseIamMembers) UpdateStatus(googleSpannerDatabaseIamMember *v1alpha1.GoogleSpannerDatabaseIamMember) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	result = &v1alpha1.GoogleSpannerDatabaseIamMember{}
	err = c.client.Put().
		Resource("googlespannerdatabaseiammembers").
		Name(googleSpannerDatabaseIamMember.Name).
		SubResource("status").
		Body(googleSpannerDatabaseIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleSpannerDatabaseIamMember and deletes it. Returns an error if one occurs.
func (c *googleSpannerDatabaseIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlespannerdatabaseiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleSpannerDatabaseIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlespannerdatabaseiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleSpannerDatabaseIamMember.
func (c *googleSpannerDatabaseIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerDatabaseIamMember, err error) {
	result = &v1alpha1.GoogleSpannerDatabaseIamMember{}
	err = c.client.Patch(pt).
		Resource("googlespannerdatabaseiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
