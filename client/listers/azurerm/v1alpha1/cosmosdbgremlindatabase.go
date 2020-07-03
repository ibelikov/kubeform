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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CosmosdbGremlinDatabaseLister helps list CosmosdbGremlinDatabases.
type CosmosdbGremlinDatabaseLister interface {
	// List lists all CosmosdbGremlinDatabases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CosmosdbGremlinDatabase, err error)
	// CosmosdbGremlinDatabases returns an object that can list and get CosmosdbGremlinDatabases.
	CosmosdbGremlinDatabases(namespace string) CosmosdbGremlinDatabaseNamespaceLister
	CosmosdbGremlinDatabaseListerExpansion
}

// cosmosdbGremlinDatabaseLister implements the CosmosdbGremlinDatabaseLister interface.
type cosmosdbGremlinDatabaseLister struct {
	indexer cache.Indexer
}

// NewCosmosdbGremlinDatabaseLister returns a new CosmosdbGremlinDatabaseLister.
func NewCosmosdbGremlinDatabaseLister(indexer cache.Indexer) CosmosdbGremlinDatabaseLister {
	return &cosmosdbGremlinDatabaseLister{indexer: indexer}
}

// List lists all CosmosdbGremlinDatabases in the indexer.
func (s *cosmosdbGremlinDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.CosmosdbGremlinDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CosmosdbGremlinDatabase))
	})
	return ret, err
}

// CosmosdbGremlinDatabases returns an object that can list and get CosmosdbGremlinDatabases.
func (s *cosmosdbGremlinDatabaseLister) CosmosdbGremlinDatabases(namespace string) CosmosdbGremlinDatabaseNamespaceLister {
	return cosmosdbGremlinDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CosmosdbGremlinDatabaseNamespaceLister helps list and get CosmosdbGremlinDatabases.
type CosmosdbGremlinDatabaseNamespaceLister interface {
	// List lists all CosmosdbGremlinDatabases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CosmosdbGremlinDatabase, err error)
	// Get retrieves the CosmosdbGremlinDatabase from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CosmosdbGremlinDatabase, error)
	CosmosdbGremlinDatabaseNamespaceListerExpansion
}

// cosmosdbGremlinDatabaseNamespaceLister implements the CosmosdbGremlinDatabaseNamespaceLister
// interface.
type cosmosdbGremlinDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CosmosdbGremlinDatabases in the indexer for a given namespace.
func (s cosmosdbGremlinDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CosmosdbGremlinDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CosmosdbGremlinDatabase))
	})
	return ret, err
}

// Get retrieves the CosmosdbGremlinDatabase from the indexer for a given namespace and name.
func (s cosmosdbGremlinDatabaseNamespaceLister) Get(name string) (*v1alpha1.CosmosdbGremlinDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cosmosdbgremlindatabase"), name)
	}
	return obj.(*v1alpha1.CosmosdbGremlinDatabase), nil
}