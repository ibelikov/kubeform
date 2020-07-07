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

// FakeServiceAccountIamMembers implements ServiceAccountIamMemberInterface
type FakeServiceAccountIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var serviceaccountiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "serviceaccountiammembers"}

var serviceaccountiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ServiceAccountIamMember"}

// Get takes name of the serviceAccountIamMember, and returns the corresponding serviceAccountIamMember object, and an error if there is any.
func (c *FakeServiceAccountIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceaccountiammembersResource, c.ns, name), &v1alpha1.ServiceAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountIamMember), err
}

// List takes label and field selectors, and returns the list of ServiceAccountIamMembers that match those selectors.
func (c *FakeServiceAccountIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceAccountIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceaccountiammembersResource, serviceaccountiammembersKind, c.ns, opts), &v1alpha1.ServiceAccountIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceAccountIamMemberList{ListMeta: obj.(*v1alpha1.ServiceAccountIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceAccountIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceAccountIamMembers.
func (c *FakeServiceAccountIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceaccountiammembersResource, c.ns, opts))

}

// Create takes the representation of a serviceAccountIamMember and creates it.  Returns the server's representation of the serviceAccountIamMember, and an error, if there is any.
func (c *FakeServiceAccountIamMembers) Create(ctx context.Context, serviceAccountIamMember *v1alpha1.ServiceAccountIamMember, opts v1.CreateOptions) (result *v1alpha1.ServiceAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceaccountiammembersResource, c.ns, serviceAccountIamMember), &v1alpha1.ServiceAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountIamMember), err
}

// Update takes the representation of a serviceAccountIamMember and updates it. Returns the server's representation of the serviceAccountIamMember, and an error, if there is any.
func (c *FakeServiceAccountIamMembers) Update(ctx context.Context, serviceAccountIamMember *v1alpha1.ServiceAccountIamMember, opts v1.UpdateOptions) (result *v1alpha1.ServiceAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceaccountiammembersResource, c.ns, serviceAccountIamMember), &v1alpha1.ServiceAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceAccountIamMembers) UpdateStatus(ctx context.Context, serviceAccountIamMember *v1alpha1.ServiceAccountIamMember, opts v1.UpdateOptions) (*v1alpha1.ServiceAccountIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceaccountiammembersResource, "status", c.ns, serviceAccountIamMember), &v1alpha1.ServiceAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountIamMember), err
}

// Delete takes name of the serviceAccountIamMember and deletes it. Returns an error if one occurs.
func (c *FakeServiceAccountIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceaccountiammembersResource, c.ns, name), &v1alpha1.ServiceAccountIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceAccountIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceaccountiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceAccountIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched serviceAccountIamMember.
func (c *FakeServiceAccountIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceAccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceaccountiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceAccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAccountIamMember), err
}
