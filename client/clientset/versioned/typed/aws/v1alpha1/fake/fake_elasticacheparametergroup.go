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

// FakeElasticacheParameterGroups implements ElasticacheParameterGroupInterface
type FakeElasticacheParameterGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var elasticacheparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "elasticacheparametergroups"}

var elasticacheparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ElasticacheParameterGroup"}

// Get takes name of the elasticacheParameterGroup, and returns the corresponding elasticacheParameterGroup object, and an error if there is any.
func (c *FakeElasticacheParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elasticacheparametergroupsResource, c.ns, name), &v1alpha1.ElasticacheParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticacheParameterGroup), err
}

// List takes label and field selectors, and returns the list of ElasticacheParameterGroups that match those selectors.
func (c *FakeElasticacheParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.ElasticacheParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elasticacheparametergroupsResource, elasticacheparametergroupsKind, c.ns, opts), &v1alpha1.ElasticacheParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElasticacheParameterGroupList{ListMeta: obj.(*v1alpha1.ElasticacheParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElasticacheParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticacheParameterGroups.
func (c *FakeElasticacheParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elasticacheparametergroupsResource, c.ns, opts))

}

// Create takes the representation of a elasticacheParameterGroup and creates it.  Returns the server's representation of the elasticacheParameterGroup, and an error, if there is any.
func (c *FakeElasticacheParameterGroups) Create(elasticacheParameterGroup *v1alpha1.ElasticacheParameterGroup) (result *v1alpha1.ElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elasticacheparametergroupsResource, c.ns, elasticacheParameterGroup), &v1alpha1.ElasticacheParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticacheParameterGroup), err
}

// Update takes the representation of a elasticacheParameterGroup and updates it. Returns the server's representation of the elasticacheParameterGroup, and an error, if there is any.
func (c *FakeElasticacheParameterGroups) Update(elasticacheParameterGroup *v1alpha1.ElasticacheParameterGroup) (result *v1alpha1.ElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elasticacheparametergroupsResource, c.ns, elasticacheParameterGroup), &v1alpha1.ElasticacheParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticacheParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticacheParameterGroups) UpdateStatus(elasticacheParameterGroup *v1alpha1.ElasticacheParameterGroup) (*v1alpha1.ElasticacheParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elasticacheparametergroupsResource, "status", c.ns, elasticacheParameterGroup), &v1alpha1.ElasticacheParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticacheParameterGroup), err
}

// Delete takes name of the elasticacheParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeElasticacheParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(elasticacheparametergroupsResource, c.ns, name), &v1alpha1.ElasticacheParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticacheParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elasticacheparametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElasticacheParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched elasticacheParameterGroup.
func (c *FakeElasticacheParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElasticacheParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elasticacheparametergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ElasticacheParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticacheParameterGroup), err
}
