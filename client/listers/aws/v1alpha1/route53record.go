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

// Route53RecordLister helps list Route53Records.
type Route53RecordLister interface {
	// List lists all Route53Records in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Route53Record, err error)
	// Route53Records returns an object that can list and get Route53Records.
	Route53Records(namespace string) Route53RecordNamespaceLister
	Route53RecordListerExpansion
}

// route53RecordLister implements the Route53RecordLister interface.
type route53RecordLister struct {
	indexer cache.Indexer
}

// NewRoute53RecordLister returns a new Route53RecordLister.
func NewRoute53RecordLister(indexer cache.Indexer) Route53RecordLister {
	return &route53RecordLister{indexer: indexer}
}

// List lists all Route53Records in the indexer.
func (s *route53RecordLister) List(selector labels.Selector) (ret []*v1alpha1.Route53Record, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route53Record))
	})
	return ret, err
}

// Route53Records returns an object that can list and get Route53Records.
func (s *route53RecordLister) Route53Records(namespace string) Route53RecordNamespaceLister {
	return route53RecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Route53RecordNamespaceLister helps list and get Route53Records.
type Route53RecordNamespaceLister interface {
	// List lists all Route53Records in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Route53Record, err error)
	// Get retrieves the Route53Record from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Route53Record, error)
	Route53RecordNamespaceListerExpansion
}

// route53RecordNamespaceLister implements the Route53RecordNamespaceLister
// interface.
type route53RecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Route53Records in the indexer for a given namespace.
func (s route53RecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Route53Record, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route53Record))
	})
	return ret, err
}

// Get retrieves the Route53Record from the indexer for a given namespace and name.
func (s route53RecordNamespaceLister) Get(name string) (*v1alpha1.Route53Record, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("route53record"), name)
	}
	return obj.(*v1alpha1.Route53Record), nil
}
