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

// DataprocJobLister helps list DataprocJobs.
type DataprocJobLister interface {
	// List lists all DataprocJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataprocJob, err error)
	// DataprocJobs returns an object that can list and get DataprocJobs.
	DataprocJobs(namespace string) DataprocJobNamespaceLister
	DataprocJobListerExpansion
}

// dataprocJobLister implements the DataprocJobLister interface.
type dataprocJobLister struct {
	indexer cache.Indexer
}

// NewDataprocJobLister returns a new DataprocJobLister.
func NewDataprocJobLister(indexer cache.Indexer) DataprocJobLister {
	return &dataprocJobLister{indexer: indexer}
}

// List lists all DataprocJobs in the indexer.
func (s *dataprocJobLister) List(selector labels.Selector) (ret []*v1alpha1.DataprocJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataprocJob))
	})
	return ret, err
}

// DataprocJobs returns an object that can list and get DataprocJobs.
func (s *dataprocJobLister) DataprocJobs(namespace string) DataprocJobNamespaceLister {
	return dataprocJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataprocJobNamespaceLister helps list and get DataprocJobs.
type DataprocJobNamespaceLister interface {
	// List lists all DataprocJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataprocJob, err error)
	// Get retrieves the DataprocJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataprocJob, error)
	DataprocJobNamespaceListerExpansion
}

// dataprocJobNamespaceLister implements the DataprocJobNamespaceLister
// interface.
type dataprocJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataprocJobs in the indexer for a given namespace.
func (s dataprocJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataprocJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataprocJob))
	})
	return ret, err
}

// Get retrieves the DataprocJob from the indexer for a given namespace and name.
func (s dataprocJobNamespaceLister) Get(name string) (*v1alpha1.DataprocJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dataprocjob"), name)
	}
	return obj.(*v1alpha1.DataprocJob), nil
}
