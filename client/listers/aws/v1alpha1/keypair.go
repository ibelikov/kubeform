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

// KeyPairLister helps list KeyPairs.
type KeyPairLister interface {
	// List lists all KeyPairs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KeyPair, err error)
	// KeyPairs returns an object that can list and get KeyPairs.
	KeyPairs(namespace string) KeyPairNamespaceLister
	KeyPairListerExpansion
}

// keyPairLister implements the KeyPairLister interface.
type keyPairLister struct {
	indexer cache.Indexer
}

// NewKeyPairLister returns a new KeyPairLister.
func NewKeyPairLister(indexer cache.Indexer) KeyPairLister {
	return &keyPairLister{indexer: indexer}
}

// List lists all KeyPairs in the indexer.
func (s *keyPairLister) List(selector labels.Selector) (ret []*v1alpha1.KeyPair, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyPair))
	})
	return ret, err
}

// KeyPairs returns an object that can list and get KeyPairs.
func (s *keyPairLister) KeyPairs(namespace string) KeyPairNamespaceLister {
	return keyPairNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KeyPairNamespaceLister helps list and get KeyPairs.
type KeyPairNamespaceLister interface {
	// List lists all KeyPairs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KeyPair, err error)
	// Get retrieves the KeyPair from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KeyPair, error)
	KeyPairNamespaceListerExpansion
}

// keyPairNamespaceLister implements the KeyPairNamespaceLister
// interface.
type keyPairNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KeyPairs in the indexer for a given namespace.
func (s keyPairNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KeyPair, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyPair))
	})
	return ret, err
}

// Get retrieves the KeyPair from the indexer for a given namespace and name.
func (s keyPairNamespaceLister) Get(name string) (*v1alpha1.KeyPair, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("keypair"), name)
	}
	return obj.(*v1alpha1.KeyPair), nil
}
