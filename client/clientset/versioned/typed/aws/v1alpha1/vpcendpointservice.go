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

// VpcEndpointServicesGetter has a method to return a VpcEndpointServiceInterface.
// A group's client should implement this interface.
type VpcEndpointServicesGetter interface {
	VpcEndpointServices(namespace string) VpcEndpointServiceInterface
}

// VpcEndpointServiceInterface has methods to work with VpcEndpointService resources.
type VpcEndpointServiceInterface interface {
	Create(*v1alpha1.VpcEndpointService) (*v1alpha1.VpcEndpointService, error)
	Update(*v1alpha1.VpcEndpointService) (*v1alpha1.VpcEndpointService, error)
	UpdateStatus(*v1alpha1.VpcEndpointService) (*v1alpha1.VpcEndpointService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VpcEndpointService, error)
	List(opts v1.ListOptions) (*v1alpha1.VpcEndpointServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcEndpointService, err error)
	VpcEndpointServiceExpansion
}

// vpcEndpointServices implements VpcEndpointServiceInterface
type vpcEndpointServices struct {
	client rest.Interface
	ns     string
}

// newVpcEndpointServices returns a VpcEndpointServices
func newVpcEndpointServices(c *AwsV1alpha1Client, namespace string) *vpcEndpointServices {
	return &vpcEndpointServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vpcEndpointService, and returns the corresponding vpcEndpointService object, and an error if there is any.
func (c *vpcEndpointServices) Get(name string, options v1.GetOptions) (result *v1alpha1.VpcEndpointService, err error) {
	result = &v1alpha1.VpcEndpointService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpcEndpointServices that match those selectors.
func (c *vpcEndpointServices) List(opts v1.ListOptions) (result *v1alpha1.VpcEndpointServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpcEndpointServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpcEndpointServices.
func (c *vpcEndpointServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vpcEndpointService and creates it.  Returns the server's representation of the vpcEndpointService, and an error, if there is any.
func (c *vpcEndpointServices) Create(vpcEndpointService *v1alpha1.VpcEndpointService) (result *v1alpha1.VpcEndpointService, err error) {
	result = &v1alpha1.VpcEndpointService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		Body(vpcEndpointService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vpcEndpointService and updates it. Returns the server's representation of the vpcEndpointService, and an error, if there is any.
func (c *vpcEndpointServices) Update(vpcEndpointService *v1alpha1.VpcEndpointService) (result *v1alpha1.VpcEndpointService, err error) {
	result = &v1alpha1.VpcEndpointService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		Name(vpcEndpointService.Name).
		Body(vpcEndpointService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vpcEndpointServices) UpdateStatus(vpcEndpointService *v1alpha1.VpcEndpointService) (result *v1alpha1.VpcEndpointService, err error) {
	result = &v1alpha1.VpcEndpointService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		Name(vpcEndpointService.Name).
		SubResource("status").
		Body(vpcEndpointService).
		Do().
		Into(result)
	return
}

// Delete takes name of the vpcEndpointService and deletes it. Returns an error if one occurs.
func (c *vpcEndpointServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpcEndpointServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcendpointservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vpcEndpointService.
func (c *vpcEndpointServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcEndpointService, err error) {
	result = &v1alpha1.VpcEndpointService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpcendpointservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
