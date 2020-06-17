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

// IotCertificatesGetter has a method to return a IotCertificateInterface.
// A group's client should implement this interface.
type IotCertificatesGetter interface {
	IotCertificates(namespace string) IotCertificateInterface
}

// IotCertificateInterface has methods to work with IotCertificate resources.
type IotCertificateInterface interface {
	Create(*v1alpha1.IotCertificate) (*v1alpha1.IotCertificate, error)
	Update(*v1alpha1.IotCertificate) (*v1alpha1.IotCertificate, error)
	UpdateStatus(*v1alpha1.IotCertificate) (*v1alpha1.IotCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IotCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.IotCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotCertificate, err error)
	IotCertificateExpansion
}

// iotCertificates implements IotCertificateInterface
type iotCertificates struct {
	client rest.Interface
	ns     string
}

// newIotCertificates returns a IotCertificates
func newIotCertificates(c *AwsV1alpha1Client, namespace string) *iotCertificates {
	return &iotCertificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iotCertificate, and returns the corresponding iotCertificate object, and an error if there is any.
func (c *iotCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.IotCertificate, err error) {
	result = &v1alpha1.IotCertificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iotcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IotCertificates that match those selectors.
func (c *iotCertificates) List(opts v1.ListOptions) (result *v1alpha1.IotCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IotCertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iotcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iotCertificates.
func (c *iotCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iotcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iotCertificate and creates it.  Returns the server's representation of the iotCertificate, and an error, if there is any.
func (c *iotCertificates) Create(iotCertificate *v1alpha1.IotCertificate) (result *v1alpha1.IotCertificate, err error) {
	result = &v1alpha1.IotCertificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iotcertificates").
		Body(iotCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iotCertificate and updates it. Returns the server's representation of the iotCertificate, and an error, if there is any.
func (c *iotCertificates) Update(iotCertificate *v1alpha1.IotCertificate) (result *v1alpha1.IotCertificate, err error) {
	result = &v1alpha1.IotCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iotcertificates").
		Name(iotCertificate.Name).
		Body(iotCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iotCertificates) UpdateStatus(iotCertificate *v1alpha1.IotCertificate) (result *v1alpha1.IotCertificate, err error) {
	result = &v1alpha1.IotCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iotcertificates").
		Name(iotCertificate.Name).
		SubResource("status").
		Body(iotCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the iotCertificate and deletes it. Returns an error if one occurs.
func (c *iotCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iotcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iotCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iotcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iotCertificate.
func (c *iotCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotCertificate, err error) {
	result = &v1alpha1.IotCertificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iotcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
