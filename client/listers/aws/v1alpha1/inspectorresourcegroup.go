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

// InspectorResourceGroupLister helps list InspectorResourceGroups.
type InspectorResourceGroupLister interface {
	// List lists all InspectorResourceGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.InspectorResourceGroup, err error)
	// InspectorResourceGroups returns an object that can list and get InspectorResourceGroups.
	InspectorResourceGroups(namespace string) InspectorResourceGroupNamespaceLister
	InspectorResourceGroupListerExpansion
}

// inspectorResourceGroupLister implements the InspectorResourceGroupLister interface.
type inspectorResourceGroupLister struct {
	indexer cache.Indexer
}

// NewInspectorResourceGroupLister returns a new InspectorResourceGroupLister.
func NewInspectorResourceGroupLister(indexer cache.Indexer) InspectorResourceGroupLister {
	return &inspectorResourceGroupLister{indexer: indexer}
}

// List lists all InspectorResourceGroups in the indexer.
func (s *inspectorResourceGroupLister) List(selector labels.Selector) (ret []*v1alpha1.InspectorResourceGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InspectorResourceGroup))
	})
	return ret, err
}

// InspectorResourceGroups returns an object that can list and get InspectorResourceGroups.
func (s *inspectorResourceGroupLister) InspectorResourceGroups(namespace string) InspectorResourceGroupNamespaceLister {
	return inspectorResourceGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InspectorResourceGroupNamespaceLister helps list and get InspectorResourceGroups.
type InspectorResourceGroupNamespaceLister interface {
	// List lists all InspectorResourceGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.InspectorResourceGroup, err error)
	// Get retrieves the InspectorResourceGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.InspectorResourceGroup, error)
	InspectorResourceGroupNamespaceListerExpansion
}

// inspectorResourceGroupNamespaceLister implements the InspectorResourceGroupNamespaceLister
// interface.
type inspectorResourceGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InspectorResourceGroups in the indexer for a given namespace.
func (s inspectorResourceGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InspectorResourceGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InspectorResourceGroup))
	})
	return ret, err
}

// Get retrieves the InspectorResourceGroup from the indexer for a given namespace and name.
func (s inspectorResourceGroupNamespaceLister) Get(name string) (*v1alpha1.InspectorResourceGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("inspectorresourcegroup"), name)
	}
	return obj.(*v1alpha1.InspectorResourceGroup), nil
}
