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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// CloudwatchDashboardLister helps list CloudwatchDashboards.
type CloudwatchDashboardLister interface {
	// List lists all CloudwatchDashboards in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchDashboard, err error)
	// CloudwatchDashboards returns an object that can list and get CloudwatchDashboards.
	CloudwatchDashboards(namespace string) CloudwatchDashboardNamespaceLister
	CloudwatchDashboardListerExpansion
}

// cloudwatchDashboardLister implements the CloudwatchDashboardLister interface.
type cloudwatchDashboardLister struct {
	indexer cache.Indexer
}

// NewCloudwatchDashboardLister returns a new CloudwatchDashboardLister.
func NewCloudwatchDashboardLister(indexer cache.Indexer) CloudwatchDashboardLister {
	return &cloudwatchDashboardLister{indexer: indexer}
}

// List lists all CloudwatchDashboards in the indexer.
func (s *cloudwatchDashboardLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchDashboard, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchDashboard))
	})
	return ret, err
}

// CloudwatchDashboards returns an object that can list and get CloudwatchDashboards.
func (s *cloudwatchDashboardLister) CloudwatchDashboards(namespace string) CloudwatchDashboardNamespaceLister {
	return cloudwatchDashboardNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudwatchDashboardNamespaceLister helps list and get CloudwatchDashboards.
type CloudwatchDashboardNamespaceLister interface {
	// List lists all CloudwatchDashboards in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchDashboard, err error)
	// Get retrieves the CloudwatchDashboard from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudwatchDashboard, error)
	CloudwatchDashboardNamespaceListerExpansion
}

// cloudwatchDashboardNamespaceLister implements the CloudwatchDashboardNamespaceLister
// interface.
type cloudwatchDashboardNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudwatchDashboards in the indexer for a given namespace.
func (s cloudwatchDashboardNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchDashboard, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchDashboard))
	})
	return ret, err
}

// Get retrieves the CloudwatchDashboard from the indexer for a given namespace and name.
func (s cloudwatchDashboardNamespaceLister) Get(name string) (*v1alpha1.CloudwatchDashboard, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudwatchdashboard"), name)
	}
	return obj.(*v1alpha1.CloudwatchDashboard), nil
}
