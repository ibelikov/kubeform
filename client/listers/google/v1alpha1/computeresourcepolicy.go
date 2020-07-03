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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComputeResourcePolicyLister helps list ComputeResourcePolicies.
type ComputeResourcePolicyLister interface {
	// List lists all ComputeResourcePolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeResourcePolicy, err error)
	// ComputeResourcePolicies returns an object that can list and get ComputeResourcePolicies.
	ComputeResourcePolicies(namespace string) ComputeResourcePolicyNamespaceLister
	ComputeResourcePolicyListerExpansion
}

// computeResourcePolicyLister implements the ComputeResourcePolicyLister interface.
type computeResourcePolicyLister struct {
	indexer cache.Indexer
}

// NewComputeResourcePolicyLister returns a new ComputeResourcePolicyLister.
func NewComputeResourcePolicyLister(indexer cache.Indexer) ComputeResourcePolicyLister {
	return &computeResourcePolicyLister{indexer: indexer}
}

// List lists all ComputeResourcePolicies in the indexer.
func (s *computeResourcePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeResourcePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeResourcePolicy))
	})
	return ret, err
}

// ComputeResourcePolicies returns an object that can list and get ComputeResourcePolicies.
func (s *computeResourcePolicyLister) ComputeResourcePolicies(namespace string) ComputeResourcePolicyNamespaceLister {
	return computeResourcePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeResourcePolicyNamespaceLister helps list and get ComputeResourcePolicies.
type ComputeResourcePolicyNamespaceLister interface {
	// List lists all ComputeResourcePolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeResourcePolicy, err error)
	// Get retrieves the ComputeResourcePolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeResourcePolicy, error)
	ComputeResourcePolicyNamespaceListerExpansion
}

// computeResourcePolicyNamespaceLister implements the ComputeResourcePolicyNamespaceLister
// interface.
type computeResourcePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeResourcePolicies in the indexer for a given namespace.
func (s computeResourcePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeResourcePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeResourcePolicy))
	})
	return ret, err
}

// Get retrieves the ComputeResourcePolicy from the indexer for a given namespace and name.
func (s computeResourcePolicyNamespaceLister) Get(name string) (*v1alpha1.ComputeResourcePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeresourcepolicy"), name)
	}
	return obj.(*v1alpha1.ComputeResourcePolicy), nil
}