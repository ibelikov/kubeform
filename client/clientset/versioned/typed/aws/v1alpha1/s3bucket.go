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

// S3BucketsGetter has a method to return a S3BucketInterface.
// A group's client should implement this interface.
type S3BucketsGetter interface {
	S3Buckets(namespace string) S3BucketInterface
}

// S3BucketInterface has methods to work with S3Bucket resources.
type S3BucketInterface interface {
	Create(ctx context.Context, s3Bucket *v1alpha1.S3Bucket, opts v1.CreateOptions) (*v1alpha1.S3Bucket, error)
	Update(ctx context.Context, s3Bucket *v1alpha1.S3Bucket, opts v1.UpdateOptions) (*v1alpha1.S3Bucket, error)
	UpdateStatus(ctx context.Context, s3Bucket *v1alpha1.S3Bucket, opts v1.UpdateOptions) (*v1alpha1.S3Bucket, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.S3Bucket, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.S3BucketList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S3Bucket, err error)
	S3BucketExpansion
}

// s3Buckets implements S3BucketInterface
type s3Buckets struct {
	client rest.Interface
	ns     string
}

// newS3Buckets returns a S3Buckets
func newS3Buckets(c *AwsV1alpha1Client, namespace string) *s3Buckets {
	return &s3Buckets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3Bucket, and returns the corresponding s3Bucket object, and an error if there is any.
func (c *s3Buckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.S3Bucket, err error) {
	result = &v1alpha1.S3Bucket{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3buckets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3Buckets that match those selectors.
func (c *s3Buckets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.S3BucketList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S3BucketList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3buckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3Buckets.
func (c *s3Buckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3buckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a s3Bucket and creates it.  Returns the server's representation of the s3Bucket, and an error, if there is any.
func (c *s3Buckets) Create(ctx context.Context, s3Bucket *v1alpha1.S3Bucket, opts v1.CreateOptions) (result *v1alpha1.S3Bucket, err error) {
	result = &v1alpha1.S3Bucket{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3buckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3Bucket).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a s3Bucket and updates it. Returns the server's representation of the s3Bucket, and an error, if there is any.
func (c *s3Buckets) Update(ctx context.Context, s3Bucket *v1alpha1.S3Bucket, opts v1.UpdateOptions) (result *v1alpha1.S3Bucket, err error) {
	result = &v1alpha1.S3Bucket{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3buckets").
		Name(s3Bucket.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3Bucket).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *s3Buckets) UpdateStatus(ctx context.Context, s3Bucket *v1alpha1.S3Bucket, opts v1.UpdateOptions) (result *v1alpha1.S3Bucket, err error) {
	result = &v1alpha1.S3Bucket{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3buckets").
		Name(s3Bucket.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3Bucket).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the s3Bucket and deletes it. Returns an error if one occurs.
func (c *s3Buckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3buckets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3Buckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3buckets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched s3Bucket.
func (c *s3Buckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S3Bucket, err error) {
	result = &v1alpha1.S3Bucket{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3buckets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
