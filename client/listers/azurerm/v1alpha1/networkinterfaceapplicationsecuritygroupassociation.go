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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// NetworkInterfaceApplicationSecurityGroupAssociationLister helps list NetworkInterfaceApplicationSecurityGroupAssociations.
type NetworkInterfaceApplicationSecurityGroupAssociationLister interface {
	// List lists all NetworkInterfaceApplicationSecurityGroupAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error)
	// NetworkInterfaceApplicationSecurityGroupAssociations returns an object that can list and get NetworkInterfaceApplicationSecurityGroupAssociations.
	NetworkInterfaceApplicationSecurityGroupAssociations(namespace string) NetworkInterfaceApplicationSecurityGroupAssociationNamespaceLister
	NetworkInterfaceApplicationSecurityGroupAssociationListerExpansion
}

// networkInterfaceApplicationSecurityGroupAssociationLister implements the NetworkInterfaceApplicationSecurityGroupAssociationLister interface.
type networkInterfaceApplicationSecurityGroupAssociationLister struct {
	indexer cache.Indexer
}

// NewNetworkInterfaceApplicationSecurityGroupAssociationLister returns a new NetworkInterfaceApplicationSecurityGroupAssociationLister.
func NewNetworkInterfaceApplicationSecurityGroupAssociationLister(indexer cache.Indexer) NetworkInterfaceApplicationSecurityGroupAssociationLister {
	return &networkInterfaceApplicationSecurityGroupAssociationLister{indexer: indexer}
}

// List lists all NetworkInterfaceApplicationSecurityGroupAssociations in the indexer.
func (s *networkInterfaceApplicationSecurityGroupAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation))
	})
	return ret, err
}

// NetworkInterfaceApplicationSecurityGroupAssociations returns an object that can list and get NetworkInterfaceApplicationSecurityGroupAssociations.
func (s *networkInterfaceApplicationSecurityGroupAssociationLister) NetworkInterfaceApplicationSecurityGroupAssociations(namespace string) NetworkInterfaceApplicationSecurityGroupAssociationNamespaceLister {
	return networkInterfaceApplicationSecurityGroupAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkInterfaceApplicationSecurityGroupAssociationNamespaceLister helps list and get NetworkInterfaceApplicationSecurityGroupAssociations.
type NetworkInterfaceApplicationSecurityGroupAssociationNamespaceLister interface {
	// List lists all NetworkInterfaceApplicationSecurityGroupAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error)
	// Get retrieves the NetworkInterfaceApplicationSecurityGroupAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, error)
	NetworkInterfaceApplicationSecurityGroupAssociationNamespaceListerExpansion
}

// networkInterfaceApplicationSecurityGroupAssociationNamespaceLister implements the NetworkInterfaceApplicationSecurityGroupAssociationNamespaceLister
// interface.
type networkInterfaceApplicationSecurityGroupAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkInterfaceApplicationSecurityGroupAssociations in the indexer for a given namespace.
func (s networkInterfaceApplicationSecurityGroupAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation))
	})
	return ret, err
}

// Get retrieves the NetworkInterfaceApplicationSecurityGroupAssociation from the indexer for a given namespace and name.
func (s networkInterfaceApplicationSecurityGroupAssociationNamespaceLister) Get(name string) (*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkinterfaceapplicationsecuritygroupassociation"), name)
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationSecurityGroupAssociation), nil
}
