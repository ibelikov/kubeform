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

// FakeGlueCrawlers implements GlueCrawlerInterface
type FakeGlueCrawlers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var gluecrawlersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "gluecrawlers"}

var gluecrawlersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GlueCrawler"}

// Get takes name of the glueCrawler, and returns the corresponding glueCrawler object, and an error if there is any.
func (c *FakeGlueCrawlers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gluecrawlersResource, c.ns, name), &v1alpha1.GlueCrawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCrawler), err
}

// List takes label and field selectors, and returns the list of GlueCrawlers that match those selectors.
func (c *FakeGlueCrawlers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GlueCrawlerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gluecrawlersResource, gluecrawlersKind, c.ns, opts), &v1alpha1.GlueCrawlerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlueCrawlerList{ListMeta: obj.(*v1alpha1.GlueCrawlerList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlueCrawlerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested glueCrawlers.
func (c *FakeGlueCrawlers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gluecrawlersResource, c.ns, opts))

}

// Create takes the representation of a glueCrawler and creates it.  Returns the server's representation of the glueCrawler, and an error, if there is any.
func (c *FakeGlueCrawlers) Create(ctx context.Context, glueCrawler *v1alpha1.GlueCrawler, opts v1.CreateOptions) (result *v1alpha1.GlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gluecrawlersResource, c.ns, glueCrawler), &v1alpha1.GlueCrawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCrawler), err
}

// Update takes the representation of a glueCrawler and updates it. Returns the server's representation of the glueCrawler, and an error, if there is any.
func (c *FakeGlueCrawlers) Update(ctx context.Context, glueCrawler *v1alpha1.GlueCrawler, opts v1.UpdateOptions) (result *v1alpha1.GlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gluecrawlersResource, c.ns, glueCrawler), &v1alpha1.GlueCrawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCrawler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlueCrawlers) UpdateStatus(ctx context.Context, glueCrawler *v1alpha1.GlueCrawler, opts v1.UpdateOptions) (*v1alpha1.GlueCrawler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gluecrawlersResource, "status", c.ns, glueCrawler), &v1alpha1.GlueCrawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCrawler), err
}

// Delete takes name of the glueCrawler and deletes it. Returns an error if one occurs.
func (c *FakeGlueCrawlers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gluecrawlersResource, c.ns, name), &v1alpha1.GlueCrawler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlueCrawlers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gluecrawlersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlueCrawlerList{})
	return err
}

// Patch applies the patch and returns the patched glueCrawler.
func (c *FakeGlueCrawlers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gluecrawlersResource, c.ns, name, pt, data, subresources...), &v1alpha1.GlueCrawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueCrawler), err
}
