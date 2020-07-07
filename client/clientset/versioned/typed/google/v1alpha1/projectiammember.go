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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ProjectIamMembersGetter has a method to return a ProjectIamMemberInterface.
// A group's client should implement this interface.
type ProjectIamMembersGetter interface {
	ProjectIamMembers(namespace string) ProjectIamMemberInterface
}

// ProjectIamMemberInterface has methods to work with ProjectIamMember resources.
type ProjectIamMemberInterface interface {
	Create(ctx context.Context, projectIamMember *v1alpha1.ProjectIamMember, opts v1.CreateOptions) (*v1alpha1.ProjectIamMember, error)
	Update(ctx context.Context, projectIamMember *v1alpha1.ProjectIamMember, opts v1.UpdateOptions) (*v1alpha1.ProjectIamMember, error)
	UpdateStatus(ctx context.Context, projectIamMember *v1alpha1.ProjectIamMember, opts v1.UpdateOptions) (*v1alpha1.ProjectIamMember, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ProjectIamMember, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ProjectIamMemberList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProjectIamMember, err error)
	ProjectIamMemberExpansion
}

// projectIamMembers implements ProjectIamMemberInterface
type projectIamMembers struct {
	client rest.Interface
	ns     string
}

// newProjectIamMembers returns a ProjectIamMembers
func newProjectIamMembers(c *GoogleV1alpha1Client, namespace string) *projectIamMembers {
	return &projectIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the projectIamMember, and returns the corresponding projectIamMember object, and an error if there is any.
func (c *projectIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProjectIamMember, err error) {
	result = &v1alpha1.ProjectIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProjectIamMembers that match those selectors.
func (c *projectIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProjectIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ProjectIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested projectIamMembers.
func (c *projectIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("projectiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a projectIamMember and creates it.  Returns the server's representation of the projectIamMember, and an error, if there is any.
func (c *projectIamMembers) Create(ctx context.Context, projectIamMember *v1alpha1.ProjectIamMember, opts v1.CreateOptions) (result *v1alpha1.ProjectIamMember, err error) {
	result = &v1alpha1.ProjectIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("projectiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectIamMember).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a projectIamMember and updates it. Returns the server's representation of the projectIamMember, and an error, if there is any.
func (c *projectIamMembers) Update(ctx context.Context, projectIamMember *v1alpha1.ProjectIamMember, opts v1.UpdateOptions) (result *v1alpha1.ProjectIamMember, err error) {
	result = &v1alpha1.ProjectIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("projectiammembers").
		Name(projectIamMember.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectIamMember).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *projectIamMembers) UpdateStatus(ctx context.Context, projectIamMember *v1alpha1.ProjectIamMember, opts v1.UpdateOptions) (result *v1alpha1.ProjectIamMember, err error) {
	result = &v1alpha1.ProjectIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("projectiammembers").
		Name(projectIamMember.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(projectIamMember).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the projectIamMember and deletes it. Returns an error if one occurs.
func (c *projectIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectiammembers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *projectIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectiammembers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched projectIamMember.
func (c *projectIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProjectIamMember, err error) {
	result = &v1alpha1.ProjectIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("projectiammembers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
