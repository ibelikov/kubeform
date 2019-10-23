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

// EmrClusterLister helps list EmrClusters.
type EmrClusterLister interface {
	// List lists all EmrClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EmrCluster, err error)
	// EmrClusters returns an object that can list and get EmrClusters.
	EmrClusters(namespace string) EmrClusterNamespaceLister
	EmrClusterListerExpansion
}

// emrClusterLister implements the EmrClusterLister interface.
type emrClusterLister struct {
	indexer cache.Indexer
}

// NewEmrClusterLister returns a new EmrClusterLister.
func NewEmrClusterLister(indexer cache.Indexer) EmrClusterLister {
	return &emrClusterLister{indexer: indexer}
}

// List lists all EmrClusters in the indexer.
func (s *emrClusterLister) List(selector labels.Selector) (ret []*v1alpha1.EmrCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EmrCluster))
	})
	return ret, err
}

// EmrClusters returns an object that can list and get EmrClusters.
func (s *emrClusterLister) EmrClusters(namespace string) EmrClusterNamespaceLister {
	return emrClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EmrClusterNamespaceLister helps list and get EmrClusters.
type EmrClusterNamespaceLister interface {
	// List lists all EmrClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EmrCluster, err error)
	// Get retrieves the EmrCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EmrCluster, error)
	EmrClusterNamespaceListerExpansion
}

// emrClusterNamespaceLister implements the EmrClusterNamespaceLister
// interface.
type emrClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EmrClusters in the indexer for a given namespace.
func (s emrClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EmrCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EmrCluster))
	})
	return ret, err
}

// Get retrieves the EmrCluster from the indexer for a given namespace and name.
func (s emrClusterNamespaceLister) Get(name string) (*v1alpha1.EmrCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("emrcluster"), name)
	}
	return obj.(*v1alpha1.EmrCluster), nil
}
