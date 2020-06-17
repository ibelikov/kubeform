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

// ApiGatewayClientCertificateLister helps list ApiGatewayClientCertificates.
type ApiGatewayClientCertificateLister interface {
	// List lists all ApiGatewayClientCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayClientCertificate, err error)
	// ApiGatewayClientCertificates returns an object that can list and get ApiGatewayClientCertificates.
	ApiGatewayClientCertificates(namespace string) ApiGatewayClientCertificateNamespaceLister
	ApiGatewayClientCertificateListerExpansion
}

// apiGatewayClientCertificateLister implements the ApiGatewayClientCertificateLister interface.
type apiGatewayClientCertificateLister struct {
	indexer cache.Indexer
}

// NewApiGatewayClientCertificateLister returns a new ApiGatewayClientCertificateLister.
func NewApiGatewayClientCertificateLister(indexer cache.Indexer) ApiGatewayClientCertificateLister {
	return &apiGatewayClientCertificateLister{indexer: indexer}
}

// List lists all ApiGatewayClientCertificates in the indexer.
func (s *apiGatewayClientCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayClientCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayClientCertificate))
	})
	return ret, err
}

// ApiGatewayClientCertificates returns an object that can list and get ApiGatewayClientCertificates.
func (s *apiGatewayClientCertificateLister) ApiGatewayClientCertificates(namespace string) ApiGatewayClientCertificateNamespaceLister {
	return apiGatewayClientCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayClientCertificateNamespaceLister helps list and get ApiGatewayClientCertificates.
type ApiGatewayClientCertificateNamespaceLister interface {
	// List lists all ApiGatewayClientCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayClientCertificate, err error)
	// Get retrieves the ApiGatewayClientCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayClientCertificate, error)
	ApiGatewayClientCertificateNamespaceListerExpansion
}

// apiGatewayClientCertificateNamespaceLister implements the ApiGatewayClientCertificateNamespaceLister
// interface.
type apiGatewayClientCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayClientCertificates in the indexer for a given namespace.
func (s apiGatewayClientCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayClientCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayClientCertificate))
	})
	return ret, err
}

// Get retrieves the ApiGatewayClientCertificate from the indexer for a given namespace and name.
func (s apiGatewayClientCertificateNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayClientCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewayclientcertificate"), name)
	}
	return obj.(*v1alpha1.ApiGatewayClientCertificate), nil
}
