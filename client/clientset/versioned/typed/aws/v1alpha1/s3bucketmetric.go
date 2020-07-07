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

// S3BucketMetricsGetter has a method to return a S3BucketMetricInterface.
// A group's client should implement this interface.
type S3BucketMetricsGetter interface {
	S3BucketMetrics(namespace string) S3BucketMetricInterface
}

// S3BucketMetricInterface has methods to work with S3BucketMetric resources.
type S3BucketMetricInterface interface {
	Create(ctx context.Context, s3BucketMetric *v1alpha1.S3BucketMetric, opts v1.CreateOptions) (*v1alpha1.S3BucketMetric, error)
	Update(ctx context.Context, s3BucketMetric *v1alpha1.S3BucketMetric, opts v1.UpdateOptions) (*v1alpha1.S3BucketMetric, error)
	UpdateStatus(ctx context.Context, s3BucketMetric *v1alpha1.S3BucketMetric, opts v1.UpdateOptions) (*v1alpha1.S3BucketMetric, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.S3BucketMetric, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.S3BucketMetricList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S3BucketMetric, err error)
	S3BucketMetricExpansion
}

// s3BucketMetrics implements S3BucketMetricInterface
type s3BucketMetrics struct {
	client rest.Interface
	ns     string
}

// newS3BucketMetrics returns a S3BucketMetrics
func newS3BucketMetrics(c *AwsV1alpha1Client, namespace string) *s3BucketMetrics {
	return &s3BucketMetrics{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3BucketMetric, and returns the corresponding s3BucketMetric object, and an error if there is any.
func (c *s3BucketMetrics) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.S3BucketMetric, err error) {
	result = &v1alpha1.S3BucketMetric{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3BucketMetrics that match those selectors.
func (c *s3BucketMetrics) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.S3BucketMetricList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S3BucketMetricList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3BucketMetrics.
func (c *s3BucketMetrics) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a s3BucketMetric and creates it.  Returns the server's representation of the s3BucketMetric, and an error, if there is any.
func (c *s3BucketMetrics) Create(ctx context.Context, s3BucketMetric *v1alpha1.S3BucketMetric, opts v1.CreateOptions) (result *v1alpha1.S3BucketMetric, err error) {
	result = &v1alpha1.S3BucketMetric{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3BucketMetric).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a s3BucketMetric and updates it. Returns the server's representation of the s3BucketMetric, and an error, if there is any.
func (c *s3BucketMetrics) Update(ctx context.Context, s3BucketMetric *v1alpha1.S3BucketMetric, opts v1.UpdateOptions) (result *v1alpha1.S3BucketMetric, err error) {
	result = &v1alpha1.S3BucketMetric{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		Name(s3BucketMetric.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3BucketMetric).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *s3BucketMetrics) UpdateStatus(ctx context.Context, s3BucketMetric *v1alpha1.S3BucketMetric, opts v1.UpdateOptions) (result *v1alpha1.S3BucketMetric, err error) {
	result = &v1alpha1.S3BucketMetric{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		Name(s3BucketMetric.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3BucketMetric).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the s3BucketMetric and deletes it. Returns an error if one occurs.
func (c *s3BucketMetrics) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3BucketMetrics) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched s3BucketMetric.
func (c *s3BucketMetrics) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S3BucketMetric, err error) {
	result = &v1alpha1.S3BucketMetric{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3bucketmetrics").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
