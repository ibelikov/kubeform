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

// IotTopicRuleLister helps list IotTopicRules.
type IotTopicRuleLister interface {
	// List lists all IotTopicRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IotTopicRule, err error)
	// IotTopicRules returns an object that can list and get IotTopicRules.
	IotTopicRules(namespace string) IotTopicRuleNamespaceLister
	IotTopicRuleListerExpansion
}

// iotTopicRuleLister implements the IotTopicRuleLister interface.
type iotTopicRuleLister struct {
	indexer cache.Indexer
}

// NewIotTopicRuleLister returns a new IotTopicRuleLister.
func NewIotTopicRuleLister(indexer cache.Indexer) IotTopicRuleLister {
	return &iotTopicRuleLister{indexer: indexer}
}

// List lists all IotTopicRules in the indexer.
func (s *iotTopicRuleLister) List(selector labels.Selector) (ret []*v1alpha1.IotTopicRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotTopicRule))
	})
	return ret, err
}

// IotTopicRules returns an object that can list and get IotTopicRules.
func (s *iotTopicRuleLister) IotTopicRules(namespace string) IotTopicRuleNamespaceLister {
	return iotTopicRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IotTopicRuleNamespaceLister helps list and get IotTopicRules.
type IotTopicRuleNamespaceLister interface {
	// List lists all IotTopicRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IotTopicRule, err error)
	// Get retrieves the IotTopicRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IotTopicRule, error)
	IotTopicRuleNamespaceListerExpansion
}

// iotTopicRuleNamespaceLister implements the IotTopicRuleNamespaceLister
// interface.
type iotTopicRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IotTopicRules in the indexer for a given namespace.
func (s iotTopicRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IotTopicRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotTopicRule))
	})
	return ret, err
}

// Get retrieves the IotTopicRule from the indexer for a given namespace and name.
func (s iotTopicRuleNamespaceLister) Get(name string) (*v1alpha1.IotTopicRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iottopicrule"), name)
	}
	return obj.(*v1alpha1.IotTopicRule), nil
}
