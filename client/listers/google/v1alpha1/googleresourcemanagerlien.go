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

// GoogleResourceManagerLienLister helps list GoogleResourceManagerLiens.
type GoogleResourceManagerLienLister interface {
	// List lists all GoogleResourceManagerLiens in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GoogleResourceManagerLien, err error)
	// Get retrieves the GoogleResourceManagerLien from the index for a given name.
	Get(name string) (*v1alpha1.GoogleResourceManagerLien, error)
	GoogleResourceManagerLienListerExpansion
}

// googleResourceManagerLienLister implements the GoogleResourceManagerLienLister interface.
type googleResourceManagerLienLister struct {
	indexer cache.Indexer
}

// NewGoogleResourceManagerLienLister returns a new GoogleResourceManagerLienLister.
func NewGoogleResourceManagerLienLister(indexer cache.Indexer) GoogleResourceManagerLienLister {
	return &googleResourceManagerLienLister{indexer: indexer}
}

// List lists all GoogleResourceManagerLiens in the indexer.
func (s *googleResourceManagerLienLister) List(selector labels.Selector) (ret []*v1alpha1.GoogleResourceManagerLien, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GoogleResourceManagerLien))
	})
	return ret, err
}

// Get retrieves the GoogleResourceManagerLien from the index for a given name.
func (s *googleResourceManagerLienLister) Get(name string) (*v1alpha1.GoogleResourceManagerLien, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("googleresourcemanagerlien"), name)
	}
	return obj.(*v1alpha1.GoogleResourceManagerLien), nil
}
