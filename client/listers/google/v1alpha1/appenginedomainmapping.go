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

// AppEngineDomainMappingLister helps list AppEngineDomainMappings.
type AppEngineDomainMappingLister interface {
	// List lists all AppEngineDomainMappings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineDomainMapping, err error)
	// AppEngineDomainMappings returns an object that can list and get AppEngineDomainMappings.
	AppEngineDomainMappings(namespace string) AppEngineDomainMappingNamespaceLister
	AppEngineDomainMappingListerExpansion
}

// appEngineDomainMappingLister implements the AppEngineDomainMappingLister interface.
type appEngineDomainMappingLister struct {
	indexer cache.Indexer
}

// NewAppEngineDomainMappingLister returns a new AppEngineDomainMappingLister.
func NewAppEngineDomainMappingLister(indexer cache.Indexer) AppEngineDomainMappingLister {
	return &appEngineDomainMappingLister{indexer: indexer}
}

// List lists all AppEngineDomainMappings in the indexer.
func (s *appEngineDomainMappingLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineDomainMapping, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineDomainMapping))
	})
	return ret, err
}

// AppEngineDomainMappings returns an object that can list and get AppEngineDomainMappings.
func (s *appEngineDomainMappingLister) AppEngineDomainMappings(namespace string) AppEngineDomainMappingNamespaceLister {
	return appEngineDomainMappingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppEngineDomainMappingNamespaceLister helps list and get AppEngineDomainMappings.
type AppEngineDomainMappingNamespaceLister interface {
	// List lists all AppEngineDomainMappings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineDomainMapping, err error)
	// Get retrieves the AppEngineDomainMapping from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppEngineDomainMapping, error)
	AppEngineDomainMappingNamespaceListerExpansion
}

// appEngineDomainMappingNamespaceLister implements the AppEngineDomainMappingNamespaceLister
// interface.
type appEngineDomainMappingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppEngineDomainMappings in the indexer for a given namespace.
func (s appEngineDomainMappingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineDomainMapping, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineDomainMapping))
	})
	return ret, err
}

// Get retrieves the AppEngineDomainMapping from the indexer for a given namespace and name.
func (s appEngineDomainMappingNamespaceLister) Get(name string) (*v1alpha1.AppEngineDomainMapping, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appenginedomainmapping"), name)
	}
	return obj.(*v1alpha1.AppEngineDomainMapping), nil
}