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

// CodecommitRepositoriesGetter has a method to return a CodecommitRepositoryInterface.
// A group's client should implement this interface.
type CodecommitRepositoriesGetter interface {
	CodecommitRepositories() CodecommitRepositoryInterface
}

// CodecommitRepositoryInterface has methods to work with CodecommitRepository resources.
type CodecommitRepositoryInterface interface {
	Create(*v1alpha1.CodecommitRepository) (*v1alpha1.CodecommitRepository, error)
	Update(*v1alpha1.CodecommitRepository) (*v1alpha1.CodecommitRepository, error)
	UpdateStatus(*v1alpha1.CodecommitRepository) (*v1alpha1.CodecommitRepository, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CodecommitRepository, error)
	List(opts v1.ListOptions) (*v1alpha1.CodecommitRepositoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodecommitRepository, err error)
	CodecommitRepositoryExpansion
}

// codecommitRepositories implements CodecommitRepositoryInterface
type codecommitRepositories struct {
	client rest.Interface
}

// newCodecommitRepositories returns a CodecommitRepositories
func newCodecommitRepositories(c *AwsV1alpha1Client) *codecommitRepositories {
	return &codecommitRepositories{
		client: c.RESTClient(),
	}
}

// Get takes name of the codecommitRepository, and returns the corresponding codecommitRepository object, and an error if there is any.
func (c *codecommitRepositories) Get(name string, options v1.GetOptions) (result *v1alpha1.CodecommitRepository, err error) {
	result = &v1alpha1.CodecommitRepository{}
	err = c.client.Get().
		Resource("codecommitrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CodecommitRepositories that match those selectors.
func (c *codecommitRepositories) List(opts v1.ListOptions) (result *v1alpha1.CodecommitRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CodecommitRepositoryList{}
	err = c.client.Get().
		Resource("codecommitrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested codecommitRepositories.
func (c *codecommitRepositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("codecommitrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a codecommitRepository and creates it.  Returns the server's representation of the codecommitRepository, and an error, if there is any.
func (c *codecommitRepositories) Create(codecommitRepository *v1alpha1.CodecommitRepository) (result *v1alpha1.CodecommitRepository, err error) {
	result = &v1alpha1.CodecommitRepository{}
	err = c.client.Post().
		Resource("codecommitrepositories").
		Body(codecommitRepository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a codecommitRepository and updates it. Returns the server's representation of the codecommitRepository, and an error, if there is any.
func (c *codecommitRepositories) Update(codecommitRepository *v1alpha1.CodecommitRepository) (result *v1alpha1.CodecommitRepository, err error) {
	result = &v1alpha1.CodecommitRepository{}
	err = c.client.Put().
		Resource("codecommitrepositories").
		Name(codecommitRepository.Name).
		Body(codecommitRepository).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *codecommitRepositories) UpdateStatus(codecommitRepository *v1alpha1.CodecommitRepository) (result *v1alpha1.CodecommitRepository, err error) {
	result = &v1alpha1.CodecommitRepository{}
	err = c.client.Put().
		Resource("codecommitrepositories").
		Name(codecommitRepository.Name).
		SubResource("status").
		Body(codecommitRepository).
		Do().
		Into(result)
	return
}

// Delete takes name of the codecommitRepository and deletes it. Returns an error if one occurs.
func (c *codecommitRepositories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("codecommitrepositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *codecommitRepositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("codecommitrepositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched codecommitRepository.
func (c *codecommitRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodecommitRepository, err error) {
	result = &v1alpha1.CodecommitRepository{}
	err = c.client.Patch(pt).
		Resource("codecommitrepositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
