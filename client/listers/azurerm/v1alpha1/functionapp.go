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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FunctionAppLister helps list FunctionApps.
type FunctionAppLister interface {
	// List lists all FunctionApps in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FunctionApp, err error)
	// FunctionApps returns an object that can list and get FunctionApps.
	FunctionApps(namespace string) FunctionAppNamespaceLister
	FunctionAppListerExpansion
}

// functionAppLister implements the FunctionAppLister interface.
type functionAppLister struct {
	indexer cache.Indexer
}

// NewFunctionAppLister returns a new FunctionAppLister.
func NewFunctionAppLister(indexer cache.Indexer) FunctionAppLister {
	return &functionAppLister{indexer: indexer}
}

// List lists all FunctionApps in the indexer.
func (s *functionAppLister) List(selector labels.Selector) (ret []*v1alpha1.FunctionApp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FunctionApp))
	})
	return ret, err
}

// FunctionApps returns an object that can list and get FunctionApps.
func (s *functionAppLister) FunctionApps(namespace string) FunctionAppNamespaceLister {
	return functionAppNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FunctionAppNamespaceLister helps list and get FunctionApps.
type FunctionAppNamespaceLister interface {
	// List lists all FunctionApps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FunctionApp, err error)
	// Get retrieves the FunctionApp from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FunctionApp, error)
	FunctionAppNamespaceListerExpansion
}

// functionAppNamespaceLister implements the FunctionAppNamespaceLister
// interface.
type functionAppNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FunctionApps in the indexer for a given namespace.
func (s functionAppNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FunctionApp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FunctionApp))
	})
	return ret, err
}

// Get retrieves the FunctionApp from the indexer for a given namespace and name.
func (s functionAppNamespaceLister) Get(name string) (*v1alpha1.FunctionApp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("functionapp"), name)
	}
	return obj.(*v1alpha1.FunctionApp), nil
}
