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

// HdinsightInteractiveQueryClusterLister helps list HdinsightInteractiveQueryClusters.
type HdinsightInteractiveQueryClusterLister interface {
	// List lists all HdinsightInteractiveQueryClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightInteractiveQueryCluster, err error)
	// HdinsightInteractiveQueryClusters returns an object that can list and get HdinsightInteractiveQueryClusters.
	HdinsightInteractiveQueryClusters(namespace string) HdinsightInteractiveQueryClusterNamespaceLister
	HdinsightInteractiveQueryClusterListerExpansion
}

// hdinsightInteractiveQueryClusterLister implements the HdinsightInteractiveQueryClusterLister interface.
type hdinsightInteractiveQueryClusterLister struct {
	indexer cache.Indexer
}

// NewHdinsightInteractiveQueryClusterLister returns a new HdinsightInteractiveQueryClusterLister.
func NewHdinsightInteractiveQueryClusterLister(indexer cache.Indexer) HdinsightInteractiveQueryClusterLister {
	return &hdinsightInteractiveQueryClusterLister{indexer: indexer}
}

// List lists all HdinsightInteractiveQueryClusters in the indexer.
func (s *hdinsightInteractiveQueryClusterLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightInteractiveQueryCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightInteractiveQueryCluster))
	})
	return ret, err
}

// HdinsightInteractiveQueryClusters returns an object that can list and get HdinsightInteractiveQueryClusters.
func (s *hdinsightInteractiveQueryClusterLister) HdinsightInteractiveQueryClusters(namespace string) HdinsightInteractiveQueryClusterNamespaceLister {
	return hdinsightInteractiveQueryClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HdinsightInteractiveQueryClusterNamespaceLister helps list and get HdinsightInteractiveQueryClusters.
type HdinsightInteractiveQueryClusterNamespaceLister interface {
	// List lists all HdinsightInteractiveQueryClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightInteractiveQueryCluster, err error)
	// Get retrieves the HdinsightInteractiveQueryCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HdinsightInteractiveQueryCluster, error)
	HdinsightInteractiveQueryClusterNamespaceListerExpansion
}

// hdinsightInteractiveQueryClusterNamespaceLister implements the HdinsightInteractiveQueryClusterNamespaceLister
// interface.
type hdinsightInteractiveQueryClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HdinsightInteractiveQueryClusters in the indexer for a given namespace.
func (s hdinsightInteractiveQueryClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightInteractiveQueryCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightInteractiveQueryCluster))
	})
	return ret, err
}

// Get retrieves the HdinsightInteractiveQueryCluster from the indexer for a given namespace and name.
func (s hdinsightInteractiveQueryClusterNamespaceLister) Get(name string) (*v1alpha1.HdinsightInteractiveQueryCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hdinsightinteractivequerycluster"), name)
	}
	return obj.(*v1alpha1.HdinsightInteractiveQueryCluster), nil
}
