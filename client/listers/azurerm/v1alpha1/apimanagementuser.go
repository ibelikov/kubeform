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

// ApiManagementUserLister helps list ApiManagementUsers.
type ApiManagementUserLister interface {
	// List lists all ApiManagementUsers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementUser, err error)
	// ApiManagementUsers returns an object that can list and get ApiManagementUsers.
	ApiManagementUsers(namespace string) ApiManagementUserNamespaceLister
	ApiManagementUserListerExpansion
}

// apiManagementUserLister implements the ApiManagementUserLister interface.
type apiManagementUserLister struct {
	indexer cache.Indexer
}

// NewApiManagementUserLister returns a new ApiManagementUserLister.
func NewApiManagementUserLister(indexer cache.Indexer) ApiManagementUserLister {
	return &apiManagementUserLister{indexer: indexer}
}

// List lists all ApiManagementUsers in the indexer.
func (s *apiManagementUserLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementUser))
	})
	return ret, err
}

// ApiManagementUsers returns an object that can list and get ApiManagementUsers.
func (s *apiManagementUserLister) ApiManagementUsers(namespace string) ApiManagementUserNamespaceLister {
	return apiManagementUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementUserNamespaceLister helps list and get ApiManagementUsers.
type ApiManagementUserNamespaceLister interface {
	// List lists all ApiManagementUsers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementUser, err error)
	// Get retrieves the ApiManagementUser from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagementUser, error)
	ApiManagementUserNamespaceListerExpansion
}

// apiManagementUserNamespaceLister implements the ApiManagementUserNamespaceLister
// interface.
type apiManagementUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagementUsers in the indexer for a given namespace.
func (s apiManagementUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementUser))
	})
	return ret, err
}

// Get retrieves the ApiManagementUser from the indexer for a given namespace and name.
func (s apiManagementUserNamespaceLister) Get(name string) (*v1alpha1.ApiManagementUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementuser"), name)
	}
	return obj.(*v1alpha1.ApiManagementUser), nil
}
