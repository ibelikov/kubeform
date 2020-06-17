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

// ComputeAddressLister helps list ComputeAddresses.
type ComputeAddressLister interface {
	// List lists all ComputeAddresses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeAddress, err error)
	// ComputeAddresses returns an object that can list and get ComputeAddresses.
	ComputeAddresses(namespace string) ComputeAddressNamespaceLister
	ComputeAddressListerExpansion
}

// computeAddressLister implements the ComputeAddressLister interface.
type computeAddressLister struct {
	indexer cache.Indexer
}

// NewComputeAddressLister returns a new ComputeAddressLister.
func NewComputeAddressLister(indexer cache.Indexer) ComputeAddressLister {
	return &computeAddressLister{indexer: indexer}
}

// List lists all ComputeAddresses in the indexer.
func (s *computeAddressLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeAddress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeAddress))
	})
	return ret, err
}

// ComputeAddresses returns an object that can list and get ComputeAddresses.
func (s *computeAddressLister) ComputeAddresses(namespace string) ComputeAddressNamespaceLister {
	return computeAddressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeAddressNamespaceLister helps list and get ComputeAddresses.
type ComputeAddressNamespaceLister interface {
	// List lists all ComputeAddresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeAddress, err error)
	// Get retrieves the ComputeAddress from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeAddress, error)
	ComputeAddressNamespaceListerExpansion
}

// computeAddressNamespaceLister implements the ComputeAddressNamespaceLister
// interface.
type computeAddressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeAddresses in the indexer for a given namespace.
func (s computeAddressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeAddress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeAddress))
	})
	return ret, err
}

// Get retrieves the ComputeAddress from the indexer for a given namespace and name.
func (s computeAddressNamespaceLister) Get(name string) (*v1alpha1.ComputeAddress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeaddress"), name)
	}
	return obj.(*v1alpha1.ComputeAddress), nil
}
