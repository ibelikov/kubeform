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

// FakeAppmeshMeshes implements AppmeshMeshInterface
type FakeAppmeshMeshes struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var appmeshmeshesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "appmeshmeshes"}

var appmeshmeshesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AppmeshMesh"}

// Get takes name of the appmeshMesh, and returns the corresponding appmeshMesh object, and an error if there is any.
func (c *FakeAppmeshMeshes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appmeshmeshesResource, c.ns, name), &v1alpha1.AppmeshMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// List takes label and field selectors, and returns the list of AppmeshMeshes that match those selectors.
func (c *FakeAppmeshMeshes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppmeshMeshList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appmeshmeshesResource, appmeshmeshesKind, c.ns, opts), &v1alpha1.AppmeshMeshList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppmeshMeshList{ListMeta: obj.(*v1alpha1.AppmeshMeshList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppmeshMeshList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appmeshMeshes.
func (c *FakeAppmeshMeshes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appmeshmeshesResource, c.ns, opts))

}

// Create takes the representation of a appmeshMesh and creates it.  Returns the server's representation of the appmeshMesh, and an error, if there is any.
func (c *FakeAppmeshMeshes) Create(ctx context.Context, appmeshMesh *v1alpha1.AppmeshMesh, opts v1.CreateOptions) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appmeshmeshesResource, c.ns, appmeshMesh), &v1alpha1.AppmeshMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// Update takes the representation of a appmeshMesh and updates it. Returns the server's representation of the appmeshMesh, and an error, if there is any.
func (c *FakeAppmeshMeshes) Update(ctx context.Context, appmeshMesh *v1alpha1.AppmeshMesh, opts v1.UpdateOptions) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appmeshmeshesResource, c.ns, appmeshMesh), &v1alpha1.AppmeshMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppmeshMeshes) UpdateStatus(ctx context.Context, appmeshMesh *v1alpha1.AppmeshMesh, opts v1.UpdateOptions) (*v1alpha1.AppmeshMesh, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appmeshmeshesResource, "status", c.ns, appmeshMesh), &v1alpha1.AppmeshMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// Delete takes name of the appmeshMesh and deletes it. Returns an error if one occurs.
func (c *FakeAppmeshMeshes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appmeshmeshesResource, c.ns, name), &v1alpha1.AppmeshMesh{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppmeshMeshes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appmeshmeshesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppmeshMeshList{})
	return err
}

// Patch applies the patch and returns the patched appmeshMesh.
func (c *FakeAppmeshMeshes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appmeshmeshesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppmeshMesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}
