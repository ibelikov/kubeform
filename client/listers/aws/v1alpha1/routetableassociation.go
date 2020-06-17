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

// RouteTableAssociationLister helps list RouteTableAssociations.
type RouteTableAssociationLister interface {
	// List lists all RouteTableAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RouteTableAssociation, err error)
	// RouteTableAssociations returns an object that can list and get RouteTableAssociations.
	RouteTableAssociations(namespace string) RouteTableAssociationNamespaceLister
	RouteTableAssociationListerExpansion
}

// routeTableAssociationLister implements the RouteTableAssociationLister interface.
type routeTableAssociationLister struct {
	indexer cache.Indexer
}

// NewRouteTableAssociationLister returns a new RouteTableAssociationLister.
func NewRouteTableAssociationLister(indexer cache.Indexer) RouteTableAssociationLister {
	return &routeTableAssociationLister{indexer: indexer}
}

// List lists all RouteTableAssociations in the indexer.
func (s *routeTableAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.RouteTableAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RouteTableAssociation))
	})
	return ret, err
}

// RouteTableAssociations returns an object that can list and get RouteTableAssociations.
func (s *routeTableAssociationLister) RouteTableAssociations(namespace string) RouteTableAssociationNamespaceLister {
	return routeTableAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RouteTableAssociationNamespaceLister helps list and get RouteTableAssociations.
type RouteTableAssociationNamespaceLister interface {
	// List lists all RouteTableAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RouteTableAssociation, err error)
	// Get retrieves the RouteTableAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RouteTableAssociation, error)
	RouteTableAssociationNamespaceListerExpansion
}

// routeTableAssociationNamespaceLister implements the RouteTableAssociationNamespaceLister
// interface.
type routeTableAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RouteTableAssociations in the indexer for a given namespace.
func (s routeTableAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RouteTableAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RouteTableAssociation))
	})
	return ret, err
}

// Get retrieves the RouteTableAssociation from the indexer for a given namespace and name.
func (s routeTableAssociationNamespaceLister) Get(name string) (*v1alpha1.RouteTableAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("routetableassociation"), name)
	}
	return obj.(*v1alpha1.RouteTableAssociation), nil
}
