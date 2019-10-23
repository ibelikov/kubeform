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

// WafregionalRuleLister helps list WafregionalRules.
type WafregionalRuleLister interface {
	// List lists all WafregionalRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalRule, err error)
	// WafregionalRules returns an object that can list and get WafregionalRules.
	WafregionalRules(namespace string) WafregionalRuleNamespaceLister
	WafregionalRuleListerExpansion
}

// wafregionalRuleLister implements the WafregionalRuleLister interface.
type wafregionalRuleLister struct {
	indexer cache.Indexer
}

// NewWafregionalRuleLister returns a new WafregionalRuleLister.
func NewWafregionalRuleLister(indexer cache.Indexer) WafregionalRuleLister {
	return &wafregionalRuleLister{indexer: indexer}
}

// List lists all WafregionalRules in the indexer.
func (s *wafregionalRuleLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalRule))
	})
	return ret, err
}

// WafregionalRules returns an object that can list and get WafregionalRules.
func (s *wafregionalRuleLister) WafregionalRules(namespace string) WafregionalRuleNamespaceLister {
	return wafregionalRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafregionalRuleNamespaceLister helps list and get WafregionalRules.
type WafregionalRuleNamespaceLister interface {
	// List lists all WafregionalRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalRule, err error)
	// Get retrieves the WafregionalRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafregionalRule, error)
	WafregionalRuleNamespaceListerExpansion
}

// wafregionalRuleNamespaceLister implements the WafregionalRuleNamespaceLister
// interface.
type wafregionalRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafregionalRules in the indexer for a given namespace.
func (s wafregionalRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalRule))
	})
	return ret, err
}

// Get retrieves the WafregionalRule from the indexer for a given namespace and name.
func (s wafregionalRuleNamespaceLister) Get(name string) (*v1alpha1.WafregionalRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafregionalrule"), name)
	}
	return obj.(*v1alpha1.WafregionalRule), nil
}
