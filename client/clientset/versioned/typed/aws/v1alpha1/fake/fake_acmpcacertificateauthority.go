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

// FakeAcmpcaCertificateAuthorities implements AcmpcaCertificateAuthorityInterface
type FakeAcmpcaCertificateAuthorities struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var acmpcacertificateauthoritiesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "acmpcacertificateauthorities"}

var acmpcacertificateauthoritiesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AcmpcaCertificateAuthority"}

// Get takes name of the acmpcaCertificateAuthority, and returns the corresponding acmpcaCertificateAuthority object, and an error if there is any.
func (c *FakeAcmpcaCertificateAuthorities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AcmpcaCertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(acmpcacertificateauthoritiesResource, c.ns, name), &v1alpha1.AcmpcaCertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmpcaCertificateAuthority), err
}

// List takes label and field selectors, and returns the list of AcmpcaCertificateAuthorities that match those selectors.
func (c *FakeAcmpcaCertificateAuthorities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AcmpcaCertificateAuthorityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(acmpcacertificateauthoritiesResource, acmpcacertificateauthoritiesKind, c.ns, opts), &v1alpha1.AcmpcaCertificateAuthorityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AcmpcaCertificateAuthorityList{ListMeta: obj.(*v1alpha1.AcmpcaCertificateAuthorityList).ListMeta}
	for _, item := range obj.(*v1alpha1.AcmpcaCertificateAuthorityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested acmpcaCertificateAuthorities.
func (c *FakeAcmpcaCertificateAuthorities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(acmpcacertificateauthoritiesResource, c.ns, opts))

}

// Create takes the representation of a acmpcaCertificateAuthority and creates it.  Returns the server's representation of the acmpcaCertificateAuthority, and an error, if there is any.
func (c *FakeAcmpcaCertificateAuthorities) Create(ctx context.Context, acmpcaCertificateAuthority *v1alpha1.AcmpcaCertificateAuthority, opts v1.CreateOptions) (result *v1alpha1.AcmpcaCertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(acmpcacertificateauthoritiesResource, c.ns, acmpcaCertificateAuthority), &v1alpha1.AcmpcaCertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmpcaCertificateAuthority), err
}

// Update takes the representation of a acmpcaCertificateAuthority and updates it. Returns the server's representation of the acmpcaCertificateAuthority, and an error, if there is any.
func (c *FakeAcmpcaCertificateAuthorities) Update(ctx context.Context, acmpcaCertificateAuthority *v1alpha1.AcmpcaCertificateAuthority, opts v1.UpdateOptions) (result *v1alpha1.AcmpcaCertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(acmpcacertificateauthoritiesResource, c.ns, acmpcaCertificateAuthority), &v1alpha1.AcmpcaCertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmpcaCertificateAuthority), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAcmpcaCertificateAuthorities) UpdateStatus(ctx context.Context, acmpcaCertificateAuthority *v1alpha1.AcmpcaCertificateAuthority, opts v1.UpdateOptions) (*v1alpha1.AcmpcaCertificateAuthority, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(acmpcacertificateauthoritiesResource, "status", c.ns, acmpcaCertificateAuthority), &v1alpha1.AcmpcaCertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmpcaCertificateAuthority), err
}

// Delete takes name of the acmpcaCertificateAuthority and deletes it. Returns an error if one occurs.
func (c *FakeAcmpcaCertificateAuthorities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(acmpcacertificateauthoritiesResource, c.ns, name), &v1alpha1.AcmpcaCertificateAuthority{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAcmpcaCertificateAuthorities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(acmpcacertificateauthoritiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AcmpcaCertificateAuthorityList{})
	return err
}

// Patch applies the patch and returns the patched acmpcaCertificateAuthority.
func (c *FakeAcmpcaCertificateAuthorities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AcmpcaCertificateAuthority, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(acmpcacertificateauthoritiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AcmpcaCertificateAuthority{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AcmpcaCertificateAuthority), err
}
