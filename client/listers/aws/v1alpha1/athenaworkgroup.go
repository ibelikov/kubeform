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

// AthenaWorkgroupLister helps list AthenaWorkgroups.
type AthenaWorkgroupLister interface {
	// List lists all AthenaWorkgroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AthenaWorkgroup, err error)
	// AthenaWorkgroups returns an object that can list and get AthenaWorkgroups.
	AthenaWorkgroups(namespace string) AthenaWorkgroupNamespaceLister
	AthenaWorkgroupListerExpansion
}

// athenaWorkgroupLister implements the AthenaWorkgroupLister interface.
type athenaWorkgroupLister struct {
	indexer cache.Indexer
}

// NewAthenaWorkgroupLister returns a new AthenaWorkgroupLister.
func NewAthenaWorkgroupLister(indexer cache.Indexer) AthenaWorkgroupLister {
	return &athenaWorkgroupLister{indexer: indexer}
}

// List lists all AthenaWorkgroups in the indexer.
func (s *athenaWorkgroupLister) List(selector labels.Selector) (ret []*v1alpha1.AthenaWorkgroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AthenaWorkgroup))
	})
	return ret, err
}

// AthenaWorkgroups returns an object that can list and get AthenaWorkgroups.
func (s *athenaWorkgroupLister) AthenaWorkgroups(namespace string) AthenaWorkgroupNamespaceLister {
	return athenaWorkgroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AthenaWorkgroupNamespaceLister helps list and get AthenaWorkgroups.
type AthenaWorkgroupNamespaceLister interface {
	// List lists all AthenaWorkgroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AthenaWorkgroup, err error)
	// Get retrieves the AthenaWorkgroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AthenaWorkgroup, error)
	AthenaWorkgroupNamespaceListerExpansion
}

// athenaWorkgroupNamespaceLister implements the AthenaWorkgroupNamespaceLister
// interface.
type athenaWorkgroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AthenaWorkgroups in the indexer for a given namespace.
func (s athenaWorkgroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AthenaWorkgroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AthenaWorkgroup))
	})
	return ret, err
}

// Get retrieves the AthenaWorkgroup from the indexer for a given namespace and name.
func (s athenaWorkgroupNamespaceLister) Get(name string) (*v1alpha1.AthenaWorkgroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("athenaworkgroup"), name)
	}
	return obj.(*v1alpha1.AthenaWorkgroup), nil
}
