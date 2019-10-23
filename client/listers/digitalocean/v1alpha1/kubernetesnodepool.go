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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubernetesNodePoolLister helps list KubernetesNodePools.
type KubernetesNodePoolLister interface {
	// List lists all KubernetesNodePools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KubernetesNodePool, err error)
	// KubernetesNodePools returns an object that can list and get KubernetesNodePools.
	KubernetesNodePools(namespace string) KubernetesNodePoolNamespaceLister
	KubernetesNodePoolListerExpansion
}

// kubernetesNodePoolLister implements the KubernetesNodePoolLister interface.
type kubernetesNodePoolLister struct {
	indexer cache.Indexer
}

// NewKubernetesNodePoolLister returns a new KubernetesNodePoolLister.
func NewKubernetesNodePoolLister(indexer cache.Indexer) KubernetesNodePoolLister {
	return &kubernetesNodePoolLister{indexer: indexer}
}

// List lists all KubernetesNodePools in the indexer.
func (s *kubernetesNodePoolLister) List(selector labels.Selector) (ret []*v1alpha1.KubernetesNodePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubernetesNodePool))
	})
	return ret, err
}

// KubernetesNodePools returns an object that can list and get KubernetesNodePools.
func (s *kubernetesNodePoolLister) KubernetesNodePools(namespace string) KubernetesNodePoolNamespaceLister {
	return kubernetesNodePoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubernetesNodePoolNamespaceLister helps list and get KubernetesNodePools.
type KubernetesNodePoolNamespaceLister interface {
	// List lists all KubernetesNodePools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KubernetesNodePool, err error)
	// Get retrieves the KubernetesNodePool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KubernetesNodePool, error)
	KubernetesNodePoolNamespaceListerExpansion
}

// kubernetesNodePoolNamespaceLister implements the KubernetesNodePoolNamespaceLister
// interface.
type kubernetesNodePoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KubernetesNodePools in the indexer for a given namespace.
func (s kubernetesNodePoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KubernetesNodePool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubernetesNodePool))
	})
	return ret, err
}

// Get retrieves the KubernetesNodePool from the indexer for a given namespace and name.
func (s kubernetesNodePoolNamespaceLister) Get(name string) (*v1alpha1.KubernetesNodePool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubernetesnodepool"), name)
	}
	return obj.(*v1alpha1.KubernetesNodePool), nil
}
