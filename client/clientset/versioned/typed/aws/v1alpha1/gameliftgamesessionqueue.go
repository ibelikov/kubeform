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

// GameliftGameSessionQueuesGetter has a method to return a GameliftGameSessionQueueInterface.
// A group's client should implement this interface.
type GameliftGameSessionQueuesGetter interface {
	GameliftGameSessionQueues() GameliftGameSessionQueueInterface
}

// GameliftGameSessionQueueInterface has methods to work with GameliftGameSessionQueue resources.
type GameliftGameSessionQueueInterface interface {
	Create(*v1alpha1.GameliftGameSessionQueue) (*v1alpha1.GameliftGameSessionQueue, error)
	Update(*v1alpha1.GameliftGameSessionQueue) (*v1alpha1.GameliftGameSessionQueue, error)
	UpdateStatus(*v1alpha1.GameliftGameSessionQueue) (*v1alpha1.GameliftGameSessionQueue, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GameliftGameSessionQueue, error)
	List(opts v1.ListOptions) (*v1alpha1.GameliftGameSessionQueueList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameliftGameSessionQueue, err error)
	GameliftGameSessionQueueExpansion
}

// gameliftGameSessionQueues implements GameliftGameSessionQueueInterface
type gameliftGameSessionQueues struct {
	client rest.Interface
}

// newGameliftGameSessionQueues returns a GameliftGameSessionQueues
func newGameliftGameSessionQueues(c *AwsV1alpha1Client) *gameliftGameSessionQueues {
	return &gameliftGameSessionQueues{
		client: c.RESTClient(),
	}
}

// Get takes name of the gameliftGameSessionQueue, and returns the corresponding gameliftGameSessionQueue object, and an error if there is any.
func (c *gameliftGameSessionQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	result = &v1alpha1.GameliftGameSessionQueue{}
	err = c.client.Get().
		Resource("gameliftgamesessionqueues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GameliftGameSessionQueues that match those selectors.
func (c *gameliftGameSessionQueues) List(opts v1.ListOptions) (result *v1alpha1.GameliftGameSessionQueueList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GameliftGameSessionQueueList{}
	err = c.client.Get().
		Resource("gameliftgamesessionqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gameliftGameSessionQueues.
func (c *gameliftGameSessionQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("gameliftgamesessionqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gameliftGameSessionQueue and creates it.  Returns the server's representation of the gameliftGameSessionQueue, and an error, if there is any.
func (c *gameliftGameSessionQueues) Create(gameliftGameSessionQueue *v1alpha1.GameliftGameSessionQueue) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	result = &v1alpha1.GameliftGameSessionQueue{}
	err = c.client.Post().
		Resource("gameliftgamesessionqueues").
		Body(gameliftGameSessionQueue).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gameliftGameSessionQueue and updates it. Returns the server's representation of the gameliftGameSessionQueue, and an error, if there is any.
func (c *gameliftGameSessionQueues) Update(gameliftGameSessionQueue *v1alpha1.GameliftGameSessionQueue) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	result = &v1alpha1.GameliftGameSessionQueue{}
	err = c.client.Put().
		Resource("gameliftgamesessionqueues").
		Name(gameliftGameSessionQueue.Name).
		Body(gameliftGameSessionQueue).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gameliftGameSessionQueues) UpdateStatus(gameliftGameSessionQueue *v1alpha1.GameliftGameSessionQueue) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	result = &v1alpha1.GameliftGameSessionQueue{}
	err = c.client.Put().
		Resource("gameliftgamesessionqueues").
		Name(gameliftGameSessionQueue.Name).
		SubResource("status").
		Body(gameliftGameSessionQueue).
		Do().
		Into(result)
	return
}

// Delete takes name of the gameliftGameSessionQueue and deletes it. Returns an error if one occurs.
func (c *gameliftGameSessionQueues) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("gameliftgamesessionqueues").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gameliftGameSessionQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("gameliftgamesessionqueues").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gameliftGameSessionQueue.
func (c *gameliftGameSessionQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	result = &v1alpha1.GameliftGameSessionQueue{}
	err = c.client.Patch(pt).
		Resource("gameliftgamesessionqueues").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
