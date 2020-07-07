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

// ElasticBeanstalkApplicationLister helps list ElasticBeanstalkApplications.
type ElasticBeanstalkApplicationLister interface {
	// List lists all ElasticBeanstalkApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticBeanstalkApplication, err error)
	// ElasticBeanstalkApplications returns an object that can list and get ElasticBeanstalkApplications.
	ElasticBeanstalkApplications(namespace string) ElasticBeanstalkApplicationNamespaceLister
	ElasticBeanstalkApplicationListerExpansion
}

// elasticBeanstalkApplicationLister implements the ElasticBeanstalkApplicationLister interface.
type elasticBeanstalkApplicationLister struct {
	indexer cache.Indexer
}

// NewElasticBeanstalkApplicationLister returns a new ElasticBeanstalkApplicationLister.
func NewElasticBeanstalkApplicationLister(indexer cache.Indexer) ElasticBeanstalkApplicationLister {
	return &elasticBeanstalkApplicationLister{indexer: indexer}
}

// List lists all ElasticBeanstalkApplications in the indexer.
func (s *elasticBeanstalkApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticBeanstalkApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticBeanstalkApplication))
	})
	return ret, err
}

// ElasticBeanstalkApplications returns an object that can list and get ElasticBeanstalkApplications.
func (s *elasticBeanstalkApplicationLister) ElasticBeanstalkApplications(namespace string) ElasticBeanstalkApplicationNamespaceLister {
	return elasticBeanstalkApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElasticBeanstalkApplicationNamespaceLister helps list and get ElasticBeanstalkApplications.
type ElasticBeanstalkApplicationNamespaceLister interface {
	// List lists all ElasticBeanstalkApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticBeanstalkApplication, err error)
	// Get retrieves the ElasticBeanstalkApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ElasticBeanstalkApplication, error)
	ElasticBeanstalkApplicationNamespaceListerExpansion
}

// elasticBeanstalkApplicationNamespaceLister implements the ElasticBeanstalkApplicationNamespaceLister
// interface.
type elasticBeanstalkApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElasticBeanstalkApplications in the indexer for a given namespace.
func (s elasticBeanstalkApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticBeanstalkApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticBeanstalkApplication))
	})
	return ret, err
}

// Get retrieves the ElasticBeanstalkApplication from the indexer for a given namespace and name.
func (s elasticBeanstalkApplicationNamespaceLister) Get(name string) (*v1alpha1.ElasticBeanstalkApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elasticbeanstalkapplication"), name)
	}
	return obj.(*v1alpha1.ElasticBeanstalkApplication), nil
}
