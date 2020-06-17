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

// FakeDocdbClusterParameterGroups implements DocdbClusterParameterGroupInterface
type FakeDocdbClusterParameterGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var docdbclusterparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "docdbclusterparametergroups"}

var docdbclusterparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DocdbClusterParameterGroup"}

// Get takes name of the docdbClusterParameterGroup, and returns the corresponding docdbClusterParameterGroup object, and an error if there is any.
func (c *FakeDocdbClusterParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(docdbclusterparametergroupsResource, c.ns, name), &v1alpha1.DocdbClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterParameterGroup), err
}

// List takes label and field selectors, and returns the list of DocdbClusterParameterGroups that match those selectors.
func (c *FakeDocdbClusterParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.DocdbClusterParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(docdbclusterparametergroupsResource, docdbclusterparametergroupsKind, c.ns, opts), &v1alpha1.DocdbClusterParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DocdbClusterParameterGroupList{ListMeta: obj.(*v1alpha1.DocdbClusterParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DocdbClusterParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested docdbClusterParameterGroups.
func (c *FakeDocdbClusterParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(docdbclusterparametergroupsResource, c.ns, opts))

}

// Create takes the representation of a docdbClusterParameterGroup and creates it.  Returns the server's representation of the docdbClusterParameterGroup, and an error, if there is any.
func (c *FakeDocdbClusterParameterGroups) Create(docdbClusterParameterGroup *v1alpha1.DocdbClusterParameterGroup) (result *v1alpha1.DocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(docdbclusterparametergroupsResource, c.ns, docdbClusterParameterGroup), &v1alpha1.DocdbClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterParameterGroup), err
}

// Update takes the representation of a docdbClusterParameterGroup and updates it. Returns the server's representation of the docdbClusterParameterGroup, and an error, if there is any.
func (c *FakeDocdbClusterParameterGroups) Update(docdbClusterParameterGroup *v1alpha1.DocdbClusterParameterGroup) (result *v1alpha1.DocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(docdbclusterparametergroupsResource, c.ns, docdbClusterParameterGroup), &v1alpha1.DocdbClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDocdbClusterParameterGroups) UpdateStatus(docdbClusterParameterGroup *v1alpha1.DocdbClusterParameterGroup) (*v1alpha1.DocdbClusterParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(docdbclusterparametergroupsResource, "status", c.ns, docdbClusterParameterGroup), &v1alpha1.DocdbClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterParameterGroup), err
}

// Delete takes name of the docdbClusterParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeDocdbClusterParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(docdbclusterparametergroupsResource, c.ns, name), &v1alpha1.DocdbClusterParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDocdbClusterParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(docdbclusterparametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DocdbClusterParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched docdbClusterParameterGroup.
func (c *FakeDocdbClusterParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DocdbClusterParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(docdbclusterparametergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DocdbClusterParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbClusterParameterGroup), err
}
