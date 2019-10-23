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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MssqlElasticpoolLister helps list MssqlElasticpools.
type MssqlElasticpoolLister interface {
	// List lists all MssqlElasticpools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MssqlElasticpool, err error)
	// MssqlElasticpools returns an object that can list and get MssqlElasticpools.
	MssqlElasticpools(namespace string) MssqlElasticpoolNamespaceLister
	MssqlElasticpoolListerExpansion
}

// mssqlElasticpoolLister implements the MssqlElasticpoolLister interface.
type mssqlElasticpoolLister struct {
	indexer cache.Indexer
}

// NewMssqlElasticpoolLister returns a new MssqlElasticpoolLister.
func NewMssqlElasticpoolLister(indexer cache.Indexer) MssqlElasticpoolLister {
	return &mssqlElasticpoolLister{indexer: indexer}
}

// List lists all MssqlElasticpools in the indexer.
func (s *mssqlElasticpoolLister) List(selector labels.Selector) (ret []*v1alpha1.MssqlElasticpool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MssqlElasticpool))
	})
	return ret, err
}

// MssqlElasticpools returns an object that can list and get MssqlElasticpools.
func (s *mssqlElasticpoolLister) MssqlElasticpools(namespace string) MssqlElasticpoolNamespaceLister {
	return mssqlElasticpoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MssqlElasticpoolNamespaceLister helps list and get MssqlElasticpools.
type MssqlElasticpoolNamespaceLister interface {
	// List lists all MssqlElasticpools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MssqlElasticpool, err error)
	// Get retrieves the MssqlElasticpool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MssqlElasticpool, error)
	MssqlElasticpoolNamespaceListerExpansion
}

// mssqlElasticpoolNamespaceLister implements the MssqlElasticpoolNamespaceLister
// interface.
type mssqlElasticpoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MssqlElasticpools in the indexer for a given namespace.
func (s mssqlElasticpoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MssqlElasticpool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MssqlElasticpool))
	})
	return ret, err
}

// Get retrieves the MssqlElasticpool from the indexer for a given namespace and name.
func (s mssqlElasticpoolNamespaceLister) Get(name string) (*v1alpha1.MssqlElasticpool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mssqlelasticpool"), name)
	}
	return obj.(*v1alpha1.MssqlElasticpool), nil
}
