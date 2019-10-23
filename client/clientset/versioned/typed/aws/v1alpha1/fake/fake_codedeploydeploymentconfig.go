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

// FakeCodedeployDeploymentConfigs implements CodedeployDeploymentConfigInterface
type FakeCodedeployDeploymentConfigs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var codedeploydeploymentconfigsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "codedeploydeploymentconfigs"}

var codedeploydeploymentconfigsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CodedeployDeploymentConfig"}

// Get takes name of the codedeployDeploymentConfig, and returns the corresponding codedeployDeploymentConfig object, and an error if there is any.
func (c *FakeCodedeployDeploymentConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(codedeploydeploymentconfigsResource, c.ns, name), &v1alpha1.CodedeployDeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentConfig), err
}

// List takes label and field selectors, and returns the list of CodedeployDeploymentConfigs that match those selectors.
func (c *FakeCodedeployDeploymentConfigs) List(opts v1.ListOptions) (result *v1alpha1.CodedeployDeploymentConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(codedeploydeploymentconfigsResource, codedeploydeploymentconfigsKind, c.ns, opts), &v1alpha1.CodedeployDeploymentConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CodedeployDeploymentConfigList{ListMeta: obj.(*v1alpha1.CodedeployDeploymentConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.CodedeployDeploymentConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested codedeployDeploymentConfigs.
func (c *FakeCodedeployDeploymentConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(codedeploydeploymentconfigsResource, c.ns, opts))

}

// Create takes the representation of a codedeployDeploymentConfig and creates it.  Returns the server's representation of the codedeployDeploymentConfig, and an error, if there is any.
func (c *FakeCodedeployDeploymentConfigs) Create(codedeployDeploymentConfig *v1alpha1.CodedeployDeploymentConfig) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(codedeploydeploymentconfigsResource, c.ns, codedeployDeploymentConfig), &v1alpha1.CodedeployDeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentConfig), err
}

// Update takes the representation of a codedeployDeploymentConfig and updates it. Returns the server's representation of the codedeployDeploymentConfig, and an error, if there is any.
func (c *FakeCodedeployDeploymentConfigs) Update(codedeployDeploymentConfig *v1alpha1.CodedeployDeploymentConfig) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(codedeploydeploymentconfigsResource, c.ns, codedeployDeploymentConfig), &v1alpha1.CodedeployDeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCodedeployDeploymentConfigs) UpdateStatus(codedeployDeploymentConfig *v1alpha1.CodedeployDeploymentConfig) (*v1alpha1.CodedeployDeploymentConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(codedeploydeploymentconfigsResource, "status", c.ns, codedeployDeploymentConfig), &v1alpha1.CodedeployDeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentConfig), err
}

// Delete takes name of the codedeployDeploymentConfig and deletes it. Returns an error if one occurs.
func (c *FakeCodedeployDeploymentConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(codedeploydeploymentconfigsResource, c.ns, name), &v1alpha1.CodedeployDeploymentConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCodedeployDeploymentConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(codedeploydeploymentconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CodedeployDeploymentConfigList{})
	return err
}

// Patch applies the patch and returns the patched codedeployDeploymentConfig.
func (c *FakeCodedeployDeploymentConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodedeployDeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(codedeploydeploymentconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CodedeployDeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployDeploymentConfig), err
}
