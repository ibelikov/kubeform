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

// DxPrivateVirtualInterfaceLister helps list DxPrivateVirtualInterfaces.
type DxPrivateVirtualInterfaceLister interface {
	// List lists all DxPrivateVirtualInterfaces in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DxPrivateVirtualInterface, err error)
	// DxPrivateVirtualInterfaces returns an object that can list and get DxPrivateVirtualInterfaces.
	DxPrivateVirtualInterfaces(namespace string) DxPrivateVirtualInterfaceNamespaceLister
	DxPrivateVirtualInterfaceListerExpansion
}

// dxPrivateVirtualInterfaceLister implements the DxPrivateVirtualInterfaceLister interface.
type dxPrivateVirtualInterfaceLister struct {
	indexer cache.Indexer
}

// NewDxPrivateVirtualInterfaceLister returns a new DxPrivateVirtualInterfaceLister.
func NewDxPrivateVirtualInterfaceLister(indexer cache.Indexer) DxPrivateVirtualInterfaceLister {
	return &dxPrivateVirtualInterfaceLister{indexer: indexer}
}

// List lists all DxPrivateVirtualInterfaces in the indexer.
func (s *dxPrivateVirtualInterfaceLister) List(selector labels.Selector) (ret []*v1alpha1.DxPrivateVirtualInterface, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxPrivateVirtualInterface))
	})
	return ret, err
}

// DxPrivateVirtualInterfaces returns an object that can list and get DxPrivateVirtualInterfaces.
func (s *dxPrivateVirtualInterfaceLister) DxPrivateVirtualInterfaces(namespace string) DxPrivateVirtualInterfaceNamespaceLister {
	return dxPrivateVirtualInterfaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DxPrivateVirtualInterfaceNamespaceLister helps list and get DxPrivateVirtualInterfaces.
type DxPrivateVirtualInterfaceNamespaceLister interface {
	// List lists all DxPrivateVirtualInterfaces in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DxPrivateVirtualInterface, err error)
	// Get retrieves the DxPrivateVirtualInterface from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DxPrivateVirtualInterface, error)
	DxPrivateVirtualInterfaceNamespaceListerExpansion
}

// dxPrivateVirtualInterfaceNamespaceLister implements the DxPrivateVirtualInterfaceNamespaceLister
// interface.
type dxPrivateVirtualInterfaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DxPrivateVirtualInterfaces in the indexer for a given namespace.
func (s dxPrivateVirtualInterfaceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DxPrivateVirtualInterface, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxPrivateVirtualInterface))
	})
	return ret, err
}

// Get retrieves the DxPrivateVirtualInterface from the indexer for a given namespace and name.
func (s dxPrivateVirtualInterfaceNamespaceLister) Get(name string) (*v1alpha1.DxPrivateVirtualInterface, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dxprivatevirtualinterface"), name)
	}
	return obj.(*v1alpha1.DxPrivateVirtualInterface), nil
}
