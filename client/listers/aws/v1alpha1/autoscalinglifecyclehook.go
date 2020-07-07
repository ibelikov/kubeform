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

// AutoscalingLifecycleHookLister helps list AutoscalingLifecycleHooks.
type AutoscalingLifecycleHookLister interface {
	// List lists all AutoscalingLifecycleHooks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AutoscalingLifecycleHook, err error)
	// AutoscalingLifecycleHooks returns an object that can list and get AutoscalingLifecycleHooks.
	AutoscalingLifecycleHooks(namespace string) AutoscalingLifecycleHookNamespaceLister
	AutoscalingLifecycleHookListerExpansion
}

// autoscalingLifecycleHookLister implements the AutoscalingLifecycleHookLister interface.
type autoscalingLifecycleHookLister struct {
	indexer cache.Indexer
}

// NewAutoscalingLifecycleHookLister returns a new AutoscalingLifecycleHookLister.
func NewAutoscalingLifecycleHookLister(indexer cache.Indexer) AutoscalingLifecycleHookLister {
	return &autoscalingLifecycleHookLister{indexer: indexer}
}

// List lists all AutoscalingLifecycleHooks in the indexer.
func (s *autoscalingLifecycleHookLister) List(selector labels.Selector) (ret []*v1alpha1.AutoscalingLifecycleHook, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoscalingLifecycleHook))
	})
	return ret, err
}

// AutoscalingLifecycleHooks returns an object that can list and get AutoscalingLifecycleHooks.
func (s *autoscalingLifecycleHookLister) AutoscalingLifecycleHooks(namespace string) AutoscalingLifecycleHookNamespaceLister {
	return autoscalingLifecycleHookNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutoscalingLifecycleHookNamespaceLister helps list and get AutoscalingLifecycleHooks.
type AutoscalingLifecycleHookNamespaceLister interface {
	// List lists all AutoscalingLifecycleHooks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AutoscalingLifecycleHook, err error)
	// Get retrieves the AutoscalingLifecycleHook from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AutoscalingLifecycleHook, error)
	AutoscalingLifecycleHookNamespaceListerExpansion
}

// autoscalingLifecycleHookNamespaceLister implements the AutoscalingLifecycleHookNamespaceLister
// interface.
type autoscalingLifecycleHookNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutoscalingLifecycleHooks in the indexer for a given namespace.
func (s autoscalingLifecycleHookNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutoscalingLifecycleHook, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoscalingLifecycleHook))
	})
	return ret, err
}

// Get retrieves the AutoscalingLifecycleHook from the indexer for a given namespace and name.
func (s autoscalingLifecycleHookNamespaceLister) Get(name string) (*v1alpha1.AutoscalingLifecycleHook, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("autoscalinglifecyclehook"), name)
	}
	return obj.(*v1alpha1.AutoscalingLifecycleHook), nil
}
