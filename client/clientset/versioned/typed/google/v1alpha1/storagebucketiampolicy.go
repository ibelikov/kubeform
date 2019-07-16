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

// StorageBucketIamPoliciesGetter has a method to return a StorageBucketIamPolicyInterface.
// A group's client should implement this interface.
type StorageBucketIamPoliciesGetter interface {
	StorageBucketIamPolicies() StorageBucketIamPolicyInterface
}

// StorageBucketIamPolicyInterface has methods to work with StorageBucketIamPolicy resources.
type StorageBucketIamPolicyInterface interface {
	Create(*v1alpha1.StorageBucketIamPolicy) (*v1alpha1.StorageBucketIamPolicy, error)
	Update(*v1alpha1.StorageBucketIamPolicy) (*v1alpha1.StorageBucketIamPolicy, error)
	UpdateStatus(*v1alpha1.StorageBucketIamPolicy) (*v1alpha1.StorageBucketIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageBucketIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageBucketIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageBucketIamPolicy, err error)
	StorageBucketIamPolicyExpansion
}

// storageBucketIamPolicies implements StorageBucketIamPolicyInterface
type storageBucketIamPolicies struct {
	client rest.Interface
}

// newStorageBucketIamPolicies returns a StorageBucketIamPolicies
func newStorageBucketIamPolicies(c *GoogleV1alpha1Client) *storageBucketIamPolicies {
	return &storageBucketIamPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the storageBucketIamPolicy, and returns the corresponding storageBucketIamPolicy object, and an error if there is any.
func (c *storageBucketIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	result = &v1alpha1.StorageBucketIamPolicy{}
	err = c.client.Get().
		Resource("storagebucketiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageBucketIamPolicies that match those selectors.
func (c *storageBucketIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.StorageBucketIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StorageBucketIamPolicyList{}
	err = c.client.Get().
		Resource("storagebucketiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageBucketIamPolicies.
func (c *storageBucketIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("storagebucketiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a storageBucketIamPolicy and creates it.  Returns the server's representation of the storageBucketIamPolicy, and an error, if there is any.
func (c *storageBucketIamPolicies) Create(storageBucketIamPolicy *v1alpha1.StorageBucketIamPolicy) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	result = &v1alpha1.StorageBucketIamPolicy{}
	err = c.client.Post().
		Resource("storagebucketiampolicies").
		Body(storageBucketIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageBucketIamPolicy and updates it. Returns the server's representation of the storageBucketIamPolicy, and an error, if there is any.
func (c *storageBucketIamPolicies) Update(storageBucketIamPolicy *v1alpha1.StorageBucketIamPolicy) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	result = &v1alpha1.StorageBucketIamPolicy{}
	err = c.client.Put().
		Resource("storagebucketiampolicies").
		Name(storageBucketIamPolicy.Name).
		Body(storageBucketIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *storageBucketIamPolicies) UpdateStatus(storageBucketIamPolicy *v1alpha1.StorageBucketIamPolicy) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	result = &v1alpha1.StorageBucketIamPolicy{}
	err = c.client.Put().
		Resource("storagebucketiampolicies").
		Name(storageBucketIamPolicy.Name).
		SubResource("status").
		Body(storageBucketIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageBucketIamPolicy and deletes it. Returns an error if one occurs.
func (c *storageBucketIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storagebucketiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageBucketIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("storagebucketiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageBucketIamPolicy.
func (c *storageBucketIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	result = &v1alpha1.StorageBucketIamPolicy{}
	err = c.client.Patch(pt).
		Resource("storagebucketiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
