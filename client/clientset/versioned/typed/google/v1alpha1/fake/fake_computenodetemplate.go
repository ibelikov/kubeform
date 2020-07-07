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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeComputeNodeTemplates implements ComputeNodeTemplateInterface
type FakeComputeNodeTemplates struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computenodetemplatesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computenodetemplates"}

var computenodetemplatesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeNodeTemplate"}

// Get takes name of the computeNodeTemplate, and returns the corresponding computeNodeTemplate object, and an error if there is any.
func (c *FakeComputeNodeTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeNodeTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computenodetemplatesResource, c.ns, name), &v1alpha1.ComputeNodeTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNodeTemplate), err
}

// List takes label and field selectors, and returns the list of ComputeNodeTemplates that match those selectors.
func (c *FakeComputeNodeTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeNodeTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computenodetemplatesResource, computenodetemplatesKind, c.ns, opts), &v1alpha1.ComputeNodeTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeNodeTemplateList{ListMeta: obj.(*v1alpha1.ComputeNodeTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeNodeTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeNodeTemplates.
func (c *FakeComputeNodeTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computenodetemplatesResource, c.ns, opts))

}

// Create takes the representation of a computeNodeTemplate and creates it.  Returns the server's representation of the computeNodeTemplate, and an error, if there is any.
func (c *FakeComputeNodeTemplates) Create(ctx context.Context, computeNodeTemplate *v1alpha1.ComputeNodeTemplate, opts v1.CreateOptions) (result *v1alpha1.ComputeNodeTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computenodetemplatesResource, c.ns, computeNodeTemplate), &v1alpha1.ComputeNodeTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNodeTemplate), err
}

// Update takes the representation of a computeNodeTemplate and updates it. Returns the server's representation of the computeNodeTemplate, and an error, if there is any.
func (c *FakeComputeNodeTemplates) Update(ctx context.Context, computeNodeTemplate *v1alpha1.ComputeNodeTemplate, opts v1.UpdateOptions) (result *v1alpha1.ComputeNodeTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computenodetemplatesResource, c.ns, computeNodeTemplate), &v1alpha1.ComputeNodeTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNodeTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeNodeTemplates) UpdateStatus(ctx context.Context, computeNodeTemplate *v1alpha1.ComputeNodeTemplate, opts v1.UpdateOptions) (*v1alpha1.ComputeNodeTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computenodetemplatesResource, "status", c.ns, computeNodeTemplate), &v1alpha1.ComputeNodeTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNodeTemplate), err
}

// Delete takes name of the computeNodeTemplate and deletes it. Returns an error if one occurs.
func (c *FakeComputeNodeTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computenodetemplatesResource, c.ns, name), &v1alpha1.ComputeNodeTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeNodeTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computenodetemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeNodeTemplateList{})
	return err
}

// Patch applies the patch and returns the patched computeNodeTemplate.
func (c *FakeComputeNodeTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeNodeTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computenodetemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeNodeTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNodeTemplate), err
}
