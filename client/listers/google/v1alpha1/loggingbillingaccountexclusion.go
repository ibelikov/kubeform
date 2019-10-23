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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LoggingBillingAccountExclusionLister helps list LoggingBillingAccountExclusions.
type LoggingBillingAccountExclusionLister interface {
	// List lists all LoggingBillingAccountExclusions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountExclusion, err error)
	// LoggingBillingAccountExclusions returns an object that can list and get LoggingBillingAccountExclusions.
	LoggingBillingAccountExclusions(namespace string) LoggingBillingAccountExclusionNamespaceLister
	LoggingBillingAccountExclusionListerExpansion
}

// loggingBillingAccountExclusionLister implements the LoggingBillingAccountExclusionLister interface.
type loggingBillingAccountExclusionLister struct {
	indexer cache.Indexer
}

// NewLoggingBillingAccountExclusionLister returns a new LoggingBillingAccountExclusionLister.
func NewLoggingBillingAccountExclusionLister(indexer cache.Indexer) LoggingBillingAccountExclusionLister {
	return &loggingBillingAccountExclusionLister{indexer: indexer}
}

// List lists all LoggingBillingAccountExclusions in the indexer.
func (s *loggingBillingAccountExclusionLister) List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountExclusion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoggingBillingAccountExclusion))
	})
	return ret, err
}

// LoggingBillingAccountExclusions returns an object that can list and get LoggingBillingAccountExclusions.
func (s *loggingBillingAccountExclusionLister) LoggingBillingAccountExclusions(namespace string) LoggingBillingAccountExclusionNamespaceLister {
	return loggingBillingAccountExclusionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LoggingBillingAccountExclusionNamespaceLister helps list and get LoggingBillingAccountExclusions.
type LoggingBillingAccountExclusionNamespaceLister interface {
	// List lists all LoggingBillingAccountExclusions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountExclusion, err error)
	// Get retrieves the LoggingBillingAccountExclusion from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LoggingBillingAccountExclusion, error)
	LoggingBillingAccountExclusionNamespaceListerExpansion
}

// loggingBillingAccountExclusionNamespaceLister implements the LoggingBillingAccountExclusionNamespaceLister
// interface.
type loggingBillingAccountExclusionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LoggingBillingAccountExclusions in the indexer for a given namespace.
func (s loggingBillingAccountExclusionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LoggingBillingAccountExclusion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoggingBillingAccountExclusion))
	})
	return ret, err
}

// Get retrieves the LoggingBillingAccountExclusion from the indexer for a given namespace and name.
func (s loggingBillingAccountExclusionNamespaceLister) Get(name string) (*v1alpha1.LoggingBillingAccountExclusion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("loggingbillingaccountexclusion"), name)
	}
	return obj.(*v1alpha1.LoggingBillingAccountExclusion), nil
}
