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

// FakeDataprocJobs implements DataprocJobInterface
type FakeDataprocJobs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var dataprocjobsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "dataprocjobs"}

var dataprocjobsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "DataprocJob"}

// Get takes name of the dataprocJob, and returns the corresponding dataprocJob object, and an error if there is any.
func (c *FakeDataprocJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataprocjobsResource, c.ns, name), &v1alpha1.DataprocJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJob), err
}

// List takes label and field selectors, and returns the list of DataprocJobs that match those selectors.
func (c *FakeDataprocJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataprocJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataprocjobsResource, dataprocjobsKind, c.ns, opts), &v1alpha1.DataprocJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataprocJobList{ListMeta: obj.(*v1alpha1.DataprocJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataprocJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataprocJobs.
func (c *FakeDataprocJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataprocjobsResource, c.ns, opts))

}

// Create takes the representation of a dataprocJob and creates it.  Returns the server's representation of the dataprocJob, and an error, if there is any.
func (c *FakeDataprocJobs) Create(ctx context.Context, dataprocJob *v1alpha1.DataprocJob, opts v1.CreateOptions) (result *v1alpha1.DataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataprocjobsResource, c.ns, dataprocJob), &v1alpha1.DataprocJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJob), err
}

// Update takes the representation of a dataprocJob and updates it. Returns the server's representation of the dataprocJob, and an error, if there is any.
func (c *FakeDataprocJobs) Update(ctx context.Context, dataprocJob *v1alpha1.DataprocJob, opts v1.UpdateOptions) (result *v1alpha1.DataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataprocjobsResource, c.ns, dataprocJob), &v1alpha1.DataprocJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataprocJobs) UpdateStatus(ctx context.Context, dataprocJob *v1alpha1.DataprocJob, opts v1.UpdateOptions) (*v1alpha1.DataprocJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataprocjobsResource, "status", c.ns, dataprocJob), &v1alpha1.DataprocJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJob), err
}

// Delete takes name of the dataprocJob and deletes it. Returns an error if one occurs.
func (c *FakeDataprocJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataprocjobsResource, c.ns, name), &v1alpha1.DataprocJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataprocJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataprocjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataprocJobList{})
	return err
}

// Patch applies the patch and returns the patched dataprocJob.
func (c *FakeDataprocJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataprocJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataprocjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataprocJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocJob), err
}
