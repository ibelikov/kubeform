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

// IotPolicyAttachmentLister helps list IotPolicyAttachments.
type IotPolicyAttachmentLister interface {
	// List lists all IotPolicyAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IotPolicyAttachment, err error)
	// IotPolicyAttachments returns an object that can list and get IotPolicyAttachments.
	IotPolicyAttachments(namespace string) IotPolicyAttachmentNamespaceLister
	IotPolicyAttachmentListerExpansion
}

// iotPolicyAttachmentLister implements the IotPolicyAttachmentLister interface.
type iotPolicyAttachmentLister struct {
	indexer cache.Indexer
}

// NewIotPolicyAttachmentLister returns a new IotPolicyAttachmentLister.
func NewIotPolicyAttachmentLister(indexer cache.Indexer) IotPolicyAttachmentLister {
	return &iotPolicyAttachmentLister{indexer: indexer}
}

// List lists all IotPolicyAttachments in the indexer.
func (s *iotPolicyAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.IotPolicyAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotPolicyAttachment))
	})
	return ret, err
}

// IotPolicyAttachments returns an object that can list and get IotPolicyAttachments.
func (s *iotPolicyAttachmentLister) IotPolicyAttachments(namespace string) IotPolicyAttachmentNamespaceLister {
	return iotPolicyAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IotPolicyAttachmentNamespaceLister helps list and get IotPolicyAttachments.
type IotPolicyAttachmentNamespaceLister interface {
	// List lists all IotPolicyAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IotPolicyAttachment, err error)
	// Get retrieves the IotPolicyAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IotPolicyAttachment, error)
	IotPolicyAttachmentNamespaceListerExpansion
}

// iotPolicyAttachmentNamespaceLister implements the IotPolicyAttachmentNamespaceLister
// interface.
type iotPolicyAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IotPolicyAttachments in the indexer for a given namespace.
func (s iotPolicyAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IotPolicyAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotPolicyAttachment))
	})
	return ret, err
}

// Get retrieves the IotPolicyAttachment from the indexer for a given namespace and name.
func (s iotPolicyAttachmentNamespaceLister) Get(name string) (*v1alpha1.IotPolicyAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iotpolicyattachment"), name)
	}
	return obj.(*v1alpha1.IotPolicyAttachment), nil
}
