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

// IotDpsCertificateLister helps list IotDpsCertificates.
type IotDpsCertificateLister interface {
	// List lists all IotDpsCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IotDpsCertificate, err error)
	// IotDpsCertificates returns an object that can list and get IotDpsCertificates.
	IotDpsCertificates(namespace string) IotDpsCertificateNamespaceLister
	IotDpsCertificateListerExpansion
}

// iotDpsCertificateLister implements the IotDpsCertificateLister interface.
type iotDpsCertificateLister struct {
	indexer cache.Indexer
}

// NewIotDpsCertificateLister returns a new IotDpsCertificateLister.
func NewIotDpsCertificateLister(indexer cache.Indexer) IotDpsCertificateLister {
	return &iotDpsCertificateLister{indexer: indexer}
}

// List lists all IotDpsCertificates in the indexer.
func (s *iotDpsCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.IotDpsCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotDpsCertificate))
	})
	return ret, err
}

// IotDpsCertificates returns an object that can list and get IotDpsCertificates.
func (s *iotDpsCertificateLister) IotDpsCertificates(namespace string) IotDpsCertificateNamespaceLister {
	return iotDpsCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IotDpsCertificateNamespaceLister helps list and get IotDpsCertificates.
type IotDpsCertificateNamespaceLister interface {
	// List lists all IotDpsCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IotDpsCertificate, err error)
	// Get retrieves the IotDpsCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IotDpsCertificate, error)
	IotDpsCertificateNamespaceListerExpansion
}

// iotDpsCertificateNamespaceLister implements the IotDpsCertificateNamespaceLister
// interface.
type iotDpsCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IotDpsCertificates in the indexer for a given namespace.
func (s iotDpsCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IotDpsCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotDpsCertificate))
	})
	return ret, err
}

// Get retrieves the IotDpsCertificate from the indexer for a given namespace and name.
func (s iotDpsCertificateNamespaceLister) Get(name string) (*v1alpha1.IotDpsCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iotdpscertificate"), name)
	}
	return obj.(*v1alpha1.IotDpsCertificate), nil
}
