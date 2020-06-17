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

// S3BucketPolicyLister helps list S3BucketPolicies.
type S3BucketPolicyLister interface {
	// List lists all S3BucketPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.S3BucketPolicy, err error)
	// S3BucketPolicies returns an object that can list and get S3BucketPolicies.
	S3BucketPolicies(namespace string) S3BucketPolicyNamespaceLister
	S3BucketPolicyListerExpansion
}

// s3BucketPolicyLister implements the S3BucketPolicyLister interface.
type s3BucketPolicyLister struct {
	indexer cache.Indexer
}

// NewS3BucketPolicyLister returns a new S3BucketPolicyLister.
func NewS3BucketPolicyLister(indexer cache.Indexer) S3BucketPolicyLister {
	return &s3BucketPolicyLister{indexer: indexer}
}

// List lists all S3BucketPolicies in the indexer.
func (s *s3BucketPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.S3BucketPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3BucketPolicy))
	})
	return ret, err
}

// S3BucketPolicies returns an object that can list and get S3BucketPolicies.
func (s *s3BucketPolicyLister) S3BucketPolicies(namespace string) S3BucketPolicyNamespaceLister {
	return s3BucketPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// S3BucketPolicyNamespaceLister helps list and get S3BucketPolicies.
type S3BucketPolicyNamespaceLister interface {
	// List lists all S3BucketPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.S3BucketPolicy, err error)
	// Get retrieves the S3BucketPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.S3BucketPolicy, error)
	S3BucketPolicyNamespaceListerExpansion
}

// s3BucketPolicyNamespaceLister implements the S3BucketPolicyNamespaceLister
// interface.
type s3BucketPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all S3BucketPolicies in the indexer for a given namespace.
func (s s3BucketPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.S3BucketPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3BucketPolicy))
	})
	return ret, err
}

// Get retrieves the S3BucketPolicy from the indexer for a given namespace and name.
func (s s3BucketPolicyNamespaceLister) Get(name string) (*v1alpha1.S3BucketPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("s3bucketpolicy"), name)
	}
	return obj.(*v1alpha1.S3BucketPolicy), nil
}
