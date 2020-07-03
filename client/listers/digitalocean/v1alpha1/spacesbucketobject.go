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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SpacesBucketObjectLister helps list SpacesBucketObjects.
type SpacesBucketObjectLister interface {
	// List lists all SpacesBucketObjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SpacesBucketObject, err error)
	// SpacesBucketObjects returns an object that can list and get SpacesBucketObjects.
	SpacesBucketObjects(namespace string) SpacesBucketObjectNamespaceLister
	SpacesBucketObjectListerExpansion
}

// spacesBucketObjectLister implements the SpacesBucketObjectLister interface.
type spacesBucketObjectLister struct {
	indexer cache.Indexer
}

// NewSpacesBucketObjectLister returns a new SpacesBucketObjectLister.
func NewSpacesBucketObjectLister(indexer cache.Indexer) SpacesBucketObjectLister {
	return &spacesBucketObjectLister{indexer: indexer}
}

// List lists all SpacesBucketObjects in the indexer.
func (s *spacesBucketObjectLister) List(selector labels.Selector) (ret []*v1alpha1.SpacesBucketObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpacesBucketObject))
	})
	return ret, err
}

// SpacesBucketObjects returns an object that can list and get SpacesBucketObjects.
func (s *spacesBucketObjectLister) SpacesBucketObjects(namespace string) SpacesBucketObjectNamespaceLister {
	return spacesBucketObjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SpacesBucketObjectNamespaceLister helps list and get SpacesBucketObjects.
type SpacesBucketObjectNamespaceLister interface {
	// List lists all SpacesBucketObjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SpacesBucketObject, err error)
	// Get retrieves the SpacesBucketObject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SpacesBucketObject, error)
	SpacesBucketObjectNamespaceListerExpansion
}

// spacesBucketObjectNamespaceLister implements the SpacesBucketObjectNamespaceLister
// interface.
type spacesBucketObjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SpacesBucketObjects in the indexer for a given namespace.
func (s spacesBucketObjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SpacesBucketObject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpacesBucketObject))
	})
	return ret, err
}

// Get retrieves the SpacesBucketObject from the indexer for a given namespace and name.
func (s spacesBucketObjectNamespaceLister) Get(name string) (*v1alpha1.SpacesBucketObject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("spacesbucketobject"), name)
	}
	return obj.(*v1alpha1.SpacesBucketObject), nil
}