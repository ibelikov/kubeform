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

// FakePubsubSubscriptions implements PubsubSubscriptionInterface
type FakePubsubSubscriptions struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var pubsubsubscriptionsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "pubsubsubscriptions"}

var pubsubsubscriptionsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "PubsubSubscription"}

// Get takes name of the pubsubSubscription, and returns the corresponding pubsubSubscription object, and an error if there is any.
func (c *FakePubsubSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PubsubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pubsubsubscriptionsResource, c.ns, name), &v1alpha1.PubsubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscription), err
}

// List takes label and field selectors, and returns the list of PubsubSubscriptions that match those selectors.
func (c *FakePubsubSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PubsubSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pubsubsubscriptionsResource, pubsubsubscriptionsKind, c.ns, opts), &v1alpha1.PubsubSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PubsubSubscriptionList{ListMeta: obj.(*v1alpha1.PubsubSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PubsubSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubsubSubscriptions.
func (c *FakePubsubSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pubsubsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a pubsubSubscription and creates it.  Returns the server's representation of the pubsubSubscription, and an error, if there is any.
func (c *FakePubsubSubscriptions) Create(ctx context.Context, pubsubSubscription *v1alpha1.PubsubSubscription, opts v1.CreateOptions) (result *v1alpha1.PubsubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pubsubsubscriptionsResource, c.ns, pubsubSubscription), &v1alpha1.PubsubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscription), err
}

// Update takes the representation of a pubsubSubscription and updates it. Returns the server's representation of the pubsubSubscription, and an error, if there is any.
func (c *FakePubsubSubscriptions) Update(ctx context.Context, pubsubSubscription *v1alpha1.PubsubSubscription, opts v1.UpdateOptions) (result *v1alpha1.PubsubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pubsubsubscriptionsResource, c.ns, pubsubSubscription), &v1alpha1.PubsubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubsubSubscriptions) UpdateStatus(ctx context.Context, pubsubSubscription *v1alpha1.PubsubSubscription, opts v1.UpdateOptions) (*v1alpha1.PubsubSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pubsubsubscriptionsResource, "status", c.ns, pubsubSubscription), &v1alpha1.PubsubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscription), err
}

// Delete takes name of the pubsubSubscription and deletes it. Returns an error if one occurs.
func (c *FakePubsubSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pubsubsubscriptionsResource, c.ns, name), &v1alpha1.PubsubSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubsubSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pubsubsubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PubsubSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched pubsubSubscription.
func (c *FakePubsubSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PubsubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pubsubsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PubsubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscription), err
}
