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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// NetappVolumeLister helps list NetappVolumes.
type NetappVolumeLister interface {
	// List lists all NetappVolumes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetappVolume, err error)
	// NetappVolumes returns an object that can list and get NetappVolumes.
	NetappVolumes(namespace string) NetappVolumeNamespaceLister
	NetappVolumeListerExpansion
}

// netappVolumeLister implements the NetappVolumeLister interface.
type netappVolumeLister struct {
	indexer cache.Indexer
}

// NewNetappVolumeLister returns a new NetappVolumeLister.
func NewNetappVolumeLister(indexer cache.Indexer) NetappVolumeLister {
	return &netappVolumeLister{indexer: indexer}
}

// List lists all NetappVolumes in the indexer.
func (s *netappVolumeLister) List(selector labels.Selector) (ret []*v1alpha1.NetappVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetappVolume))
	})
	return ret, err
}

// NetappVolumes returns an object that can list and get NetappVolumes.
func (s *netappVolumeLister) NetappVolumes(namespace string) NetappVolumeNamespaceLister {
	return netappVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetappVolumeNamespaceLister helps list and get NetappVolumes.
type NetappVolumeNamespaceLister interface {
	// List lists all NetappVolumes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetappVolume, err error)
	// Get retrieves the NetappVolume from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetappVolume, error)
	NetappVolumeNamespaceListerExpansion
}

// netappVolumeNamespaceLister implements the NetappVolumeNamespaceLister
// interface.
type netappVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetappVolumes in the indexer for a given namespace.
func (s netappVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetappVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetappVolume))
	})
	return ret, err
}

// Get retrieves the NetappVolume from the indexer for a given namespace and name.
func (s netappVolumeNamespaceLister) Get(name string) (*v1alpha1.NetappVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("netappvolume"), name)
	}
	return obj.(*v1alpha1.NetappVolume), nil
}
