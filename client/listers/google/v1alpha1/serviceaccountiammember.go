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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceAccountIamMemberLister helps list ServiceAccountIamMembers.
type ServiceAccountIamMemberLister interface {
	// List lists all ServiceAccountIamMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountIamMember, err error)
	// ServiceAccountIamMembers returns an object that can list and get ServiceAccountIamMembers.
	ServiceAccountIamMembers(namespace string) ServiceAccountIamMemberNamespaceLister
	ServiceAccountIamMemberListerExpansion
}

// serviceAccountIamMemberLister implements the ServiceAccountIamMemberLister interface.
type serviceAccountIamMemberLister struct {
	indexer cache.Indexer
}

// NewServiceAccountIamMemberLister returns a new ServiceAccountIamMemberLister.
func NewServiceAccountIamMemberLister(indexer cache.Indexer) ServiceAccountIamMemberLister {
	return &serviceAccountIamMemberLister{indexer: indexer}
}

// List lists all ServiceAccountIamMembers in the indexer.
func (s *serviceAccountIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceAccountIamMember))
	})
	return ret, err
}

// ServiceAccountIamMembers returns an object that can list and get ServiceAccountIamMembers.
func (s *serviceAccountIamMemberLister) ServiceAccountIamMembers(namespace string) ServiceAccountIamMemberNamespaceLister {
	return serviceAccountIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceAccountIamMemberNamespaceLister helps list and get ServiceAccountIamMembers.
type ServiceAccountIamMemberNamespaceLister interface {
	// List lists all ServiceAccountIamMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountIamMember, err error)
	// Get retrieves the ServiceAccountIamMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServiceAccountIamMember, error)
	ServiceAccountIamMemberNamespaceListerExpansion
}

// serviceAccountIamMemberNamespaceLister implements the ServiceAccountIamMemberNamespaceLister
// interface.
type serviceAccountIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceAccountIamMembers in the indexer for a given namespace.
func (s serviceAccountIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceAccountIamMember))
	})
	return ret, err
}

// Get retrieves the ServiceAccountIamMember from the indexer for a given namespace and name.
func (s serviceAccountIamMemberNamespaceLister) Get(name string) (*v1alpha1.ServiceAccountIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serviceaccountiammember"), name)
	}
	return obj.(*v1alpha1.ServiceAccountIamMember), nil
}
