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

// IamUserLister helps list IamUsers.
type IamUserLister interface {
	// List lists all IamUsers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamUser, err error)
	// IamUsers returns an object that can list and get IamUsers.
	IamUsers(namespace string) IamUserNamespaceLister
	IamUserListerExpansion
}

// iamUserLister implements the IamUserLister interface.
type iamUserLister struct {
	indexer cache.Indexer
}

// NewIamUserLister returns a new IamUserLister.
func NewIamUserLister(indexer cache.Indexer) IamUserLister {
	return &iamUserLister{indexer: indexer}
}

// List lists all IamUsers in the indexer.
func (s *iamUserLister) List(selector labels.Selector) (ret []*v1alpha1.IamUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUser))
	})
	return ret, err
}

// IamUsers returns an object that can list and get IamUsers.
func (s *iamUserLister) IamUsers(namespace string) IamUserNamespaceLister {
	return iamUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamUserNamespaceLister helps list and get IamUsers.
type IamUserNamespaceLister interface {
	// List lists all IamUsers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamUser, err error)
	// Get retrieves the IamUser from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamUser, error)
	IamUserNamespaceListerExpansion
}

// iamUserNamespaceLister implements the IamUserNamespaceLister
// interface.
type iamUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamUsers in the indexer for a given namespace.
func (s iamUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUser))
	})
	return ret, err
}

// Get retrieves the IamUser from the indexer for a given namespace and name.
func (s iamUserNamespaceLister) Get(name string) (*v1alpha1.IamUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamuser"), name)
	}
	return obj.(*v1alpha1.IamUser), nil
}
