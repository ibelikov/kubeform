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

// AwsCloudwatchDashboardsGetter has a method to return a AwsCloudwatchDashboardInterface.
// A group's client should implement this interface.
type AwsCloudwatchDashboardsGetter interface {
	AwsCloudwatchDashboards() AwsCloudwatchDashboardInterface
}

// AwsCloudwatchDashboardInterface has methods to work with AwsCloudwatchDashboard resources.
type AwsCloudwatchDashboardInterface interface {
	Create(*v1alpha1.AwsCloudwatchDashboard) (*v1alpha1.AwsCloudwatchDashboard, error)
	Update(*v1alpha1.AwsCloudwatchDashboard) (*v1alpha1.AwsCloudwatchDashboard, error)
	UpdateStatus(*v1alpha1.AwsCloudwatchDashboard) (*v1alpha1.AwsCloudwatchDashboard, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCloudwatchDashboard, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCloudwatchDashboardList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchDashboard, err error)
	AwsCloudwatchDashboardExpansion
}

// awsCloudwatchDashboards implements AwsCloudwatchDashboardInterface
type awsCloudwatchDashboards struct {
	client rest.Interface
}

// newAwsCloudwatchDashboards returns a AwsCloudwatchDashboards
func newAwsCloudwatchDashboards(c *AwsV1alpha1Client) *awsCloudwatchDashboards {
	return &awsCloudwatchDashboards{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCloudwatchDashboard, and returns the corresponding awsCloudwatchDashboard object, and an error if there is any.
func (c *awsCloudwatchDashboards) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchDashboard, err error) {
	result = &v1alpha1.AwsCloudwatchDashboard{}
	err = c.client.Get().
		Resource("awscloudwatchdashboards").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudwatchDashboards that match those selectors.
func (c *awsCloudwatchDashboards) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchDashboardList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCloudwatchDashboardList{}
	err = c.client.Get().
		Resource("awscloudwatchdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchDashboards.
func (c *awsCloudwatchDashboards) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscloudwatchdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCloudwatchDashboard and creates it.  Returns the server's representation of the awsCloudwatchDashboard, and an error, if there is any.
func (c *awsCloudwatchDashboards) Create(awsCloudwatchDashboard *v1alpha1.AwsCloudwatchDashboard) (result *v1alpha1.AwsCloudwatchDashboard, err error) {
	result = &v1alpha1.AwsCloudwatchDashboard{}
	err = c.client.Post().
		Resource("awscloudwatchdashboards").
		Body(awsCloudwatchDashboard).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudwatchDashboard and updates it. Returns the server's representation of the awsCloudwatchDashboard, and an error, if there is any.
func (c *awsCloudwatchDashboards) Update(awsCloudwatchDashboard *v1alpha1.AwsCloudwatchDashboard) (result *v1alpha1.AwsCloudwatchDashboard, err error) {
	result = &v1alpha1.AwsCloudwatchDashboard{}
	err = c.client.Put().
		Resource("awscloudwatchdashboards").
		Name(awsCloudwatchDashboard.Name).
		Body(awsCloudwatchDashboard).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCloudwatchDashboards) UpdateStatus(awsCloudwatchDashboard *v1alpha1.AwsCloudwatchDashboard) (result *v1alpha1.AwsCloudwatchDashboard, err error) {
	result = &v1alpha1.AwsCloudwatchDashboard{}
	err = c.client.Put().
		Resource("awscloudwatchdashboards").
		Name(awsCloudwatchDashboard.Name).
		SubResource("status").
		Body(awsCloudwatchDashboard).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudwatchDashboard and deletes it. Returns an error if one occurs.
func (c *awsCloudwatchDashboards) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscloudwatchdashboards").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudwatchDashboards) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscloudwatchdashboards").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudwatchDashboard.
func (c *awsCloudwatchDashboards) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchDashboard, err error) {
	result = &v1alpha1.AwsCloudwatchDashboard{}
	err = c.client.Patch(pt).
		Resource("awscloudwatchdashboards").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
