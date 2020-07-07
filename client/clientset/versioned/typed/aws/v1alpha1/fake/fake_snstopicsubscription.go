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

// FakeSnsTopicSubscriptions implements SnsTopicSubscriptionInterface
type FakeSnsTopicSubscriptions struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var snstopicsubscriptionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "snstopicsubscriptions"}

var snstopicsubscriptionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SnsTopicSubscription"}

// Get takes name of the snsTopicSubscription, and returns the corresponding snsTopicSubscription object, and an error if there is any.
func (c *FakeSnsTopicSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snstopicsubscriptionsResource, c.ns, name), &v1alpha1.SnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicSubscription), err
}

// List takes label and field selectors, and returns the list of SnsTopicSubscriptions that match those selectors.
func (c *FakeSnsTopicSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SnsTopicSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snstopicsubscriptionsResource, snstopicsubscriptionsKind, c.ns, opts), &v1alpha1.SnsTopicSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SnsTopicSubscriptionList{ListMeta: obj.(*v1alpha1.SnsTopicSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SnsTopicSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snsTopicSubscriptions.
func (c *FakeSnsTopicSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snstopicsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a snsTopicSubscription and creates it.  Returns the server's representation of the snsTopicSubscription, and an error, if there is any.
func (c *FakeSnsTopicSubscriptions) Create(ctx context.Context, snsTopicSubscription *v1alpha1.SnsTopicSubscription, opts v1.CreateOptions) (result *v1alpha1.SnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snstopicsubscriptionsResource, c.ns, snsTopicSubscription), &v1alpha1.SnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicSubscription), err
}

// Update takes the representation of a snsTopicSubscription and updates it. Returns the server's representation of the snsTopicSubscription, and an error, if there is any.
func (c *FakeSnsTopicSubscriptions) Update(ctx context.Context, snsTopicSubscription *v1alpha1.SnsTopicSubscription, opts v1.UpdateOptions) (result *v1alpha1.SnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snstopicsubscriptionsResource, c.ns, snsTopicSubscription), &v1alpha1.SnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSnsTopicSubscriptions) UpdateStatus(ctx context.Context, snsTopicSubscription *v1alpha1.SnsTopicSubscription, opts v1.UpdateOptions) (*v1alpha1.SnsTopicSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(snstopicsubscriptionsResource, "status", c.ns, snsTopicSubscription), &v1alpha1.SnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicSubscription), err
}

// Delete takes name of the snsTopicSubscription and deletes it. Returns an error if one occurs.
func (c *FakeSnsTopicSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snstopicsubscriptionsResource, c.ns, name), &v1alpha1.SnsTopicSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnsTopicSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snstopicsubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SnsTopicSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched snsTopicSubscription.
func (c *FakeSnsTopicSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SnsTopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snstopicsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SnsTopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicSubscription), err
}
