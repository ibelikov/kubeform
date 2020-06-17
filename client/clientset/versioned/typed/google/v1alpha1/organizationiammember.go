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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// OrganizationIamMembersGetter has a method to return a OrganizationIamMemberInterface.
// A group's client should implement this interface.
type OrganizationIamMembersGetter interface {
	OrganizationIamMembers(namespace string) OrganizationIamMemberInterface
}

// OrganizationIamMemberInterface has methods to work with OrganizationIamMember resources.
type OrganizationIamMemberInterface interface {
	Create(*v1alpha1.OrganizationIamMember) (*v1alpha1.OrganizationIamMember, error)
	Update(*v1alpha1.OrganizationIamMember) (*v1alpha1.OrganizationIamMember, error)
	UpdateStatus(*v1alpha1.OrganizationIamMember) (*v1alpha1.OrganizationIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OrganizationIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.OrganizationIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationIamMember, err error)
	OrganizationIamMemberExpansion
}

// organizationIamMembers implements OrganizationIamMemberInterface
type organizationIamMembers struct {
	client rest.Interface
	ns     string
}

// newOrganizationIamMembers returns a OrganizationIamMembers
func newOrganizationIamMembers(c *GoogleV1alpha1Client, namespace string) *organizationIamMembers {
	return &organizationIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the organizationIamMember, and returns the corresponding organizationIamMember object, and an error if there is any.
func (c *organizationIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.OrganizationIamMember, err error) {
	result = &v1alpha1.OrganizationIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizationiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OrganizationIamMembers that match those selectors.
func (c *organizationIamMembers) List(opts v1.ListOptions) (result *v1alpha1.OrganizationIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OrganizationIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizationiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested organizationIamMembers.
func (c *organizationIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("organizationiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a organizationIamMember and creates it.  Returns the server's representation of the organizationIamMember, and an error, if there is any.
func (c *organizationIamMembers) Create(organizationIamMember *v1alpha1.OrganizationIamMember) (result *v1alpha1.OrganizationIamMember, err error) {
	result = &v1alpha1.OrganizationIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("organizationiammembers").
		Body(organizationIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a organizationIamMember and updates it. Returns the server's representation of the organizationIamMember, and an error, if there is any.
func (c *organizationIamMembers) Update(organizationIamMember *v1alpha1.OrganizationIamMember) (result *v1alpha1.OrganizationIamMember, err error) {
	result = &v1alpha1.OrganizationIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizationiammembers").
		Name(organizationIamMember.Name).
		Body(organizationIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *organizationIamMembers) UpdateStatus(organizationIamMember *v1alpha1.OrganizationIamMember) (result *v1alpha1.OrganizationIamMember, err error) {
	result = &v1alpha1.OrganizationIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizationiammembers").
		Name(organizationIamMember.Name).
		SubResource("status").
		Body(organizationIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the organizationIamMember and deletes it. Returns an error if one occurs.
func (c *organizationIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizationiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *organizationIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizationiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched organizationIamMember.
func (c *organizationIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationIamMember, err error) {
	result = &v1alpha1.OrganizationIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("organizationiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
