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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppServicePlans implements AppServicePlanInterface
type FakeAppServicePlans struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var appserviceplansResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "appserviceplans"}

var appserviceplansKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AppServicePlan"}

// Get takes name of the appServicePlan, and returns the corresponding appServicePlan object, and an error if there is any.
func (c *FakeAppServicePlans) Get(name string, options v1.GetOptions) (result *v1alpha1.AppServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appserviceplansResource, c.ns, name), &v1alpha1.AppServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServicePlan), err
}

// List takes label and field selectors, and returns the list of AppServicePlans that match those selectors.
func (c *FakeAppServicePlans) List(opts v1.ListOptions) (result *v1alpha1.AppServicePlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appserviceplansResource, appserviceplansKind, c.ns, opts), &v1alpha1.AppServicePlanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppServicePlanList{ListMeta: obj.(*v1alpha1.AppServicePlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppServicePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appServicePlans.
func (c *FakeAppServicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appserviceplansResource, c.ns, opts))

}

// Create takes the representation of a appServicePlan and creates it.  Returns the server's representation of the appServicePlan, and an error, if there is any.
func (c *FakeAppServicePlans) Create(appServicePlan *v1alpha1.AppServicePlan) (result *v1alpha1.AppServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appserviceplansResource, c.ns, appServicePlan), &v1alpha1.AppServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServicePlan), err
}

// Update takes the representation of a appServicePlan and updates it. Returns the server's representation of the appServicePlan, and an error, if there is any.
func (c *FakeAppServicePlans) Update(appServicePlan *v1alpha1.AppServicePlan) (result *v1alpha1.AppServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appserviceplansResource, c.ns, appServicePlan), &v1alpha1.AppServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServicePlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppServicePlans) UpdateStatus(appServicePlan *v1alpha1.AppServicePlan) (*v1alpha1.AppServicePlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appserviceplansResource, "status", c.ns, appServicePlan), &v1alpha1.AppServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServicePlan), err
}

// Delete takes name of the appServicePlan and deletes it. Returns an error if one occurs.
func (c *FakeAppServicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appserviceplansResource, c.ns, name), &v1alpha1.AppServicePlan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppServicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appserviceplansResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppServicePlanList{})
	return err
}

// Patch applies the patch and returns the patched appServicePlan.
func (c *FakeAppServicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServicePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appserviceplansResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppServicePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServicePlan), err
}
