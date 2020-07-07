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

// ComputeDiskLister helps list ComputeDisks.
type ComputeDiskLister interface {
	// List lists all ComputeDisks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeDisk, err error)
	// ComputeDisks returns an object that can list and get ComputeDisks.
	ComputeDisks(namespace string) ComputeDiskNamespaceLister
	ComputeDiskListerExpansion
}

// computeDiskLister implements the ComputeDiskLister interface.
type computeDiskLister struct {
	indexer cache.Indexer
}

// NewComputeDiskLister returns a new ComputeDiskLister.
func NewComputeDiskLister(indexer cache.Indexer) ComputeDiskLister {
	return &computeDiskLister{indexer: indexer}
}

// List lists all ComputeDisks in the indexer.
func (s *computeDiskLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeDisk, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeDisk))
	})
	return ret, err
}

// ComputeDisks returns an object that can list and get ComputeDisks.
func (s *computeDiskLister) ComputeDisks(namespace string) ComputeDiskNamespaceLister {
	return computeDiskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeDiskNamespaceLister helps list and get ComputeDisks.
type ComputeDiskNamespaceLister interface {
	// List lists all ComputeDisks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeDisk, err error)
	// Get retrieves the ComputeDisk from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeDisk, error)
	ComputeDiskNamespaceListerExpansion
}

// computeDiskNamespaceLister implements the ComputeDiskNamespaceLister
// interface.
type computeDiskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeDisks in the indexer for a given namespace.
func (s computeDiskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeDisk, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeDisk))
	})
	return ret, err
}

// Get retrieves the ComputeDisk from the indexer for a given namespace and name.
func (s computeDiskNamespaceLister) Get(name string) (*v1alpha1.ComputeDisk, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computedisk"), name)
	}
	return obj.(*v1alpha1.ComputeDisk), nil
}
