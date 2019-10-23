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

	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SshkeysGetter has a method to return a SshkeyInterface.
// A group's client should implement this interface.
type SshkeysGetter interface {
	Sshkeys(namespace string) SshkeyInterface
}

// SshkeyInterface has methods to work with Sshkey resources.
type SshkeyInterface interface {
	Create(*v1alpha1.Sshkey) (*v1alpha1.Sshkey, error)
	Update(*v1alpha1.Sshkey) (*v1alpha1.Sshkey, error)
	UpdateStatus(*v1alpha1.Sshkey) (*v1alpha1.Sshkey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Sshkey, error)
	List(opts v1.ListOptions) (*v1alpha1.SshkeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Sshkey, err error)
	SshkeyExpansion
}

// sshkeys implements SshkeyInterface
type sshkeys struct {
	client rest.Interface
	ns     string
}

// newSshkeys returns a Sshkeys
func newSshkeys(c *LinodeV1alpha1Client, namespace string) *sshkeys {
	return &sshkeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sshkey, and returns the corresponding sshkey object, and an error if there is any.
func (c *sshkeys) Get(name string, options v1.GetOptions) (result *v1alpha1.Sshkey, err error) {
	result = &v1alpha1.Sshkey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sshkeys that match those selectors.
func (c *sshkeys) List(opts v1.ListOptions) (result *v1alpha1.SshkeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SshkeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sshkeys.
func (c *sshkeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sshkey and creates it.  Returns the server's representation of the sshkey, and an error, if there is any.
func (c *sshkeys) Create(sshkey *v1alpha1.Sshkey) (result *v1alpha1.Sshkey, err error) {
	result = &v1alpha1.Sshkey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sshkeys").
		Body(sshkey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sshkey and updates it. Returns the server's representation of the sshkey, and an error, if there is any.
func (c *sshkeys) Update(sshkey *v1alpha1.Sshkey) (result *v1alpha1.Sshkey, err error) {
	result = &v1alpha1.Sshkey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(sshkey.Name).
		Body(sshkey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sshkeys) UpdateStatus(sshkey *v1alpha1.Sshkey) (result *v1alpha1.Sshkey, err error) {
	result = &v1alpha1.Sshkey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(sshkey.Name).
		SubResource("status").
		Body(sshkey).
		Do().
		Into(result)
	return
}

// Delete takes name of the sshkey and deletes it. Returns an error if one occurs.
func (c *sshkeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sshkeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshkeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sshkey.
func (c *sshkeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Sshkey, err error) {
	result = &v1alpha1.Sshkey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sshkeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
