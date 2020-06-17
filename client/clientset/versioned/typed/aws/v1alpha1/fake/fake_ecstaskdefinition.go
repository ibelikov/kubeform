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

// FakeEcsTaskDefinitions implements EcsTaskDefinitionInterface
type FakeEcsTaskDefinitions struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ecstaskdefinitionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ecstaskdefinitions"}

var ecstaskdefinitionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EcsTaskDefinition"}

// Get takes name of the ecsTaskDefinition, and returns the corresponding ecsTaskDefinition object, and an error if there is any.
func (c *FakeEcsTaskDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.EcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ecstaskdefinitionsResource, c.ns, name), &v1alpha1.EcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsTaskDefinition), err
}

// List takes label and field selectors, and returns the list of EcsTaskDefinitions that match those selectors.
func (c *FakeEcsTaskDefinitions) List(opts v1.ListOptions) (result *v1alpha1.EcsTaskDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ecstaskdefinitionsResource, ecstaskdefinitionsKind, c.ns, opts), &v1alpha1.EcsTaskDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EcsTaskDefinitionList{ListMeta: obj.(*v1alpha1.EcsTaskDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.EcsTaskDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ecsTaskDefinitions.
func (c *FakeEcsTaskDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ecstaskdefinitionsResource, c.ns, opts))

}

// Create takes the representation of a ecsTaskDefinition and creates it.  Returns the server's representation of the ecsTaskDefinition, and an error, if there is any.
func (c *FakeEcsTaskDefinitions) Create(ecsTaskDefinition *v1alpha1.EcsTaskDefinition) (result *v1alpha1.EcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ecstaskdefinitionsResource, c.ns, ecsTaskDefinition), &v1alpha1.EcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsTaskDefinition), err
}

// Update takes the representation of a ecsTaskDefinition and updates it. Returns the server's representation of the ecsTaskDefinition, and an error, if there is any.
func (c *FakeEcsTaskDefinitions) Update(ecsTaskDefinition *v1alpha1.EcsTaskDefinition) (result *v1alpha1.EcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ecstaskdefinitionsResource, c.ns, ecsTaskDefinition), &v1alpha1.EcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsTaskDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEcsTaskDefinitions) UpdateStatus(ecsTaskDefinition *v1alpha1.EcsTaskDefinition) (*v1alpha1.EcsTaskDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ecstaskdefinitionsResource, "status", c.ns, ecsTaskDefinition), &v1alpha1.EcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsTaskDefinition), err
}

// Delete takes name of the ecsTaskDefinition and deletes it. Returns an error if one occurs.
func (c *FakeEcsTaskDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ecstaskdefinitionsResource, c.ns, name), &v1alpha1.EcsTaskDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEcsTaskDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ecstaskdefinitionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EcsTaskDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched ecsTaskDefinition.
func (c *FakeEcsTaskDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcsTaskDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ecstaskdefinitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EcsTaskDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsTaskDefinition), err
}
