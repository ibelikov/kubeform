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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SagemakerNotebookInstanceLifecycleConfigurationLister helps list SagemakerNotebookInstanceLifecycleConfigurations.
type SagemakerNotebookInstanceLifecycleConfigurationLister interface {
	// List lists all SagemakerNotebookInstanceLifecycleConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error)
	// SagemakerNotebookInstanceLifecycleConfigurations returns an object that can list and get SagemakerNotebookInstanceLifecycleConfigurations.
	SagemakerNotebookInstanceLifecycleConfigurations(namespace string) SagemakerNotebookInstanceLifecycleConfigurationNamespaceLister
	SagemakerNotebookInstanceLifecycleConfigurationListerExpansion
}

// sagemakerNotebookInstanceLifecycleConfigurationLister implements the SagemakerNotebookInstanceLifecycleConfigurationLister interface.
type sagemakerNotebookInstanceLifecycleConfigurationLister struct {
	indexer cache.Indexer
}

// NewSagemakerNotebookInstanceLifecycleConfigurationLister returns a new SagemakerNotebookInstanceLifecycleConfigurationLister.
func NewSagemakerNotebookInstanceLifecycleConfigurationLister(indexer cache.Indexer) SagemakerNotebookInstanceLifecycleConfigurationLister {
	return &sagemakerNotebookInstanceLifecycleConfigurationLister{indexer: indexer}
}

// List lists all SagemakerNotebookInstanceLifecycleConfigurations in the indexer.
func (s *sagemakerNotebookInstanceLifecycleConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration))
	})
	return ret, err
}

// SagemakerNotebookInstanceLifecycleConfigurations returns an object that can list and get SagemakerNotebookInstanceLifecycleConfigurations.
func (s *sagemakerNotebookInstanceLifecycleConfigurationLister) SagemakerNotebookInstanceLifecycleConfigurations(namespace string) SagemakerNotebookInstanceLifecycleConfigurationNamespaceLister {
	return sagemakerNotebookInstanceLifecycleConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SagemakerNotebookInstanceLifecycleConfigurationNamespaceLister helps list and get SagemakerNotebookInstanceLifecycleConfigurations.
type SagemakerNotebookInstanceLifecycleConfigurationNamespaceLister interface {
	// List lists all SagemakerNotebookInstanceLifecycleConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error)
	// Get retrieves the SagemakerNotebookInstanceLifecycleConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, error)
	SagemakerNotebookInstanceLifecycleConfigurationNamespaceListerExpansion
}

// sagemakerNotebookInstanceLifecycleConfigurationNamespaceLister implements the SagemakerNotebookInstanceLifecycleConfigurationNamespaceLister
// interface.
type sagemakerNotebookInstanceLifecycleConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SagemakerNotebookInstanceLifecycleConfigurations in the indexer for a given namespace.
func (s sagemakerNotebookInstanceLifecycleConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration))
	})
	return ret, err
}

// Get retrieves the SagemakerNotebookInstanceLifecycleConfiguration from the indexer for a given namespace and name.
func (s sagemakerNotebookInstanceLifecycleConfigurationNamespaceLister) Get(name string) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sagemakernotebookinstancelifecycleconfiguration"), name)
	}
	return obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration), nil
}
