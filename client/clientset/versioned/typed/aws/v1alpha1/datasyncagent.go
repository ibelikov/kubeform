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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatasyncAgentsGetter has a method to return a DatasyncAgentInterface.
// A group's client should implement this interface.
type DatasyncAgentsGetter interface {
	DatasyncAgents(namespace string) DatasyncAgentInterface
}

// DatasyncAgentInterface has methods to work with DatasyncAgent resources.
type DatasyncAgentInterface interface {
	Create(*v1alpha1.DatasyncAgent) (*v1alpha1.DatasyncAgent, error)
	Update(*v1alpha1.DatasyncAgent) (*v1alpha1.DatasyncAgent, error)
	UpdateStatus(*v1alpha1.DatasyncAgent) (*v1alpha1.DatasyncAgent, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DatasyncAgent, error)
	List(opts v1.ListOptions) (*v1alpha1.DatasyncAgentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatasyncAgent, err error)
	DatasyncAgentExpansion
}

// datasyncAgents implements DatasyncAgentInterface
type datasyncAgents struct {
	client rest.Interface
	ns     string
}

// newDatasyncAgents returns a DatasyncAgents
func newDatasyncAgents(c *AwsV1alpha1Client, namespace string) *datasyncAgents {
	return &datasyncAgents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the datasyncAgent, and returns the corresponding datasyncAgent object, and an error if there is any.
func (c *datasyncAgents) Get(name string, options v1.GetOptions) (result *v1alpha1.DatasyncAgent, err error) {
	result = &v1alpha1.DatasyncAgent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasyncagents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DatasyncAgents that match those selectors.
func (c *datasyncAgents) List(opts v1.ListOptions) (result *v1alpha1.DatasyncAgentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatasyncAgentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasyncagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested datasyncAgents.
func (c *datasyncAgents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datasyncagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a datasyncAgent and creates it.  Returns the server's representation of the datasyncAgent, and an error, if there is any.
func (c *datasyncAgents) Create(datasyncAgent *v1alpha1.DatasyncAgent) (result *v1alpha1.DatasyncAgent, err error) {
	result = &v1alpha1.DatasyncAgent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datasyncagents").
		Body(datasyncAgent).
		Do().
		Into(result)
	return
}

// Update takes the representation of a datasyncAgent and updates it. Returns the server's representation of the datasyncAgent, and an error, if there is any.
func (c *datasyncAgents) Update(datasyncAgent *v1alpha1.DatasyncAgent) (result *v1alpha1.DatasyncAgent, err error) {
	result = &v1alpha1.DatasyncAgent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasyncagents").
		Name(datasyncAgent.Name).
		Body(datasyncAgent).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *datasyncAgents) UpdateStatus(datasyncAgent *v1alpha1.DatasyncAgent) (result *v1alpha1.DatasyncAgent, err error) {
	result = &v1alpha1.DatasyncAgent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasyncagents").
		Name(datasyncAgent.Name).
		SubResource("status").
		Body(datasyncAgent).
		Do().
		Into(result)
	return
}

// Delete takes name of the datasyncAgent and deletes it. Returns an error if one occurs.
func (c *datasyncAgents) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasyncagents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *datasyncAgents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasyncagents").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched datasyncAgent.
func (c *datasyncAgents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatasyncAgent, err error) {
	result = &v1alpha1.DatasyncAgent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datasyncagents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
