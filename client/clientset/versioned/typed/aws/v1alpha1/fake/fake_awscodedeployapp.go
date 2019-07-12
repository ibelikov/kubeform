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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsCodedeployApps implements AwsCodedeployAppInterface
type FakeAwsCodedeployApps struct {
	Fake *FakeAwsV1alpha1
}

var awscodedeployappsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awscodedeployapps"}

var awscodedeployappsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsCodedeployApp"}

// Get takes name of the awsCodedeployApp, and returns the corresponding awsCodedeployApp object, and an error if there is any.
func (c *FakeAwsCodedeployApps) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awscodedeployappsResource, name), &v1alpha1.AwsCodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployApp), err
}

// List takes label and field selectors, and returns the list of AwsCodedeployApps that match those selectors.
func (c *FakeAwsCodedeployApps) List(opts v1.ListOptions) (result *v1alpha1.AwsCodedeployAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awscodedeployappsResource, awscodedeployappsKind, opts), &v1alpha1.AwsCodedeployAppList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsCodedeployAppList{ListMeta: obj.(*v1alpha1.AwsCodedeployAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsCodedeployAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCodedeployApps.
func (c *FakeAwsCodedeployApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awscodedeployappsResource, opts))
}

// Create takes the representation of a awsCodedeployApp and creates it.  Returns the server's representation of the awsCodedeployApp, and an error, if there is any.
func (c *FakeAwsCodedeployApps) Create(awsCodedeployApp *v1alpha1.AwsCodedeployApp) (result *v1alpha1.AwsCodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awscodedeployappsResource, awsCodedeployApp), &v1alpha1.AwsCodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployApp), err
}

// Update takes the representation of a awsCodedeployApp and updates it. Returns the server's representation of the awsCodedeployApp, and an error, if there is any.
func (c *FakeAwsCodedeployApps) Update(awsCodedeployApp *v1alpha1.AwsCodedeployApp) (result *v1alpha1.AwsCodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awscodedeployappsResource, awsCodedeployApp), &v1alpha1.AwsCodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsCodedeployApps) UpdateStatus(awsCodedeployApp *v1alpha1.AwsCodedeployApp) (*v1alpha1.AwsCodedeployApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awscodedeployappsResource, "status", awsCodedeployApp), &v1alpha1.AwsCodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployApp), err
}

// Delete takes name of the awsCodedeployApp and deletes it. Returns an error if one occurs.
func (c *FakeAwsCodedeployApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awscodedeployappsResource, name), &v1alpha1.AwsCodedeployApp{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCodedeployApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awscodedeployappsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsCodedeployAppList{})
	return err
}

// Patch applies the patch and returns the patched awsCodedeployApp.
func (c *FakeAwsCodedeployApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awscodedeployappsResource, name, pt, data, subresources...), &v1alpha1.AwsCodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCodedeployApp), err
}
