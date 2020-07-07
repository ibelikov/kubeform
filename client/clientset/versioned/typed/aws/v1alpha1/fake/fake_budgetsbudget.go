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

// FakeBudgetsBudgets implements BudgetsBudgetInterface
type FakeBudgetsBudgets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var budgetsbudgetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "budgetsbudgets"}

var budgetsbudgetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "BudgetsBudget"}

// Get takes name of the budgetsBudget, and returns the corresponding budgetsBudget object, and an error if there is any.
func (c *FakeBudgetsBudgets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(budgetsbudgetsResource, c.ns, name), &v1alpha1.BudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BudgetsBudget), err
}

// List takes label and field selectors, and returns the list of BudgetsBudgets that match those selectors.
func (c *FakeBudgetsBudgets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BudgetsBudgetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(budgetsbudgetsResource, budgetsbudgetsKind, c.ns, opts), &v1alpha1.BudgetsBudgetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BudgetsBudgetList{ListMeta: obj.(*v1alpha1.BudgetsBudgetList).ListMeta}
	for _, item := range obj.(*v1alpha1.BudgetsBudgetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested budgetsBudgets.
func (c *FakeBudgetsBudgets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(budgetsbudgetsResource, c.ns, opts))

}

// Create takes the representation of a budgetsBudget and creates it.  Returns the server's representation of the budgetsBudget, and an error, if there is any.
func (c *FakeBudgetsBudgets) Create(ctx context.Context, budgetsBudget *v1alpha1.BudgetsBudget, opts v1.CreateOptions) (result *v1alpha1.BudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(budgetsbudgetsResource, c.ns, budgetsBudget), &v1alpha1.BudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BudgetsBudget), err
}

// Update takes the representation of a budgetsBudget and updates it. Returns the server's representation of the budgetsBudget, and an error, if there is any.
func (c *FakeBudgetsBudgets) Update(ctx context.Context, budgetsBudget *v1alpha1.BudgetsBudget, opts v1.UpdateOptions) (result *v1alpha1.BudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(budgetsbudgetsResource, c.ns, budgetsBudget), &v1alpha1.BudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BudgetsBudget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBudgetsBudgets) UpdateStatus(ctx context.Context, budgetsBudget *v1alpha1.BudgetsBudget, opts v1.UpdateOptions) (*v1alpha1.BudgetsBudget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(budgetsbudgetsResource, "status", c.ns, budgetsBudget), &v1alpha1.BudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BudgetsBudget), err
}

// Delete takes name of the budgetsBudget and deletes it. Returns an error if one occurs.
func (c *FakeBudgetsBudgets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(budgetsbudgetsResource, c.ns, name), &v1alpha1.BudgetsBudget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBudgetsBudgets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(budgetsbudgetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BudgetsBudgetList{})
	return err
}

// Patch applies the patch and returns the patched budgetsBudget.
func (c *FakeBudgetsBudgets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BudgetsBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(budgetsbudgetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BudgetsBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BudgetsBudget), err
}
