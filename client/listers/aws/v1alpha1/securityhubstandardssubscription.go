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

// SecurityhubStandardsSubscriptionLister helps list SecurityhubStandardsSubscriptions.
type SecurityhubStandardsSubscriptionLister interface {
	// List lists all SecurityhubStandardsSubscriptions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityhubStandardsSubscription, err error)
	// SecurityhubStandardsSubscriptions returns an object that can list and get SecurityhubStandardsSubscriptions.
	SecurityhubStandardsSubscriptions(namespace string) SecurityhubStandardsSubscriptionNamespaceLister
	SecurityhubStandardsSubscriptionListerExpansion
}

// securityhubStandardsSubscriptionLister implements the SecurityhubStandardsSubscriptionLister interface.
type securityhubStandardsSubscriptionLister struct {
	indexer cache.Indexer
}

// NewSecurityhubStandardsSubscriptionLister returns a new SecurityhubStandardsSubscriptionLister.
func NewSecurityhubStandardsSubscriptionLister(indexer cache.Indexer) SecurityhubStandardsSubscriptionLister {
	return &securityhubStandardsSubscriptionLister{indexer: indexer}
}

// List lists all SecurityhubStandardsSubscriptions in the indexer.
func (s *securityhubStandardsSubscriptionLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityhubStandardsSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityhubStandardsSubscription))
	})
	return ret, err
}

// SecurityhubStandardsSubscriptions returns an object that can list and get SecurityhubStandardsSubscriptions.
func (s *securityhubStandardsSubscriptionLister) SecurityhubStandardsSubscriptions(namespace string) SecurityhubStandardsSubscriptionNamespaceLister {
	return securityhubStandardsSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecurityhubStandardsSubscriptionNamespaceLister helps list and get SecurityhubStandardsSubscriptions.
type SecurityhubStandardsSubscriptionNamespaceLister interface {
	// List lists all SecurityhubStandardsSubscriptions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityhubStandardsSubscription, err error)
	// Get retrieves the SecurityhubStandardsSubscription from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecurityhubStandardsSubscription, error)
	SecurityhubStandardsSubscriptionNamespaceListerExpansion
}

// securityhubStandardsSubscriptionNamespaceLister implements the SecurityhubStandardsSubscriptionNamespaceLister
// interface.
type securityhubStandardsSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecurityhubStandardsSubscriptions in the indexer for a given namespace.
func (s securityhubStandardsSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityhubStandardsSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityhubStandardsSubscription))
	})
	return ret, err
}

// Get retrieves the SecurityhubStandardsSubscription from the indexer for a given namespace and name.
func (s securityhubStandardsSubscriptionNamespaceLister) Get(name string) (*v1alpha1.SecurityhubStandardsSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("securityhubstandardssubscription"), name)
	}
	return obj.(*v1alpha1.SecurityhubStandardsSubscription), nil
}
