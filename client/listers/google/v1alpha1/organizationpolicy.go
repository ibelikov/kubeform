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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// OrganizationPolicyLister helps list OrganizationPolicies.
type OrganizationPolicyLister interface {
	// List lists all OrganizationPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationPolicy, err error)
	// Get retrieves the OrganizationPolicy from the index for a given name.
	Get(name string) (*v1alpha1.OrganizationPolicy, error)
	OrganizationPolicyListerExpansion
}

// organizationPolicyLister implements the OrganizationPolicyLister interface.
type organizationPolicyLister struct {
	indexer cache.Indexer
}

// NewOrganizationPolicyLister returns a new OrganizationPolicyLister.
func NewOrganizationPolicyLister(indexer cache.Indexer) OrganizationPolicyLister {
	return &organizationPolicyLister{indexer: indexer}
}

// List lists all OrganizationPolicies in the indexer.
func (s *organizationPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationPolicy))
	})
	return ret, err
}

// Get retrieves the OrganizationPolicy from the index for a given name.
func (s *organizationPolicyLister) Get(name string) (*v1alpha1.OrganizationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("organizationpolicy"), name)
	}
	return obj.(*v1alpha1.OrganizationPolicy), nil
}
