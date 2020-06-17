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

// DbInstanceRoleAssociationsGetter has a method to return a DbInstanceRoleAssociationInterface.
// A group's client should implement this interface.
type DbInstanceRoleAssociationsGetter interface {
	DbInstanceRoleAssociations(namespace string) DbInstanceRoleAssociationInterface
}

// DbInstanceRoleAssociationInterface has methods to work with DbInstanceRoleAssociation resources.
type DbInstanceRoleAssociationInterface interface {
	Create(*v1alpha1.DbInstanceRoleAssociation) (*v1alpha1.DbInstanceRoleAssociation, error)
	Update(*v1alpha1.DbInstanceRoleAssociation) (*v1alpha1.DbInstanceRoleAssociation, error)
	UpdateStatus(*v1alpha1.DbInstanceRoleAssociation) (*v1alpha1.DbInstanceRoleAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DbInstanceRoleAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.DbInstanceRoleAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbInstanceRoleAssociation, err error)
	DbInstanceRoleAssociationExpansion
}

// dbInstanceRoleAssociations implements DbInstanceRoleAssociationInterface
type dbInstanceRoleAssociations struct {
	client rest.Interface
	ns     string
}

// newDbInstanceRoleAssociations returns a DbInstanceRoleAssociations
func newDbInstanceRoleAssociations(c *AwsV1alpha1Client, namespace string) *dbInstanceRoleAssociations {
	return &dbInstanceRoleAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dbInstanceRoleAssociation, and returns the corresponding dbInstanceRoleAssociation object, and an error if there is any.
func (c *dbInstanceRoleAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.DbInstanceRoleAssociation, err error) {
	result = &v1alpha1.DbInstanceRoleAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DbInstanceRoleAssociations that match those selectors.
func (c *dbInstanceRoleAssociations) List(opts v1.ListOptions) (result *v1alpha1.DbInstanceRoleAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DbInstanceRoleAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dbInstanceRoleAssociations.
func (c *dbInstanceRoleAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dbInstanceRoleAssociation and creates it.  Returns the server's representation of the dbInstanceRoleAssociation, and an error, if there is any.
func (c *dbInstanceRoleAssociations) Create(dbInstanceRoleAssociation *v1alpha1.DbInstanceRoleAssociation) (result *v1alpha1.DbInstanceRoleAssociation, err error) {
	result = &v1alpha1.DbInstanceRoleAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		Body(dbInstanceRoleAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dbInstanceRoleAssociation and updates it. Returns the server's representation of the dbInstanceRoleAssociation, and an error, if there is any.
func (c *dbInstanceRoleAssociations) Update(dbInstanceRoleAssociation *v1alpha1.DbInstanceRoleAssociation) (result *v1alpha1.DbInstanceRoleAssociation, err error) {
	result = &v1alpha1.DbInstanceRoleAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		Name(dbInstanceRoleAssociation.Name).
		Body(dbInstanceRoleAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dbInstanceRoleAssociations) UpdateStatus(dbInstanceRoleAssociation *v1alpha1.DbInstanceRoleAssociation) (result *v1alpha1.DbInstanceRoleAssociation, err error) {
	result = &v1alpha1.DbInstanceRoleAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		Name(dbInstanceRoleAssociation.Name).
		SubResource("status").
		Body(dbInstanceRoleAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the dbInstanceRoleAssociation and deletes it. Returns an error if one occurs.
func (c *dbInstanceRoleAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dbInstanceRoleAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dbInstanceRoleAssociation.
func (c *dbInstanceRoleAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbInstanceRoleAssociation, err error) {
	result = &v1alpha1.DbInstanceRoleAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dbinstanceroleassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
