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

// FakeDmsCertificates implements DmsCertificateInterface
type FakeDmsCertificates struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dmscertificatesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dmscertificates"}

var dmscertificatesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DmsCertificate"}

// Get takes name of the dmsCertificate, and returns the corresponding dmsCertificate object, and an error if there is any.
func (c *FakeDmsCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DmsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dmscertificatesResource, c.ns, name), &v1alpha1.DmsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsCertificate), err
}

// List takes label and field selectors, and returns the list of DmsCertificates that match those selectors.
func (c *FakeDmsCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DmsCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dmscertificatesResource, dmscertificatesKind, c.ns, opts), &v1alpha1.DmsCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DmsCertificateList{ListMeta: obj.(*v1alpha1.DmsCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.DmsCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dmsCertificates.
func (c *FakeDmsCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dmscertificatesResource, c.ns, opts))

}

// Create takes the representation of a dmsCertificate and creates it.  Returns the server's representation of the dmsCertificate, and an error, if there is any.
func (c *FakeDmsCertificates) Create(ctx context.Context, dmsCertificate *v1alpha1.DmsCertificate, opts v1.CreateOptions) (result *v1alpha1.DmsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dmscertificatesResource, c.ns, dmsCertificate), &v1alpha1.DmsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsCertificate), err
}

// Update takes the representation of a dmsCertificate and updates it. Returns the server's representation of the dmsCertificate, and an error, if there is any.
func (c *FakeDmsCertificates) Update(ctx context.Context, dmsCertificate *v1alpha1.DmsCertificate, opts v1.UpdateOptions) (result *v1alpha1.DmsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dmscertificatesResource, c.ns, dmsCertificate), &v1alpha1.DmsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDmsCertificates) UpdateStatus(ctx context.Context, dmsCertificate *v1alpha1.DmsCertificate, opts v1.UpdateOptions) (*v1alpha1.DmsCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dmscertificatesResource, "status", c.ns, dmsCertificate), &v1alpha1.DmsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsCertificate), err
}

// Delete takes name of the dmsCertificate and deletes it. Returns an error if one occurs.
func (c *FakeDmsCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dmscertificatesResource, c.ns, name), &v1alpha1.DmsCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDmsCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dmscertificatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DmsCertificateList{})
	return err
}

// Patch applies the patch and returns the patched dmsCertificate.
func (c *FakeDmsCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DmsCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dmscertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DmsCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsCertificate), err
}
