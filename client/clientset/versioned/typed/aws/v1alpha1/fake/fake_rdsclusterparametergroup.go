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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeRdsClusterParameterGroups implements RdsClusterParameterGroupInterface
type FakeRdsClusterParameterGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var rdsclusterparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "rdsclusterparametergroups"}

var rdsclusterparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "RdsClusterParameterGroup"}

// Get takes name of the rdsClusterParameterGroup, and returns the corresponding rdsClusterParameterGroup object, and an error if there is any.
func (c *FakeRdsClusterParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rdsclusterparametergroupsResource, c.ns, name), &v1alpha1.RdsClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsClusterParameterGroup), err
}

// List takes label and field selectors, and returns the list of RdsClusterParameterGroups that match those selectors.
func (c *FakeRdsClusterParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.RdsClusterParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rdsclusterparametergroupsResource, rdsclusterparametergroupsKind, c.ns, opts), &v1alpha1.RdsClusterParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RdsClusterParameterGroupList{ListMeta: obj.(*v1alpha1.RdsClusterParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.RdsClusterParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rdsClusterParameterGroups.
func (c *FakeRdsClusterParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rdsclusterparametergroupsResource, c.ns, opts))

}

// Create takes the representation of a rdsClusterParameterGroup and creates it.  Returns the server's representation of the rdsClusterParameterGroup, and an error, if there is any.
func (c *FakeRdsClusterParameterGroups) Create(rdsClusterParameterGroup *v1alpha1.RdsClusterParameterGroup) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rdsclusterparametergroupsResource, c.ns, rdsClusterParameterGroup), &v1alpha1.RdsClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsClusterParameterGroup), err
}

// Update takes the representation of a rdsClusterParameterGroup and updates it. Returns the server's representation of the rdsClusterParameterGroup, and an error, if there is any.
func (c *FakeRdsClusterParameterGroups) Update(rdsClusterParameterGroup *v1alpha1.RdsClusterParameterGroup) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rdsclusterparametergroupsResource, c.ns, rdsClusterParameterGroup), &v1alpha1.RdsClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsClusterParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRdsClusterParameterGroups) UpdateStatus(rdsClusterParameterGroup *v1alpha1.RdsClusterParameterGroup) (*v1alpha1.RdsClusterParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rdsclusterparametergroupsResource, "status", c.ns, rdsClusterParameterGroup), &v1alpha1.RdsClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsClusterParameterGroup), err
}

// Delete takes name of the rdsClusterParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeRdsClusterParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rdsclusterparametergroupsResource, c.ns, name), &v1alpha1.RdsClusterParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRdsClusterParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rdsclusterparametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RdsClusterParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched rdsClusterParameterGroup.
func (c *FakeRdsClusterParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rdsclusterparametergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RdsClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsClusterParameterGroup), err
}
