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

// FakeEfsMountTargets implements EfsMountTargetInterface
type FakeEfsMountTargets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var efsmounttargetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "efsmounttargets"}

var efsmounttargetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EfsMountTarget"}

// Get takes name of the efsMountTarget, and returns the corresponding efsMountTarget object, and an error if there is any.
func (c *FakeEfsMountTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.EfsMountTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(efsmounttargetsResource, c.ns, name), &v1alpha1.EfsMountTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsMountTarget), err
}

// List takes label and field selectors, and returns the list of EfsMountTargets that match those selectors.
func (c *FakeEfsMountTargets) List(opts v1.ListOptions) (result *v1alpha1.EfsMountTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(efsmounttargetsResource, efsmounttargetsKind, c.ns, opts), &v1alpha1.EfsMountTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EfsMountTargetList{ListMeta: obj.(*v1alpha1.EfsMountTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.EfsMountTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested efsMountTargets.
func (c *FakeEfsMountTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(efsmounttargetsResource, c.ns, opts))

}

// Create takes the representation of a efsMountTarget and creates it.  Returns the server's representation of the efsMountTarget, and an error, if there is any.
func (c *FakeEfsMountTargets) Create(efsMountTarget *v1alpha1.EfsMountTarget) (result *v1alpha1.EfsMountTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(efsmounttargetsResource, c.ns, efsMountTarget), &v1alpha1.EfsMountTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsMountTarget), err
}

// Update takes the representation of a efsMountTarget and updates it. Returns the server's representation of the efsMountTarget, and an error, if there is any.
func (c *FakeEfsMountTargets) Update(efsMountTarget *v1alpha1.EfsMountTarget) (result *v1alpha1.EfsMountTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(efsmounttargetsResource, c.ns, efsMountTarget), &v1alpha1.EfsMountTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsMountTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEfsMountTargets) UpdateStatus(efsMountTarget *v1alpha1.EfsMountTarget) (*v1alpha1.EfsMountTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(efsmounttargetsResource, "status", c.ns, efsMountTarget), &v1alpha1.EfsMountTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsMountTarget), err
}

// Delete takes name of the efsMountTarget and deletes it. Returns an error if one occurs.
func (c *FakeEfsMountTargets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(efsmounttargetsResource, c.ns, name), &v1alpha1.EfsMountTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEfsMountTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(efsmounttargetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EfsMountTargetList{})
	return err
}

// Patch applies the patch and returns the patched efsMountTarget.
func (c *FakeEfsMountTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EfsMountTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(efsmounttargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EfsMountTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EfsMountTarget), err
}
