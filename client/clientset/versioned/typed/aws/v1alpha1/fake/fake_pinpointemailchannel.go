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

// FakePinpointEmailChannels implements PinpointEmailChannelInterface
type FakePinpointEmailChannels struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var pinpointemailchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointemailchannels"}

var pinpointemailchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointEmailChannel"}

// Get takes name of the pinpointEmailChannel, and returns the corresponding pinpointEmailChannel object, and an error if there is any.
func (c *FakePinpointEmailChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PinpointEmailChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinpointemailchannelsResource, c.ns, name), &v1alpha1.PinpointEmailChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointEmailChannel), err
}

// List takes label and field selectors, and returns the list of PinpointEmailChannels that match those selectors.
func (c *FakePinpointEmailChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PinpointEmailChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinpointemailchannelsResource, pinpointemailchannelsKind, c.ns, opts), &v1alpha1.PinpointEmailChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointEmailChannelList{ListMeta: obj.(*v1alpha1.PinpointEmailChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointEmailChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointEmailChannels.
func (c *FakePinpointEmailChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinpointemailchannelsResource, c.ns, opts))

}

// Create takes the representation of a pinpointEmailChannel and creates it.  Returns the server's representation of the pinpointEmailChannel, and an error, if there is any.
func (c *FakePinpointEmailChannels) Create(ctx context.Context, pinpointEmailChannel *v1alpha1.PinpointEmailChannel, opts v1.CreateOptions) (result *v1alpha1.PinpointEmailChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinpointemailchannelsResource, c.ns, pinpointEmailChannel), &v1alpha1.PinpointEmailChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointEmailChannel), err
}

// Update takes the representation of a pinpointEmailChannel and updates it. Returns the server's representation of the pinpointEmailChannel, and an error, if there is any.
func (c *FakePinpointEmailChannels) Update(ctx context.Context, pinpointEmailChannel *v1alpha1.PinpointEmailChannel, opts v1.UpdateOptions) (result *v1alpha1.PinpointEmailChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinpointemailchannelsResource, c.ns, pinpointEmailChannel), &v1alpha1.PinpointEmailChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointEmailChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointEmailChannels) UpdateStatus(ctx context.Context, pinpointEmailChannel *v1alpha1.PinpointEmailChannel, opts v1.UpdateOptions) (*v1alpha1.PinpointEmailChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pinpointemailchannelsResource, "status", c.ns, pinpointEmailChannel), &v1alpha1.PinpointEmailChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointEmailChannel), err
}

// Delete takes name of the pinpointEmailChannel and deletes it. Returns an error if one occurs.
func (c *FakePinpointEmailChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinpointemailchannelsResource, c.ns, name), &v1alpha1.PinpointEmailChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointEmailChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinpointemailchannelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointEmailChannelList{})
	return err
}

// Patch applies the patch and returns the patched pinpointEmailChannel.
func (c *FakePinpointEmailChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PinpointEmailChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinpointemailchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinpointEmailChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointEmailChannel), err
}
