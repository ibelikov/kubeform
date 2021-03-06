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

// DbSecurityGroupLister helps list DbSecurityGroups.
type DbSecurityGroupLister interface {
	// List lists all DbSecurityGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DbSecurityGroup, err error)
	// DbSecurityGroups returns an object that can list and get DbSecurityGroups.
	DbSecurityGroups(namespace string) DbSecurityGroupNamespaceLister
	DbSecurityGroupListerExpansion
}

// dbSecurityGroupLister implements the DbSecurityGroupLister interface.
type dbSecurityGroupLister struct {
	indexer cache.Indexer
}

// NewDbSecurityGroupLister returns a new DbSecurityGroupLister.
func NewDbSecurityGroupLister(indexer cache.Indexer) DbSecurityGroupLister {
	return &dbSecurityGroupLister{indexer: indexer}
}

// List lists all DbSecurityGroups in the indexer.
func (s *dbSecurityGroupLister) List(selector labels.Selector) (ret []*v1alpha1.DbSecurityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbSecurityGroup))
	})
	return ret, err
}

// DbSecurityGroups returns an object that can list and get DbSecurityGroups.
func (s *dbSecurityGroupLister) DbSecurityGroups(namespace string) DbSecurityGroupNamespaceLister {
	return dbSecurityGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbSecurityGroupNamespaceLister helps list and get DbSecurityGroups.
type DbSecurityGroupNamespaceLister interface {
	// List lists all DbSecurityGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DbSecurityGroup, err error)
	// Get retrieves the DbSecurityGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DbSecurityGroup, error)
	DbSecurityGroupNamespaceListerExpansion
}

// dbSecurityGroupNamespaceLister implements the DbSecurityGroupNamespaceLister
// interface.
type dbSecurityGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbSecurityGroups in the indexer for a given namespace.
func (s dbSecurityGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DbSecurityGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbSecurityGroup))
	})
	return ret, err
}

// Get retrieves the DbSecurityGroup from the indexer for a given namespace and name.
func (s dbSecurityGroupNamespaceLister) Get(name string) (*v1alpha1.DbSecurityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dbsecuritygroup"), name)
	}
	return obj.(*v1alpha1.DbSecurityGroup), nil
}
