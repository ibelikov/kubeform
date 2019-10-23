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

// DataFactoryLinkedServiceDataLakeStorageGen2Lister helps list DataFactoryLinkedServiceDataLakeStorageGen2s.
type DataFactoryLinkedServiceDataLakeStorageGen2Lister interface {
	// List lists all DataFactoryLinkedServiceDataLakeStorageGen2s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2, err error)
	// DataFactoryLinkedServiceDataLakeStorageGen2s returns an object that can list and get DataFactoryLinkedServiceDataLakeStorageGen2s.
	DataFactoryLinkedServiceDataLakeStorageGen2s(namespace string) DataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister
	DataFactoryLinkedServiceDataLakeStorageGen2ListerExpansion
}

// dataFactoryLinkedServiceDataLakeStorageGen2Lister implements the DataFactoryLinkedServiceDataLakeStorageGen2Lister interface.
type dataFactoryLinkedServiceDataLakeStorageGen2Lister struct {
	indexer cache.Indexer
}

// NewDataFactoryLinkedServiceDataLakeStorageGen2Lister returns a new DataFactoryLinkedServiceDataLakeStorageGen2Lister.
func NewDataFactoryLinkedServiceDataLakeStorageGen2Lister(indexer cache.Indexer) DataFactoryLinkedServiceDataLakeStorageGen2Lister {
	return &dataFactoryLinkedServiceDataLakeStorageGen2Lister{indexer: indexer}
}

// List lists all DataFactoryLinkedServiceDataLakeStorageGen2s in the indexer.
func (s *dataFactoryLinkedServiceDataLakeStorageGen2Lister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2))
	})
	return ret, err
}

// DataFactoryLinkedServiceDataLakeStorageGen2s returns an object that can list and get DataFactoryLinkedServiceDataLakeStorageGen2s.
func (s *dataFactoryLinkedServiceDataLakeStorageGen2Lister) DataFactoryLinkedServiceDataLakeStorageGen2s(namespace string) DataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister {
	return dataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister helps list and get DataFactoryLinkedServiceDataLakeStorageGen2s.
type DataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister interface {
	// List lists all DataFactoryLinkedServiceDataLakeStorageGen2s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2, err error)
	// Get retrieves the DataFactoryLinkedServiceDataLakeStorageGen2 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2, error)
	DataFactoryLinkedServiceDataLakeStorageGen2NamespaceListerExpansion
}

// dataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister implements the DataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister
// interface.
type dataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataFactoryLinkedServiceDataLakeStorageGen2s in the indexer for a given namespace.
func (s dataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2))
	})
	return ret, err
}

// Get retrieves the DataFactoryLinkedServiceDataLakeStorageGen2 from the indexer for a given namespace and name.
func (s dataFactoryLinkedServiceDataLakeStorageGen2NamespaceLister) Get(name string) (*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datafactorylinkedservicedatalakestoragegen2"), name)
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceDataLakeStorageGen2), nil
}
