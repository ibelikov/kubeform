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

// SesTemplateLister helps list SesTemplates.
type SesTemplateLister interface {
	// List lists all SesTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SesTemplate, err error)
	// SesTemplates returns an object that can list and get SesTemplates.
	SesTemplates(namespace string) SesTemplateNamespaceLister
	SesTemplateListerExpansion
}

// sesTemplateLister implements the SesTemplateLister interface.
type sesTemplateLister struct {
	indexer cache.Indexer
}

// NewSesTemplateLister returns a new SesTemplateLister.
func NewSesTemplateLister(indexer cache.Indexer) SesTemplateLister {
	return &sesTemplateLister{indexer: indexer}
}

// List lists all SesTemplates in the indexer.
func (s *sesTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.SesTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SesTemplate))
	})
	return ret, err
}

// SesTemplates returns an object that can list and get SesTemplates.
func (s *sesTemplateLister) SesTemplates(namespace string) SesTemplateNamespaceLister {
	return sesTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SesTemplateNamespaceLister helps list and get SesTemplates.
type SesTemplateNamespaceLister interface {
	// List lists all SesTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SesTemplate, err error)
	// Get retrieves the SesTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SesTemplate, error)
	SesTemplateNamespaceListerExpansion
}

// sesTemplateNamespaceLister implements the SesTemplateNamespaceLister
// interface.
type sesTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SesTemplates in the indexer for a given namespace.
func (s sesTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SesTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SesTemplate))
	})
	return ret, err
}

// Get retrieves the SesTemplate from the indexer for a given namespace and name.
func (s sesTemplateNamespaceLister) Get(name string) (*v1alpha1.SesTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sestemplate"), name)
	}
	return obj.(*v1alpha1.SesTemplate), nil
}
