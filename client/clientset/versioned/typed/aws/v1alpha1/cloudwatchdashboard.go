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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// CloudwatchDashboardsGetter has a method to return a CloudwatchDashboardInterface.
// A group's client should implement this interface.
type CloudwatchDashboardsGetter interface {
	CloudwatchDashboards(namespace string) CloudwatchDashboardInterface
}

// CloudwatchDashboardInterface has methods to work with CloudwatchDashboard resources.
type CloudwatchDashboardInterface interface {
	Create(*v1alpha1.CloudwatchDashboard) (*v1alpha1.CloudwatchDashboard, error)
	Update(*v1alpha1.CloudwatchDashboard) (*v1alpha1.CloudwatchDashboard, error)
	UpdateStatus(*v1alpha1.CloudwatchDashboard) (*v1alpha1.CloudwatchDashboard, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudwatchDashboard, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudwatchDashboardList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchDashboard, err error)
	CloudwatchDashboardExpansion
}

// cloudwatchDashboards implements CloudwatchDashboardInterface
type cloudwatchDashboards struct {
	client rest.Interface
	ns     string
}

// newCloudwatchDashboards returns a CloudwatchDashboards
func newCloudwatchDashboards(c *AwsV1alpha1Client, namespace string) *cloudwatchDashboards {
	return &cloudwatchDashboards{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudwatchDashboard, and returns the corresponding cloudwatchDashboard object, and an error if there is any.
func (c *cloudwatchDashboards) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchDashboard, err error) {
	result = &v1alpha1.CloudwatchDashboard{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudwatchDashboards that match those selectors.
func (c *cloudwatchDashboards) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchDashboardList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudwatchDashboardList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudwatchDashboards.
func (c *cloudwatchDashboards) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudwatchDashboard and creates it.  Returns the server's representation of the cloudwatchDashboard, and an error, if there is any.
func (c *cloudwatchDashboards) Create(cloudwatchDashboard *v1alpha1.CloudwatchDashboard) (result *v1alpha1.CloudwatchDashboard, err error) {
	result = &v1alpha1.CloudwatchDashboard{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		Body(cloudwatchDashboard).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudwatchDashboard and updates it. Returns the server's representation of the cloudwatchDashboard, and an error, if there is any.
func (c *cloudwatchDashboards) Update(cloudwatchDashboard *v1alpha1.CloudwatchDashboard) (result *v1alpha1.CloudwatchDashboard, err error) {
	result = &v1alpha1.CloudwatchDashboard{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		Name(cloudwatchDashboard.Name).
		Body(cloudwatchDashboard).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudwatchDashboards) UpdateStatus(cloudwatchDashboard *v1alpha1.CloudwatchDashboard) (result *v1alpha1.CloudwatchDashboard, err error) {
	result = &v1alpha1.CloudwatchDashboard{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		Name(cloudwatchDashboard.Name).
		SubResource("status").
		Body(cloudwatchDashboard).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudwatchDashboard and deletes it. Returns an error if one occurs.
func (c *cloudwatchDashboards) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudwatchDashboards) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudwatchDashboard.
func (c *cloudwatchDashboards) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchDashboard, err error) {
	result = &v1alpha1.CloudwatchDashboard{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudwatchdashboards").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
