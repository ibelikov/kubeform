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

// FakeApiGatewayDocumentationParts implements ApiGatewayDocumentationPartInterface
type FakeApiGatewayDocumentationParts struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewaydocumentationpartsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewaydocumentationparts"}

var apigatewaydocumentationpartsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayDocumentationPart"}

// Get takes name of the apiGatewayDocumentationPart, and returns the corresponding apiGatewayDocumentationPart object, and an error if there is any.
func (c *FakeApiGatewayDocumentationParts) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayDocumentationPart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewaydocumentationpartsResource, c.ns, name), &v1alpha1.ApiGatewayDocumentationPart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayDocumentationPart), err
}

// List takes label and field selectors, and returns the list of ApiGatewayDocumentationParts that match those selectors.
func (c *FakeApiGatewayDocumentationParts) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayDocumentationPartList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewaydocumentationpartsResource, apigatewaydocumentationpartsKind, c.ns, opts), &v1alpha1.ApiGatewayDocumentationPartList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayDocumentationPartList{ListMeta: obj.(*v1alpha1.ApiGatewayDocumentationPartList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayDocumentationPartList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayDocumentationParts.
func (c *FakeApiGatewayDocumentationParts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewaydocumentationpartsResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayDocumentationPart and creates it.  Returns the server's representation of the apiGatewayDocumentationPart, and an error, if there is any.
func (c *FakeApiGatewayDocumentationParts) Create(apiGatewayDocumentationPart *v1alpha1.ApiGatewayDocumentationPart) (result *v1alpha1.ApiGatewayDocumentationPart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewaydocumentationpartsResource, c.ns, apiGatewayDocumentationPart), &v1alpha1.ApiGatewayDocumentationPart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayDocumentationPart), err
}

// Update takes the representation of a apiGatewayDocumentationPart and updates it. Returns the server's representation of the apiGatewayDocumentationPart, and an error, if there is any.
func (c *FakeApiGatewayDocumentationParts) Update(apiGatewayDocumentationPart *v1alpha1.ApiGatewayDocumentationPart) (result *v1alpha1.ApiGatewayDocumentationPart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewaydocumentationpartsResource, c.ns, apiGatewayDocumentationPart), &v1alpha1.ApiGatewayDocumentationPart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayDocumentationPart), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayDocumentationParts) UpdateStatus(apiGatewayDocumentationPart *v1alpha1.ApiGatewayDocumentationPart) (*v1alpha1.ApiGatewayDocumentationPart, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewaydocumentationpartsResource, "status", c.ns, apiGatewayDocumentationPart), &v1alpha1.ApiGatewayDocumentationPart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayDocumentationPart), err
}

// Delete takes name of the apiGatewayDocumentationPart and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayDocumentationParts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewaydocumentationpartsResource, c.ns, name), &v1alpha1.ApiGatewayDocumentationPart{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayDocumentationParts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewaydocumentationpartsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayDocumentationPartList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayDocumentationPart.
func (c *FakeApiGatewayDocumentationParts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayDocumentationPart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewaydocumentationpartsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayDocumentationPart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayDocumentationPart), err
}
