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

// FakeOpsworksRdsDbInstances implements OpsworksRdsDbInstanceInterface
type FakeOpsworksRdsDbInstances struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworksrdsdbinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksrdsdbinstances"}

var opsworksrdsdbinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksRdsDbInstance"}

// Get takes name of the opsworksRdsDbInstance, and returns the corresponding opsworksRdsDbInstance object, and an error if there is any.
func (c *FakeOpsworksRdsDbInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworksrdsdbinstancesResource, c.ns, name), &v1alpha1.OpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksRdsDbInstance), err
}

// List takes label and field selectors, and returns the list of OpsworksRdsDbInstances that match those selectors.
func (c *FakeOpsworksRdsDbInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpsworksRdsDbInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworksrdsdbinstancesResource, opsworksrdsdbinstancesKind, c.ns, opts), &v1alpha1.OpsworksRdsDbInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksRdsDbInstanceList{ListMeta: obj.(*v1alpha1.OpsworksRdsDbInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksRdsDbInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksRdsDbInstances.
func (c *FakeOpsworksRdsDbInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworksrdsdbinstancesResource, c.ns, opts))

}

// Create takes the representation of a opsworksRdsDbInstance and creates it.  Returns the server's representation of the opsworksRdsDbInstance, and an error, if there is any.
func (c *FakeOpsworksRdsDbInstances) Create(ctx context.Context, opsworksRdsDbInstance *v1alpha1.OpsworksRdsDbInstance, opts v1.CreateOptions) (result *v1alpha1.OpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworksrdsdbinstancesResource, c.ns, opsworksRdsDbInstance), &v1alpha1.OpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksRdsDbInstance), err
}

// Update takes the representation of a opsworksRdsDbInstance and updates it. Returns the server's representation of the opsworksRdsDbInstance, and an error, if there is any.
func (c *FakeOpsworksRdsDbInstances) Update(ctx context.Context, opsworksRdsDbInstance *v1alpha1.OpsworksRdsDbInstance, opts v1.UpdateOptions) (result *v1alpha1.OpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworksrdsdbinstancesResource, c.ns, opsworksRdsDbInstance), &v1alpha1.OpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksRdsDbInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksRdsDbInstances) UpdateStatus(ctx context.Context, opsworksRdsDbInstance *v1alpha1.OpsworksRdsDbInstance, opts v1.UpdateOptions) (*v1alpha1.OpsworksRdsDbInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworksrdsdbinstancesResource, "status", c.ns, opsworksRdsDbInstance), &v1alpha1.OpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksRdsDbInstance), err
}

// Delete takes name of the opsworksRdsDbInstance and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksRdsDbInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworksrdsdbinstancesResource, c.ns, name), &v1alpha1.OpsworksRdsDbInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksRdsDbInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworksrdsdbinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksRdsDbInstanceList{})
	return err
}

// Patch applies the patch and returns the patched opsworksRdsDbInstance.
func (c *FakeOpsworksRdsDbInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpsworksRdsDbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworksrdsdbinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksRdsDbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksRdsDbInstance), err
}
