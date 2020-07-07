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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// ContainerClusterLister helps list ContainerClusters.
type ContainerClusterLister interface {
	// List lists all ContainerClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerCluster, err error)
	// ContainerClusters returns an object that can list and get ContainerClusters.
	ContainerClusters(namespace string) ContainerClusterNamespaceLister
	ContainerClusterListerExpansion
}

// containerClusterLister implements the ContainerClusterLister interface.
type containerClusterLister struct {
	indexer cache.Indexer
}

// NewContainerClusterLister returns a new ContainerClusterLister.
func NewContainerClusterLister(indexer cache.Indexer) ContainerClusterLister {
	return &containerClusterLister{indexer: indexer}
}

// List lists all ContainerClusters in the indexer.
func (s *containerClusterLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerCluster))
	})
	return ret, err
}

// ContainerClusters returns an object that can list and get ContainerClusters.
func (s *containerClusterLister) ContainerClusters(namespace string) ContainerClusterNamespaceLister {
	return containerClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ContainerClusterNamespaceLister helps list and get ContainerClusters.
type ContainerClusterNamespaceLister interface {
	// List lists all ContainerClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerCluster, err error)
	// Get retrieves the ContainerCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ContainerCluster, error)
	ContainerClusterNamespaceListerExpansion
}

// containerClusterNamespaceLister implements the ContainerClusterNamespaceLister
// interface.
type containerClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ContainerClusters in the indexer for a given namespace.
func (s containerClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerCluster))
	})
	return ret, err
}

// Get retrieves the ContainerCluster from the indexer for a given namespace and name.
func (s containerClusterNamespaceLister) Get(name string) (*v1alpha1.ContainerCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("containercluster"), name)
	}
	return obj.(*v1alpha1.ContainerCluster), nil
}
