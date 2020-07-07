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

// AppEngineStandardAppVersionLister helps list AppEngineStandardAppVersions.
type AppEngineStandardAppVersionLister interface {
	// List lists all AppEngineStandardAppVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineStandardAppVersion, err error)
	// AppEngineStandardAppVersions returns an object that can list and get AppEngineStandardAppVersions.
	AppEngineStandardAppVersions(namespace string) AppEngineStandardAppVersionNamespaceLister
	AppEngineStandardAppVersionListerExpansion
}

// appEngineStandardAppVersionLister implements the AppEngineStandardAppVersionLister interface.
type appEngineStandardAppVersionLister struct {
	indexer cache.Indexer
}

// NewAppEngineStandardAppVersionLister returns a new AppEngineStandardAppVersionLister.
func NewAppEngineStandardAppVersionLister(indexer cache.Indexer) AppEngineStandardAppVersionLister {
	return &appEngineStandardAppVersionLister{indexer: indexer}
}

// List lists all AppEngineStandardAppVersions in the indexer.
func (s *appEngineStandardAppVersionLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineStandardAppVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineStandardAppVersion))
	})
	return ret, err
}

// AppEngineStandardAppVersions returns an object that can list and get AppEngineStandardAppVersions.
func (s *appEngineStandardAppVersionLister) AppEngineStandardAppVersions(namespace string) AppEngineStandardAppVersionNamespaceLister {
	return appEngineStandardAppVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppEngineStandardAppVersionNamespaceLister helps list and get AppEngineStandardAppVersions.
type AppEngineStandardAppVersionNamespaceLister interface {
	// List lists all AppEngineStandardAppVersions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineStandardAppVersion, err error)
	// Get retrieves the AppEngineStandardAppVersion from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppEngineStandardAppVersion, error)
	AppEngineStandardAppVersionNamespaceListerExpansion
}

// appEngineStandardAppVersionNamespaceLister implements the AppEngineStandardAppVersionNamespaceLister
// interface.
type appEngineStandardAppVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppEngineStandardAppVersions in the indexer for a given namespace.
func (s appEngineStandardAppVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineStandardAppVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineStandardAppVersion))
	})
	return ret, err
}

// Get retrieves the AppEngineStandardAppVersion from the indexer for a given namespace and name.
func (s appEngineStandardAppVersionNamespaceLister) Get(name string) (*v1alpha1.AppEngineStandardAppVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appenginestandardappversion"), name)
	}
	return obj.(*v1alpha1.AppEngineStandardAppVersion), nil
}
