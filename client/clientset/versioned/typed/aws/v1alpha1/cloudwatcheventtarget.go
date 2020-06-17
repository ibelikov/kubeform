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

// CloudwatchEventTargetsGetter has a method to return a CloudwatchEventTargetInterface.
// A group's client should implement this interface.
type CloudwatchEventTargetsGetter interface {
	CloudwatchEventTargets(namespace string) CloudwatchEventTargetInterface
}

// CloudwatchEventTargetInterface has methods to work with CloudwatchEventTarget resources.
type CloudwatchEventTargetInterface interface {
	Create(*v1alpha1.CloudwatchEventTarget) (*v1alpha1.CloudwatchEventTarget, error)
	Update(*v1alpha1.CloudwatchEventTarget) (*v1alpha1.CloudwatchEventTarget, error)
	UpdateStatus(*v1alpha1.CloudwatchEventTarget) (*v1alpha1.CloudwatchEventTarget, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudwatchEventTarget, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudwatchEventTargetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchEventTarget, err error)
	CloudwatchEventTargetExpansion
}

// cloudwatchEventTargets implements CloudwatchEventTargetInterface
type cloudwatchEventTargets struct {
	client rest.Interface
	ns     string
}

// newCloudwatchEventTargets returns a CloudwatchEventTargets
func newCloudwatchEventTargets(c *AwsV1alpha1Client, namespace string) *cloudwatchEventTargets {
	return &cloudwatchEventTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudwatchEventTarget, and returns the corresponding cloudwatchEventTarget object, and an error if there is any.
func (c *cloudwatchEventTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchEventTarget, err error) {
	result = &v1alpha1.CloudwatchEventTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudwatchEventTargets that match those selectors.
func (c *cloudwatchEventTargets) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchEventTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudwatchEventTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudwatchEventTargets.
func (c *cloudwatchEventTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudwatchEventTarget and creates it.  Returns the server's representation of the cloudwatchEventTarget, and an error, if there is any.
func (c *cloudwatchEventTargets) Create(cloudwatchEventTarget *v1alpha1.CloudwatchEventTarget) (result *v1alpha1.CloudwatchEventTarget, err error) {
	result = &v1alpha1.CloudwatchEventTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		Body(cloudwatchEventTarget).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudwatchEventTarget and updates it. Returns the server's representation of the cloudwatchEventTarget, and an error, if there is any.
func (c *cloudwatchEventTargets) Update(cloudwatchEventTarget *v1alpha1.CloudwatchEventTarget) (result *v1alpha1.CloudwatchEventTarget, err error) {
	result = &v1alpha1.CloudwatchEventTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		Name(cloudwatchEventTarget.Name).
		Body(cloudwatchEventTarget).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudwatchEventTargets) UpdateStatus(cloudwatchEventTarget *v1alpha1.CloudwatchEventTarget) (result *v1alpha1.CloudwatchEventTarget, err error) {
	result = &v1alpha1.CloudwatchEventTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		Name(cloudwatchEventTarget.Name).
		SubResource("status").
		Body(cloudwatchEventTarget).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudwatchEventTarget and deletes it. Returns an error if one occurs.
func (c *cloudwatchEventTargets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudwatchEventTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudwatchEventTarget.
func (c *cloudwatchEventTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchEventTarget, err error) {
	result = &v1alpha1.CloudwatchEventTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudwatcheventtargets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
