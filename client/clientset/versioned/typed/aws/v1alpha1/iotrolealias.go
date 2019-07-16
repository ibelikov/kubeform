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

// IotRoleAliasesGetter has a method to return a IotRoleAliasInterface.
// A group's client should implement this interface.
type IotRoleAliasesGetter interface {
	IotRoleAliases() IotRoleAliasInterface
}

// IotRoleAliasInterface has methods to work with IotRoleAlias resources.
type IotRoleAliasInterface interface {
	Create(*v1alpha1.IotRoleAlias) (*v1alpha1.IotRoleAlias, error)
	Update(*v1alpha1.IotRoleAlias) (*v1alpha1.IotRoleAlias, error)
	UpdateStatus(*v1alpha1.IotRoleAlias) (*v1alpha1.IotRoleAlias, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IotRoleAlias, error)
	List(opts v1.ListOptions) (*v1alpha1.IotRoleAliasList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotRoleAlias, err error)
	IotRoleAliasExpansion
}

// iotRoleAliases implements IotRoleAliasInterface
type iotRoleAliases struct {
	client rest.Interface
}

// newIotRoleAliases returns a IotRoleAliases
func newIotRoleAliases(c *AwsV1alpha1Client) *iotRoleAliases {
	return &iotRoleAliases{
		client: c.RESTClient(),
	}
}

// Get takes name of the iotRoleAlias, and returns the corresponding iotRoleAlias object, and an error if there is any.
func (c *iotRoleAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.IotRoleAlias, err error) {
	result = &v1alpha1.IotRoleAlias{}
	err = c.client.Get().
		Resource("iotrolealiases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IotRoleAliases that match those selectors.
func (c *iotRoleAliases) List(opts v1.ListOptions) (result *v1alpha1.IotRoleAliasList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IotRoleAliasList{}
	err = c.client.Get().
		Resource("iotrolealiases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iotRoleAliases.
func (c *iotRoleAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("iotrolealiases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iotRoleAlias and creates it.  Returns the server's representation of the iotRoleAlias, and an error, if there is any.
func (c *iotRoleAliases) Create(iotRoleAlias *v1alpha1.IotRoleAlias) (result *v1alpha1.IotRoleAlias, err error) {
	result = &v1alpha1.IotRoleAlias{}
	err = c.client.Post().
		Resource("iotrolealiases").
		Body(iotRoleAlias).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iotRoleAlias and updates it. Returns the server's representation of the iotRoleAlias, and an error, if there is any.
func (c *iotRoleAliases) Update(iotRoleAlias *v1alpha1.IotRoleAlias) (result *v1alpha1.IotRoleAlias, err error) {
	result = &v1alpha1.IotRoleAlias{}
	err = c.client.Put().
		Resource("iotrolealiases").
		Name(iotRoleAlias.Name).
		Body(iotRoleAlias).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iotRoleAliases) UpdateStatus(iotRoleAlias *v1alpha1.IotRoleAlias) (result *v1alpha1.IotRoleAlias, err error) {
	result = &v1alpha1.IotRoleAlias{}
	err = c.client.Put().
		Resource("iotrolealiases").
		Name(iotRoleAlias.Name).
		SubResource("status").
		Body(iotRoleAlias).
		Do().
		Into(result)
	return
}

// Delete takes name of the iotRoleAlias and deletes it. Returns an error if one occurs.
func (c *iotRoleAliases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("iotrolealiases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iotRoleAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("iotrolealiases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iotRoleAlias.
func (c *iotRoleAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotRoleAlias, err error) {
	result = &v1alpha1.IotRoleAlias{}
	err = c.client.Patch(pt).
		Resource("iotrolealiases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
