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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeCloudiotRegistries implements CloudiotRegistryInterface
type FakeCloudiotRegistries struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var cloudiotregistriesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "cloudiotregistries"}

var cloudiotregistriesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "CloudiotRegistry"}

// Get takes name of the cloudiotRegistry, and returns the corresponding cloudiotRegistry object, and an error if there is any.
func (c *FakeCloudiotRegistries) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudiotRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudiotregistriesResource, c.ns, name), &v1alpha1.CloudiotRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudiotRegistry), err
}

// List takes label and field selectors, and returns the list of CloudiotRegistries that match those selectors.
func (c *FakeCloudiotRegistries) List(opts v1.ListOptions) (result *v1alpha1.CloudiotRegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudiotregistriesResource, cloudiotregistriesKind, c.ns, opts), &v1alpha1.CloudiotRegistryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudiotRegistryList{ListMeta: obj.(*v1alpha1.CloudiotRegistryList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudiotRegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudiotRegistries.
func (c *FakeCloudiotRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudiotregistriesResource, c.ns, opts))

}

// Create takes the representation of a cloudiotRegistry and creates it.  Returns the server's representation of the cloudiotRegistry, and an error, if there is any.
func (c *FakeCloudiotRegistries) Create(cloudiotRegistry *v1alpha1.CloudiotRegistry) (result *v1alpha1.CloudiotRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudiotregistriesResource, c.ns, cloudiotRegistry), &v1alpha1.CloudiotRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudiotRegistry), err
}

// Update takes the representation of a cloudiotRegistry and updates it. Returns the server's representation of the cloudiotRegistry, and an error, if there is any.
func (c *FakeCloudiotRegistries) Update(cloudiotRegistry *v1alpha1.CloudiotRegistry) (result *v1alpha1.CloudiotRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudiotregistriesResource, c.ns, cloudiotRegistry), &v1alpha1.CloudiotRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudiotRegistry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudiotRegistries) UpdateStatus(cloudiotRegistry *v1alpha1.CloudiotRegistry) (*v1alpha1.CloudiotRegistry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudiotregistriesResource, "status", c.ns, cloudiotRegistry), &v1alpha1.CloudiotRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudiotRegistry), err
}

// Delete takes name of the cloudiotRegistry and deletes it. Returns an error if one occurs.
func (c *FakeCloudiotRegistries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudiotregistriesResource, c.ns, name), &v1alpha1.CloudiotRegistry{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudiotRegistries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudiotregistriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudiotRegistryList{})
	return err
}

// Patch applies the patch and returns the patched cloudiotRegistry.
func (c *FakeCloudiotRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudiotRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudiotregistriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudiotRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudiotRegistry), err
}
