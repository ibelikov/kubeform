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

// RedshiftClusterLister helps list RedshiftClusters.
type RedshiftClusterLister interface {
	// List lists all RedshiftClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftCluster, err error)
	// RedshiftClusters returns an object that can list and get RedshiftClusters.
	RedshiftClusters(namespace string) RedshiftClusterNamespaceLister
	RedshiftClusterListerExpansion
}

// redshiftClusterLister implements the RedshiftClusterLister interface.
type redshiftClusterLister struct {
	indexer cache.Indexer
}

// NewRedshiftClusterLister returns a new RedshiftClusterLister.
func NewRedshiftClusterLister(indexer cache.Indexer) RedshiftClusterLister {
	return &redshiftClusterLister{indexer: indexer}
}

// List lists all RedshiftClusters in the indexer.
func (s *redshiftClusterLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftCluster))
	})
	return ret, err
}

// RedshiftClusters returns an object that can list and get RedshiftClusters.
func (s *redshiftClusterLister) RedshiftClusters(namespace string) RedshiftClusterNamespaceLister {
	return redshiftClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedshiftClusterNamespaceLister helps list and get RedshiftClusters.
type RedshiftClusterNamespaceLister interface {
	// List lists all RedshiftClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftCluster, err error)
	// Get retrieves the RedshiftCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RedshiftCluster, error)
	RedshiftClusterNamespaceListerExpansion
}

// redshiftClusterNamespaceLister implements the RedshiftClusterNamespaceLister
// interface.
type redshiftClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedshiftClusters in the indexer for a given namespace.
func (s redshiftClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftCluster))
	})
	return ret, err
}

// Get retrieves the RedshiftCluster from the indexer for a given namespace and name.
func (s redshiftClusterNamespaceLister) Get(name string) (*v1alpha1.RedshiftCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("redshiftcluster"), name)
	}
	return obj.(*v1alpha1.RedshiftCluster), nil
}
