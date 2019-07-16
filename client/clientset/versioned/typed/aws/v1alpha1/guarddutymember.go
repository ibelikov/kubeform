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

// GuarddutyMembersGetter has a method to return a GuarddutyMemberInterface.
// A group's client should implement this interface.
type GuarddutyMembersGetter interface {
	GuarddutyMembers() GuarddutyMemberInterface
}

// GuarddutyMemberInterface has methods to work with GuarddutyMember resources.
type GuarddutyMemberInterface interface {
	Create(*v1alpha1.GuarddutyMember) (*v1alpha1.GuarddutyMember, error)
	Update(*v1alpha1.GuarddutyMember) (*v1alpha1.GuarddutyMember, error)
	UpdateStatus(*v1alpha1.GuarddutyMember) (*v1alpha1.GuarddutyMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GuarddutyMember, error)
	List(opts v1.ListOptions) (*v1alpha1.GuarddutyMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyMember, err error)
	GuarddutyMemberExpansion
}

// guarddutyMembers implements GuarddutyMemberInterface
type guarddutyMembers struct {
	client rest.Interface
}

// newGuarddutyMembers returns a GuarddutyMembers
func newGuarddutyMembers(c *AwsV1alpha1Client) *guarddutyMembers {
	return &guarddutyMembers{
		client: c.RESTClient(),
	}
}

// Get takes name of the guarddutyMember, and returns the corresponding guarddutyMember object, and an error if there is any.
func (c *guarddutyMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GuarddutyMember, err error) {
	result = &v1alpha1.GuarddutyMember{}
	err = c.client.Get().
		Resource("guarddutymembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GuarddutyMembers that match those selectors.
func (c *guarddutyMembers) List(opts v1.ListOptions) (result *v1alpha1.GuarddutyMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GuarddutyMemberList{}
	err = c.client.Get().
		Resource("guarddutymembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested guarddutyMembers.
func (c *guarddutyMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("guarddutymembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a guarddutyMember and creates it.  Returns the server's representation of the guarddutyMember, and an error, if there is any.
func (c *guarddutyMembers) Create(guarddutyMember *v1alpha1.GuarddutyMember) (result *v1alpha1.GuarddutyMember, err error) {
	result = &v1alpha1.GuarddutyMember{}
	err = c.client.Post().
		Resource("guarddutymembers").
		Body(guarddutyMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a guarddutyMember and updates it. Returns the server's representation of the guarddutyMember, and an error, if there is any.
func (c *guarddutyMembers) Update(guarddutyMember *v1alpha1.GuarddutyMember) (result *v1alpha1.GuarddutyMember, err error) {
	result = &v1alpha1.GuarddutyMember{}
	err = c.client.Put().
		Resource("guarddutymembers").
		Name(guarddutyMember.Name).
		Body(guarddutyMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *guarddutyMembers) UpdateStatus(guarddutyMember *v1alpha1.GuarddutyMember) (result *v1alpha1.GuarddutyMember, err error) {
	result = &v1alpha1.GuarddutyMember{}
	err = c.client.Put().
		Resource("guarddutymembers").
		Name(guarddutyMember.Name).
		SubResource("status").
		Body(guarddutyMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the guarddutyMember and deletes it. Returns an error if one occurs.
func (c *guarddutyMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("guarddutymembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *guarddutyMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("guarddutymembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched guarddutyMember.
func (c *guarddutyMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyMember, err error) {
	result = &v1alpha1.GuarddutyMember{}
	err = c.client.Patch(pt).
		Resource("guarddutymembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
