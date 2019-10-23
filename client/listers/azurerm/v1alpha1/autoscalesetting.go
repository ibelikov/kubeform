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

// AutoscaleSettingLister helps list AutoscaleSettings.
type AutoscaleSettingLister interface {
	// List lists all AutoscaleSettings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AutoscaleSetting, err error)
	// AutoscaleSettings returns an object that can list and get AutoscaleSettings.
	AutoscaleSettings(namespace string) AutoscaleSettingNamespaceLister
	AutoscaleSettingListerExpansion
}

// autoscaleSettingLister implements the AutoscaleSettingLister interface.
type autoscaleSettingLister struct {
	indexer cache.Indexer
}

// NewAutoscaleSettingLister returns a new AutoscaleSettingLister.
func NewAutoscaleSettingLister(indexer cache.Indexer) AutoscaleSettingLister {
	return &autoscaleSettingLister{indexer: indexer}
}

// List lists all AutoscaleSettings in the indexer.
func (s *autoscaleSettingLister) List(selector labels.Selector) (ret []*v1alpha1.AutoscaleSetting, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoscaleSetting))
	})
	return ret, err
}

// AutoscaleSettings returns an object that can list and get AutoscaleSettings.
func (s *autoscaleSettingLister) AutoscaleSettings(namespace string) AutoscaleSettingNamespaceLister {
	return autoscaleSettingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutoscaleSettingNamespaceLister helps list and get AutoscaleSettings.
type AutoscaleSettingNamespaceLister interface {
	// List lists all AutoscaleSettings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AutoscaleSetting, err error)
	// Get retrieves the AutoscaleSetting from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AutoscaleSetting, error)
	AutoscaleSettingNamespaceListerExpansion
}

// autoscaleSettingNamespaceLister implements the AutoscaleSettingNamespaceLister
// interface.
type autoscaleSettingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutoscaleSettings in the indexer for a given namespace.
func (s autoscaleSettingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutoscaleSetting, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoscaleSetting))
	})
	return ret, err
}

// Get retrieves the AutoscaleSetting from the indexer for a given namespace and name.
func (s autoscaleSettingNamespaceLister) Get(name string) (*v1alpha1.AutoscaleSetting, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("autoscalesetting"), name)
	}
	return obj.(*v1alpha1.AutoscaleSetting), nil
}
