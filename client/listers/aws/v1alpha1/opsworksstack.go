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

// OpsworksStackLister helps list OpsworksStacks.
type OpsworksStackLister interface {
	// List lists all OpsworksStacks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksStack, err error)
	// OpsworksStacks returns an object that can list and get OpsworksStacks.
	OpsworksStacks(namespace string) OpsworksStackNamespaceLister
	OpsworksStackListerExpansion
}

// opsworksStackLister implements the OpsworksStackLister interface.
type opsworksStackLister struct {
	indexer cache.Indexer
}

// NewOpsworksStackLister returns a new OpsworksStackLister.
func NewOpsworksStackLister(indexer cache.Indexer) OpsworksStackLister {
	return &opsworksStackLister{indexer: indexer}
}

// List lists all OpsworksStacks in the indexer.
func (s *opsworksStackLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksStack, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksStack))
	})
	return ret, err
}

// OpsworksStacks returns an object that can list and get OpsworksStacks.
func (s *opsworksStackLister) OpsworksStacks(namespace string) OpsworksStackNamespaceLister {
	return opsworksStackNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpsworksStackNamespaceLister helps list and get OpsworksStacks.
type OpsworksStackNamespaceLister interface {
	// List lists all OpsworksStacks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksStack, err error)
	// Get retrieves the OpsworksStack from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OpsworksStack, error)
	OpsworksStackNamespaceListerExpansion
}

// opsworksStackNamespaceLister implements the OpsworksStackNamespaceLister
// interface.
type opsworksStackNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpsworksStacks in the indexer for a given namespace.
func (s opsworksStackNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksStack, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksStack))
	})
	return ret, err
}

// Get retrieves the OpsworksStack from the indexer for a given namespace and name.
func (s opsworksStackNamespaceLister) Get(name string) (*v1alpha1.OpsworksStack, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("opsworksstack"), name)
	}
	return obj.(*v1alpha1.OpsworksStack), nil
}
