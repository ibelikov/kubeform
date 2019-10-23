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

// BatchPoolLister helps list BatchPools.
type BatchPoolLister interface {
	// List lists all BatchPools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BatchPool, err error)
	// BatchPools returns an object that can list and get BatchPools.
	BatchPools(namespace string) BatchPoolNamespaceLister
	BatchPoolListerExpansion
}

// batchPoolLister implements the BatchPoolLister interface.
type batchPoolLister struct {
	indexer cache.Indexer
}

// NewBatchPoolLister returns a new BatchPoolLister.
func NewBatchPoolLister(indexer cache.Indexer) BatchPoolLister {
	return &batchPoolLister{indexer: indexer}
}

// List lists all BatchPools in the indexer.
func (s *batchPoolLister) List(selector labels.Selector) (ret []*v1alpha1.BatchPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BatchPool))
	})
	return ret, err
}

// BatchPools returns an object that can list and get BatchPools.
func (s *batchPoolLister) BatchPools(namespace string) BatchPoolNamespaceLister {
	return batchPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BatchPoolNamespaceLister helps list and get BatchPools.
type BatchPoolNamespaceLister interface {
	// List lists all BatchPools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BatchPool, err error)
	// Get retrieves the BatchPool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BatchPool, error)
	BatchPoolNamespaceListerExpansion
}

// batchPoolNamespaceLister implements the BatchPoolNamespaceLister
// interface.
type batchPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BatchPools in the indexer for a given namespace.
func (s batchPoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BatchPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BatchPool))
	})
	return ret, err
}

// Get retrieves the BatchPool from the indexer for a given namespace and name.
func (s batchPoolNamespaceLister) Get(name string) (*v1alpha1.BatchPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("batchpool"), name)
	}
	return obj.(*v1alpha1.BatchPool), nil
}
