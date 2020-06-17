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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeLbListeners implements LbListenerInterface
type FakeLbListeners struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var lblistenersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lblisteners"}

var lblistenersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LbListener"}

// Get takes name of the lbListener, and returns the corresponding lbListener object, and an error if there is any.
func (c *FakeLbListeners) Get(name string, options v1.GetOptions) (result *v1alpha1.LbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lblistenersResource, c.ns, name), &v1alpha1.LbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListener), err
}

// List takes label and field selectors, and returns the list of LbListeners that match those selectors.
func (c *FakeLbListeners) List(opts v1.ListOptions) (result *v1alpha1.LbListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lblistenersResource, lblistenersKind, c.ns, opts), &v1alpha1.LbListenerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbListenerList{ListMeta: obj.(*v1alpha1.LbListenerList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbListeners.
func (c *FakeLbListeners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lblistenersResource, c.ns, opts))

}

// Create takes the representation of a lbListener and creates it.  Returns the server's representation of the lbListener, and an error, if there is any.
func (c *FakeLbListeners) Create(lbListener *v1alpha1.LbListener) (result *v1alpha1.LbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lblistenersResource, c.ns, lbListener), &v1alpha1.LbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListener), err
}

// Update takes the representation of a lbListener and updates it. Returns the server's representation of the lbListener, and an error, if there is any.
func (c *FakeLbListeners) Update(lbListener *v1alpha1.LbListener) (result *v1alpha1.LbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lblistenersResource, c.ns, lbListener), &v1alpha1.LbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListener), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbListeners) UpdateStatus(lbListener *v1alpha1.LbListener) (*v1alpha1.LbListener, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lblistenersResource, "status", c.ns, lbListener), &v1alpha1.LbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListener), err
}

// Delete takes name of the lbListener and deletes it. Returns an error if one occurs.
func (c *FakeLbListeners) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lblistenersResource, c.ns, name), &v1alpha1.LbListener{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbListeners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lblistenersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbListenerList{})
	return err
}

// Patch applies the patch and returns the patched lbListener.
func (c *FakeLbListeners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lblistenersResource, c.ns, name, pt, data, subresources...), &v1alpha1.LbListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListener), err
}
