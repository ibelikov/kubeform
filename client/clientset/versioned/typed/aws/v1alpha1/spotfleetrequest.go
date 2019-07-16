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

// SpotFleetRequestsGetter has a method to return a SpotFleetRequestInterface.
// A group's client should implement this interface.
type SpotFleetRequestsGetter interface {
	SpotFleetRequests() SpotFleetRequestInterface
}

// SpotFleetRequestInterface has methods to work with SpotFleetRequest resources.
type SpotFleetRequestInterface interface {
	Create(*v1alpha1.SpotFleetRequest) (*v1alpha1.SpotFleetRequest, error)
	Update(*v1alpha1.SpotFleetRequest) (*v1alpha1.SpotFleetRequest, error)
	UpdateStatus(*v1alpha1.SpotFleetRequest) (*v1alpha1.SpotFleetRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SpotFleetRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.SpotFleetRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotFleetRequest, err error)
	SpotFleetRequestExpansion
}

// spotFleetRequests implements SpotFleetRequestInterface
type spotFleetRequests struct {
	client rest.Interface
}

// newSpotFleetRequests returns a SpotFleetRequests
func newSpotFleetRequests(c *AwsV1alpha1Client) *spotFleetRequests {
	return &spotFleetRequests{
		client: c.RESTClient(),
	}
}

// Get takes name of the spotFleetRequest, and returns the corresponding spotFleetRequest object, and an error if there is any.
func (c *spotFleetRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.SpotFleetRequest, err error) {
	result = &v1alpha1.SpotFleetRequest{}
	err = c.client.Get().
		Resource("spotfleetrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpotFleetRequests that match those selectors.
func (c *spotFleetRequests) List(opts v1.ListOptions) (result *v1alpha1.SpotFleetRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpotFleetRequestList{}
	err = c.client.Get().
		Resource("spotfleetrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spotFleetRequests.
func (c *spotFleetRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("spotfleetrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a spotFleetRequest and creates it.  Returns the server's representation of the spotFleetRequest, and an error, if there is any.
func (c *spotFleetRequests) Create(spotFleetRequest *v1alpha1.SpotFleetRequest) (result *v1alpha1.SpotFleetRequest, err error) {
	result = &v1alpha1.SpotFleetRequest{}
	err = c.client.Post().
		Resource("spotfleetrequests").
		Body(spotFleetRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a spotFleetRequest and updates it. Returns the server's representation of the spotFleetRequest, and an error, if there is any.
func (c *spotFleetRequests) Update(spotFleetRequest *v1alpha1.SpotFleetRequest) (result *v1alpha1.SpotFleetRequest, err error) {
	result = &v1alpha1.SpotFleetRequest{}
	err = c.client.Put().
		Resource("spotfleetrequests").
		Name(spotFleetRequest.Name).
		Body(spotFleetRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *spotFleetRequests) UpdateStatus(spotFleetRequest *v1alpha1.SpotFleetRequest) (result *v1alpha1.SpotFleetRequest, err error) {
	result = &v1alpha1.SpotFleetRequest{}
	err = c.client.Put().
		Resource("spotfleetrequests").
		Name(spotFleetRequest.Name).
		SubResource("status").
		Body(spotFleetRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the spotFleetRequest and deletes it. Returns an error if one occurs.
func (c *spotFleetRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("spotfleetrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spotFleetRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("spotfleetrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched spotFleetRequest.
func (c *spotFleetRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotFleetRequest, err error) {
	result = &v1alpha1.SpotFleetRequest{}
	err = c.client.Patch(pt).
		Resource("spotfleetrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
