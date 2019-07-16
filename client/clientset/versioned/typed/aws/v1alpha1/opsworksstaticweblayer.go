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

// OpsworksStaticWebLayersGetter has a method to return a OpsworksStaticWebLayerInterface.
// A group's client should implement this interface.
type OpsworksStaticWebLayersGetter interface {
	OpsworksStaticWebLayers() OpsworksStaticWebLayerInterface
}

// OpsworksStaticWebLayerInterface has methods to work with OpsworksStaticWebLayer resources.
type OpsworksStaticWebLayerInterface interface {
	Create(*v1alpha1.OpsworksStaticWebLayer) (*v1alpha1.OpsworksStaticWebLayer, error)
	Update(*v1alpha1.OpsworksStaticWebLayer) (*v1alpha1.OpsworksStaticWebLayer, error)
	UpdateStatus(*v1alpha1.OpsworksStaticWebLayer) (*v1alpha1.OpsworksStaticWebLayer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpsworksStaticWebLayer, error)
	List(opts v1.ListOptions) (*v1alpha1.OpsworksStaticWebLayerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksStaticWebLayer, err error)
	OpsworksStaticWebLayerExpansion
}

// opsworksStaticWebLayers implements OpsworksStaticWebLayerInterface
type opsworksStaticWebLayers struct {
	client rest.Interface
}

// newOpsworksStaticWebLayers returns a OpsworksStaticWebLayers
func newOpsworksStaticWebLayers(c *AwsV1alpha1Client) *opsworksStaticWebLayers {
	return &opsworksStaticWebLayers{
		client: c.RESTClient(),
	}
}

// Get takes name of the opsworksStaticWebLayer, and returns the corresponding opsworksStaticWebLayer object, and an error if there is any.
func (c *opsworksStaticWebLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksStaticWebLayer, err error) {
	result = &v1alpha1.OpsworksStaticWebLayer{}
	err = c.client.Get().
		Resource("opsworksstaticweblayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpsworksStaticWebLayers that match those selectors.
func (c *opsworksStaticWebLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksStaticWebLayerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpsworksStaticWebLayerList{}
	err = c.client.Get().
		Resource("opsworksstaticweblayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested opsworksStaticWebLayers.
func (c *opsworksStaticWebLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("opsworksstaticweblayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a opsworksStaticWebLayer and creates it.  Returns the server's representation of the opsworksStaticWebLayer, and an error, if there is any.
func (c *opsworksStaticWebLayers) Create(opsworksStaticWebLayer *v1alpha1.OpsworksStaticWebLayer) (result *v1alpha1.OpsworksStaticWebLayer, err error) {
	result = &v1alpha1.OpsworksStaticWebLayer{}
	err = c.client.Post().
		Resource("opsworksstaticweblayers").
		Body(opsworksStaticWebLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a opsworksStaticWebLayer and updates it. Returns the server's representation of the opsworksStaticWebLayer, and an error, if there is any.
func (c *opsworksStaticWebLayers) Update(opsworksStaticWebLayer *v1alpha1.OpsworksStaticWebLayer) (result *v1alpha1.OpsworksStaticWebLayer, err error) {
	result = &v1alpha1.OpsworksStaticWebLayer{}
	err = c.client.Put().
		Resource("opsworksstaticweblayers").
		Name(opsworksStaticWebLayer.Name).
		Body(opsworksStaticWebLayer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *opsworksStaticWebLayers) UpdateStatus(opsworksStaticWebLayer *v1alpha1.OpsworksStaticWebLayer) (result *v1alpha1.OpsworksStaticWebLayer, err error) {
	result = &v1alpha1.OpsworksStaticWebLayer{}
	err = c.client.Put().
		Resource("opsworksstaticweblayers").
		Name(opsworksStaticWebLayer.Name).
		SubResource("status").
		Body(opsworksStaticWebLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the opsworksStaticWebLayer and deletes it. Returns an error if one occurs.
func (c *opsworksStaticWebLayers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("opsworksstaticweblayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *opsworksStaticWebLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("opsworksstaticweblayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched opsworksStaticWebLayer.
func (c *opsworksStaticWebLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksStaticWebLayer, err error) {
	result = &v1alpha1.OpsworksStaticWebLayer{}
	err = c.client.Patch(pt).
		Resource("opsworksstaticweblayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
