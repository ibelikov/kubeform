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

// FakeDatasyncLocationS3s implements DatasyncLocationS3Interface
type FakeDatasyncLocationS3s struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var datasynclocations3sResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "datasynclocations3s"}

var datasynclocations3sKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DatasyncLocationS3"}

// Get takes name of the datasyncLocationS3, and returns the corresponding datasyncLocationS3 object, and an error if there is any.
func (c *FakeDatasyncLocationS3s) Get(name string, options v1.GetOptions) (result *v1alpha1.DatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasynclocations3sResource, c.ns, name), &v1alpha1.DatasyncLocationS3{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationS3), err
}

// List takes label and field selectors, and returns the list of DatasyncLocationS3s that match those selectors.
func (c *FakeDatasyncLocationS3s) List(opts v1.ListOptions) (result *v1alpha1.DatasyncLocationS3List, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasynclocations3sResource, datasynclocations3sKind, c.ns, opts), &v1alpha1.DatasyncLocationS3List{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatasyncLocationS3List{ListMeta: obj.(*v1alpha1.DatasyncLocationS3List).ListMeta}
	for _, item := range obj.(*v1alpha1.DatasyncLocationS3List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasyncLocationS3s.
func (c *FakeDatasyncLocationS3s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasynclocations3sResource, c.ns, opts))

}

// Create takes the representation of a datasyncLocationS3 and creates it.  Returns the server's representation of the datasyncLocationS3, and an error, if there is any.
func (c *FakeDatasyncLocationS3s) Create(datasyncLocationS3 *v1alpha1.DatasyncLocationS3) (result *v1alpha1.DatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasynclocations3sResource, c.ns, datasyncLocationS3), &v1alpha1.DatasyncLocationS3{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationS3), err
}

// Update takes the representation of a datasyncLocationS3 and updates it. Returns the server's representation of the datasyncLocationS3, and an error, if there is any.
func (c *FakeDatasyncLocationS3s) Update(datasyncLocationS3 *v1alpha1.DatasyncLocationS3) (result *v1alpha1.DatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasynclocations3sResource, c.ns, datasyncLocationS3), &v1alpha1.DatasyncLocationS3{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationS3), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatasyncLocationS3s) UpdateStatus(datasyncLocationS3 *v1alpha1.DatasyncLocationS3) (*v1alpha1.DatasyncLocationS3, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datasynclocations3sResource, "status", c.ns, datasyncLocationS3), &v1alpha1.DatasyncLocationS3{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationS3), err
}

// Delete takes name of the datasyncLocationS3 and deletes it. Returns an error if one occurs.
func (c *FakeDatasyncLocationS3s) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasynclocations3sResource, c.ns, name), &v1alpha1.DatasyncLocationS3{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasyncLocationS3s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasynclocations3sResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatasyncLocationS3List{})
	return err
}

// Patch applies the patch and returns the patched datasyncLocationS3.
func (c *FakeDatasyncLocationS3s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasynclocations3sResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatasyncLocationS3{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationS3), err
}
