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

// LambdaLayerVersionLister helps list LambdaLayerVersions.
type LambdaLayerVersionLister interface {
	// List lists all LambdaLayerVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LambdaLayerVersion, err error)
	// LambdaLayerVersions returns an object that can list and get LambdaLayerVersions.
	LambdaLayerVersions(namespace string) LambdaLayerVersionNamespaceLister
	LambdaLayerVersionListerExpansion
}

// lambdaLayerVersionLister implements the LambdaLayerVersionLister interface.
type lambdaLayerVersionLister struct {
	indexer cache.Indexer
}

// NewLambdaLayerVersionLister returns a new LambdaLayerVersionLister.
func NewLambdaLayerVersionLister(indexer cache.Indexer) LambdaLayerVersionLister {
	return &lambdaLayerVersionLister{indexer: indexer}
}

// List lists all LambdaLayerVersions in the indexer.
func (s *lambdaLayerVersionLister) List(selector labels.Selector) (ret []*v1alpha1.LambdaLayerVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LambdaLayerVersion))
	})
	return ret, err
}

// LambdaLayerVersions returns an object that can list and get LambdaLayerVersions.
func (s *lambdaLayerVersionLister) LambdaLayerVersions(namespace string) LambdaLayerVersionNamespaceLister {
	return lambdaLayerVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LambdaLayerVersionNamespaceLister helps list and get LambdaLayerVersions.
type LambdaLayerVersionNamespaceLister interface {
	// List lists all LambdaLayerVersions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LambdaLayerVersion, err error)
	// Get retrieves the LambdaLayerVersion from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LambdaLayerVersion, error)
	LambdaLayerVersionNamespaceListerExpansion
}

// lambdaLayerVersionNamespaceLister implements the LambdaLayerVersionNamespaceLister
// interface.
type lambdaLayerVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LambdaLayerVersions in the indexer for a given namespace.
func (s lambdaLayerVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LambdaLayerVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LambdaLayerVersion))
	})
	return ret, err
}

// Get retrieves the LambdaLayerVersion from the indexer for a given namespace and name.
func (s lambdaLayerVersionNamespaceLister) Get(name string) (*v1alpha1.LambdaLayerVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lambdalayerversion"), name)
	}
	return obj.(*v1alpha1.LambdaLayerVersion), nil
}
