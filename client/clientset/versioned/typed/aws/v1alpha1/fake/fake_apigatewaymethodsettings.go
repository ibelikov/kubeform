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

// FakeApiGatewayMethodSettingses implements ApiGatewayMethodSettingsInterface
type FakeApiGatewayMethodSettingses struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewaymethodsettingsesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewaymethodsettingses"}

var apigatewaymethodsettingsesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayMethodSettings"}

// Get takes name of the apiGatewayMethodSettings, and returns the corresponding apiGatewayMethodSettings object, and an error if there is any.
func (c *FakeApiGatewayMethodSettingses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayMethodSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewaymethodsettingsesResource, c.ns, name), &v1alpha1.ApiGatewayMethodSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethodSettings), err
}

// List takes label and field selectors, and returns the list of ApiGatewayMethodSettingses that match those selectors.
func (c *FakeApiGatewayMethodSettingses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiGatewayMethodSettingsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewaymethodsettingsesResource, apigatewaymethodsettingsesKind, c.ns, opts), &v1alpha1.ApiGatewayMethodSettingsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayMethodSettingsList{ListMeta: obj.(*v1alpha1.ApiGatewayMethodSettingsList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayMethodSettingsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayMethodSettingses.
func (c *FakeApiGatewayMethodSettingses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewaymethodsettingsesResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayMethodSettings and creates it.  Returns the server's representation of the apiGatewayMethodSettings, and an error, if there is any.
func (c *FakeApiGatewayMethodSettingses) Create(ctx context.Context, apiGatewayMethodSettings *v1alpha1.ApiGatewayMethodSettings, opts v1.CreateOptions) (result *v1alpha1.ApiGatewayMethodSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewaymethodsettingsesResource, c.ns, apiGatewayMethodSettings), &v1alpha1.ApiGatewayMethodSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethodSettings), err
}

// Update takes the representation of a apiGatewayMethodSettings and updates it. Returns the server's representation of the apiGatewayMethodSettings, and an error, if there is any.
func (c *FakeApiGatewayMethodSettingses) Update(ctx context.Context, apiGatewayMethodSettings *v1alpha1.ApiGatewayMethodSettings, opts v1.UpdateOptions) (result *v1alpha1.ApiGatewayMethodSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewaymethodsettingsesResource, c.ns, apiGatewayMethodSettings), &v1alpha1.ApiGatewayMethodSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethodSettings), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayMethodSettingses) UpdateStatus(ctx context.Context, apiGatewayMethodSettings *v1alpha1.ApiGatewayMethodSettings, opts v1.UpdateOptions) (*v1alpha1.ApiGatewayMethodSettings, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewaymethodsettingsesResource, "status", c.ns, apiGatewayMethodSettings), &v1alpha1.ApiGatewayMethodSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethodSettings), err
}

// Delete takes name of the apiGatewayMethodSettings and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayMethodSettingses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewaymethodsettingsesResource, c.ns, name), &v1alpha1.ApiGatewayMethodSettings{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayMethodSettingses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewaymethodsettingsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayMethodSettingsList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayMethodSettings.
func (c *FakeApiGatewayMethodSettingses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiGatewayMethodSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewaymethodsettingsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayMethodSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayMethodSettings), err
}
