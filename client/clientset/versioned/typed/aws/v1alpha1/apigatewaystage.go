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

// ApiGatewayStagesGetter has a method to return a ApiGatewayStageInterface.
// A group's client should implement this interface.
type ApiGatewayStagesGetter interface {
	ApiGatewayStages() ApiGatewayStageInterface
}

// ApiGatewayStageInterface has methods to work with ApiGatewayStage resources.
type ApiGatewayStageInterface interface {
	Create(*v1alpha1.ApiGatewayStage) (*v1alpha1.ApiGatewayStage, error)
	Update(*v1alpha1.ApiGatewayStage) (*v1alpha1.ApiGatewayStage, error)
	UpdateStatus(*v1alpha1.ApiGatewayStage) (*v1alpha1.ApiGatewayStage, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiGatewayStage, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiGatewayStageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayStage, err error)
	ApiGatewayStageExpansion
}

// apiGatewayStages implements ApiGatewayStageInterface
type apiGatewayStages struct {
	client rest.Interface
}

// newApiGatewayStages returns a ApiGatewayStages
func newApiGatewayStages(c *AwsV1alpha1Client) *apiGatewayStages {
	return &apiGatewayStages{
		client: c.RESTClient(),
	}
}

// Get takes name of the apiGatewayStage, and returns the corresponding apiGatewayStage object, and an error if there is any.
func (c *apiGatewayStages) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayStage, err error) {
	result = &v1alpha1.ApiGatewayStage{}
	err = c.client.Get().
		Resource("apigatewaystages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiGatewayStages that match those selectors.
func (c *apiGatewayStages) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayStageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiGatewayStageList{}
	err = c.client.Get().
		Resource("apigatewaystages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiGatewayStages.
func (c *apiGatewayStages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apigatewaystages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiGatewayStage and creates it.  Returns the server's representation of the apiGatewayStage, and an error, if there is any.
func (c *apiGatewayStages) Create(apiGatewayStage *v1alpha1.ApiGatewayStage) (result *v1alpha1.ApiGatewayStage, err error) {
	result = &v1alpha1.ApiGatewayStage{}
	err = c.client.Post().
		Resource("apigatewaystages").
		Body(apiGatewayStage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiGatewayStage and updates it. Returns the server's representation of the apiGatewayStage, and an error, if there is any.
func (c *apiGatewayStages) Update(apiGatewayStage *v1alpha1.ApiGatewayStage) (result *v1alpha1.ApiGatewayStage, err error) {
	result = &v1alpha1.ApiGatewayStage{}
	err = c.client.Put().
		Resource("apigatewaystages").
		Name(apiGatewayStage.Name).
		Body(apiGatewayStage).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiGatewayStages) UpdateStatus(apiGatewayStage *v1alpha1.ApiGatewayStage) (result *v1alpha1.ApiGatewayStage, err error) {
	result = &v1alpha1.ApiGatewayStage{}
	err = c.client.Put().
		Resource("apigatewaystages").
		Name(apiGatewayStage.Name).
		SubResource("status").
		Body(apiGatewayStage).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiGatewayStage and deletes it. Returns an error if one occurs.
func (c *apiGatewayStages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apigatewaystages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiGatewayStages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apigatewaystages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiGatewayStage.
func (c *apiGatewayStages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayStage, err error) {
	result = &v1alpha1.ApiGatewayStage{}
	err = c.client.Patch(pt).
		Resource("apigatewaystages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
