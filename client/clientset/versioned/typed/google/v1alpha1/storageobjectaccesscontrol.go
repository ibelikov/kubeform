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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// StorageObjectAccessControlsGetter has a method to return a StorageObjectAccessControlInterface.
// A group's client should implement this interface.
type StorageObjectAccessControlsGetter interface {
	StorageObjectAccessControls() StorageObjectAccessControlInterface
}

// StorageObjectAccessControlInterface has methods to work with StorageObjectAccessControl resources.
type StorageObjectAccessControlInterface interface {
	Create(*v1alpha1.StorageObjectAccessControl) (*v1alpha1.StorageObjectAccessControl, error)
	Update(*v1alpha1.StorageObjectAccessControl) (*v1alpha1.StorageObjectAccessControl, error)
	UpdateStatus(*v1alpha1.StorageObjectAccessControl) (*v1alpha1.StorageObjectAccessControl, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageObjectAccessControl, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageObjectAccessControlList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageObjectAccessControl, err error)
	StorageObjectAccessControlExpansion
}

// storageObjectAccessControls implements StorageObjectAccessControlInterface
type storageObjectAccessControls struct {
	client rest.Interface
}

// newStorageObjectAccessControls returns a StorageObjectAccessControls
func newStorageObjectAccessControls(c *GoogleV1alpha1Client) *storageObjectAccessControls {
	return &storageObjectAccessControls{
		client: c.RESTClient(),
	}
}

// Get takes name of the storageObjectAccessControl, and returns the corresponding storageObjectAccessControl object, and an error if there is any.
func (c *storageObjectAccessControls) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageObjectAccessControl, err error) {
	result = &v1alpha1.StorageObjectAccessControl{}
	err = c.client.Get().
		Resource("storageobjectaccesscontrols").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageObjectAccessControls that match those selectors.
func (c *storageObjectAccessControls) List(opts v1.ListOptions) (result *v1alpha1.StorageObjectAccessControlList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageObjectAccessControlList{}
	err = c.client.Get().
		Resource("storageobjectaccesscontrols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageObjectAccessControls.
func (c *storageObjectAccessControls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storageobjectaccesscontrols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storageObjectAccessControl and creates it.  Returns the server's representation of the storageObjectAccessControl, and an error, if there is any.
func (c *storageObjectAccessControls) Create(storageObjectAccessControl *v1alpha1.StorageObjectAccessControl) (result *v1alpha1.StorageObjectAccessControl, err error) {
	result = &v1alpha1.StorageObjectAccessControl{}
	err = c.client.Post().
		Resource("storageobjectaccesscontrols").
		Body(storageObjectAccessControl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageObjectAccessControl and updates it. Returns the server's representation of the storageObjectAccessControl, and an error, if there is any.
func (c *storageObjectAccessControls) Update(storageObjectAccessControl *v1alpha1.StorageObjectAccessControl) (result *v1alpha1.StorageObjectAccessControl, err error) {
	result = &v1alpha1.StorageObjectAccessControl{}
	err = c.client.Put().
		Resource("storageobjectaccesscontrols").
		Name(storageObjectAccessControl.Name).
		Body(storageObjectAccessControl).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *storageObjectAccessControls) UpdateStatus(storageObjectAccessControl *v1alpha1.StorageObjectAccessControl) (result *v1alpha1.StorageObjectAccessControl, err error) {
	result = &v1alpha1.StorageObjectAccessControl{}
	err = c.client.Put().
		Resource("storageobjectaccesscontrols").
		Name(storageObjectAccessControl.Name).
		SubResource("status").
		Body(storageObjectAccessControl).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageObjectAccessControl and deletes it. Returns an error if one occurs.
func (c *storageObjectAccessControls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storageobjectaccesscontrols").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageObjectAccessControls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storageobjectaccesscontrols").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageObjectAccessControl.
func (c *storageObjectAccessControls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageObjectAccessControl, err error) {
	result = &v1alpha1.StorageObjectAccessControl{}
	err = c.client.Patch(pt).
		Resource("storageobjectaccesscontrols").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
