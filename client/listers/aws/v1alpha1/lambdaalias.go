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

// LambdaAliasLister helps list LambdaAliases.
type LambdaAliasLister interface {
	// List lists all LambdaAliases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LambdaAlias, err error)
	// LambdaAliases returns an object that can list and get LambdaAliases.
	LambdaAliases(namespace string) LambdaAliasNamespaceLister
	LambdaAliasListerExpansion
}

// lambdaAliasLister implements the LambdaAliasLister interface.
type lambdaAliasLister struct {
	indexer cache.Indexer
}

// NewLambdaAliasLister returns a new LambdaAliasLister.
func NewLambdaAliasLister(indexer cache.Indexer) LambdaAliasLister {
	return &lambdaAliasLister{indexer: indexer}
}

// List lists all LambdaAliases in the indexer.
func (s *lambdaAliasLister) List(selector labels.Selector) (ret []*v1alpha1.LambdaAlias, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LambdaAlias))
	})
	return ret, err
}

// LambdaAliases returns an object that can list and get LambdaAliases.
func (s *lambdaAliasLister) LambdaAliases(namespace string) LambdaAliasNamespaceLister {
	return lambdaAliasNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LambdaAliasNamespaceLister helps list and get LambdaAliases.
type LambdaAliasNamespaceLister interface {
	// List lists all LambdaAliases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LambdaAlias, err error)
	// Get retrieves the LambdaAlias from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LambdaAlias, error)
	LambdaAliasNamespaceListerExpansion
}

// lambdaAliasNamespaceLister implements the LambdaAliasNamespaceLister
// interface.
type lambdaAliasNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LambdaAliases in the indexer for a given namespace.
func (s lambdaAliasNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LambdaAlias, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LambdaAlias))
	})
	return ret, err
}

// Get retrieves the LambdaAlias from the indexer for a given namespace and name.
func (s lambdaAliasNamespaceLister) Get(name string) (*v1alpha1.LambdaAlias, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lambdaalias"), name)
	}
	return obj.(*v1alpha1.LambdaAlias), nil
}
