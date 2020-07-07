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

// FakeApiGatewayAuthorizers implements ApiGatewayAuthorizerInterface
type FakeApiGatewayAuthorizers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewayauthorizersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewayauthorizers"}

var apigatewayauthorizersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayAuthorizer"}

// Get takes name of the apiGatewayAuthorizer, and returns the corresponding apiGatewayAuthorizer object, and an error if there is any.
func (c *FakeApiGatewayAuthorizers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayAuthorizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewayauthorizersResource, c.ns, name), &v1alpha1.ApiGatewayAuthorizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayAuthorizer), err
}

// List takes label and field selectors, and returns the list of ApiGatewayAuthorizers that match those selectors.
func (c *FakeApiGatewayAuthorizers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiGatewayAuthorizerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewayauthorizersResource, apigatewayauthorizersKind, c.ns, opts), &v1alpha1.ApiGatewayAuthorizerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayAuthorizerList{ListMeta: obj.(*v1alpha1.ApiGatewayAuthorizerList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayAuthorizerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayAuthorizers.
func (c *FakeApiGatewayAuthorizers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewayauthorizersResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayAuthorizer and creates it.  Returns the server's representation of the apiGatewayAuthorizer, and an error, if there is any.
func (c *FakeApiGatewayAuthorizers) Create(ctx context.Context, apiGatewayAuthorizer *v1alpha1.ApiGatewayAuthorizer, opts v1.CreateOptions) (result *v1alpha1.ApiGatewayAuthorizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewayauthorizersResource, c.ns, apiGatewayAuthorizer), &v1alpha1.ApiGatewayAuthorizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayAuthorizer), err
}

// Update takes the representation of a apiGatewayAuthorizer and updates it. Returns the server's representation of the apiGatewayAuthorizer, and an error, if there is any.
func (c *FakeApiGatewayAuthorizers) Update(ctx context.Context, apiGatewayAuthorizer *v1alpha1.ApiGatewayAuthorizer, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayAuthorizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewayauthorizersResource, c.ns, apiGatewayAuthorizer), &v1alpha1.ApiGatewayAuthorizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayAuthorizer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayAuthorizers) UpdateStatus(ctx context.Context, apiGatewayAuthorizer *v1alpha1.ApiGatewayAuthorizer, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayAuthorizer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewayauthorizersResource, "status", c.ns, apiGatewayAuthorizer), &v1alpha1.ApiGatewayAuthorizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayAuthorizer), err
}

// Delete takes name of the apiGatewayAuthorizer and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayAuthorizers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewayauthorizersResource, c.ns, name), &v1alpha1.ApiGatewayAuthorizer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayAuthorizers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewayauthorizersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayAuthorizerList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayAuthorizer.
func (c *FakeApiGatewayAuthorizers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayAuthorizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewayauthorizersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayAuthorizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayAuthorizer), err
}
