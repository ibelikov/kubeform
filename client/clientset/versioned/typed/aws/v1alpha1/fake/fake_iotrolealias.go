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

// FakeIotRoleAliases implements IotRoleAliasInterface
type FakeIotRoleAliases struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iotrolealiasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iotrolealiases"}

var iotrolealiasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IotRoleAlias"}

// Get takes name of the iotRoleAlias, and returns the corresponding iotRoleAlias object, and an error if there is any.
func (c *FakeIotRoleAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.IotRoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iotrolealiasesResource, c.ns, name), &v1alpha1.IotRoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotRoleAlias), err
}

// List takes label and field selectors, and returns the list of IotRoleAliases that match those selectors.
func (c *FakeIotRoleAliases) List(opts v1.ListOptions) (result *v1alpha1.IotRoleAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iotrolealiasesResource, iotrolealiasesKind, c.ns, opts), &v1alpha1.IotRoleAliasList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IotRoleAliasList{ListMeta: obj.(*v1alpha1.IotRoleAliasList).ListMeta}
	for _, item := range obj.(*v1alpha1.IotRoleAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iotRoleAliases.
func (c *FakeIotRoleAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iotrolealiasesResource, c.ns, opts))

}

// Create takes the representation of a iotRoleAlias and creates it.  Returns the server's representation of the iotRoleAlias, and an error, if there is any.
func (c *FakeIotRoleAliases) Create(iotRoleAlias *v1alpha1.IotRoleAlias) (result *v1alpha1.IotRoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iotrolealiasesResource, c.ns, iotRoleAlias), &v1alpha1.IotRoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotRoleAlias), err
}

// Update takes the representation of a iotRoleAlias and updates it. Returns the server's representation of the iotRoleAlias, and an error, if there is any.
func (c *FakeIotRoleAliases) Update(iotRoleAlias *v1alpha1.IotRoleAlias) (result *v1alpha1.IotRoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iotrolealiasesResource, c.ns, iotRoleAlias), &v1alpha1.IotRoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotRoleAlias), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIotRoleAliases) UpdateStatus(iotRoleAlias *v1alpha1.IotRoleAlias) (*v1alpha1.IotRoleAlias, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iotrolealiasesResource, "status", c.ns, iotRoleAlias), &v1alpha1.IotRoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotRoleAlias), err
}

// Delete takes name of the iotRoleAlias and deletes it. Returns an error if one occurs.
func (c *FakeIotRoleAliases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iotrolealiasesResource, c.ns, name), &v1alpha1.IotRoleAlias{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIotRoleAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iotrolealiasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IotRoleAliasList{})
	return err
}

// Patch applies the patch and returns the patched iotRoleAlias.
func (c *FakeIotRoleAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotRoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iotrolealiasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IotRoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IotRoleAlias), err
}
