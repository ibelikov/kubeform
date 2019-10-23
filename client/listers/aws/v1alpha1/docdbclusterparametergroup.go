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

// DocdbClusterParameterGroupLister helps list DocdbClusterParameterGroups.
type DocdbClusterParameterGroupLister interface {
	// List lists all DocdbClusterParameterGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterParameterGroup, err error)
	// DocdbClusterParameterGroups returns an object that can list and get DocdbClusterParameterGroups.
	DocdbClusterParameterGroups(namespace string) DocdbClusterParameterGroupNamespaceLister
	DocdbClusterParameterGroupListerExpansion
}

// docdbClusterParameterGroupLister implements the DocdbClusterParameterGroupLister interface.
type docdbClusterParameterGroupLister struct {
	indexer cache.Indexer
}

// NewDocdbClusterParameterGroupLister returns a new DocdbClusterParameterGroupLister.
func NewDocdbClusterParameterGroupLister(indexer cache.Indexer) DocdbClusterParameterGroupLister {
	return &docdbClusterParameterGroupLister{indexer: indexer}
}

// List lists all DocdbClusterParameterGroups in the indexer.
func (s *docdbClusterParameterGroupLister) List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterParameterGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DocdbClusterParameterGroup))
	})
	return ret, err
}

// DocdbClusterParameterGroups returns an object that can list and get DocdbClusterParameterGroups.
func (s *docdbClusterParameterGroupLister) DocdbClusterParameterGroups(namespace string) DocdbClusterParameterGroupNamespaceLister {
	return docdbClusterParameterGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DocdbClusterParameterGroupNamespaceLister helps list and get DocdbClusterParameterGroups.
type DocdbClusterParameterGroupNamespaceLister interface {
	// List lists all DocdbClusterParameterGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterParameterGroup, err error)
	// Get retrieves the DocdbClusterParameterGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DocdbClusterParameterGroup, error)
	DocdbClusterParameterGroupNamespaceListerExpansion
}

// docdbClusterParameterGroupNamespaceLister implements the DocdbClusterParameterGroupNamespaceLister
// interface.
type docdbClusterParameterGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DocdbClusterParameterGroups in the indexer for a given namespace.
func (s docdbClusterParameterGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterParameterGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DocdbClusterParameterGroup))
	})
	return ret, err
}

// Get retrieves the DocdbClusterParameterGroup from the indexer for a given namespace and name.
func (s docdbClusterParameterGroupNamespaceLister) Get(name string) (*v1alpha1.DocdbClusterParameterGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("docdbclusterparametergroup"), name)
	}
	return obj.(*v1alpha1.DocdbClusterParameterGroup), nil
}
