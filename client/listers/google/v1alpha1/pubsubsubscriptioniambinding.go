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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PubsubSubscriptionIamBindingLister helps list PubsubSubscriptionIamBindings.
type PubsubSubscriptionIamBindingLister interface {
	// List lists all PubsubSubscriptionIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamBinding, err error)
	// PubsubSubscriptionIamBindings returns an object that can list and get PubsubSubscriptionIamBindings.
	PubsubSubscriptionIamBindings(namespace string) PubsubSubscriptionIamBindingNamespaceLister
	PubsubSubscriptionIamBindingListerExpansion
}

// pubsubSubscriptionIamBindingLister implements the PubsubSubscriptionIamBindingLister interface.
type pubsubSubscriptionIamBindingLister struct {
	indexer cache.Indexer
}

// NewPubsubSubscriptionIamBindingLister returns a new PubsubSubscriptionIamBindingLister.
func NewPubsubSubscriptionIamBindingLister(indexer cache.Indexer) PubsubSubscriptionIamBindingLister {
	return &pubsubSubscriptionIamBindingLister{indexer: indexer}
}

// List lists all PubsubSubscriptionIamBindings in the indexer.
func (s *pubsubSubscriptionIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubSubscriptionIamBinding))
	})
	return ret, err
}

// PubsubSubscriptionIamBindings returns an object that can list and get PubsubSubscriptionIamBindings.
func (s *pubsubSubscriptionIamBindingLister) PubsubSubscriptionIamBindings(namespace string) PubsubSubscriptionIamBindingNamespaceLister {
	return pubsubSubscriptionIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PubsubSubscriptionIamBindingNamespaceLister helps list and get PubsubSubscriptionIamBindings.
type PubsubSubscriptionIamBindingNamespaceLister interface {
	// List lists all PubsubSubscriptionIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamBinding, err error)
	// Get retrieves the PubsubSubscriptionIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PubsubSubscriptionIamBinding, error)
	PubsubSubscriptionIamBindingNamespaceListerExpansion
}

// pubsubSubscriptionIamBindingNamespaceLister implements the PubsubSubscriptionIamBindingNamespaceLister
// interface.
type pubsubSubscriptionIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PubsubSubscriptionIamBindings in the indexer for a given namespace.
func (s pubsubSubscriptionIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubSubscriptionIamBinding))
	})
	return ret, err
}

// Get retrieves the PubsubSubscriptionIamBinding from the indexer for a given namespace and name.
func (s pubsubSubscriptionIamBindingNamespaceLister) Get(name string) (*v1alpha1.PubsubSubscriptionIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pubsubsubscriptioniambinding"), name)
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamBinding), nil
}
