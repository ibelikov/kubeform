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

// FakePinpointApps implements PinpointAppInterface
type FakePinpointApps struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var pinpointappsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointapps"}

var pinpointappsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointApp"}

// Get takes name of the pinpointApp, and returns the corresponding pinpointApp object, and an error if there is any.
func (c *FakePinpointApps) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinpointappsResource, c.ns, name), &v1alpha1.PinpointApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApp), err
}

// List takes label and field selectors, and returns the list of PinpointApps that match those selectors.
func (c *FakePinpointApps) List(opts v1.ListOptions) (result *v1alpha1.PinpointAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinpointappsResource, pinpointappsKind, c.ns, opts), &v1alpha1.PinpointAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointAppList{ListMeta: obj.(*v1alpha1.PinpointAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointApps.
func (c *FakePinpointApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinpointappsResource, c.ns, opts))

}

// Create takes the representation of a pinpointApp and creates it.  Returns the server's representation of the pinpointApp, and an error, if there is any.
func (c *FakePinpointApps) Create(pinpointApp *v1alpha1.PinpointApp) (result *v1alpha1.PinpointApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinpointappsResource, c.ns, pinpointApp), &v1alpha1.PinpointApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApp), err
}

// Update takes the representation of a pinpointApp and updates it. Returns the server's representation of the pinpointApp, and an error, if there is any.
func (c *FakePinpointApps) Update(pinpointApp *v1alpha1.PinpointApp) (result *v1alpha1.PinpointApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinpointappsResource, c.ns, pinpointApp), &v1alpha1.PinpointApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointApps) UpdateStatus(pinpointApp *v1alpha1.PinpointApp) (*v1alpha1.PinpointApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pinpointappsResource, "status", c.ns, pinpointApp), &v1alpha1.PinpointApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApp), err
}

// Delete takes name of the pinpointApp and deletes it. Returns an error if one occurs.
func (c *FakePinpointApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinpointappsResource, c.ns, name), &v1alpha1.PinpointApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinpointappsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointAppList{})
	return err
}

// Patch applies the patch and returns the patched pinpointApp.
func (c *FakePinpointApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinpointappsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinpointApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApp), err
}
