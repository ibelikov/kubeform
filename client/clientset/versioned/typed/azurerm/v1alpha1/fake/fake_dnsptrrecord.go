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

// FakeDnsPtrRecords implements DnsPtrRecordInterface
type FakeDnsPtrRecords struct {
	Fake *FakeAzurermV1alpha1
}

var dnsptrrecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "dnsptrrecords"}

var dnsptrrecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DnsPtrRecord"}

// Get takes name of the dnsPtrRecord, and returns the corresponding dnsPtrRecord object, and an error if there is any.
func (c *FakeDnsPtrRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dnsptrrecordsResource, name), &v1alpha1.DnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsPtrRecord), err
}

// List takes label and field selectors, and returns the list of DnsPtrRecords that match those selectors.
func (c *FakeDnsPtrRecords) List(opts v1.ListOptions) (result *v1alpha1.DnsPtrRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dnsptrrecordsResource, dnsptrrecordsKind, opts), &v1alpha1.DnsPtrRecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DnsPtrRecordList{ListMeta: obj.(*v1alpha1.DnsPtrRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.DnsPtrRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnsPtrRecords.
func (c *FakeDnsPtrRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dnsptrrecordsResource, opts))
}

// Create takes the representation of a dnsPtrRecord and creates it.  Returns the server's representation of the dnsPtrRecord, and an error, if there is any.
func (c *FakeDnsPtrRecords) Create(dnsPtrRecord *v1alpha1.DnsPtrRecord) (result *v1alpha1.DnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dnsptrrecordsResource, dnsPtrRecord), &v1alpha1.DnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsPtrRecord), err
}

// Update takes the representation of a dnsPtrRecord and updates it. Returns the server's representation of the dnsPtrRecord, and an error, if there is any.
func (c *FakeDnsPtrRecords) Update(dnsPtrRecord *v1alpha1.DnsPtrRecord) (result *v1alpha1.DnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dnsptrrecordsResource, dnsPtrRecord), &v1alpha1.DnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsPtrRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnsPtrRecords) UpdateStatus(dnsPtrRecord *v1alpha1.DnsPtrRecord) (*v1alpha1.DnsPtrRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dnsptrrecordsResource, "status", dnsPtrRecord), &v1alpha1.DnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsPtrRecord), err
}

// Delete takes name of the dnsPtrRecord and deletes it. Returns an error if one occurs.
func (c *FakeDnsPtrRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dnsptrrecordsResource, name), &v1alpha1.DnsPtrRecord{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnsPtrRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dnsptrrecordsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DnsPtrRecordList{})
	return err
}

// Patch applies the patch and returns the patched dnsPtrRecord.
func (c *FakeDnsPtrRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dnsptrrecordsResource, name, pt, data, subresources...), &v1alpha1.DnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsPtrRecord), err
}
