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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeLambdaAliases implements LambdaAliasInterface
type FakeLambdaAliases struct {
	Fake *FakeAwsV1alpha1
}

var lambdaaliasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lambdaaliases"}

var lambdaaliasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LambdaAlias"}

// Get takes name of the lambdaAlias, and returns the corresponding lambdaAlias object, and an error if there is any.
func (c *FakeLambdaAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.LambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(lambdaaliasesResource, name), &v1alpha1.LambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LambdaAlias), err
}

// List takes label and field selectors, and returns the list of LambdaAliases that match those selectors.
func (c *FakeLambdaAliases) List(opts v1.ListOptions) (result *v1alpha1.LambdaAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(lambdaaliasesResource, lambdaaliasesKind, opts), &v1alpha1.LambdaAliasList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LambdaAliasList{ListMeta: obj.(*v1alpha1.LambdaAliasList).ListMeta}
	for _, item := range obj.(*v1alpha1.LambdaAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lambdaAliases.
func (c *FakeLambdaAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(lambdaaliasesResource, opts))
}

// Create takes the representation of a lambdaAlias and creates it.  Returns the server's representation of the lambdaAlias, and an error, if there is any.
func (c *FakeLambdaAliases) Create(lambdaAlias *v1alpha1.LambdaAlias) (result *v1alpha1.LambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(lambdaaliasesResource, lambdaAlias), &v1alpha1.LambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LambdaAlias), err
}

// Update takes the representation of a lambdaAlias and updates it. Returns the server's representation of the lambdaAlias, and an error, if there is any.
func (c *FakeLambdaAliases) Update(lambdaAlias *v1alpha1.LambdaAlias) (result *v1alpha1.LambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(lambdaaliasesResource, lambdaAlias), &v1alpha1.LambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LambdaAlias), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLambdaAliases) UpdateStatus(lambdaAlias *v1alpha1.LambdaAlias) (*v1alpha1.LambdaAlias, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(lambdaaliasesResource, "status", lambdaAlias), &v1alpha1.LambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LambdaAlias), err
}

// Delete takes name of the lambdaAlias and deletes it. Returns an error if one occurs.
func (c *FakeLambdaAliases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(lambdaaliasesResource, name), &v1alpha1.LambdaAlias{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLambdaAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(lambdaaliasesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LambdaAliasList{})
	return err
}

// Patch applies the patch and returns the patched lambdaAlias.
func (c *FakeLambdaAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LambdaAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(lambdaaliasesResource, name, pt, data, subresources...), &v1alpha1.LambdaAlias{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LambdaAlias), err
}
