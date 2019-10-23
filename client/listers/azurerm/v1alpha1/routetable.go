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

// RouteTableLister helps list RouteTables.
type RouteTableLister interface {
	// List lists all RouteTables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RouteTable, err error)
	// RouteTables returns an object that can list and get RouteTables.
	RouteTables(namespace string) RouteTableNamespaceLister
	RouteTableListerExpansion
}

// routeTableLister implements the RouteTableLister interface.
type routeTableLister struct {
	indexer cache.Indexer
}

// NewRouteTableLister returns a new RouteTableLister.
func NewRouteTableLister(indexer cache.Indexer) RouteTableLister {
	return &routeTableLister{indexer: indexer}
}

// List lists all RouteTables in the indexer.
func (s *routeTableLister) List(selector labels.Selector) (ret []*v1alpha1.RouteTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RouteTable))
	})
	return ret, err
}

// RouteTables returns an object that can list and get RouteTables.
func (s *routeTableLister) RouteTables(namespace string) RouteTableNamespaceLister {
	return routeTableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RouteTableNamespaceLister helps list and get RouteTables.
type RouteTableNamespaceLister interface {
	// List lists all RouteTables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RouteTable, err error)
	// Get retrieves the RouteTable from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RouteTable, error)
	RouteTableNamespaceListerExpansion
}

// routeTableNamespaceLister implements the RouteTableNamespaceLister
// interface.
type routeTableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RouteTables in the indexer for a given namespace.
func (s routeTableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RouteTable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RouteTable))
	})
	return ret, err
}

// Get retrieves the RouteTable from the indexer for a given namespace and name.
func (s routeTableNamespaceLister) Get(name string) (*v1alpha1.RouteTable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("routetable"), name)
	}
	return obj.(*v1alpha1.RouteTable), nil
}
