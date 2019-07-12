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

// AwsResourcegroupsGroupsGetter has a method to return a AwsResourcegroupsGroupInterface.
// A group's client should implement this interface.
type AwsResourcegroupsGroupsGetter interface {
	AwsResourcegroupsGroups() AwsResourcegroupsGroupInterface
}

// AwsResourcegroupsGroupInterface has methods to work with AwsResourcegroupsGroup resources.
type AwsResourcegroupsGroupInterface interface {
	Create(*v1alpha1.AwsResourcegroupsGroup) (*v1alpha1.AwsResourcegroupsGroup, error)
	Update(*v1alpha1.AwsResourcegroupsGroup) (*v1alpha1.AwsResourcegroupsGroup, error)
	UpdateStatus(*v1alpha1.AwsResourcegroupsGroup) (*v1alpha1.AwsResourcegroupsGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsResourcegroupsGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsResourcegroupsGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsResourcegroupsGroup, err error)
	AwsResourcegroupsGroupExpansion
}

// awsResourcegroupsGroups implements AwsResourcegroupsGroupInterface
type awsResourcegroupsGroups struct {
	client rest.Interface
}

// newAwsResourcegroupsGroups returns a AwsResourcegroupsGroups
func newAwsResourcegroupsGroups(c *AwsV1alpha1Client) *awsResourcegroupsGroups {
	return &awsResourcegroupsGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsResourcegroupsGroup, and returns the corresponding awsResourcegroupsGroup object, and an error if there is any.
func (c *awsResourcegroupsGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsResourcegroupsGroup, err error) {
	result = &v1alpha1.AwsResourcegroupsGroup{}
	err = c.client.Get().
		Resource("awsresourcegroupsgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsResourcegroupsGroups that match those selectors.
func (c *awsResourcegroupsGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsResourcegroupsGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsResourcegroupsGroupList{}
	err = c.client.Get().
		Resource("awsresourcegroupsgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsResourcegroupsGroups.
func (c *awsResourcegroupsGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsresourcegroupsgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsResourcegroupsGroup and creates it.  Returns the server's representation of the awsResourcegroupsGroup, and an error, if there is any.
func (c *awsResourcegroupsGroups) Create(awsResourcegroupsGroup *v1alpha1.AwsResourcegroupsGroup) (result *v1alpha1.AwsResourcegroupsGroup, err error) {
	result = &v1alpha1.AwsResourcegroupsGroup{}
	err = c.client.Post().
		Resource("awsresourcegroupsgroups").
		Body(awsResourcegroupsGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsResourcegroupsGroup and updates it. Returns the server's representation of the awsResourcegroupsGroup, and an error, if there is any.
func (c *awsResourcegroupsGroups) Update(awsResourcegroupsGroup *v1alpha1.AwsResourcegroupsGroup) (result *v1alpha1.AwsResourcegroupsGroup, err error) {
	result = &v1alpha1.AwsResourcegroupsGroup{}
	err = c.client.Put().
		Resource("awsresourcegroupsgroups").
		Name(awsResourcegroupsGroup.Name).
		Body(awsResourcegroupsGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsResourcegroupsGroups) UpdateStatus(awsResourcegroupsGroup *v1alpha1.AwsResourcegroupsGroup) (result *v1alpha1.AwsResourcegroupsGroup, err error) {
	result = &v1alpha1.AwsResourcegroupsGroup{}
	err = c.client.Put().
		Resource("awsresourcegroupsgroups").
		Name(awsResourcegroupsGroup.Name).
		SubResource("status").
		Body(awsResourcegroupsGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsResourcegroupsGroup and deletes it. Returns an error if one occurs.
func (c *awsResourcegroupsGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsresourcegroupsgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsResourcegroupsGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsresourcegroupsgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsResourcegroupsGroup.
func (c *awsResourcegroupsGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsResourcegroupsGroup, err error) {
	result = &v1alpha1.AwsResourcegroupsGroup{}
	err = c.client.Patch(pt).
		Resource("awsresourcegroupsgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
