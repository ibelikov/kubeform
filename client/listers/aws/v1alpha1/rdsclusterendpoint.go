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

// RdsClusterEndpointLister helps list RdsClusterEndpoints.
type RdsClusterEndpointLister interface {
	// List lists all RdsClusterEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RdsClusterEndpoint, err error)
	// RdsClusterEndpoints returns an object that can list and get RdsClusterEndpoints.
	RdsClusterEndpoints(namespace string) RdsClusterEndpointNamespaceLister
	RdsClusterEndpointListerExpansion
}

// rdsClusterEndpointLister implements the RdsClusterEndpointLister interface.
type rdsClusterEndpointLister struct {
	indexer cache.Indexer
}

// NewRdsClusterEndpointLister returns a new RdsClusterEndpointLister.
func NewRdsClusterEndpointLister(indexer cache.Indexer) RdsClusterEndpointLister {
	return &rdsClusterEndpointLister{indexer: indexer}
}

// List lists all RdsClusterEndpoints in the indexer.
func (s *rdsClusterEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.RdsClusterEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RdsClusterEndpoint))
	})
	return ret, err
}

// RdsClusterEndpoints returns an object that can list and get RdsClusterEndpoints.
func (s *rdsClusterEndpointLister) RdsClusterEndpoints(namespace string) RdsClusterEndpointNamespaceLister {
	return rdsClusterEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RdsClusterEndpointNamespaceLister helps list and get RdsClusterEndpoints.
type RdsClusterEndpointNamespaceLister interface {
	// List lists all RdsClusterEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RdsClusterEndpoint, err error)
	// Get retrieves the RdsClusterEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RdsClusterEndpoint, error)
	RdsClusterEndpointNamespaceListerExpansion
}

// rdsClusterEndpointNamespaceLister implements the RdsClusterEndpointNamespaceLister
// interface.
type rdsClusterEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RdsClusterEndpoints in the indexer for a given namespace.
func (s rdsClusterEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RdsClusterEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RdsClusterEndpoint))
	})
	return ret, err
}

// Get retrieves the RdsClusterEndpoint from the indexer for a given namespace and name.
func (s rdsClusterEndpointNamespaceLister) Get(name string) (*v1alpha1.RdsClusterEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rdsclusterendpoint"), name)
	}
	return obj.(*v1alpha1.RdsClusterEndpoint), nil
}
