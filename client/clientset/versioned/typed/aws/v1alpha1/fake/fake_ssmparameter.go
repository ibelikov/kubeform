/*
Copyright 2019 The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSsmParameters implements SsmParameterInterface
type FakeSsmParameters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ssmparametersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ssmparameters"}

var ssmparametersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SsmParameter"}

// Get takes name of the ssmParameter, and returns the corresponding ssmParameter object, and an error if there is any.
func (c *FakeSsmParameters) Get(name string, options v1.GetOptions) (result *v1alpha1.SsmParameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ssmparametersResource, c.ns, name), &v1alpha1.SsmParameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmParameter), err
}

// List takes label and field selectors, and returns the list of SsmParameters that match those selectors.
func (c *FakeSsmParameters) List(opts v1.ListOptions) (result *v1alpha1.SsmParameterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ssmparametersResource, ssmparametersKind, c.ns, opts), &v1alpha1.SsmParameterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SsmParameterList{ListMeta: obj.(*v1alpha1.SsmParameterList).ListMeta}
	for _, item := range obj.(*v1alpha1.SsmParameterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ssmParameters.
func (c *FakeSsmParameters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ssmparametersResource, c.ns, opts))

}

// Create takes the representation of a ssmParameter and creates it.  Returns the server's representation of the ssmParameter, and an error, if there is any.
func (c *FakeSsmParameters) Create(ssmParameter *v1alpha1.SsmParameter) (result *v1alpha1.SsmParameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ssmparametersResource, c.ns, ssmParameter), &v1alpha1.SsmParameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmParameter), err
}

// Update takes the representation of a ssmParameter and updates it. Returns the server's representation of the ssmParameter, and an error, if there is any.
func (c *FakeSsmParameters) Update(ssmParameter *v1alpha1.SsmParameter) (result *v1alpha1.SsmParameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ssmparametersResource, c.ns, ssmParameter), &v1alpha1.SsmParameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmParameter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsmParameters) UpdateStatus(ssmParameter *v1alpha1.SsmParameter) (*v1alpha1.SsmParameter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ssmparametersResource, "status", c.ns, ssmParameter), &v1alpha1.SsmParameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmParameter), err
}

// Delete takes name of the ssmParameter and deletes it. Returns an error if one occurs.
func (c *FakeSsmParameters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ssmparametersResource, c.ns, name), &v1alpha1.SsmParameter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsmParameters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ssmparametersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsmParameterList{})
	return err
}

// Patch applies the patch and returns the patched ssmParameter.
func (c *FakeSsmParameters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SsmParameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ssmparametersResource, c.ns, name, pt, data, subresources...), &v1alpha1.SsmParameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmParameter), err
}
