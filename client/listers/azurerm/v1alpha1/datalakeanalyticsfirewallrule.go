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

// DataLakeAnalyticsFirewallRuleLister helps list DataLakeAnalyticsFirewallRules.
type DataLakeAnalyticsFirewallRuleLister interface {
	// List lists all DataLakeAnalyticsFirewallRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataLakeAnalyticsFirewallRule, err error)
	// DataLakeAnalyticsFirewallRules returns an object that can list and get DataLakeAnalyticsFirewallRules.
	DataLakeAnalyticsFirewallRules(namespace string) DataLakeAnalyticsFirewallRuleNamespaceLister
	DataLakeAnalyticsFirewallRuleListerExpansion
}

// dataLakeAnalyticsFirewallRuleLister implements the DataLakeAnalyticsFirewallRuleLister interface.
type dataLakeAnalyticsFirewallRuleLister struct {
	indexer cache.Indexer
}

// NewDataLakeAnalyticsFirewallRuleLister returns a new DataLakeAnalyticsFirewallRuleLister.
func NewDataLakeAnalyticsFirewallRuleLister(indexer cache.Indexer) DataLakeAnalyticsFirewallRuleLister {
	return &dataLakeAnalyticsFirewallRuleLister{indexer: indexer}
}

// List lists all DataLakeAnalyticsFirewallRules in the indexer.
func (s *dataLakeAnalyticsFirewallRuleLister) List(selector labels.Selector) (ret []*v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataLakeAnalyticsFirewallRule))
	})
	return ret, err
}

// DataLakeAnalyticsFirewallRules returns an object that can list and get DataLakeAnalyticsFirewallRules.
func (s *dataLakeAnalyticsFirewallRuleLister) DataLakeAnalyticsFirewallRules(namespace string) DataLakeAnalyticsFirewallRuleNamespaceLister {
	return dataLakeAnalyticsFirewallRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataLakeAnalyticsFirewallRuleNamespaceLister helps list and get DataLakeAnalyticsFirewallRules.
type DataLakeAnalyticsFirewallRuleNamespaceLister interface {
	// List lists all DataLakeAnalyticsFirewallRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataLakeAnalyticsFirewallRule, err error)
	// Get retrieves the DataLakeAnalyticsFirewallRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataLakeAnalyticsFirewallRule, error)
	DataLakeAnalyticsFirewallRuleNamespaceListerExpansion
}

// dataLakeAnalyticsFirewallRuleNamespaceLister implements the DataLakeAnalyticsFirewallRuleNamespaceLister
// interface.
type dataLakeAnalyticsFirewallRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataLakeAnalyticsFirewallRules in the indexer for a given namespace.
func (s dataLakeAnalyticsFirewallRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataLakeAnalyticsFirewallRule))
	})
	return ret, err
}

// Get retrieves the DataLakeAnalyticsFirewallRule from the indexer for a given namespace and name.
func (s dataLakeAnalyticsFirewallRuleNamespaceLister) Get(name string) (*v1alpha1.DataLakeAnalyticsFirewallRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datalakeanalyticsfirewallrule"), name)
	}
	return obj.(*v1alpha1.DataLakeAnalyticsFirewallRule), nil
}
