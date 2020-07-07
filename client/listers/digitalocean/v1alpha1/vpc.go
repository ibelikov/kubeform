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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// VpcLister helps list Vpcs.
type VpcLister interface {
	// List lists all Vpcs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Vpc, err error)
	// Vpcs returns an object that can list and get Vpcs.
	Vpcs(namespace string) VpcNamespaceLister
	VpcListerExpansion
}

// vpcLister implements the VpcLister interface.
type vpcLister struct {
	indexer cache.Indexer
}

// NewVpcLister returns a new VpcLister.
func NewVpcLister(indexer cache.Indexer) VpcLister {
	return &vpcLister{indexer: indexer}
}

// List lists all Vpcs in the indexer.
func (s *vpcLister) List(selector labels.Selector) (ret []*v1alpha1.Vpc, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Vpc))
	})
	return ret, err
}

// Vpcs returns an object that can list and get Vpcs.
func (s *vpcLister) Vpcs(namespace string) VpcNamespaceLister {
	return vpcNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcNamespaceLister helps list and get Vpcs.
type VpcNamespaceLister interface {
	// List lists all Vpcs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Vpc, err error)
	// Get retrieves the Vpc from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Vpc, error)
	VpcNamespaceListerExpansion
}

// vpcNamespaceLister implements the VpcNamespaceLister
// interface.
type vpcNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Vpcs in the indexer for a given namespace.
func (s vpcNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Vpc, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Vpc))
	})
	return ret, err
}

// Get retrieves the Vpc from the indexer for a given namespace and name.
func (s vpcNamespaceLister) Get(name string) (*v1alpha1.Vpc, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpc"), name)
	}
	return obj.(*v1alpha1.Vpc), nil
}
