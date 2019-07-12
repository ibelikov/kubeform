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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermLogicAppTriggerHttpRequests implements AzurermLogicAppTriggerHttpRequestInterface
type FakeAzurermLogicAppTriggerHttpRequests struct {
	Fake *FakeAzurermV1alpha1
}

var azurermlogicapptriggerhttprequestsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermlogicapptriggerhttprequests"}

var azurermlogicapptriggerhttprequestsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermLogicAppTriggerHttpRequest"}

// Get takes name of the azurermLogicAppTriggerHttpRequest, and returns the corresponding azurermLogicAppTriggerHttpRequest object, and an error if there is any.
func (c *FakeAzurermLogicAppTriggerHttpRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLogicAppTriggerHttpRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermlogicapptriggerhttprequestsResource, name), &v1alpha1.AzurermLogicAppTriggerHttpRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerHttpRequest), err
}

// List takes label and field selectors, and returns the list of AzurermLogicAppTriggerHttpRequests that match those selectors.
func (c *FakeAzurermLogicAppTriggerHttpRequests) List(opts v1.ListOptions) (result *v1alpha1.AzurermLogicAppTriggerHttpRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermlogicapptriggerhttprequestsResource, azurermlogicapptriggerhttprequestsKind, opts), &v1alpha1.AzurermLogicAppTriggerHttpRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermLogicAppTriggerHttpRequestList{ListMeta: obj.(*v1alpha1.AzurermLogicAppTriggerHttpRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermLogicAppTriggerHttpRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermLogicAppTriggerHttpRequests.
func (c *FakeAzurermLogicAppTriggerHttpRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermlogicapptriggerhttprequestsResource, opts))
}

// Create takes the representation of a azurermLogicAppTriggerHttpRequest and creates it.  Returns the server's representation of the azurermLogicAppTriggerHttpRequest, and an error, if there is any.
func (c *FakeAzurermLogicAppTriggerHttpRequests) Create(azurermLogicAppTriggerHttpRequest *v1alpha1.AzurermLogicAppTriggerHttpRequest) (result *v1alpha1.AzurermLogicAppTriggerHttpRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermlogicapptriggerhttprequestsResource, azurermLogicAppTriggerHttpRequest), &v1alpha1.AzurermLogicAppTriggerHttpRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerHttpRequest), err
}

// Update takes the representation of a azurermLogicAppTriggerHttpRequest and updates it. Returns the server's representation of the azurermLogicAppTriggerHttpRequest, and an error, if there is any.
func (c *FakeAzurermLogicAppTriggerHttpRequests) Update(azurermLogicAppTriggerHttpRequest *v1alpha1.AzurermLogicAppTriggerHttpRequest) (result *v1alpha1.AzurermLogicAppTriggerHttpRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermlogicapptriggerhttprequestsResource, azurermLogicAppTriggerHttpRequest), &v1alpha1.AzurermLogicAppTriggerHttpRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerHttpRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermLogicAppTriggerHttpRequests) UpdateStatus(azurermLogicAppTriggerHttpRequest *v1alpha1.AzurermLogicAppTriggerHttpRequest) (*v1alpha1.AzurermLogicAppTriggerHttpRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermlogicapptriggerhttprequestsResource, "status", azurermLogicAppTriggerHttpRequest), &v1alpha1.AzurermLogicAppTriggerHttpRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerHttpRequest), err
}

// Delete takes name of the azurermLogicAppTriggerHttpRequest and deletes it. Returns an error if one occurs.
func (c *FakeAzurermLogicAppTriggerHttpRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermlogicapptriggerhttprequestsResource, name), &v1alpha1.AzurermLogicAppTriggerHttpRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermLogicAppTriggerHttpRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermlogicapptriggerhttprequestsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermLogicAppTriggerHttpRequestList{})
	return err
}

// Patch applies the patch and returns the patched azurermLogicAppTriggerHttpRequest.
func (c *FakeAzurermLogicAppTriggerHttpRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppTriggerHttpRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermlogicapptriggerhttprequestsResource, name, pt, data, subresources...), &v1alpha1.AzurermLogicAppTriggerHttpRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerHttpRequest), err
}
