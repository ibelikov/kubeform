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

// CodebuildSourceCredentialLister helps list CodebuildSourceCredentials.
type CodebuildSourceCredentialLister interface {
	// List lists all CodebuildSourceCredentials in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CodebuildSourceCredential, err error)
	// CodebuildSourceCredentials returns an object that can list and get CodebuildSourceCredentials.
	CodebuildSourceCredentials(namespace string) CodebuildSourceCredentialNamespaceLister
	CodebuildSourceCredentialListerExpansion
}

// codebuildSourceCredentialLister implements the CodebuildSourceCredentialLister interface.
type codebuildSourceCredentialLister struct {
	indexer cache.Indexer
}

// NewCodebuildSourceCredentialLister returns a new CodebuildSourceCredentialLister.
func NewCodebuildSourceCredentialLister(indexer cache.Indexer) CodebuildSourceCredentialLister {
	return &codebuildSourceCredentialLister{indexer: indexer}
}

// List lists all CodebuildSourceCredentials in the indexer.
func (s *codebuildSourceCredentialLister) List(selector labels.Selector) (ret []*v1alpha1.CodebuildSourceCredential, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodebuildSourceCredential))
	})
	return ret, err
}

// CodebuildSourceCredentials returns an object that can list and get CodebuildSourceCredentials.
func (s *codebuildSourceCredentialLister) CodebuildSourceCredentials(namespace string) CodebuildSourceCredentialNamespaceLister {
	return codebuildSourceCredentialNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CodebuildSourceCredentialNamespaceLister helps list and get CodebuildSourceCredentials.
type CodebuildSourceCredentialNamespaceLister interface {
	// List lists all CodebuildSourceCredentials in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CodebuildSourceCredential, err error)
	// Get retrieves the CodebuildSourceCredential from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CodebuildSourceCredential, error)
	CodebuildSourceCredentialNamespaceListerExpansion
}

// codebuildSourceCredentialNamespaceLister implements the CodebuildSourceCredentialNamespaceLister
// interface.
type codebuildSourceCredentialNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CodebuildSourceCredentials in the indexer for a given namespace.
func (s codebuildSourceCredentialNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CodebuildSourceCredential, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodebuildSourceCredential))
	})
	return ret, err
}

// Get retrieves the CodebuildSourceCredential from the indexer for a given namespace and name.
func (s codebuildSourceCredentialNamespaceLister) Get(name string) (*v1alpha1.CodebuildSourceCredential, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("codebuildsourcecredential"), name)
	}
	return obj.(*v1alpha1.CodebuildSourceCredential), nil
}
