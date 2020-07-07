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

// FakeCloudwatchLogGroups implements CloudwatchLogGroupInterface
type FakeCloudwatchLogGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudwatchloggroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudwatchloggroups"}

var cloudwatchloggroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudwatchLogGroup"}

// Get takes name of the cloudwatchLogGroup, and returns the corresponding cloudwatchLogGroup object, and an error if there is any.
func (c *FakeCloudwatchLogGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudwatchloggroupsResource, c.ns, name), &v1alpha1.CloudwatchLogGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogGroup), err
}

// List takes label and field selectors, and returns the list of CloudwatchLogGroups that match those selectors.
func (c *FakeCloudwatchLogGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudwatchLogGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudwatchloggroupsResource, cloudwatchloggroupsKind, c.ns, opts), &v1alpha1.CloudwatchLogGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudwatchLogGroupList{ListMeta: obj.(*v1alpha1.CloudwatchLogGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudwatchLogGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogGroups.
func (c *FakeCloudwatchLogGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudwatchloggroupsResource, c.ns, opts))

}

// Create takes the representation of a cloudwatchLogGroup and creates it.  Returns the server's representation of the cloudwatchLogGroup, and an error, if there is any.
func (c *FakeCloudwatchLogGroups) Create(ctx context.Context, cloudwatchLogGroup *v1alpha1.CloudwatchLogGroup, opts v1.CreateOptions) (result *v1alpha1.CloudwatchLogGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudwatchloggroupsResource, c.ns, cloudwatchLogGroup), &v1alpha1.CloudwatchLogGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogGroup), err
}

// Update takes the representation of a cloudwatchLogGroup and updates it. Returns the server's representation of the cloudwatchLogGroup, and an error, if there is any.
func (c *FakeCloudwatchLogGroups) Update(ctx context.Context, cloudwatchLogGroup *v1alpha1.CloudwatchLogGroup, opts v1.UpdateOptions) (result *v1alpha1.CloudwatchLogGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudwatchloggroupsResource, c.ns, cloudwatchLogGroup), &v1alpha1.CloudwatchLogGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudwatchLogGroups) UpdateStatus(ctx context.Context, cloudwatchLogGroup *v1alpha1.CloudwatchLogGroup, opts v1.UpdateOptions) (*v1alpha1.CloudwatchLogGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudwatchloggroupsResource, "status", c.ns, cloudwatchLogGroup), &v1alpha1.CloudwatchLogGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogGroup), err
}

// Delete takes name of the cloudwatchLogGroup and deletes it. Returns an error if one occurs.
func (c *FakeCloudwatchLogGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudwatchloggroupsResource, c.ns, name), &v1alpha1.CloudwatchLogGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudwatchLogGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudwatchloggroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudwatchLogGroupList{})
	return err
}

// Patch applies the patch and returns the patched cloudwatchLogGroup.
func (c *FakeCloudwatchLogGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudwatchLogGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudwatchloggroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudwatchLogGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogGroup), err
}
