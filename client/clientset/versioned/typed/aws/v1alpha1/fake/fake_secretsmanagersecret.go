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

// FakeSecretsmanagerSecrets implements SecretsmanagerSecretInterface
type FakeSecretsmanagerSecrets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var secretsmanagersecretsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "secretsmanagersecrets"}

var secretsmanagersecretsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SecretsmanagerSecret"}

// Get takes name of the secretsmanagerSecret, and returns the corresponding secretsmanagerSecret object, and an error if there is any.
func (c *FakeSecretsmanagerSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.SecretsmanagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretsmanagersecretsResource, c.ns, name), &v1alpha1.SecretsmanagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretsmanagerSecret), err
}

// List takes label and field selectors, and returns the list of SecretsmanagerSecrets that match those selectors.
func (c *FakeSecretsmanagerSecrets) List(opts v1.ListOptions) (result *v1alpha1.SecretsmanagerSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretsmanagersecretsResource, secretsmanagersecretsKind, c.ns, opts), &v1alpha1.SecretsmanagerSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecretsmanagerSecretList{ListMeta: obj.(*v1alpha1.SecretsmanagerSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecretsmanagerSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretsmanagerSecrets.
func (c *FakeSecretsmanagerSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretsmanagersecretsResource, c.ns, opts))

}

// Create takes the representation of a secretsmanagerSecret and creates it.  Returns the server's representation of the secretsmanagerSecret, and an error, if there is any.
func (c *FakeSecretsmanagerSecrets) Create(secretsmanagerSecret *v1alpha1.SecretsmanagerSecret) (result *v1alpha1.SecretsmanagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretsmanagersecretsResource, c.ns, secretsmanagerSecret), &v1alpha1.SecretsmanagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretsmanagerSecret), err
}

// Update takes the representation of a secretsmanagerSecret and updates it. Returns the server's representation of the secretsmanagerSecret, and an error, if there is any.
func (c *FakeSecretsmanagerSecrets) Update(secretsmanagerSecret *v1alpha1.SecretsmanagerSecret) (result *v1alpha1.SecretsmanagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretsmanagersecretsResource, c.ns, secretsmanagerSecret), &v1alpha1.SecretsmanagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretsmanagerSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecretsmanagerSecrets) UpdateStatus(secretsmanagerSecret *v1alpha1.SecretsmanagerSecret) (*v1alpha1.SecretsmanagerSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretsmanagersecretsResource, "status", c.ns, secretsmanagerSecret), &v1alpha1.SecretsmanagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretsmanagerSecret), err
}

// Delete takes name of the secretsmanagerSecret and deletes it. Returns an error if one occurs.
func (c *FakeSecretsmanagerSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretsmanagersecretsResource, c.ns, name), &v1alpha1.SecretsmanagerSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretsmanagerSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretsmanagersecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecretsmanagerSecretList{})
	return err
}

// Patch applies the patch and returns the patched secretsmanagerSecret.
func (c *FakeSecretsmanagerSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretsmanagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretsmanagersecretsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecretsmanagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretsmanagerSecret), err
}
