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

// IamUserSSHKeyLister helps list IamUserSSHKeys.
type IamUserSSHKeyLister interface {
	// List lists all IamUserSSHKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserSSHKey, err error)
	// IamUserSSHKeys returns an object that can list and get IamUserSSHKeys.
	IamUserSSHKeys(namespace string) IamUserSSHKeyNamespaceLister
	IamUserSSHKeyListerExpansion
}

// iamUserSSHKeyLister implements the IamUserSSHKeyLister interface.
type iamUserSSHKeyLister struct {
	indexer cache.Indexer
}

// NewIamUserSSHKeyLister returns a new IamUserSSHKeyLister.
func NewIamUserSSHKeyLister(indexer cache.Indexer) IamUserSSHKeyLister {
	return &iamUserSSHKeyLister{indexer: indexer}
}

// List lists all IamUserSSHKeys in the indexer.
func (s *iamUserSSHKeyLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserSSHKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserSSHKey))
	})
	return ret, err
}

// IamUserSSHKeys returns an object that can list and get IamUserSSHKeys.
func (s *iamUserSSHKeyLister) IamUserSSHKeys(namespace string) IamUserSSHKeyNamespaceLister {
	return iamUserSSHKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamUserSSHKeyNamespaceLister helps list and get IamUserSSHKeys.
type IamUserSSHKeyNamespaceLister interface {
	// List lists all IamUserSSHKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserSSHKey, err error)
	// Get retrieves the IamUserSSHKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamUserSSHKey, error)
	IamUserSSHKeyNamespaceListerExpansion
}

// iamUserSSHKeyNamespaceLister implements the IamUserSSHKeyNamespaceLister
// interface.
type iamUserSSHKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamUserSSHKeys in the indexer for a given namespace.
func (s iamUserSSHKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserSSHKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserSSHKey))
	})
	return ret, err
}

// Get retrieves the IamUserSSHKey from the indexer for a given namespace and name.
func (s iamUserSSHKeyNamespaceLister) Get(name string) (*v1alpha1.IamUserSSHKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamusersshkey"), name)
	}
	return obj.(*v1alpha1.IamUserSSHKey), nil
}
