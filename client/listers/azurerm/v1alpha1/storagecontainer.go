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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// StorageContainerLister helps list StorageContainers.
type StorageContainerLister interface {
	// List lists all StorageContainers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageContainer, err error)
	// StorageContainers returns an object that can list and get StorageContainers.
	StorageContainers(namespace string) StorageContainerNamespaceLister
	StorageContainerListerExpansion
}

// storageContainerLister implements the StorageContainerLister interface.
type storageContainerLister struct {
	indexer cache.Indexer
}

// NewStorageContainerLister returns a new StorageContainerLister.
func NewStorageContainerLister(indexer cache.Indexer) StorageContainerLister {
	return &storageContainerLister{indexer: indexer}
}

// List lists all StorageContainers in the indexer.
func (s *storageContainerLister) List(selector labels.Selector) (ret []*v1alpha1.StorageContainer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageContainer))
	})
	return ret, err
}

// StorageContainers returns an object that can list and get StorageContainers.
func (s *storageContainerLister) StorageContainers(namespace string) StorageContainerNamespaceLister {
	return storageContainerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageContainerNamespaceLister helps list and get StorageContainers.
type StorageContainerNamespaceLister interface {
	// List lists all StorageContainers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageContainer, err error)
	// Get retrieves the StorageContainer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageContainer, error)
	StorageContainerNamespaceListerExpansion
}

// storageContainerNamespaceLister implements the StorageContainerNamespaceLister
// interface.
type storageContainerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageContainers in the indexer for a given namespace.
func (s storageContainerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageContainer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageContainer))
	})
	return ret, err
}

// Get retrieves the StorageContainer from the indexer for a given namespace and name.
func (s storageContainerNamespaceLister) Get(name string) (*v1alpha1.StorageContainer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagecontainer"), name)
	}
	return obj.(*v1alpha1.StorageContainer), nil
}
