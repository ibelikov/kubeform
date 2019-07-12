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

// AwsDefaultSecurityGroupsGetter has a method to return a AwsDefaultSecurityGroupInterface.
// A group's client should implement this interface.
type AwsDefaultSecurityGroupsGetter interface {
	AwsDefaultSecurityGroups() AwsDefaultSecurityGroupInterface
}

// AwsDefaultSecurityGroupInterface has methods to work with AwsDefaultSecurityGroup resources.
type AwsDefaultSecurityGroupInterface interface {
	Create(*v1alpha1.AwsDefaultSecurityGroup) (*v1alpha1.AwsDefaultSecurityGroup, error)
	Update(*v1alpha1.AwsDefaultSecurityGroup) (*v1alpha1.AwsDefaultSecurityGroup, error)
	UpdateStatus(*v1alpha1.AwsDefaultSecurityGroup) (*v1alpha1.AwsDefaultSecurityGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDefaultSecurityGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDefaultSecurityGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultSecurityGroup, err error)
	AwsDefaultSecurityGroupExpansion
}

// awsDefaultSecurityGroups implements AwsDefaultSecurityGroupInterface
type awsDefaultSecurityGroups struct {
	client rest.Interface
}

// newAwsDefaultSecurityGroups returns a AwsDefaultSecurityGroups
func newAwsDefaultSecurityGroups(c *AwsV1alpha1Client) *awsDefaultSecurityGroups {
	return &awsDefaultSecurityGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDefaultSecurityGroup, and returns the corresponding awsDefaultSecurityGroup object, and an error if there is any.
func (c *awsDefaultSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDefaultSecurityGroup, err error) {
	result = &v1alpha1.AwsDefaultSecurityGroup{}
	err = c.client.Get().
		Resource("awsdefaultsecuritygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDefaultSecurityGroups that match those selectors.
func (c *awsDefaultSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsDefaultSecurityGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDefaultSecurityGroupList{}
	err = c.client.Get().
		Resource("awsdefaultsecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDefaultSecurityGroups.
func (c *awsDefaultSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdefaultsecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDefaultSecurityGroup and creates it.  Returns the server's representation of the awsDefaultSecurityGroup, and an error, if there is any.
func (c *awsDefaultSecurityGroups) Create(awsDefaultSecurityGroup *v1alpha1.AwsDefaultSecurityGroup) (result *v1alpha1.AwsDefaultSecurityGroup, err error) {
	result = &v1alpha1.AwsDefaultSecurityGroup{}
	err = c.client.Post().
		Resource("awsdefaultsecuritygroups").
		Body(awsDefaultSecurityGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDefaultSecurityGroup and updates it. Returns the server's representation of the awsDefaultSecurityGroup, and an error, if there is any.
func (c *awsDefaultSecurityGroups) Update(awsDefaultSecurityGroup *v1alpha1.AwsDefaultSecurityGroup) (result *v1alpha1.AwsDefaultSecurityGroup, err error) {
	result = &v1alpha1.AwsDefaultSecurityGroup{}
	err = c.client.Put().
		Resource("awsdefaultsecuritygroups").
		Name(awsDefaultSecurityGroup.Name).
		Body(awsDefaultSecurityGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDefaultSecurityGroups) UpdateStatus(awsDefaultSecurityGroup *v1alpha1.AwsDefaultSecurityGroup) (result *v1alpha1.AwsDefaultSecurityGroup, err error) {
	result = &v1alpha1.AwsDefaultSecurityGroup{}
	err = c.client.Put().
		Resource("awsdefaultsecuritygroups").
		Name(awsDefaultSecurityGroup.Name).
		SubResource("status").
		Body(awsDefaultSecurityGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDefaultSecurityGroup and deletes it. Returns an error if one occurs.
func (c *awsDefaultSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdefaultsecuritygroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDefaultSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdefaultsecuritygroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDefaultSecurityGroup.
func (c *awsDefaultSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultSecurityGroup, err error) {
	result = &v1alpha1.AwsDefaultSecurityGroup{}
	err = c.client.Patch(pt).
		Resource("awsdefaultsecuritygroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
