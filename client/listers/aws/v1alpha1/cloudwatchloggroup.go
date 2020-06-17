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

// CloudwatchLogGroupLister helps list CloudwatchLogGroups.
type CloudwatchLogGroupLister interface {
	// List lists all CloudwatchLogGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogGroup, err error)
	// CloudwatchLogGroups returns an object that can list and get CloudwatchLogGroups.
	CloudwatchLogGroups(namespace string) CloudwatchLogGroupNamespaceLister
	CloudwatchLogGroupListerExpansion
}

// cloudwatchLogGroupLister implements the CloudwatchLogGroupLister interface.
type cloudwatchLogGroupLister struct {
	indexer cache.Indexer
}

// NewCloudwatchLogGroupLister returns a new CloudwatchLogGroupLister.
func NewCloudwatchLogGroupLister(indexer cache.Indexer) CloudwatchLogGroupLister {
	return &cloudwatchLogGroupLister{indexer: indexer}
}

// List lists all CloudwatchLogGroups in the indexer.
func (s *cloudwatchLogGroupLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchLogGroup))
	})
	return ret, err
}

// CloudwatchLogGroups returns an object that can list and get CloudwatchLogGroups.
func (s *cloudwatchLogGroupLister) CloudwatchLogGroups(namespace string) CloudwatchLogGroupNamespaceLister {
	return cloudwatchLogGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudwatchLogGroupNamespaceLister helps list and get CloudwatchLogGroups.
type CloudwatchLogGroupNamespaceLister interface {
	// List lists all CloudwatchLogGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogGroup, err error)
	// Get retrieves the CloudwatchLogGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudwatchLogGroup, error)
	CloudwatchLogGroupNamespaceListerExpansion
}

// cloudwatchLogGroupNamespaceLister implements the CloudwatchLogGroupNamespaceLister
// interface.
type cloudwatchLogGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudwatchLogGroups in the indexer for a given namespace.
func (s cloudwatchLogGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchLogGroup))
	})
	return ret, err
}

// Get retrieves the CloudwatchLogGroup from the indexer for a given namespace and name.
func (s cloudwatchLogGroupNamespaceLister) Get(name string) (*v1alpha1.CloudwatchLogGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudwatchloggroup"), name)
	}
	return obj.(*v1alpha1.CloudwatchLogGroup), nil
}
