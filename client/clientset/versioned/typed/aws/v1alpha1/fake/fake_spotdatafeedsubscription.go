/*
Copyright 2019 The Kubeform Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeSpotDatafeedSubscriptions implements SpotDatafeedSubscriptionInterface
type FakeSpotDatafeedSubscriptions struct {
	Fake *FakeAwsV1alpha1
}

var spotdatafeedsubscriptionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "spotdatafeedsubscriptions"}

var spotdatafeedsubscriptionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SpotDatafeedSubscription"}

// Get takes name of the spotDatafeedSubscription, and returns the corresponding spotDatafeedSubscription object, and an error if there is any.
func (c *FakeSpotDatafeedSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(spotdatafeedsubscriptionsResource, name), &v1alpha1.SpotDatafeedSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotDatafeedSubscription), err
}

// List takes label and field selectors, and returns the list of SpotDatafeedSubscriptions that match those selectors.
func (c *FakeSpotDatafeedSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.SpotDatafeedSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(spotdatafeedsubscriptionsResource, spotdatafeedsubscriptionsKind, opts), &v1alpha1.SpotDatafeedSubscriptionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpotDatafeedSubscriptionList{ListMeta: obj.(*v1alpha1.SpotDatafeedSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpotDatafeedSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spotDatafeedSubscriptions.
func (c *FakeSpotDatafeedSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(spotdatafeedsubscriptionsResource, opts))
}

// Create takes the representation of a spotDatafeedSubscription and creates it.  Returns the server's representation of the spotDatafeedSubscription, and an error, if there is any.
func (c *FakeSpotDatafeedSubscriptions) Create(spotDatafeedSubscription *v1alpha1.SpotDatafeedSubscription) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(spotdatafeedsubscriptionsResource, spotDatafeedSubscription), &v1alpha1.SpotDatafeedSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotDatafeedSubscription), err
}

// Update takes the representation of a spotDatafeedSubscription and updates it. Returns the server's representation of the spotDatafeedSubscription, and an error, if there is any.
func (c *FakeSpotDatafeedSubscriptions) Update(spotDatafeedSubscription *v1alpha1.SpotDatafeedSubscription) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(spotdatafeedsubscriptionsResource, spotDatafeedSubscription), &v1alpha1.SpotDatafeedSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotDatafeedSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpotDatafeedSubscriptions) UpdateStatus(spotDatafeedSubscription *v1alpha1.SpotDatafeedSubscription) (*v1alpha1.SpotDatafeedSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(spotdatafeedsubscriptionsResource, "status", spotDatafeedSubscription), &v1alpha1.SpotDatafeedSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotDatafeedSubscription), err
}

// Delete takes name of the spotDatafeedSubscription and deletes it. Returns an error if one occurs.
func (c *FakeSpotDatafeedSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(spotdatafeedsubscriptionsResource, name), &v1alpha1.SpotDatafeedSubscription{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpotDatafeedSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(spotdatafeedsubscriptionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpotDatafeedSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched spotDatafeedSubscription.
func (c *FakeSpotDatafeedSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(spotdatafeedsubscriptionsResource, name, pt, data, subresources...), &v1alpha1.SpotDatafeedSubscription{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotDatafeedSubscription), err
}
