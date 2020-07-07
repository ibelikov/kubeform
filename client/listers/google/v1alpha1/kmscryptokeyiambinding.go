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

// KmsCryptoKeyIamBindingLister helps list KmsCryptoKeyIamBindings.
type KmsCryptoKeyIamBindingLister interface {
	// List lists all KmsCryptoKeyIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamBinding, err error)
	// KmsCryptoKeyIamBindings returns an object that can list and get KmsCryptoKeyIamBindings.
	KmsCryptoKeyIamBindings(namespace string) KmsCryptoKeyIamBindingNamespaceLister
	KmsCryptoKeyIamBindingListerExpansion
}

// kmsCryptoKeyIamBindingLister implements the KmsCryptoKeyIamBindingLister interface.
type kmsCryptoKeyIamBindingLister struct {
	indexer cache.Indexer
}

// NewKmsCryptoKeyIamBindingLister returns a new KmsCryptoKeyIamBindingLister.
func NewKmsCryptoKeyIamBindingLister(indexer cache.Indexer) KmsCryptoKeyIamBindingLister {
	return &kmsCryptoKeyIamBindingLister{indexer: indexer}
}

// List lists all KmsCryptoKeyIamBindings in the indexer.
func (s *kmsCryptoKeyIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCryptoKeyIamBinding))
	})
	return ret, err
}

// KmsCryptoKeyIamBindings returns an object that can list and get KmsCryptoKeyIamBindings.
func (s *kmsCryptoKeyIamBindingLister) KmsCryptoKeyIamBindings(namespace string) KmsCryptoKeyIamBindingNamespaceLister {
	return kmsCryptoKeyIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KmsCryptoKeyIamBindingNamespaceLister helps list and get KmsCryptoKeyIamBindings.
type KmsCryptoKeyIamBindingNamespaceLister interface {
	// List lists all KmsCryptoKeyIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamBinding, err error)
	// Get retrieves the KmsCryptoKeyIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KmsCryptoKeyIamBinding, error)
	KmsCryptoKeyIamBindingNamespaceListerExpansion
}

// kmsCryptoKeyIamBindingNamespaceLister implements the KmsCryptoKeyIamBindingNamespaceLister
// interface.
type kmsCryptoKeyIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KmsCryptoKeyIamBindings in the indexer for a given namespace.
func (s kmsCryptoKeyIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCryptoKeyIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCryptoKeyIamBinding))
	})
	return ret, err
}

// Get retrieves the KmsCryptoKeyIamBinding from the indexer for a given namespace and name.
func (s kmsCryptoKeyIamBindingNamespaceLister) Get(name string) (*v1alpha1.KmsCryptoKeyIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kmscryptokeyiambinding"), name)
	}
	return obj.(*v1alpha1.KmsCryptoKeyIamBinding), nil
}
