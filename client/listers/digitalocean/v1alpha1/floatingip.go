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

// FloatingIPLister helps list FloatingIPs.
type FloatingIPLister interface {
	// List lists all FloatingIPs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error)
	// FloatingIPs returns an object that can list and get FloatingIPs.
	FloatingIPs(namespace string) FloatingIPNamespaceLister
	FloatingIPListerExpansion
}

// floatingIPLister implements the FloatingIPLister interface.
type floatingIPLister struct {
	indexer cache.Indexer
}

// NewFloatingIPLister returns a new FloatingIPLister.
func NewFloatingIPLister(indexer cache.Indexer) FloatingIPLister {
	return &floatingIPLister{indexer: indexer}
}

// List lists all FloatingIPs in the indexer.
func (s *floatingIPLister) List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FloatingIP))
	})
	return ret, err
}

// FloatingIPs returns an object that can list and get FloatingIPs.
func (s *floatingIPLister) FloatingIPs(namespace string) FloatingIPNamespaceLister {
	return floatingIPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FloatingIPNamespaceLister helps list and get FloatingIPs.
type FloatingIPNamespaceLister interface {
	// List lists all FloatingIPs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error)
	// Get retrieves the FloatingIP from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FloatingIP, error)
	FloatingIPNamespaceListerExpansion
}

// floatingIPNamespaceLister implements the FloatingIPNamespaceLister
// interface.
type floatingIPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FloatingIPs in the indexer for a given namespace.
func (s floatingIPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FloatingIP))
	})
	return ret, err
}

// Get retrieves the FloatingIP from the indexer for a given namespace and name.
func (s floatingIPNamespaceLister) Get(name string) (*v1alpha1.FloatingIP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("floatingip"), name)
	}
	return obj.(*v1alpha1.FloatingIP), nil
}
