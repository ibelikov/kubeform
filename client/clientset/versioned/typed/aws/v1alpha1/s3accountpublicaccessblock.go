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

// S3AccountPublicAccessBlocksGetter has a method to return a S3AccountPublicAccessBlockInterface.
// A group's client should implement this interface.
type S3AccountPublicAccessBlocksGetter interface {
	S3AccountPublicAccessBlocks(namespace string) S3AccountPublicAccessBlockInterface
}

// S3AccountPublicAccessBlockInterface has methods to work with S3AccountPublicAccessBlock resources.
type S3AccountPublicAccessBlockInterface interface {
	Create(ctx context.Context, s3AccountPublicAccessBlock *v1alpha1.S3AccountPublicAccessBlock, opts v1.CreateOptions) (*v1alpha1.S3AccountPublicAccessBlock, error)
	Update(ctx context.Context, s3AccountPublicAccessBlock *v1alpha1.S3AccountPublicAccessBlock, opts v1.UpdateOptions) (*v1alpha1.S3AccountPublicAccessBlock, error)
	UpdateStatus(ctx context.Context, s3AccountPublicAccessBlock *v1alpha1.S3AccountPublicAccessBlock, opts v1.UpdateOptions) (*v1alpha1.S3AccountPublicAccessBlock, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.S3AccountPublicAccessBlock, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.S3AccountPublicAccessBlockList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S3AccountPublicAccessBlock, err error)
	S3AccountPublicAccessBlockExpansion
}

// s3AccountPublicAccessBlocks implements S3AccountPublicAccessBlockInterface
type s3AccountPublicAccessBlocks struct {
	client rest.Interface
	ns     string
}

// newS3AccountPublicAccessBlocks returns a S3AccountPublicAccessBlocks
func newS3AccountPublicAccessBlocks(c *AwsV1alpha1Client, namespace string) *s3AccountPublicAccessBlocks {
	return &s3AccountPublicAccessBlocks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3AccountPublicAccessBlock, and returns the corresponding s3AccountPublicAccessBlock object, and an error if there is any.
func (c *s3AccountPublicAccessBlocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.S3AccountPublicAccessBlock, err error) {
	result = &v1alpha1.S3AccountPublicAccessBlock{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3AccountPublicAccessBlocks that match those selectors.
func (c *s3AccountPublicAccessBlocks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.S3AccountPublicAccessBlockList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S3AccountPublicAccessBlockList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3AccountPublicAccessBlocks.
func (c *s3AccountPublicAccessBlocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a s3AccountPublicAccessBlock and creates it.  Returns the server's representation of the s3AccountPublicAccessBlock, and an error, if there is any.
func (c *s3AccountPublicAccessBlocks) Create(ctx context.Context, s3AccountPublicAccessBlock *v1alpha1.S3AccountPublicAccessBlock, opts v1.CreateOptions) (result *v1alpha1.S3AccountPublicAccessBlock, err error) {
	result = &v1alpha1.S3AccountPublicAccessBlock{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3AccountPublicAccessBlock).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a s3AccountPublicAccessBlock and updates it. Returns the server's representation of the s3AccountPublicAccessBlock, and an error, if there is any.
func (c *s3AccountPublicAccessBlocks) Update(ctx context.Context, s3AccountPublicAccessBlock *v1alpha1.S3AccountPublicAccessBlock, opts v1.UpdateOptions) (result *v1alpha1.S3AccountPublicAccessBlock, err error) {
	result = &v1alpha1.S3AccountPublicAccessBlock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		Name(s3AccountPublicAccessBlock.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3AccountPublicAccessBlock).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *s3AccountPublicAccessBlocks) UpdateStatus(ctx context.Context, s3AccountPublicAccessBlock *v1alpha1.S3AccountPublicAccessBlock, opts v1.UpdateOptions) (result *v1alpha1.S3AccountPublicAccessBlock, err error) {
	result = &v1alpha1.S3AccountPublicAccessBlock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		Name(s3AccountPublicAccessBlock.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s3AccountPublicAccessBlock).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the s3AccountPublicAccessBlock and deletes it. Returns an error if one occurs.
func (c *s3AccountPublicAccessBlocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3AccountPublicAccessBlocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched s3AccountPublicAccessBlock.
func (c *s3AccountPublicAccessBlocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S3AccountPublicAccessBlock, err error) {
	result = &v1alpha1.S3AccountPublicAccessBlock{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3accountpublicaccessblocks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
