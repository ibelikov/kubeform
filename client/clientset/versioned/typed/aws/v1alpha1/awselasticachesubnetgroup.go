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

// AwsElasticacheSubnetGroupsGetter has a method to return a AwsElasticacheSubnetGroupInterface.
// A group's client should implement this interface.
type AwsElasticacheSubnetGroupsGetter interface {
	AwsElasticacheSubnetGroups() AwsElasticacheSubnetGroupInterface
}

// AwsElasticacheSubnetGroupInterface has methods to work with AwsElasticacheSubnetGroup resources.
type AwsElasticacheSubnetGroupInterface interface {
	Create(*v1alpha1.AwsElasticacheSubnetGroup) (*v1alpha1.AwsElasticacheSubnetGroup, error)
	Update(*v1alpha1.AwsElasticacheSubnetGroup) (*v1alpha1.AwsElasticacheSubnetGroup, error)
	UpdateStatus(*v1alpha1.AwsElasticacheSubnetGroup) (*v1alpha1.AwsElasticacheSubnetGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsElasticacheSubnetGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsElasticacheSubnetGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticacheSubnetGroup, err error)
	AwsElasticacheSubnetGroupExpansion
}

// awsElasticacheSubnetGroups implements AwsElasticacheSubnetGroupInterface
type awsElasticacheSubnetGroups struct {
	client rest.Interface
}

// newAwsElasticacheSubnetGroups returns a AwsElasticacheSubnetGroups
func newAwsElasticacheSubnetGroups(c *AwsV1alpha1Client) *awsElasticacheSubnetGroups {
	return &awsElasticacheSubnetGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsElasticacheSubnetGroup, and returns the corresponding awsElasticacheSubnetGroup object, and an error if there is any.
func (c *awsElasticacheSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticacheSubnetGroup, err error) {
	result = &v1alpha1.AwsElasticacheSubnetGroup{}
	err = c.client.Get().
		Resource("awselasticachesubnetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsElasticacheSubnetGroups that match those selectors.
func (c *awsElasticacheSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticacheSubnetGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsElasticacheSubnetGroupList{}
	err = c.client.Get().
		Resource("awselasticachesubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsElasticacheSubnetGroups.
func (c *awsElasticacheSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awselasticachesubnetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsElasticacheSubnetGroup and creates it.  Returns the server's representation of the awsElasticacheSubnetGroup, and an error, if there is any.
func (c *awsElasticacheSubnetGroups) Create(awsElasticacheSubnetGroup *v1alpha1.AwsElasticacheSubnetGroup) (result *v1alpha1.AwsElasticacheSubnetGroup, err error) {
	result = &v1alpha1.AwsElasticacheSubnetGroup{}
	err = c.client.Post().
		Resource("awselasticachesubnetgroups").
		Body(awsElasticacheSubnetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsElasticacheSubnetGroup and updates it. Returns the server's representation of the awsElasticacheSubnetGroup, and an error, if there is any.
func (c *awsElasticacheSubnetGroups) Update(awsElasticacheSubnetGroup *v1alpha1.AwsElasticacheSubnetGroup) (result *v1alpha1.AwsElasticacheSubnetGroup, err error) {
	result = &v1alpha1.AwsElasticacheSubnetGroup{}
	err = c.client.Put().
		Resource("awselasticachesubnetgroups").
		Name(awsElasticacheSubnetGroup.Name).
		Body(awsElasticacheSubnetGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsElasticacheSubnetGroups) UpdateStatus(awsElasticacheSubnetGroup *v1alpha1.AwsElasticacheSubnetGroup) (result *v1alpha1.AwsElasticacheSubnetGroup, err error) {
	result = &v1alpha1.AwsElasticacheSubnetGroup{}
	err = c.client.Put().
		Resource("awselasticachesubnetgroups").
		Name(awsElasticacheSubnetGroup.Name).
		SubResource("status").
		Body(awsElasticacheSubnetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsElasticacheSubnetGroup and deletes it. Returns an error if one occurs.
func (c *awsElasticacheSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awselasticachesubnetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsElasticacheSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awselasticachesubnetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsElasticacheSubnetGroup.
func (c *awsElasticacheSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticacheSubnetGroup, err error) {
	result = &v1alpha1.AwsElasticacheSubnetGroup{}
	err = c.client.Patch(pt).
		Resource("awselasticachesubnetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
