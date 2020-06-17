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

// AppEngineApplicationLister helps list AppEngineApplications.
type AppEngineApplicationLister interface {
	// List lists all AppEngineApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineApplication, err error)
	// AppEngineApplications returns an object that can list and get AppEngineApplications.
	AppEngineApplications(namespace string) AppEngineApplicationNamespaceLister
	AppEngineApplicationListerExpansion
}

// appEngineApplicationLister implements the AppEngineApplicationLister interface.
type appEngineApplicationLister struct {
	indexer cache.Indexer
}

// NewAppEngineApplicationLister returns a new AppEngineApplicationLister.
func NewAppEngineApplicationLister(indexer cache.Indexer) AppEngineApplicationLister {
	return &appEngineApplicationLister{indexer: indexer}
}

// List lists all AppEngineApplications in the indexer.
func (s *appEngineApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineApplication))
	})
	return ret, err
}

// AppEngineApplications returns an object that can list and get AppEngineApplications.
func (s *appEngineApplicationLister) AppEngineApplications(namespace string) AppEngineApplicationNamespaceLister {
	return appEngineApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppEngineApplicationNamespaceLister helps list and get AppEngineApplications.
type AppEngineApplicationNamespaceLister interface {
	// List lists all AppEngineApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineApplication, err error)
	// Get retrieves the AppEngineApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppEngineApplication, error)
	AppEngineApplicationNamespaceListerExpansion
}

// appEngineApplicationNamespaceLister implements the AppEngineApplicationNamespaceLister
// interface.
type appEngineApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppEngineApplications in the indexer for a given namespace.
func (s appEngineApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineApplication))
	})
	return ret, err
}

// Get retrieves the AppEngineApplication from the indexer for a given namespace and name.
func (s appEngineApplicationNamespaceLister) Get(name string) (*v1alpha1.AppEngineApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appengineapplication"), name)
	}
	return obj.(*v1alpha1.AppEngineApplication), nil
}
