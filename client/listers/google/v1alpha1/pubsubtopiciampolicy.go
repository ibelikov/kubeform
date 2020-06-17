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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// PubsubTopicIamPolicyLister helps list PubsubTopicIamPolicies.
type PubsubTopicIamPolicyLister interface {
	// List lists all PubsubTopicIamPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamPolicy, err error)
	// PubsubTopicIamPolicies returns an object that can list and get PubsubTopicIamPolicies.
	PubsubTopicIamPolicies(namespace string) PubsubTopicIamPolicyNamespaceLister
	PubsubTopicIamPolicyListerExpansion
}

// pubsubTopicIamPolicyLister implements the PubsubTopicIamPolicyLister interface.
type pubsubTopicIamPolicyLister struct {
	indexer cache.Indexer
}

// NewPubsubTopicIamPolicyLister returns a new PubsubTopicIamPolicyLister.
func NewPubsubTopicIamPolicyLister(indexer cache.Indexer) PubsubTopicIamPolicyLister {
	return &pubsubTopicIamPolicyLister{indexer: indexer}
}

// List lists all PubsubTopicIamPolicies in the indexer.
func (s *pubsubTopicIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubTopicIamPolicy))
	})
	return ret, err
}

// PubsubTopicIamPolicies returns an object that can list and get PubsubTopicIamPolicies.
func (s *pubsubTopicIamPolicyLister) PubsubTopicIamPolicies(namespace string) PubsubTopicIamPolicyNamespaceLister {
	return pubsubTopicIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PubsubTopicIamPolicyNamespaceLister helps list and get PubsubTopicIamPolicies.
type PubsubTopicIamPolicyNamespaceLister interface {
	// List lists all PubsubTopicIamPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamPolicy, err error)
	// Get retrieves the PubsubTopicIamPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PubsubTopicIamPolicy, error)
	PubsubTopicIamPolicyNamespaceListerExpansion
}

// pubsubTopicIamPolicyNamespaceLister implements the PubsubTopicIamPolicyNamespaceLister
// interface.
type pubsubTopicIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PubsubTopicIamPolicies in the indexer for a given namespace.
func (s pubsubTopicIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubTopicIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubTopicIamPolicy))
	})
	return ret, err
}

// Get retrieves the PubsubTopicIamPolicy from the indexer for a given namespace and name.
func (s pubsubTopicIamPolicyNamespaceLister) Get(name string) (*v1alpha1.PubsubTopicIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pubsubtopiciampolicy"), name)
	}
	return obj.(*v1alpha1.PubsubTopicIamPolicy), nil
}
