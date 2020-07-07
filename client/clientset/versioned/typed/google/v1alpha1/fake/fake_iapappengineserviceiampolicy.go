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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeIapAppEngineServiceIamPolicies implements IapAppEngineServiceIamPolicyInterface
type FakeIapAppEngineServiceIamPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var iapappengineserviceiampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "iapappengineserviceiampolicies"}

var iapappengineserviceiampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "IapAppEngineServiceIamPolicy"}

// Get takes name of the iapAppEngineServiceIamPolicy, and returns the corresponding iapAppEngineServiceIamPolicy object, and an error if there is any.
func (c *FakeIapAppEngineServiceIamPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IapAppEngineServiceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iapappengineserviceiampoliciesResource, c.ns, name), &v1alpha1.IapAppEngineServiceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamPolicy), err
}

// List takes label and field selectors, and returns the list of IapAppEngineServiceIamPolicies that match those selectors.
func (c *FakeIapAppEngineServiceIamPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IapAppEngineServiceIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iapappengineserviceiampoliciesResource, iapappengineserviceiampoliciesKind, c.ns, opts), &v1alpha1.IapAppEngineServiceIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IapAppEngineServiceIamPolicyList{ListMeta: obj.(*v1alpha1.IapAppEngineServiceIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IapAppEngineServiceIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iapAppEngineServiceIamPolicies.
func (c *FakeIapAppEngineServiceIamPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iapappengineserviceiampoliciesResource, c.ns, opts))

}

// Create takes the representation of a iapAppEngineServiceIamPolicy and creates it.  Returns the server's representation of the iapAppEngineServiceIamPolicy, and an error, if there is any.
func (c *FakeIapAppEngineServiceIamPolicies) Create(ctx context.Context, iapAppEngineServiceIamPolicy *v1alpha1.IapAppEngineServiceIamPolicy, opts v1.CreateOptions) (result *v1alpha1.IapAppEngineServiceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iapappengineserviceiampoliciesResource, c.ns, iapAppEngineServiceIamPolicy), &v1alpha1.IapAppEngineServiceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamPolicy), err
}

// Update takes the representation of a iapAppEngineServiceIamPolicy and updates it. Returns the server's representation of the iapAppEngineServiceIamPolicy, and an error, if there is any.
func (c *FakeIapAppEngineServiceIamPolicies) Update(ctx context.Context, iapAppEngineServiceIamPolicy *v1alpha1.IapAppEngineServiceIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.IapAppEngineServiceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iapappengineserviceiampoliciesResource, c.ns, iapAppEngineServiceIamPolicy), &v1alpha1.IapAppEngineServiceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIapAppEngineServiceIamPolicies) UpdateStatus(ctx context.Context, iapAppEngineServiceIamPolicy *v1alpha1.IapAppEngineServiceIamPolicy, opts v1.UpdateOptions) (*v1alpha1.IapAppEngineServiceIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iapappengineserviceiampoliciesResource, "status", c.ns, iapAppEngineServiceIamPolicy), &v1alpha1.IapAppEngineServiceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamPolicy), err
}

// Delete takes name of the iapAppEngineServiceIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIapAppEngineServiceIamPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iapappengineserviceiampoliciesResource, c.ns, name), &v1alpha1.IapAppEngineServiceIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIapAppEngineServiceIamPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iapappengineserviceiampoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IapAppEngineServiceIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched iapAppEngineServiceIamPolicy.
func (c *FakeIapAppEngineServiceIamPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IapAppEngineServiceIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iapappengineserviceiampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IapAppEngineServiceIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamPolicy), err
}
