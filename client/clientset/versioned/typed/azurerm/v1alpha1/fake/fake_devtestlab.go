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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeDevTestLabs implements DevTestLabInterface
type FakeDevTestLabs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var devtestlabsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "devtestlabs"}

var devtestlabsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DevTestLab"}

// Get takes name of the devTestLab, and returns the corresponding devTestLab object, and an error if there is any.
func (c *FakeDevTestLabs) Get(name string, options v1.GetOptions) (result *v1alpha1.DevTestLab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devtestlabsResource, c.ns, name), &v1alpha1.DevTestLab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLab), err
}

// List takes label and field selectors, and returns the list of DevTestLabs that match those selectors.
func (c *FakeDevTestLabs) List(opts v1.ListOptions) (result *v1alpha1.DevTestLabList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devtestlabsResource, devtestlabsKind, c.ns, opts), &v1alpha1.DevTestLabList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevTestLabList{ListMeta: obj.(*v1alpha1.DevTestLabList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevTestLabList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devTestLabs.
func (c *FakeDevTestLabs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devtestlabsResource, c.ns, opts))

}

// Create takes the representation of a devTestLab and creates it.  Returns the server's representation of the devTestLab, and an error, if there is any.
func (c *FakeDevTestLabs) Create(devTestLab *v1alpha1.DevTestLab) (result *v1alpha1.DevTestLab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devtestlabsResource, c.ns, devTestLab), &v1alpha1.DevTestLab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLab), err
}

// Update takes the representation of a devTestLab and updates it. Returns the server's representation of the devTestLab, and an error, if there is any.
func (c *FakeDevTestLabs) Update(devTestLab *v1alpha1.DevTestLab) (result *v1alpha1.DevTestLab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devtestlabsResource, c.ns, devTestLab), &v1alpha1.DevTestLab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLab), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevTestLabs) UpdateStatus(devTestLab *v1alpha1.DevTestLab) (*v1alpha1.DevTestLab, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devtestlabsResource, "status", c.ns, devTestLab), &v1alpha1.DevTestLab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLab), err
}

// Delete takes name of the devTestLab and deletes it. Returns an error if one occurs.
func (c *FakeDevTestLabs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devtestlabsResource, c.ns, name), &v1alpha1.DevTestLab{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevTestLabs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devtestlabsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevTestLabList{})
	return err
}

// Patch applies the patch and returns the patched devTestLab.
func (c *FakeDevTestLabs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevTestLab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devtestlabsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevTestLab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestLab), err
}
