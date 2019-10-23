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

// OpsworksRdsDbInstanceLister helps list OpsworksRdsDbInstances.
type OpsworksRdsDbInstanceLister interface {
	// List lists all OpsworksRdsDbInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksRdsDbInstance, err error)
	// OpsworksRdsDbInstances returns an object that can list and get OpsworksRdsDbInstances.
	OpsworksRdsDbInstances(namespace string) OpsworksRdsDbInstanceNamespaceLister
	OpsworksRdsDbInstanceListerExpansion
}

// opsworksRdsDbInstanceLister implements the OpsworksRdsDbInstanceLister interface.
type opsworksRdsDbInstanceLister struct {
	indexer cache.Indexer
}

// NewOpsworksRdsDbInstanceLister returns a new OpsworksRdsDbInstanceLister.
func NewOpsworksRdsDbInstanceLister(indexer cache.Indexer) OpsworksRdsDbInstanceLister {
	return &opsworksRdsDbInstanceLister{indexer: indexer}
}

// List lists all OpsworksRdsDbInstances in the indexer.
func (s *opsworksRdsDbInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksRdsDbInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksRdsDbInstance))
	})
	return ret, err
}

// OpsworksRdsDbInstances returns an object that can list and get OpsworksRdsDbInstances.
func (s *opsworksRdsDbInstanceLister) OpsworksRdsDbInstances(namespace string) OpsworksRdsDbInstanceNamespaceLister {
	return opsworksRdsDbInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpsworksRdsDbInstanceNamespaceLister helps list and get OpsworksRdsDbInstances.
type OpsworksRdsDbInstanceNamespaceLister interface {
	// List lists all OpsworksRdsDbInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksRdsDbInstance, err error)
	// Get retrieves the OpsworksRdsDbInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OpsworksRdsDbInstance, error)
	OpsworksRdsDbInstanceNamespaceListerExpansion
}

// opsworksRdsDbInstanceNamespaceLister implements the OpsworksRdsDbInstanceNamespaceLister
// interface.
type opsworksRdsDbInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpsworksRdsDbInstances in the indexer for a given namespace.
func (s opsworksRdsDbInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksRdsDbInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksRdsDbInstance))
	})
	return ret, err
}

// Get retrieves the OpsworksRdsDbInstance from the indexer for a given namespace and name.
func (s opsworksRdsDbInstanceNamespaceLister) Get(name string) (*v1alpha1.OpsworksRdsDbInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("opsworksrdsdbinstance"), name)
	}
	return obj.(*v1alpha1.OpsworksRdsDbInstance), nil
}
