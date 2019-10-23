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

// DatasyncLocationNfsLister helps list DatasyncLocationNfses.
type DatasyncLocationNfsLister interface {
	// List lists all DatasyncLocationNfses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationNfs, err error)
	// DatasyncLocationNfses returns an object that can list and get DatasyncLocationNfses.
	DatasyncLocationNfses(namespace string) DatasyncLocationNfsNamespaceLister
	DatasyncLocationNfsListerExpansion
}

// datasyncLocationNfsLister implements the DatasyncLocationNfsLister interface.
type datasyncLocationNfsLister struct {
	indexer cache.Indexer
}

// NewDatasyncLocationNfsLister returns a new DatasyncLocationNfsLister.
func NewDatasyncLocationNfsLister(indexer cache.Indexer) DatasyncLocationNfsLister {
	return &datasyncLocationNfsLister{indexer: indexer}
}

// List lists all DatasyncLocationNfses in the indexer.
func (s *datasyncLocationNfsLister) List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationNfs, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasyncLocationNfs))
	})
	return ret, err
}

// DatasyncLocationNfses returns an object that can list and get DatasyncLocationNfses.
func (s *datasyncLocationNfsLister) DatasyncLocationNfses(namespace string) DatasyncLocationNfsNamespaceLister {
	return datasyncLocationNfsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatasyncLocationNfsNamespaceLister helps list and get DatasyncLocationNfses.
type DatasyncLocationNfsNamespaceLister interface {
	// List lists all DatasyncLocationNfses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationNfs, err error)
	// Get retrieves the DatasyncLocationNfs from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DatasyncLocationNfs, error)
	DatasyncLocationNfsNamespaceListerExpansion
}

// datasyncLocationNfsNamespaceLister implements the DatasyncLocationNfsNamespaceLister
// interface.
type datasyncLocationNfsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DatasyncLocationNfses in the indexer for a given namespace.
func (s datasyncLocationNfsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DatasyncLocationNfs, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasyncLocationNfs))
	})
	return ret, err
}

// Get retrieves the DatasyncLocationNfs from the indexer for a given namespace and name.
func (s datasyncLocationNfsNamespaceLister) Get(name string) (*v1alpha1.DatasyncLocationNfs, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datasynclocationnfs"), name)
	}
	return obj.(*v1alpha1.DatasyncLocationNfs), nil
}
