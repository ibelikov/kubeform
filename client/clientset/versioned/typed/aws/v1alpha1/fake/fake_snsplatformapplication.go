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

// FakeSnsPlatformApplications implements SnsPlatformApplicationInterface
type FakeSnsPlatformApplications struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var snsplatformapplicationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "snsplatformapplications"}

var snsplatformapplicationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SnsPlatformApplication"}

// Get takes name of the snsPlatformApplication, and returns the corresponding snsPlatformApplication object, and an error if there is any.
func (c *FakeSnsPlatformApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SnsPlatformApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snsplatformapplicationsResource, c.ns, name), &v1alpha1.SnsPlatformApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsPlatformApplication), err
}

// List takes label and field selectors, and returns the list of SnsPlatformApplications that match those selectors.
func (c *FakeSnsPlatformApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SnsPlatformApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snsplatformapplicationsResource, snsplatformapplicationsKind, c.ns, opts), &v1alpha1.SnsPlatformApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SnsPlatformApplicationList{ListMeta: obj.(*v1alpha1.SnsPlatformApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SnsPlatformApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snsPlatformApplications.
func (c *FakeSnsPlatformApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snsplatformapplicationsResource, c.ns, opts))

}

// Create takes the representation of a snsPlatformApplication and creates it.  Returns the server's representation of the snsPlatformApplication, and an error, if there is any.
func (c *FakeSnsPlatformApplications) Create(ctx context.Context, snsPlatformApplication *v1alpha1.SnsPlatformApplication, opts v1.CreateOptions) (result *v1alpha1.SnsPlatformApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snsplatformapplicationsResource, c.ns, snsPlatformApplication), &v1alpha1.SnsPlatformApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsPlatformApplication), err
}

// Update takes the representation of a snsPlatformApplication and updates it. Returns the server's representation of the snsPlatformApplication, and an error, if there is any.
func (c *FakeSnsPlatformApplications) Update(ctx context.Context, snsPlatformApplication *v1alpha1.SnsPlatformApplication, opts v1.UpdateOptions) (result *v1alpha1.SnsPlatformApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snsplatformapplicationsResource, c.ns, snsPlatformApplication), &v1alpha1.SnsPlatformApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsPlatformApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSnsPlatformApplications) UpdateStatus(ctx context.Context, snsPlatformApplication *v1alpha1.SnsPlatformApplication, opts v1.UpdateOptions) (*v1alpha1.SnsPlatformApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(snsplatformapplicationsResource, "status", c.ns, snsPlatformApplication), &v1alpha1.SnsPlatformApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsPlatformApplication), err
}

// Delete takes name of the snsPlatformApplication and deletes it. Returns an error if one occurs.
func (c *FakeSnsPlatformApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snsplatformapplicationsResource, c.ns, name), &v1alpha1.SnsPlatformApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnsPlatformApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snsplatformapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SnsPlatformApplicationList{})
	return err
}

// Patch applies the patch and returns the patched snsPlatformApplication.
func (c *FakeSnsPlatformApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SnsPlatformApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snsplatformapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SnsPlatformApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsPlatformApplication), err
}
