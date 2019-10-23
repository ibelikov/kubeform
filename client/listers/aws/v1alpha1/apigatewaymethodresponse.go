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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApiGatewayMethodResponseLister helps list ApiGatewayMethodResponses.
type ApiGatewayMethodResponseLister interface {
	// List lists all ApiGatewayMethodResponses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodResponse, err error)
	// ApiGatewayMethodResponses returns an object that can list and get ApiGatewayMethodResponses.
	ApiGatewayMethodResponses(namespace string) ApiGatewayMethodResponseNamespaceLister
	ApiGatewayMethodResponseListerExpansion
}

// apiGatewayMethodResponseLister implements the ApiGatewayMethodResponseLister interface.
type apiGatewayMethodResponseLister struct {
	indexer cache.Indexer
}

// NewApiGatewayMethodResponseLister returns a new ApiGatewayMethodResponseLister.
func NewApiGatewayMethodResponseLister(indexer cache.Indexer) ApiGatewayMethodResponseLister {
	return &apiGatewayMethodResponseLister{indexer: indexer}
}

// List lists all ApiGatewayMethodResponses in the indexer.
func (s *apiGatewayMethodResponseLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodResponse, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayMethodResponse))
	})
	return ret, err
}

// ApiGatewayMethodResponses returns an object that can list and get ApiGatewayMethodResponses.
func (s *apiGatewayMethodResponseLister) ApiGatewayMethodResponses(namespace string) ApiGatewayMethodResponseNamespaceLister {
	return apiGatewayMethodResponseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayMethodResponseNamespaceLister helps list and get ApiGatewayMethodResponses.
type ApiGatewayMethodResponseNamespaceLister interface {
	// List lists all ApiGatewayMethodResponses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodResponse, err error)
	// Get retrieves the ApiGatewayMethodResponse from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayMethodResponse, error)
	ApiGatewayMethodResponseNamespaceListerExpansion
}

// apiGatewayMethodResponseNamespaceLister implements the ApiGatewayMethodResponseNamespaceLister
// interface.
type apiGatewayMethodResponseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayMethodResponses in the indexer for a given namespace.
func (s apiGatewayMethodResponseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodResponse, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayMethodResponse))
	})
	return ret, err
}

// Get retrieves the ApiGatewayMethodResponse from the indexer for a given namespace and name.
func (s apiGatewayMethodResponseNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayMethodResponse, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewaymethodresponse"), name)
	}
	return obj.(*v1alpha1.ApiGatewayMethodResponse), nil
}
