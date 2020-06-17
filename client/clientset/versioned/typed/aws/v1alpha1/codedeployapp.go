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

// CodedeployAppsGetter has a method to return a CodedeployAppInterface.
// A group's client should implement this interface.
type CodedeployAppsGetter interface {
	CodedeployApps(namespace string) CodedeployAppInterface
}

// CodedeployAppInterface has methods to work with CodedeployApp resources.
type CodedeployAppInterface interface {
	Create(*v1alpha1.CodedeployApp) (*v1alpha1.CodedeployApp, error)
	Update(*v1alpha1.CodedeployApp) (*v1alpha1.CodedeployApp, error)
	UpdateStatus(*v1alpha1.CodedeployApp) (*v1alpha1.CodedeployApp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CodedeployApp, error)
	List(opts v1.ListOptions) (*v1alpha1.CodedeployAppList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodedeployApp, err error)
	CodedeployAppExpansion
}

// codedeployApps implements CodedeployAppInterface
type codedeployApps struct {
	client rest.Interface
	ns     string
}

// newCodedeployApps returns a CodedeployApps
func newCodedeployApps(c *AwsV1alpha1Client, namespace string) *codedeployApps {
	return &codedeployApps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the codedeployApp, and returns the corresponding codedeployApp object, and an error if there is any.
func (c *codedeployApps) Get(name string, options v1.GetOptions) (result *v1alpha1.CodedeployApp, err error) {
	result = &v1alpha1.CodedeployApp{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("codedeployapps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CodedeployApps that match those selectors.
func (c *codedeployApps) List(opts v1.ListOptions) (result *v1alpha1.CodedeployAppList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CodedeployAppList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("codedeployapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested codedeployApps.
func (c *codedeployApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("codedeployapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a codedeployApp and creates it.  Returns the server's representation of the codedeployApp, and an error, if there is any.
func (c *codedeployApps) Create(codedeployApp *v1alpha1.CodedeployApp) (result *v1alpha1.CodedeployApp, err error) {
	result = &v1alpha1.CodedeployApp{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("codedeployapps").
		Body(codedeployApp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a codedeployApp and updates it. Returns the server's representation of the codedeployApp, and an error, if there is any.
func (c *codedeployApps) Update(codedeployApp *v1alpha1.CodedeployApp) (result *v1alpha1.CodedeployApp, err error) {
	result = &v1alpha1.CodedeployApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("codedeployapps").
		Name(codedeployApp.Name).
		Body(codedeployApp).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *codedeployApps) UpdateStatus(codedeployApp *v1alpha1.CodedeployApp) (result *v1alpha1.CodedeployApp, err error) {
	result = &v1alpha1.CodedeployApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("codedeployapps").
		Name(codedeployApp.Name).
		SubResource("status").
		Body(codedeployApp).
		Do().
		Into(result)
	return
}

// Delete takes name of the codedeployApp and deletes it. Returns an error if one occurs.
func (c *codedeployApps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("codedeployapps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *codedeployApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("codedeployapps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched codedeployApp.
func (c *codedeployApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodedeployApp, err error) {
	result = &v1alpha1.CodedeployApp{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("codedeployapps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
