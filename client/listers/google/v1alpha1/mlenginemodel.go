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

// MlEngineModelLister helps list MlEngineModels.
type MlEngineModelLister interface {
	// List lists all MlEngineModels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MlEngineModel, err error)
	// MlEngineModels returns an object that can list and get MlEngineModels.
	MlEngineModels(namespace string) MlEngineModelNamespaceLister
	MlEngineModelListerExpansion
}

// mlEngineModelLister implements the MlEngineModelLister interface.
type mlEngineModelLister struct {
	indexer cache.Indexer
}

// NewMlEngineModelLister returns a new MlEngineModelLister.
func NewMlEngineModelLister(indexer cache.Indexer) MlEngineModelLister {
	return &mlEngineModelLister{indexer: indexer}
}

// List lists all MlEngineModels in the indexer.
func (s *mlEngineModelLister) List(selector labels.Selector) (ret []*v1alpha1.MlEngineModel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MlEngineModel))
	})
	return ret, err
}

// MlEngineModels returns an object that can list and get MlEngineModels.
func (s *mlEngineModelLister) MlEngineModels(namespace string) MlEngineModelNamespaceLister {
	return mlEngineModelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MlEngineModelNamespaceLister helps list and get MlEngineModels.
type MlEngineModelNamespaceLister interface {
	// List lists all MlEngineModels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MlEngineModel, err error)
	// Get retrieves the MlEngineModel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MlEngineModel, error)
	MlEngineModelNamespaceListerExpansion
}

// mlEngineModelNamespaceLister implements the MlEngineModelNamespaceLister
// interface.
type mlEngineModelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MlEngineModels in the indexer for a given namespace.
func (s mlEngineModelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MlEngineModel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MlEngineModel))
	})
	return ret, err
}

// Get retrieves the MlEngineModel from the indexer for a given namespace and name.
func (s mlEngineModelNamespaceLister) Get(name string) (*v1alpha1.MlEngineModel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mlenginemodel"), name)
	}
	return obj.(*v1alpha1.MlEngineModel), nil
}
