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

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigConfigurationRecorderStatus_s implements ConfigConfigurationRecorderStatus_Interface
type FakeConfigConfigurationRecorderStatus_s struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var configconfigurationrecorderstatus_sResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "configconfigurationrecorderstatus_s"}

var configconfigurationrecorderstatus_sKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ConfigConfigurationRecorderStatus_"}

// Get takes name of the configConfigurationRecorderStatus_, and returns the corresponding configConfigurationRecorderStatus_ object, and an error if there is any.
func (c *FakeConfigConfigurationRecorderStatus_s) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configconfigurationrecorderstatus_sResource, c.ns, name), &v1alpha1.ConfigConfigurationRecorderStatus_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigConfigurationRecorderStatus_), err
}

// List takes label and field selectors, and returns the list of ConfigConfigurationRecorderStatus_s that match those selectors.
func (c *FakeConfigConfigurationRecorderStatus_s) List(opts v1.ListOptions) (result *v1alpha1.ConfigConfigurationRecorderStatus_List, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configconfigurationrecorderstatus_sResource, configconfigurationrecorderstatus_sKind, c.ns, opts), &v1alpha1.ConfigConfigurationRecorderStatus_List{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigConfigurationRecorderStatus_List{ListMeta: obj.(*v1alpha1.ConfigConfigurationRecorderStatus_List).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigConfigurationRecorderStatus_List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configConfigurationRecorderStatus_s.
func (c *FakeConfigConfigurationRecorderStatus_s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configconfigurationrecorderstatus_sResource, c.ns, opts))

}

// Create takes the representation of a configConfigurationRecorderStatus_ and creates it.  Returns the server's representation of the configConfigurationRecorderStatus_, and an error, if there is any.
func (c *FakeConfigConfigurationRecorderStatus_s) Create(configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configconfigurationrecorderstatus_sResource, c.ns, configConfigurationRecorderStatus_), &v1alpha1.ConfigConfigurationRecorderStatus_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigConfigurationRecorderStatus_), err
}

// Update takes the representation of a configConfigurationRecorderStatus_ and updates it. Returns the server's representation of the configConfigurationRecorderStatus_, and an error, if there is any.
func (c *FakeConfigConfigurationRecorderStatus_s) Update(configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configconfigurationrecorderstatus_sResource, c.ns, configConfigurationRecorderStatus_), &v1alpha1.ConfigConfigurationRecorderStatus_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigConfigurationRecorderStatus_), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigConfigurationRecorderStatus_s) UpdateStatus(configConfigurationRecorderStatus_ *v1alpha1.ConfigConfigurationRecorderStatus_) (*v1alpha1.ConfigConfigurationRecorderStatus_, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configconfigurationrecorderstatus_sResource, "status", c.ns, configConfigurationRecorderStatus_), &v1alpha1.ConfigConfigurationRecorderStatus_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigConfigurationRecorderStatus_), err
}

// Delete takes name of the configConfigurationRecorderStatus_ and deletes it. Returns an error if one occurs.
func (c *FakeConfigConfigurationRecorderStatus_s) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configconfigurationrecorderstatus_sResource, c.ns, name), &v1alpha1.ConfigConfigurationRecorderStatus_{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigConfigurationRecorderStatus_s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configconfigurationrecorderstatus_sResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigConfigurationRecorderStatus_List{})
	return err
}

// Patch applies the patch and returns the patched configConfigurationRecorderStatus_.
func (c *FakeConfigConfigurationRecorderStatus_s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigConfigurationRecorderStatus_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configconfigurationrecorderstatus_sResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigConfigurationRecorderStatus_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigConfigurationRecorderStatus_), err
}
