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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IapWebIamMemberLister helps list IapWebIamMembers.
type IapWebIamMemberLister interface {
	// List lists all IapWebIamMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IapWebIamMember, err error)
	// IapWebIamMembers returns an object that can list and get IapWebIamMembers.
	IapWebIamMembers(namespace string) IapWebIamMemberNamespaceLister
	IapWebIamMemberListerExpansion
}

// iapWebIamMemberLister implements the IapWebIamMemberLister interface.
type iapWebIamMemberLister struct {
	indexer cache.Indexer
}

// NewIapWebIamMemberLister returns a new IapWebIamMemberLister.
func NewIapWebIamMemberLister(indexer cache.Indexer) IapWebIamMemberLister {
	return &iapWebIamMemberLister{indexer: indexer}
}

// List lists all IapWebIamMembers in the indexer.
func (s *iapWebIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.IapWebIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IapWebIamMember))
	})
	return ret, err
}

// IapWebIamMembers returns an object that can list and get IapWebIamMembers.
func (s *iapWebIamMemberLister) IapWebIamMembers(namespace string) IapWebIamMemberNamespaceLister {
	return iapWebIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IapWebIamMemberNamespaceLister helps list and get IapWebIamMembers.
type IapWebIamMemberNamespaceLister interface {
	// List lists all IapWebIamMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IapWebIamMember, err error)
	// Get retrieves the IapWebIamMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IapWebIamMember, error)
	IapWebIamMemberNamespaceListerExpansion
}

// iapWebIamMemberNamespaceLister implements the IapWebIamMemberNamespaceLister
// interface.
type iapWebIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IapWebIamMembers in the indexer for a given namespace.
func (s iapWebIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IapWebIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IapWebIamMember))
	})
	return ret, err
}

// Get retrieves the IapWebIamMember from the indexer for a given namespace and name.
func (s iapWebIamMemberNamespaceLister) Get(name string) (*v1alpha1.IapWebIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iapwebiammember"), name)
	}
	return obj.(*v1alpha1.IapWebIamMember), nil
}