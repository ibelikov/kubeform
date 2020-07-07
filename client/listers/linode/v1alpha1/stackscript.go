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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// StackscriptLister helps list Stackscripts.
type StackscriptLister interface {
	// List lists all Stackscripts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Stackscript, err error)
	// Stackscripts returns an object that can list and get Stackscripts.
	Stackscripts(namespace string) StackscriptNamespaceLister
	StackscriptListerExpansion
}

// stackscriptLister implements the StackscriptLister interface.
type stackscriptLister struct {
	indexer cache.Indexer
}

// NewStackscriptLister returns a new StackscriptLister.
func NewStackscriptLister(indexer cache.Indexer) StackscriptLister {
	return &stackscriptLister{indexer: indexer}
}

// List lists all Stackscripts in the indexer.
func (s *stackscriptLister) List(selector labels.Selector) (ret []*v1alpha1.Stackscript, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Stackscript))
	})
	return ret, err
}

// Stackscripts returns an object that can list and get Stackscripts.
func (s *stackscriptLister) Stackscripts(namespace string) StackscriptNamespaceLister {
	return stackscriptNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StackscriptNamespaceLister helps list and get Stackscripts.
type StackscriptNamespaceLister interface {
	// List lists all Stackscripts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Stackscript, err error)
	// Get retrieves the Stackscript from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Stackscript, error)
	StackscriptNamespaceListerExpansion
}

// stackscriptNamespaceLister implements the StackscriptNamespaceLister
// interface.
type stackscriptNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Stackscripts in the indexer for a given namespace.
func (s stackscriptNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Stackscript, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Stackscript))
	})
	return ret, err
}

// Get retrieves the Stackscript from the indexer for a given namespace and name.
func (s stackscriptNamespaceLister) Get(name string) (*v1alpha1.Stackscript, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("stackscript"), name)
	}
	return obj.(*v1alpha1.Stackscript), nil
}
