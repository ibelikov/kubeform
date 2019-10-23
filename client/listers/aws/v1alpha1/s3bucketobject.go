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

// S3BucketObjectLister helps list S3BucketObjects.
type S3BucketObjectLister interface {
	// List lists all S3BucketObjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.S3BucketObject, err error)
	// S3BucketObjects returns an object that can list and get S3BucketObjects.
	S3BucketObjects(namespace string) S3BucketObjectNamespaceLister
	S3BucketObjectListerExpansion
}

// s3BucketObjectLister implements the S3BucketObjectLister interface.
type s3BucketObjectLister struct {
	indexer cache.Indexer
}

// NewS3BucketObjectLister returns a new S3BucketObjectLister.
func NewS3BucketObjectLister(indexer cache.Indexer) S3BucketObjectLister {
	return &s3BucketObjectLister{indexer: indexer}
}

// List lists all S3BucketObjects in the indexer.
func (s *s3BucketObjectLister) List(selector labels.Selector) (ret []*v1alpha1.S3BucketObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3BucketObject))
	})
	return ret, err
}

// S3BucketObjects returns an object that can list and get S3BucketObjects.
func (s *s3BucketObjectLister) S3BucketObjects(namespace string) S3BucketObjectNamespaceLister {
	return s3BucketObjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// S3BucketObjectNamespaceLister helps list and get S3BucketObjects.
type S3BucketObjectNamespaceLister interface {
	// List lists all S3BucketObjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.S3BucketObject, err error)
	// Get retrieves the S3BucketObject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.S3BucketObject, error)
	S3BucketObjectNamespaceListerExpansion
}

// s3BucketObjectNamespaceLister implements the S3BucketObjectNamespaceLister
// interface.
type s3BucketObjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all S3BucketObjects in the indexer for a given namespace.
func (s s3BucketObjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.S3BucketObject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3BucketObject))
	})
	return ret, err
}

// Get retrieves the S3BucketObject from the indexer for a given namespace and name.
func (s s3BucketObjectNamespaceLister) Get(name string) (*v1alpha1.S3BucketObject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("s3bucketobject"), name)
	}
	return obj.(*v1alpha1.S3BucketObject), nil
}
