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

// KubernetesClusterLister helps list KubernetesClusters.
type KubernetesClusterLister interface {
	// List lists all KubernetesClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KubernetesCluster, err error)
	// KubernetesClusters returns an object that can list and get KubernetesClusters.
	KubernetesClusters(namespace string) KubernetesClusterNamespaceLister
	KubernetesClusterListerExpansion
}

// kubernetesClusterLister implements the KubernetesClusterLister interface.
type kubernetesClusterLister struct {
	indexer cache.Indexer
}

// NewKubernetesClusterLister returns a new KubernetesClusterLister.
func NewKubernetesClusterLister(indexer cache.Indexer) KubernetesClusterLister {
	return &kubernetesClusterLister{indexer: indexer}
}

// List lists all KubernetesClusters in the indexer.
func (s *kubernetesClusterLister) List(selector labels.Selector) (ret []*v1alpha1.KubernetesCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubernetesCluster))
	})
	return ret, err
}

// KubernetesClusters returns an object that can list and get KubernetesClusters.
func (s *kubernetesClusterLister) KubernetesClusters(namespace string) KubernetesClusterNamespaceLister {
	return kubernetesClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubernetesClusterNamespaceLister helps list and get KubernetesClusters.
type KubernetesClusterNamespaceLister interface {
	// List lists all KubernetesClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KubernetesCluster, err error)
	// Get retrieves the KubernetesCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KubernetesCluster, error)
	KubernetesClusterNamespaceListerExpansion
}

// kubernetesClusterNamespaceLister implements the KubernetesClusterNamespaceLister
// interface.
type kubernetesClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KubernetesClusters in the indexer for a given namespace.
func (s kubernetesClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KubernetesCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubernetesCluster))
	})
	return ret, err
}

// Get retrieves the KubernetesCluster from the indexer for a given namespace and name.
func (s kubernetesClusterNamespaceLister) Get(name string) (*v1alpha1.KubernetesCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubernetescluster"), name)
	}
	return obj.(*v1alpha1.KubernetesCluster), nil
}
