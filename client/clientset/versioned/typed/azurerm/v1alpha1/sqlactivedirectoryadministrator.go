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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SqlActiveDirectoryAdministratorsGetter has a method to return a SqlActiveDirectoryAdministratorInterface.
// A group's client should implement this interface.
type SqlActiveDirectoryAdministratorsGetter interface {
	SqlActiveDirectoryAdministrators(namespace string) SqlActiveDirectoryAdministratorInterface
}

// SqlActiveDirectoryAdministratorInterface has methods to work with SqlActiveDirectoryAdministrator resources.
type SqlActiveDirectoryAdministratorInterface interface {
	Create(*v1alpha1.SqlActiveDirectoryAdministrator) (*v1alpha1.SqlActiveDirectoryAdministrator, error)
	Update(*v1alpha1.SqlActiveDirectoryAdministrator) (*v1alpha1.SqlActiveDirectoryAdministrator, error)
	UpdateStatus(*v1alpha1.SqlActiveDirectoryAdministrator) (*v1alpha1.SqlActiveDirectoryAdministrator, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SqlActiveDirectoryAdministrator, error)
	List(opts v1.ListOptions) (*v1alpha1.SqlActiveDirectoryAdministratorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlActiveDirectoryAdministrator, err error)
	SqlActiveDirectoryAdministratorExpansion
}

// sqlActiveDirectoryAdministrators implements SqlActiveDirectoryAdministratorInterface
type sqlActiveDirectoryAdministrators struct {
	client rest.Interface
	ns     string
}

// newSqlActiveDirectoryAdministrators returns a SqlActiveDirectoryAdministrators
func newSqlActiveDirectoryAdministrators(c *AzurermV1alpha1Client, namespace string) *sqlActiveDirectoryAdministrators {
	return &sqlActiveDirectoryAdministrators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sqlActiveDirectoryAdministrator, and returns the corresponding sqlActiveDirectoryAdministrator object, and an error if there is any.
func (c *sqlActiveDirectoryAdministrators) Get(name string, options v1.GetOptions) (result *v1alpha1.SqlActiveDirectoryAdministrator, err error) {
	result = &v1alpha1.SqlActiveDirectoryAdministrator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SqlActiveDirectoryAdministrators that match those selectors.
func (c *sqlActiveDirectoryAdministrators) List(opts v1.ListOptions) (result *v1alpha1.SqlActiveDirectoryAdministratorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SqlActiveDirectoryAdministratorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sqlActiveDirectoryAdministrators.
func (c *sqlActiveDirectoryAdministrators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sqlActiveDirectoryAdministrator and creates it.  Returns the server's representation of the sqlActiveDirectoryAdministrator, and an error, if there is any.
func (c *sqlActiveDirectoryAdministrators) Create(sqlActiveDirectoryAdministrator *v1alpha1.SqlActiveDirectoryAdministrator) (result *v1alpha1.SqlActiveDirectoryAdministrator, err error) {
	result = &v1alpha1.SqlActiveDirectoryAdministrator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		Body(sqlActiveDirectoryAdministrator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sqlActiveDirectoryAdministrator and updates it. Returns the server's representation of the sqlActiveDirectoryAdministrator, and an error, if there is any.
func (c *sqlActiveDirectoryAdministrators) Update(sqlActiveDirectoryAdministrator *v1alpha1.SqlActiveDirectoryAdministrator) (result *v1alpha1.SqlActiveDirectoryAdministrator, err error) {
	result = &v1alpha1.SqlActiveDirectoryAdministrator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		Name(sqlActiveDirectoryAdministrator.Name).
		Body(sqlActiveDirectoryAdministrator).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sqlActiveDirectoryAdministrators) UpdateStatus(sqlActiveDirectoryAdministrator *v1alpha1.SqlActiveDirectoryAdministrator) (result *v1alpha1.SqlActiveDirectoryAdministrator, err error) {
	result = &v1alpha1.SqlActiveDirectoryAdministrator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		Name(sqlActiveDirectoryAdministrator.Name).
		SubResource("status").
		Body(sqlActiveDirectoryAdministrator).
		Do().
		Into(result)
	return
}

// Delete takes name of the sqlActiveDirectoryAdministrator and deletes it. Returns an error if one occurs.
func (c *sqlActiveDirectoryAdministrators) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sqlActiveDirectoryAdministrators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sqlActiveDirectoryAdministrator.
func (c *sqlActiveDirectoryAdministrators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlActiveDirectoryAdministrator, err error) {
	result = &v1alpha1.SqlActiveDirectoryAdministrator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sqlactivedirectoryadministrators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
