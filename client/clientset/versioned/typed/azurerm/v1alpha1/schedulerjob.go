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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// SchedulerJobsGetter has a method to return a SchedulerJobInterface.
// A group's client should implement this interface.
type SchedulerJobsGetter interface {
	SchedulerJobs(namespace string) SchedulerJobInterface
}

// SchedulerJobInterface has methods to work with SchedulerJob resources.
type SchedulerJobInterface interface {
	Create(*v1alpha1.SchedulerJob) (*v1alpha1.SchedulerJob, error)
	Update(*v1alpha1.SchedulerJob) (*v1alpha1.SchedulerJob, error)
	UpdateStatus(*v1alpha1.SchedulerJob) (*v1alpha1.SchedulerJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SchedulerJob, error)
	List(opts v1.ListOptions) (*v1alpha1.SchedulerJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SchedulerJob, err error)
	SchedulerJobExpansion
}

// schedulerJobs implements SchedulerJobInterface
type schedulerJobs struct {
	client rest.Interface
	ns     string
}

// newSchedulerJobs returns a SchedulerJobs
func newSchedulerJobs(c *AzurermV1alpha1Client, namespace string) *schedulerJobs {
	return &schedulerJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the schedulerJob, and returns the corresponding schedulerJob object, and an error if there is any.
func (c *schedulerJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SchedulerJobs that match those selectors.
func (c *schedulerJobs) List(opts v1.ListOptions) (result *v1alpha1.SchedulerJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SchedulerJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulerJobs.
func (c *schedulerJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a schedulerJob and creates it.  Returns the server's representation of the schedulerJob, and an error, if there is any.
func (c *schedulerJobs) Create(schedulerJob *v1alpha1.SchedulerJob) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Body(schedulerJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a schedulerJob and updates it. Returns the server's representation of the schedulerJob, and an error, if there is any.
func (c *schedulerJobs) Update(schedulerJob *v1alpha1.SchedulerJob) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(schedulerJob.Name).
		Body(schedulerJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *schedulerJobs) UpdateStatus(schedulerJob *v1alpha1.SchedulerJob) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(schedulerJob.Name).
		SubResource("status").
		Body(schedulerJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the schedulerJob and deletes it. Returns an error if one occurs.
func (c *schedulerJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulerJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulerjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched schedulerJob.
func (c *schedulerJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("schedulerjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
