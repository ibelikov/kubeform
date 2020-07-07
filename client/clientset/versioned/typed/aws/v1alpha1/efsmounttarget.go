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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// EfsMountTargetsGetter has a method to return a EfsMountTargetInterface.
// A group's client should implement this interface.
type EfsMountTargetsGetter interface {
	EfsMountTargets(namespace string) EfsMountTargetInterface
}

// EfsMountTargetInterface has methods to work with EfsMountTarget resources.
type EfsMountTargetInterface interface {
	Create(ctx context.Context, efsMountTarget *v1alpha1.EfsMountTarget, opts v1.CreateOptions) (*v1alpha1.EfsMountTarget, error)
	Update(ctx context.Context, efsMountTarget *v1alpha1.EfsMountTarget, opts v1.UpdateOptions) (*v1alpha1.EfsMountTarget, error)
	UpdateStatus(ctx context.Context, efsMountTarget *v1alpha1.EfsMountTarget, opts v1.UpdateOptions) (*v1alpha1.EfsMountTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EfsMountTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EfsMountTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EfsMountTarget, err error)
	EfsMountTargetExpansion
}

// efsMountTargets implements EfsMountTargetInterface
type efsMountTargets struct {
	client rest.Interface
	ns     string
}

// newEfsMountTargets returns a EfsMountTargets
func newEfsMountTargets(c *AwsV1alpha1Client, namespace string) *efsMountTargets {
	return &efsMountTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the efsMountTarget, and returns the corresponding efsMountTarget object, and an error if there is any.
func (c *efsMountTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EfsMountTarget, err error) {
	result = &v1alpha1.EfsMountTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("efsmounttargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EfsMountTargets that match those selectors.
func (c *efsMountTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EfsMountTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EfsMountTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("efsmounttargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested efsMountTargets.
func (c *efsMountTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("efsmounttargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a efsMountTarget and creates it.  Returns the server's representation of the efsMountTarget, and an error, if there is any.
func (c *efsMountTargets) Create(ctx context.Context, efsMountTarget *v1alpha1.EfsMountTarget, opts v1.CreateOptions) (result *v1alpha1.EfsMountTarget, err error) {
	result = &v1alpha1.EfsMountTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("efsmounttargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(efsMountTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a efsMountTarget and updates it. Returns the server's representation of the efsMountTarget, and an error, if there is any.
func (c *efsMountTargets) Update(ctx context.Context, efsMountTarget *v1alpha1.EfsMountTarget, opts v1.UpdateOptions) (result *v1alpha1.EfsMountTarget, err error) {
	result = &v1alpha1.EfsMountTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("efsmounttargets").
		Name(efsMountTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(efsMountTarget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *efsMountTargets) UpdateStatus(ctx context.Context, efsMountTarget *v1alpha1.EfsMountTarget, opts v1.UpdateOptions) (result *v1alpha1.EfsMountTarget, err error) {
	result = &v1alpha1.EfsMountTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("efsmounttargets").
		Name(efsMountTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(efsMountTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the efsMountTarget and deletes it. Returns an error if one occurs.
func (c *efsMountTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("efsmounttargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *efsMountTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("efsmounttargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched efsMountTarget.
func (c *efsMountTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EfsMountTarget, err error) {
	result = &v1alpha1.EfsMountTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("efsmounttargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
