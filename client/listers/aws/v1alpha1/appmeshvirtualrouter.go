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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// AppmeshVirtualRouterLister helps list AppmeshVirtualRouters.
type AppmeshVirtualRouterLister interface {
	// List lists all AppmeshVirtualRouters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualRouter, err error)
	// AppmeshVirtualRouters returns an object that can list and get AppmeshVirtualRouters.
	AppmeshVirtualRouters(namespace string) AppmeshVirtualRouterNamespaceLister
	AppmeshVirtualRouterListerExpansion
}

// appmeshVirtualRouterLister implements the AppmeshVirtualRouterLister interface.
type appmeshVirtualRouterLister struct {
	indexer cache.Indexer
}

// NewAppmeshVirtualRouterLister returns a new AppmeshVirtualRouterLister.
func NewAppmeshVirtualRouterLister(indexer cache.Indexer) AppmeshVirtualRouterLister {
	return &appmeshVirtualRouterLister{indexer: indexer}
}

// List lists all AppmeshVirtualRouters in the indexer.
func (s *appmeshVirtualRouterLister) List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualRouter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppmeshVirtualRouter))
	})
	return ret, err
}

// AppmeshVirtualRouters returns an object that can list and get AppmeshVirtualRouters.
func (s *appmeshVirtualRouterLister) AppmeshVirtualRouters(namespace string) AppmeshVirtualRouterNamespaceLister {
	return appmeshVirtualRouterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppmeshVirtualRouterNamespaceLister helps list and get AppmeshVirtualRouters.
type AppmeshVirtualRouterNamespaceLister interface {
	// List lists all AppmeshVirtualRouters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualRouter, err error)
	// Get retrieves the AppmeshVirtualRouter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppmeshVirtualRouter, error)
	AppmeshVirtualRouterNamespaceListerExpansion
}

// appmeshVirtualRouterNamespaceLister implements the AppmeshVirtualRouterNamespaceLister
// interface.
type appmeshVirtualRouterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppmeshVirtualRouters in the indexer for a given namespace.
func (s appmeshVirtualRouterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppmeshVirtualRouter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppmeshVirtualRouter))
	})
	return ret, err
}

// Get retrieves the AppmeshVirtualRouter from the indexer for a given namespace and name.
func (s appmeshVirtualRouterNamespaceLister) Get(name string) (*v1alpha1.AppmeshVirtualRouter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appmeshvirtualrouter"), name)
	}
	return obj.(*v1alpha1.AppmeshVirtualRouter), nil
}
