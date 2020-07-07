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

// AlbTargetGroupLister helps list AlbTargetGroups.
type AlbTargetGroupLister interface {
	// List lists all AlbTargetGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AlbTargetGroup, err error)
	// AlbTargetGroups returns an object that can list and get AlbTargetGroups.
	AlbTargetGroups(namespace string) AlbTargetGroupNamespaceLister
	AlbTargetGroupListerExpansion
}

// albTargetGroupLister implements the AlbTargetGroupLister interface.
type albTargetGroupLister struct {
	indexer cache.Indexer
}

// NewAlbTargetGroupLister returns a new AlbTargetGroupLister.
func NewAlbTargetGroupLister(indexer cache.Indexer) AlbTargetGroupLister {
	return &albTargetGroupLister{indexer: indexer}
}

// List lists all AlbTargetGroups in the indexer.
func (s *albTargetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AlbTargetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlbTargetGroup))
	})
	return ret, err
}

// AlbTargetGroups returns an object that can list and get AlbTargetGroups.
func (s *albTargetGroupLister) AlbTargetGroups(namespace string) AlbTargetGroupNamespaceLister {
	return albTargetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlbTargetGroupNamespaceLister helps list and get AlbTargetGroups.
type AlbTargetGroupNamespaceLister interface {
	// List lists all AlbTargetGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AlbTargetGroup, err error)
	// Get retrieves the AlbTargetGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AlbTargetGroup, error)
	AlbTargetGroupNamespaceListerExpansion
}

// albTargetGroupNamespaceLister implements the AlbTargetGroupNamespaceLister
// interface.
type albTargetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlbTargetGroups in the indexer for a given namespace.
func (s albTargetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AlbTargetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlbTargetGroup))
	})
	return ret, err
}

// Get retrieves the AlbTargetGroup from the indexer for a given namespace and name.
func (s albTargetGroupNamespaceLister) Get(name string) (*v1alpha1.AlbTargetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("albtargetgroup"), name)
	}
	return obj.(*v1alpha1.AlbTargetGroup), nil
}
