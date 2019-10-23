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

// RuntimeconfigVariableLister helps list RuntimeconfigVariables.
type RuntimeconfigVariableLister interface {
	// List lists all RuntimeconfigVariables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigVariable, err error)
	// RuntimeconfigVariables returns an object that can list and get RuntimeconfigVariables.
	RuntimeconfigVariables(namespace string) RuntimeconfigVariableNamespaceLister
	RuntimeconfigVariableListerExpansion
}

// runtimeconfigVariableLister implements the RuntimeconfigVariableLister interface.
type runtimeconfigVariableLister struct {
	indexer cache.Indexer
}

// NewRuntimeconfigVariableLister returns a new RuntimeconfigVariableLister.
func NewRuntimeconfigVariableLister(indexer cache.Indexer) RuntimeconfigVariableLister {
	return &runtimeconfigVariableLister{indexer: indexer}
}

// List lists all RuntimeconfigVariables in the indexer.
func (s *runtimeconfigVariableLister) List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigVariable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuntimeconfigVariable))
	})
	return ret, err
}

// RuntimeconfigVariables returns an object that can list and get RuntimeconfigVariables.
func (s *runtimeconfigVariableLister) RuntimeconfigVariables(namespace string) RuntimeconfigVariableNamespaceLister {
	return runtimeconfigVariableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RuntimeconfigVariableNamespaceLister helps list and get RuntimeconfigVariables.
type RuntimeconfigVariableNamespaceLister interface {
	// List lists all RuntimeconfigVariables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigVariable, err error)
	// Get retrieves the RuntimeconfigVariable from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RuntimeconfigVariable, error)
	RuntimeconfigVariableNamespaceListerExpansion
}

// runtimeconfigVariableNamespaceLister implements the RuntimeconfigVariableNamespaceLister
// interface.
type runtimeconfigVariableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RuntimeconfigVariables in the indexer for a given namespace.
func (s runtimeconfigVariableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigVariable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuntimeconfigVariable))
	})
	return ret, err
}

// Get retrieves the RuntimeconfigVariable from the indexer for a given namespace and name.
func (s runtimeconfigVariableNamespaceLister) Get(name string) (*v1alpha1.RuntimeconfigVariable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("runtimeconfigvariable"), name)
	}
	return obj.(*v1alpha1.RuntimeconfigVariable), nil
}
