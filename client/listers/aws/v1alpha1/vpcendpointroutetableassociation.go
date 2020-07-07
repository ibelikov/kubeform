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

// VpcEndpointRouteTableAssociationLister helps list VpcEndpointRouteTableAssociations.
type VpcEndpointRouteTableAssociationLister interface {
	// List lists all VpcEndpointRouteTableAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointRouteTableAssociation, err error)
	// VpcEndpointRouteTableAssociations returns an object that can list and get VpcEndpointRouteTableAssociations.
	VpcEndpointRouteTableAssociations(namespace string) VpcEndpointRouteTableAssociationNamespaceLister
	VpcEndpointRouteTableAssociationListerExpansion
}

// vpcEndpointRouteTableAssociationLister implements the VpcEndpointRouteTableAssociationLister interface.
type vpcEndpointRouteTableAssociationLister struct {
	indexer cache.Indexer
}

// NewVpcEndpointRouteTableAssociationLister returns a new VpcEndpointRouteTableAssociationLister.
func NewVpcEndpointRouteTableAssociationLister(indexer cache.Indexer) VpcEndpointRouteTableAssociationLister {
	return &vpcEndpointRouteTableAssociationLister{indexer: indexer}
}

// List lists all VpcEndpointRouteTableAssociations in the indexer.
func (s *vpcEndpointRouteTableAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointRouteTableAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcEndpointRouteTableAssociation))
	})
	return ret, err
}

// VpcEndpointRouteTableAssociations returns an object that can list and get VpcEndpointRouteTableAssociations.
func (s *vpcEndpointRouteTableAssociationLister) VpcEndpointRouteTableAssociations(namespace string) VpcEndpointRouteTableAssociationNamespaceLister {
	return vpcEndpointRouteTableAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcEndpointRouteTableAssociationNamespaceLister helps list and get VpcEndpointRouteTableAssociations.
type VpcEndpointRouteTableAssociationNamespaceLister interface {
	// List lists all VpcEndpointRouteTableAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointRouteTableAssociation, err error)
	// Get retrieves the VpcEndpointRouteTableAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VpcEndpointRouteTableAssociation, error)
	VpcEndpointRouteTableAssociationNamespaceListerExpansion
}

// vpcEndpointRouteTableAssociationNamespaceLister implements the VpcEndpointRouteTableAssociationNamespaceLister
// interface.
type vpcEndpointRouteTableAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcEndpointRouteTableAssociations in the indexer for a given namespace.
func (s vpcEndpointRouteTableAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointRouteTableAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcEndpointRouteTableAssociation))
	})
	return ret, err
}

// Get retrieves the VpcEndpointRouteTableAssociation from the indexer for a given namespace and name.
func (s vpcEndpointRouteTableAssociationNamespaceLister) Get(name string) (*v1alpha1.VpcEndpointRouteTableAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcendpointroutetableassociation"), name)
	}
	return obj.(*v1alpha1.VpcEndpointRouteTableAssociation), nil
}
