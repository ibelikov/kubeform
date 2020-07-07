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

// FakeCloudformationStackSetInstances implements CloudformationStackSetInstanceInterface
type FakeCloudformationStackSetInstances struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudformationstacksetinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudformationstacksetinstances"}

var cloudformationstacksetinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudformationStackSetInstance"}

// Get takes name of the cloudformationStackSetInstance, and returns the corresponding cloudformationStackSetInstance object, and an error if there is any.
func (c *FakeCloudformationStackSetInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudformationstacksetinstancesResource, c.ns, name), &v1alpha1.CloudformationStackSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSetInstance), err
}

// List takes label and field selectors, and returns the list of CloudformationStackSetInstances that match those selectors.
func (c *FakeCloudformationStackSetInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudformationStackSetInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudformationstacksetinstancesResource, cloudformationstacksetinstancesKind, c.ns, opts), &v1alpha1.CloudformationStackSetInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudformationStackSetInstanceList{ListMeta: obj.(*v1alpha1.CloudformationStackSetInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudformationStackSetInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudformationStackSetInstances.
func (c *FakeCloudformationStackSetInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudformationstacksetinstancesResource, c.ns, opts))

}

// Create takes the representation of a cloudformationStackSetInstance and creates it.  Returns the server's representation of the cloudformationStackSetInstance, and an error, if there is any.
func (c *FakeCloudformationStackSetInstances) Create(ctx context.Context, cloudformationStackSetInstance *v1alpha1.CloudformationStackSetInstance, opts v1.CreateOptions) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudformationstacksetinstancesResource, c.ns, cloudformationStackSetInstance), &v1alpha1.CloudformationStackSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSetInstance), err
}

// Update takes the representation of a cloudformationStackSetInstance and updates it. Returns the server's representation of the cloudformationStackSetInstance, and an error, if there is any.
func (c *FakeCloudformationStackSetInstances) Update(ctx context.Context, cloudformationStackSetInstance *v1alpha1.CloudformationStackSetInstance, opts v1.UpdateOptions) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudformationstacksetinstancesResource, c.ns, cloudformationStackSetInstance), &v1alpha1.CloudformationStackSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSetInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudformationStackSetInstances) UpdateStatus(ctx context.Context, cloudformationStackSetInstance *v1alpha1.CloudformationStackSetInstance, opts v1.UpdateOptions) (*v1alpha1.CloudformationStackSetInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudformationstacksetinstancesResource, "status", c.ns, cloudformationStackSetInstance), &v1alpha1.CloudformationStackSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSetInstance), err
}

// Delete takes name of the cloudformationStackSetInstance and deletes it. Returns an error if one occurs.
func (c *FakeCloudformationStackSetInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudformationstacksetinstancesResource, c.ns, name), &v1alpha1.CloudformationStackSetInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudformationStackSetInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudformationstacksetinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudformationStackSetInstanceList{})
	return err
}

// Patch applies the patch and returns the patched cloudformationStackSetInstance.
func (c *FakeCloudformationStackSetInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudformationstacksetinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudformationStackSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSetInstance), err
}
