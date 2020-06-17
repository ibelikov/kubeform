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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// FolderIamPoliciesGetter has a method to return a FolderIamPolicyInterface.
// A group's client should implement this interface.
type FolderIamPoliciesGetter interface {
	FolderIamPolicies(namespace string) FolderIamPolicyInterface
}

// FolderIamPolicyInterface has methods to work with FolderIamPolicy resources.
type FolderIamPolicyInterface interface {
	Create(*v1alpha1.FolderIamPolicy) (*v1alpha1.FolderIamPolicy, error)
	Update(*v1alpha1.FolderIamPolicy) (*v1alpha1.FolderIamPolicy, error)
	UpdateStatus(*v1alpha1.FolderIamPolicy) (*v1alpha1.FolderIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FolderIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.FolderIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FolderIamPolicy, err error)
	FolderIamPolicyExpansion
}

// folderIamPolicies implements FolderIamPolicyInterface
type folderIamPolicies struct {
	client rest.Interface
	ns     string
}

// newFolderIamPolicies returns a FolderIamPolicies
func newFolderIamPolicies(c *GoogleV1alpha1Client, namespace string) *folderIamPolicies {
	return &folderIamPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the folderIamPolicy, and returns the corresponding folderIamPolicy object, and an error if there is any.
func (c *folderIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.FolderIamPolicy, err error) {
	result = &v1alpha1.FolderIamPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("folderiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FolderIamPolicies that match those selectors.
func (c *folderIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.FolderIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FolderIamPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("folderiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested folderIamPolicies.
func (c *folderIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("folderiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a folderIamPolicy and creates it.  Returns the server's representation of the folderIamPolicy, and an error, if there is any.
func (c *folderIamPolicies) Create(folderIamPolicy *v1alpha1.FolderIamPolicy) (result *v1alpha1.FolderIamPolicy, err error) {
	result = &v1alpha1.FolderIamPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("folderiampolicies").
		Body(folderIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a folderIamPolicy and updates it. Returns the server's representation of the folderIamPolicy, and an error, if there is any.
func (c *folderIamPolicies) Update(folderIamPolicy *v1alpha1.FolderIamPolicy) (result *v1alpha1.FolderIamPolicy, err error) {
	result = &v1alpha1.FolderIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("folderiampolicies").
		Name(folderIamPolicy.Name).
		Body(folderIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *folderIamPolicies) UpdateStatus(folderIamPolicy *v1alpha1.FolderIamPolicy) (result *v1alpha1.FolderIamPolicy, err error) {
	result = &v1alpha1.FolderIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("folderiampolicies").
		Name(folderIamPolicy.Name).
		SubResource("status").
		Body(folderIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the folderIamPolicy and deletes it. Returns an error if one occurs.
func (c *folderIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("folderiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *folderIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("folderiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched folderIamPolicy.
func (c *folderIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FolderIamPolicy, err error) {
	result = &v1alpha1.FolderIamPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("folderiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
