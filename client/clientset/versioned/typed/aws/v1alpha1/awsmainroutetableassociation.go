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

// AwsMainRouteTableAssociationsGetter has a method to return a AwsMainRouteTableAssociationInterface.
// A group's client should implement this interface.
type AwsMainRouteTableAssociationsGetter interface {
	AwsMainRouteTableAssociations() AwsMainRouteTableAssociationInterface
}

// AwsMainRouteTableAssociationInterface has methods to work with AwsMainRouteTableAssociation resources.
type AwsMainRouteTableAssociationInterface interface {
	Create(*v1alpha1.AwsMainRouteTableAssociation) (*v1alpha1.AwsMainRouteTableAssociation, error)
	Update(*v1alpha1.AwsMainRouteTableAssociation) (*v1alpha1.AwsMainRouteTableAssociation, error)
	UpdateStatus(*v1alpha1.AwsMainRouteTableAssociation) (*v1alpha1.AwsMainRouteTableAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsMainRouteTableAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsMainRouteTableAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsMainRouteTableAssociation, err error)
	AwsMainRouteTableAssociationExpansion
}

// awsMainRouteTableAssociations implements AwsMainRouteTableAssociationInterface
type awsMainRouteTableAssociations struct {
	client rest.Interface
}

// newAwsMainRouteTableAssociations returns a AwsMainRouteTableAssociations
func newAwsMainRouteTableAssociations(c *AwsV1alpha1Client) *awsMainRouteTableAssociations {
	return &awsMainRouteTableAssociations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsMainRouteTableAssociation, and returns the corresponding awsMainRouteTableAssociation object, and an error if there is any.
func (c *awsMainRouteTableAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsMainRouteTableAssociation, err error) {
	result = &v1alpha1.AwsMainRouteTableAssociation{}
	err = c.client.Get().
		Resource("awsmainroutetableassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsMainRouteTableAssociations that match those selectors.
func (c *awsMainRouteTableAssociations) List(opts v1.ListOptions) (result *v1alpha1.AwsMainRouteTableAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsMainRouteTableAssociationList{}
	err = c.client.Get().
		Resource("awsmainroutetableassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsMainRouteTableAssociations.
func (c *awsMainRouteTableAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsmainroutetableassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsMainRouteTableAssociation and creates it.  Returns the server's representation of the awsMainRouteTableAssociation, and an error, if there is any.
func (c *awsMainRouteTableAssociations) Create(awsMainRouteTableAssociation *v1alpha1.AwsMainRouteTableAssociation) (result *v1alpha1.AwsMainRouteTableAssociation, err error) {
	result = &v1alpha1.AwsMainRouteTableAssociation{}
	err = c.client.Post().
		Resource("awsmainroutetableassociations").
		Body(awsMainRouteTableAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsMainRouteTableAssociation and updates it. Returns the server's representation of the awsMainRouteTableAssociation, and an error, if there is any.
func (c *awsMainRouteTableAssociations) Update(awsMainRouteTableAssociation *v1alpha1.AwsMainRouteTableAssociation) (result *v1alpha1.AwsMainRouteTableAssociation, err error) {
	result = &v1alpha1.AwsMainRouteTableAssociation{}
	err = c.client.Put().
		Resource("awsmainroutetableassociations").
		Name(awsMainRouteTableAssociation.Name).
		Body(awsMainRouteTableAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsMainRouteTableAssociations) UpdateStatus(awsMainRouteTableAssociation *v1alpha1.AwsMainRouteTableAssociation) (result *v1alpha1.AwsMainRouteTableAssociation, err error) {
	result = &v1alpha1.AwsMainRouteTableAssociation{}
	err = c.client.Put().
		Resource("awsmainroutetableassociations").
		Name(awsMainRouteTableAssociation.Name).
		SubResource("status").
		Body(awsMainRouteTableAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsMainRouteTableAssociation and deletes it. Returns an error if one occurs.
func (c *awsMainRouteTableAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsmainroutetableassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsMainRouteTableAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsmainroutetableassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsMainRouteTableAssociation.
func (c *awsMainRouteTableAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsMainRouteTableAssociation, err error) {
	result = &v1alpha1.AwsMainRouteTableAssociation{}
	err = c.client.Patch(pt).
		Resource("awsmainroutetableassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
