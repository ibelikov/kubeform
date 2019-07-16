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

// FakeMonitorActivityLogAlerts implements MonitorActivityLogAlertInterface
type FakeMonitorActivityLogAlerts struct {
	Fake *FakeAzurermV1alpha1
}

var monitoractivitylogalertsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "monitoractivitylogalerts"}

var monitoractivitylogalertsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MonitorActivityLogAlert"}

// Get takes name of the monitorActivityLogAlert, and returns the corresponding monitorActivityLogAlert object, and an error if there is any.
func (c *FakeMonitorActivityLogAlerts) Get(name string, options v1.GetOptions) (result *v1alpha1.MonitorActivityLogAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(monitoractivitylogalertsResource, name), &v1alpha1.MonitorActivityLogAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActivityLogAlert), err
}

// List takes label and field selectors, and returns the list of MonitorActivityLogAlerts that match those selectors.
func (c *FakeMonitorActivityLogAlerts) List(opts v1.ListOptions) (result *v1alpha1.MonitorActivityLogAlertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(monitoractivitylogalertsResource, monitoractivitylogalertsKind, opts), &v1alpha1.MonitorActivityLogAlertList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitorActivityLogAlertList{ListMeta: obj.(*v1alpha1.MonitorActivityLogAlertList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitorActivityLogAlertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitorActivityLogAlerts.
func (c *FakeMonitorActivityLogAlerts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(monitoractivitylogalertsResource, opts))
}

// Create takes the representation of a monitorActivityLogAlert and creates it.  Returns the server's representation of the monitorActivityLogAlert, and an error, if there is any.
func (c *FakeMonitorActivityLogAlerts) Create(monitorActivityLogAlert *v1alpha1.MonitorActivityLogAlert) (result *v1alpha1.MonitorActivityLogAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(monitoractivitylogalertsResource, monitorActivityLogAlert), &v1alpha1.MonitorActivityLogAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActivityLogAlert), err
}

// Update takes the representation of a monitorActivityLogAlert and updates it. Returns the server's representation of the monitorActivityLogAlert, and an error, if there is any.
func (c *FakeMonitorActivityLogAlerts) Update(monitorActivityLogAlert *v1alpha1.MonitorActivityLogAlert) (result *v1alpha1.MonitorActivityLogAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(monitoractivitylogalertsResource, monitorActivityLogAlert), &v1alpha1.MonitorActivityLogAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActivityLogAlert), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitorActivityLogAlerts) UpdateStatus(monitorActivityLogAlert *v1alpha1.MonitorActivityLogAlert) (*v1alpha1.MonitorActivityLogAlert, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(monitoractivitylogalertsResource, "status", monitorActivityLogAlert), &v1alpha1.MonitorActivityLogAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActivityLogAlert), err
}

// Delete takes name of the monitorActivityLogAlert and deletes it. Returns an error if one occurs.
func (c *FakeMonitorActivityLogAlerts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(monitoractivitylogalertsResource, name), &v1alpha1.MonitorActivityLogAlert{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitorActivityLogAlerts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(monitoractivitylogalertsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitorActivityLogAlertList{})
	return err
}

// Patch applies the patch and returns the patched monitorActivityLogAlert.
func (c *FakeMonitorActivityLogAlerts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MonitorActivityLogAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(monitoractivitylogalertsResource, name, pt, data, subresources...), &v1alpha1.MonitorActivityLogAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActivityLogAlert), err
}
