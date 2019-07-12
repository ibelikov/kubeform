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

// FakeAwsApiGatewayStages implements AwsApiGatewayStageInterface
type FakeAwsApiGatewayStages struct {
	Fake *FakeAwsV1alpha1
}

var awsapigatewaystagesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsapigatewaystages"}

var awsapigatewaystagesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsApiGatewayStage"}

// Get takes name of the awsApiGatewayStage, and returns the corresponding awsApiGatewayStage object, and an error if there is any.
func (c *FakeAwsApiGatewayStages) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsapigatewaystagesResource, name), &v1alpha1.AwsApiGatewayStage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayStage), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayStages that match those selectors.
func (c *FakeAwsApiGatewayStages) List(opts v1.ListOptions) (result *v1alpha1.AwsApiGatewayStageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsapigatewaystagesResource, awsapigatewaystagesKind, opts), &v1alpha1.AwsApiGatewayStageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsApiGatewayStageList{ListMeta: obj.(*v1alpha1.AwsApiGatewayStageList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsApiGatewayStageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayStages.
func (c *FakeAwsApiGatewayStages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsapigatewaystagesResource, opts))
}

// Create takes the representation of a awsApiGatewayStage and creates it.  Returns the server's representation of the awsApiGatewayStage, and an error, if there is any.
func (c *FakeAwsApiGatewayStages) Create(awsApiGatewayStage *v1alpha1.AwsApiGatewayStage) (result *v1alpha1.AwsApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsapigatewaystagesResource, awsApiGatewayStage), &v1alpha1.AwsApiGatewayStage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayStage), err
}

// Update takes the representation of a awsApiGatewayStage and updates it. Returns the server's representation of the awsApiGatewayStage, and an error, if there is any.
func (c *FakeAwsApiGatewayStages) Update(awsApiGatewayStage *v1alpha1.AwsApiGatewayStage) (result *v1alpha1.AwsApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsapigatewaystagesResource, awsApiGatewayStage), &v1alpha1.AwsApiGatewayStage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayStage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsApiGatewayStages) UpdateStatus(awsApiGatewayStage *v1alpha1.AwsApiGatewayStage) (*v1alpha1.AwsApiGatewayStage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsapigatewaystagesResource, "status", awsApiGatewayStage), &v1alpha1.AwsApiGatewayStage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayStage), err
}

// Delete takes name of the awsApiGatewayStage and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayStages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsapigatewaystagesResource, name), &v1alpha1.AwsApiGatewayStage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayStages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsapigatewaystagesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsApiGatewayStageList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayStage.
func (c *FakeAwsApiGatewayStages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsapigatewaystagesResource, name, pt, data, subresources...), &v1alpha1.AwsApiGatewayStage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsApiGatewayStage), err
}
