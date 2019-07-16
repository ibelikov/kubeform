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

// ProjectIamCustomRolesGetter has a method to return a ProjectIamCustomRoleInterface.
// A group's client should implement this interface.
type ProjectIamCustomRolesGetter interface {
	ProjectIamCustomRoles() ProjectIamCustomRoleInterface
}

// ProjectIamCustomRoleInterface has methods to work with ProjectIamCustomRole resources.
type ProjectIamCustomRoleInterface interface {
	Create(*v1alpha1.ProjectIamCustomRole) (*v1alpha1.ProjectIamCustomRole, error)
	Update(*v1alpha1.ProjectIamCustomRole) (*v1alpha1.ProjectIamCustomRole, error)
	UpdateStatus(*v1alpha1.ProjectIamCustomRole) (*v1alpha1.ProjectIamCustomRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ProjectIamCustomRole, error)
	List(opts v1.ListOptions) (*v1alpha1.ProjectIamCustomRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProjectIamCustomRole, err error)
	ProjectIamCustomRoleExpansion
}

// projectIamCustomRoles implements ProjectIamCustomRoleInterface
type projectIamCustomRoles struct {
	client rest.Interface
}

// newProjectIamCustomRoles returns a ProjectIamCustomRoles
func newProjectIamCustomRoles(c *GoogleV1alpha1Client) *projectIamCustomRoles {
	return &projectIamCustomRoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the projectIamCustomRole, and returns the corresponding projectIamCustomRole object, and an error if there is any.
func (c *projectIamCustomRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.ProjectIamCustomRole, err error) {
	result = &v1alpha1.ProjectIamCustomRole{}
	err = c.client.Get().
		Resource("projectiamcustomroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProjectIamCustomRoles that match those selectors.
func (c *projectIamCustomRoles) List(opts v1.ListOptions) (result *v1alpha1.ProjectIamCustomRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ProjectIamCustomRoleList{}
	err = c.client.Get().
		Resource("projectiamcustomroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested projectIamCustomRoles.
func (c *projectIamCustomRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("projectiamcustomroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a projectIamCustomRole and creates it.  Returns the server's representation of the projectIamCustomRole, and an error, if there is any.
func (c *projectIamCustomRoles) Create(projectIamCustomRole *v1alpha1.ProjectIamCustomRole) (result *v1alpha1.ProjectIamCustomRole, err error) {
	result = &v1alpha1.ProjectIamCustomRole{}
	err = c.client.Post().
		Resource("projectiamcustomroles").
		Body(projectIamCustomRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a projectIamCustomRole and updates it. Returns the server's representation of the projectIamCustomRole, and an error, if there is any.
func (c *projectIamCustomRoles) Update(projectIamCustomRole *v1alpha1.ProjectIamCustomRole) (result *v1alpha1.ProjectIamCustomRole, err error) {
	result = &v1alpha1.ProjectIamCustomRole{}
	err = c.client.Put().
		Resource("projectiamcustomroles").
		Name(projectIamCustomRole.Name).
		Body(projectIamCustomRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *projectIamCustomRoles) UpdateStatus(projectIamCustomRole *v1alpha1.ProjectIamCustomRole) (result *v1alpha1.ProjectIamCustomRole, err error) {
	result = &v1alpha1.ProjectIamCustomRole{}
	err = c.client.Put().
		Resource("projectiamcustomroles").
		Name(projectIamCustomRole.Name).
		SubResource("status").
		Body(projectIamCustomRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the projectIamCustomRole and deletes it. Returns an error if one occurs.
func (c *projectIamCustomRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("projectiamcustomroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *projectIamCustomRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("projectiamcustomroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched projectIamCustomRole.
func (c *projectIamCustomRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProjectIamCustomRole, err error) {
	result = &v1alpha1.ProjectIamCustomRole{}
	err = c.client.Patch(pt).
		Resource("projectiamcustomroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
