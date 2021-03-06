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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
)

// FakeAzureAppServices implements AzureAppServiceInterface
type FakeAzureAppServices struct {
	Fake *FakeModulesV1alpha1
	ns   string
}

var azureappservicesResource = schema.GroupVersionResource{Group: "modules.kubeform.com", Version: "v1alpha1", Resource: "azureappservices"}

var azureappservicesKind = schema.GroupVersionKind{Group: "modules.kubeform.com", Version: "v1alpha1", Kind: "AzureAppService"}

// Get takes name of the azureAppService, and returns the corresponding azureAppService object, and an error if there is any.
func (c *FakeAzureAppServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureAppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureappservicesResource, c.ns, name), &v1alpha1.AzureAppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAppService), err
}

// List takes label and field selectors, and returns the list of AzureAppServices that match those selectors.
func (c *FakeAzureAppServices) List(opts v1.ListOptions) (result *v1alpha1.AzureAppServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureappservicesResource, azureappservicesKind, c.ns, opts), &v1alpha1.AzureAppServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureAppServiceList{ListMeta: obj.(*v1alpha1.AzureAppServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureAppServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureAppServices.
func (c *FakeAzureAppServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureappservicesResource, c.ns, opts))

}

// Create takes the representation of a azureAppService and creates it.  Returns the server's representation of the azureAppService, and an error, if there is any.
func (c *FakeAzureAppServices) Create(azureAppService *v1alpha1.AzureAppService) (result *v1alpha1.AzureAppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureappservicesResource, c.ns, azureAppService), &v1alpha1.AzureAppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAppService), err
}

// Update takes the representation of a azureAppService and updates it. Returns the server's representation of the azureAppService, and an error, if there is any.
func (c *FakeAzureAppServices) Update(azureAppService *v1alpha1.AzureAppService) (result *v1alpha1.AzureAppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureappservicesResource, c.ns, azureAppService), &v1alpha1.AzureAppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAppService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureAppServices) UpdateStatus(azureAppService *v1alpha1.AzureAppService) (*v1alpha1.AzureAppService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azureappservicesResource, "status", c.ns, azureAppService), &v1alpha1.AzureAppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAppService), err
}

// Delete takes name of the azureAppService and deletes it. Returns an error if one occurs.
func (c *FakeAzureAppServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azureappservicesResource, c.ns, name), &v1alpha1.AzureAppService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureAppServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureappservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureAppServiceList{})
	return err
}

// Patch applies the patch and returns the patched azureAppService.
func (c *FakeAzureAppServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureAppService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureappservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureAppService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureAppService), err
}
