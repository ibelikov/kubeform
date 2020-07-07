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

// FakeApiGatewayStages implements ApiGatewayStageInterface
type FakeApiGatewayStages struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewaystagesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewaystages"}

var apigatewaystagesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayStage"}

// Get takes name of the apiGatewayStage, and returns the corresponding apiGatewayStage object, and an error if there is any.
func (c *FakeApiGatewayStages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewaystagesResource, c.ns, name), &v1alpha1.ApiGatewayStage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayStage), err
}

// List takes label and field selectors, and returns the list of ApiGatewayStages that match those selectors.
func (c *FakeApiGatewayStages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiGatewayStageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewaystagesResource, apigatewaystagesKind, c.ns, opts), &v1alpha1.ApiGatewayStageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayStageList{ListMeta: obj.(*v1alpha1.ApiGatewayStageList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayStageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayStages.
func (c *FakeApiGatewayStages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewaystagesResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayStage and creates it.  Returns the server's representation of the apiGatewayStage, and an error, if there is any.
func (c *FakeApiGatewayStages) Create(ctx context.Context, apiGatewayStage *v1alpha1.ApiGatewayStage, opts v1.CreateOptions) (result *v1alpha1.ApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewaystagesResource, c.ns, apiGatewayStage), &v1alpha1.ApiGatewayStage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayStage), err
}

// Update takes the representation of a apiGatewayStage and updates it. Returns the server's representation of the apiGatewayStage, and an error, if there is any.
func (c *FakeApiGatewayStages) Update(ctx context.Context, apiGatewayStage *v1alpha1.ApiGatewayStage, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewaystagesResource, c.ns, apiGatewayStage), &v1alpha1.ApiGatewayStage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayStage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayStages) UpdateStatus(ctx context.Context, apiGatewayStage *v1alpha1.ApiGatewayStage, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayStage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewaystagesResource, "status", c.ns, apiGatewayStage), &v1alpha1.ApiGatewayStage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayStage), err
}

// Delete takes name of the apiGatewayStage and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayStages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewaystagesResource, c.ns, name), &v1alpha1.ApiGatewayStage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayStages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewaystagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayStageList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayStage.
func (c *FakeApiGatewayStages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayStage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewaystagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayStage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayStage), err
}
