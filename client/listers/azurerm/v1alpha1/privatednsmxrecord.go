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

// PrivateDNSMxRecordLister helps list PrivateDNSMxRecords.
type PrivateDNSMxRecordLister interface {
	// List lists all PrivateDNSMxRecords in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSMxRecord, err error)
	// PrivateDNSMxRecords returns an object that can list and get PrivateDNSMxRecords.
	PrivateDNSMxRecords(namespace string) PrivateDNSMxRecordNamespaceLister
	PrivateDNSMxRecordListerExpansion
}

// privateDNSMxRecordLister implements the PrivateDNSMxRecordLister interface.
type privateDNSMxRecordLister struct {
	indexer cache.Indexer
}

// NewPrivateDNSMxRecordLister returns a new PrivateDNSMxRecordLister.
func NewPrivateDNSMxRecordLister(indexer cache.Indexer) PrivateDNSMxRecordLister {
	return &privateDNSMxRecordLister{indexer: indexer}
}

// List lists all PrivateDNSMxRecords in the indexer.
func (s *privateDNSMxRecordLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSMxRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateDNSMxRecord))
	})
	return ret, err
}

// PrivateDNSMxRecords returns an object that can list and get PrivateDNSMxRecords.
func (s *privateDNSMxRecordLister) PrivateDNSMxRecords(namespace string) PrivateDNSMxRecordNamespaceLister {
	return privateDNSMxRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrivateDNSMxRecordNamespaceLister helps list and get PrivateDNSMxRecords.
type PrivateDNSMxRecordNamespaceLister interface {
	// List lists all PrivateDNSMxRecords in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSMxRecord, err error)
	// Get retrieves the PrivateDNSMxRecord from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PrivateDNSMxRecord, error)
	PrivateDNSMxRecordNamespaceListerExpansion
}

// privateDNSMxRecordNamespaceLister implements the PrivateDNSMxRecordNamespaceLister
// interface.
type privateDNSMxRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrivateDNSMxRecords in the indexer for a given namespace.
func (s privateDNSMxRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateDNSMxRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateDNSMxRecord))
	})
	return ret, err
}

// Get retrieves the PrivateDNSMxRecord from the indexer for a given namespace and name.
func (s privateDNSMxRecordNamespaceLister) Get(name string) (*v1alpha1.PrivateDNSMxRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("privatednsmxrecord"), name)
	}
	return obj.(*v1alpha1.PrivateDNSMxRecord), nil
}
