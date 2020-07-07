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

// FakeAppmeshVirtualServices implements AppmeshVirtualServiceInterface
type FakeAppmeshVirtualServices struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var appmeshvirtualservicesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "appmeshvirtualservices"}

var appmeshvirtualservicesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AppmeshVirtualService"}

// Get takes name of the appmeshVirtualService, and returns the corresponding appmeshVirtualService object, and an error if there is any.
func (c *FakeAppmeshVirtualServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppmeshVirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appmeshvirtualservicesResource, c.ns, name), &v1alpha1.AppmeshVirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshVirtualService), err
}

// List takes label and field selectors, and returns the list of AppmeshVirtualServices that match those selectors.
func (c *FakeAppmeshVirtualServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppmeshVirtualServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appmeshvirtualservicesResource, appmeshvirtualservicesKind, c.ns, opts), &v1alpha1.AppmeshVirtualServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppmeshVirtualServiceList{ListMeta: obj.(*v1alpha1.AppmeshVirtualServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppmeshVirtualServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appmeshVirtualServices.
func (c *FakeAppmeshVirtualServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appmeshvirtualservicesResource, c.ns, opts))

}

// Create takes the representation of a appmeshVirtualService and creates it.  Returns the server's representation of the appmeshVirtualService, and an error, if there is any.
func (c *FakeAppmeshVirtualServices) Create(ctx context.Context, appmeshVirtualService *v1alpha1.AppmeshVirtualService, opts v1.CreateOptions) (result *v1alpha1.AppmeshVirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appmeshvirtualservicesResource, c.ns, appmeshVirtualService), &v1alpha1.AppmeshVirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshVirtualService), err
}

// Update takes the representation of a appmeshVirtualService and updates it. Returns the server's representation of the appmeshVirtualService, and an error, if there is any.
func (c *FakeAppmeshVirtualServices) Update(ctx context.Context, appmeshVirtualService *v1alpha1.AppmeshVirtualService, opts v1.UpdateOptions) (result *v1alpha1.AppmeshVirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appmeshvirtualservicesResource, c.ns, appmeshVirtualService), &v1alpha1.AppmeshVirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshVirtualService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppmeshVirtualServices) UpdateStatus(ctx context.Context, appmeshVirtualService *v1alpha1.AppmeshVirtualService, opts v1.UpdateOptions) (*v1alpha1.AppmeshVirtualService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appmeshvirtualservicesResource, "status", c.ns, appmeshVirtualService), &v1alpha1.AppmeshVirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshVirtualService), err
}

// Delete takes name of the appmeshVirtualService and deletes it. Returns an error if one occurs.
func (c *FakeAppmeshVirtualServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appmeshvirtualservicesResource, c.ns, name), &v1alpha1.AppmeshVirtualService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppmeshVirtualServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appmeshvirtualservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppmeshVirtualServiceList{})
	return err
}

// Patch applies the patch and returns the patched appmeshVirtualService.
func (c *FakeAppmeshVirtualServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppmeshVirtualService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appmeshvirtualservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppmeshVirtualService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshVirtualService), err
}
