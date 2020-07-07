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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeApiGatewayMethods implements ApiGatewayMethodInterface
type FakeApiGatewayMethods struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewaymethodsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewaymethods"}

var apigatewaymethodsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayMethod"}

// Get takes name of the apiGatewayMethod, and returns the corresponding apiGatewayMethod object, and an error if there is any.
func (c *FakeApiGatewayMethods) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewaymethodsResource, c.ns, name), &v1alpha1.ApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethod), err
}

// List takes label and field selectors, and returns the list of ApiGatewayMethods that match those selectors.
func (c *FakeApiGatewayMethods) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiGatewayMethodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewaymethodsResource, apigatewaymethodsKind, c.ns, opts), &v1alpha1.ApiGatewayMethodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayMethodList{ListMeta: obj.(*v1alpha1.ApiGatewayMethodList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayMethodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayMethods.
func (c *FakeApiGatewayMethods) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewaymethodsResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayMethod and creates it.  Returns the server's representation of the apiGatewayMethod, and an error, if there is any.
func (c *FakeApiGatewayMethods) Create(ctx context.Context, apiGatewayMethod *v1alpha1.ApiGatewayMethod, opts v1.CreateOptions) (result *v1alpha1.ApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewaymethodsResource, c.ns, apiGatewayMethod), &v1alpha1.ApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethod), err
}

// Update takes the representation of a apiGatewayMethod and updates it. Returns the server's representation of the apiGatewayMethod, and an error, if there is any.
func (c *FakeApiGatewayMethods) Update(ctx context.Context, apiGatewayMethod *v1alpha1.ApiGatewayMethod, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewaymethodsResource, c.ns, apiGatewayMethod), &v1alpha1.ApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethod), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayMethods) UpdateStatus(ctx context.Context, apiGatewayMethod *v1alpha1.ApiGatewayMethod, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayMethod, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewaymethodsResource, "status", c.ns, apiGatewayMethod), &v1alpha1.ApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethod), err
}

// Delete takes name of the apiGatewayMethod and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayMethods) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewaymethodsResource, c.ns, name), &v1alpha1.ApiGatewayMethod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayMethods) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewaymethodsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayMethodList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayMethod.
func (c *FakeApiGatewayMethods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayMethod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewaymethodsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayMethod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethod), err
}
