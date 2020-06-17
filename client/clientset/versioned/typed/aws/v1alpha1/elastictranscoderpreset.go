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

// ElastictranscoderPresetsGetter has a method to return a ElastictranscoderPresetInterface.
// A group's client should implement this interface.
type ElastictranscoderPresetsGetter interface {
	ElastictranscoderPresets(namespace string) ElastictranscoderPresetInterface
}

// ElastictranscoderPresetInterface has methods to work with ElastictranscoderPreset resources.
type ElastictranscoderPresetInterface interface {
	Create(*v1alpha1.ElastictranscoderPreset) (*v1alpha1.ElastictranscoderPreset, error)
	Update(*v1alpha1.ElastictranscoderPreset) (*v1alpha1.ElastictranscoderPreset, error)
	UpdateStatus(*v1alpha1.ElastictranscoderPreset) (*v1alpha1.ElastictranscoderPreset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ElastictranscoderPreset, error)
	List(opts v1.ListOptions) (*v1alpha1.ElastictranscoderPresetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElastictranscoderPreset, err error)
	ElastictranscoderPresetExpansion
}

// elastictranscoderPresets implements ElastictranscoderPresetInterface
type elastictranscoderPresets struct {
	client rest.Interface
	ns     string
}

// newElastictranscoderPresets returns a ElastictranscoderPresets
func newElastictranscoderPresets(c *AwsV1alpha1Client, namespace string) *elastictranscoderPresets {
	return &elastictranscoderPresets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the elastictranscoderPreset, and returns the corresponding elastictranscoderPreset object, and an error if there is any.
func (c *elastictranscoderPresets) Get(name string, options v1.GetOptions) (result *v1alpha1.ElastictranscoderPreset, err error) {
	result = &v1alpha1.ElastictranscoderPreset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElastictranscoderPresets that match those selectors.
func (c *elastictranscoderPresets) List(opts v1.ListOptions) (result *v1alpha1.ElastictranscoderPresetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElastictranscoderPresetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elastictranscoderPresets.
func (c *elastictranscoderPresets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a elastictranscoderPreset and creates it.  Returns the server's representation of the elastictranscoderPreset, and an error, if there is any.
func (c *elastictranscoderPresets) Create(elastictranscoderPreset *v1alpha1.ElastictranscoderPreset) (result *v1alpha1.ElastictranscoderPreset, err error) {
	result = &v1alpha1.ElastictranscoderPreset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		Body(elastictranscoderPreset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a elastictranscoderPreset and updates it. Returns the server's representation of the elastictranscoderPreset, and an error, if there is any.
func (c *elastictranscoderPresets) Update(elastictranscoderPreset *v1alpha1.ElastictranscoderPreset) (result *v1alpha1.ElastictranscoderPreset, err error) {
	result = &v1alpha1.ElastictranscoderPreset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		Name(elastictranscoderPreset.Name).
		Body(elastictranscoderPreset).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *elastictranscoderPresets) UpdateStatus(elastictranscoderPreset *v1alpha1.ElastictranscoderPreset) (result *v1alpha1.ElastictranscoderPreset, err error) {
	result = &v1alpha1.ElastictranscoderPreset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		Name(elastictranscoderPreset.Name).
		SubResource("status").
		Body(elastictranscoderPreset).
		Do().
		Into(result)
	return
}

// Delete takes name of the elastictranscoderPreset and deletes it. Returns an error if one occurs.
func (c *elastictranscoderPresets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elastictranscoderPresets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched elastictranscoderPreset.
func (c *elastictranscoderPresets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElastictranscoderPreset, err error) {
	result = &v1alpha1.ElastictranscoderPreset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elastictranscoderpresets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
