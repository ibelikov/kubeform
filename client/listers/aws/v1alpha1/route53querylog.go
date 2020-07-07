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

// Route53QueryLogLister helps list Route53QueryLogs.
type Route53QueryLogLister interface {
	// List lists all Route53QueryLogs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Route53QueryLog, err error)
	// Route53QueryLogs returns an object that can list and get Route53QueryLogs.
	Route53QueryLogs(namespace string) Route53QueryLogNamespaceLister
	Route53QueryLogListerExpansion
}

// route53QueryLogLister implements the Route53QueryLogLister interface.
type route53QueryLogLister struct {
	indexer cache.Indexer
}

// NewRoute53QueryLogLister returns a new Route53QueryLogLister.
func NewRoute53QueryLogLister(indexer cache.Indexer) Route53QueryLogLister {
	return &route53QueryLogLister{indexer: indexer}
}

// List lists all Route53QueryLogs in the indexer.
func (s *route53QueryLogLister) List(selector labels.Selector) (ret []*v1alpha1.Route53QueryLog, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route53QueryLog))
	})
	return ret, err
}

// Route53QueryLogs returns an object that can list and get Route53QueryLogs.
func (s *route53QueryLogLister) Route53QueryLogs(namespace string) Route53QueryLogNamespaceLister {
	return route53QueryLogNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Route53QueryLogNamespaceLister helps list and get Route53QueryLogs.
type Route53QueryLogNamespaceLister interface {
	// List lists all Route53QueryLogs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Route53QueryLog, err error)
	// Get retrieves the Route53QueryLog from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Route53QueryLog, error)
	Route53QueryLogNamespaceListerExpansion
}

// route53QueryLogNamespaceLister implements the Route53QueryLogNamespaceLister
// interface.
type route53QueryLogNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Route53QueryLogs in the indexer for a given namespace.
func (s route53QueryLogNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Route53QueryLog, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route53QueryLog))
	})
	return ret, err
}

// Get retrieves the Route53QueryLog from the indexer for a given namespace and name.
func (s route53QueryLogNamespaceLister) Get(name string) (*v1alpha1.Route53QueryLog, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("route53querylog"), name)
	}
	return obj.(*v1alpha1.Route53QueryLog), nil
}
