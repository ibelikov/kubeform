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

// ProjectResourcesLister helps list ProjectResourceses.
type ProjectResourcesLister interface {
	// List lists all ProjectResourceses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectResources, err error)
	// ProjectResourceses returns an object that can list and get ProjectResourceses.
	ProjectResourceses(namespace string) ProjectResourcesNamespaceLister
	ProjectResourcesListerExpansion
}

// projectResourcesLister implements the ProjectResourcesLister interface.
type projectResourcesLister struct {
	indexer cache.Indexer
}

// NewProjectResourcesLister returns a new ProjectResourcesLister.
func NewProjectResourcesLister(indexer cache.Indexer) ProjectResourcesLister {
	return &projectResourcesLister{indexer: indexer}
}

// List lists all ProjectResourceses in the indexer.
func (s *projectResourcesLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectResources, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectResources))
	})
	return ret, err
}

// ProjectResourceses returns an object that can list and get ProjectResourceses.
func (s *projectResourcesLister) ProjectResourceses(namespace string) ProjectResourcesNamespaceLister {
	return projectResourcesNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectResourcesNamespaceLister helps list and get ProjectResourceses.
type ProjectResourcesNamespaceLister interface {
	// List lists all ProjectResourceses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectResources, err error)
	// Get retrieves the ProjectResources from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ProjectResources, error)
	ProjectResourcesNamespaceListerExpansion
}

// projectResourcesNamespaceLister implements the ProjectResourcesNamespaceLister
// interface.
type projectResourcesNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProjectResourceses in the indexer for a given namespace.
func (s projectResourcesNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectResources, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectResources))
	})
	return ret, err
}

// Get retrieves the ProjectResources from the indexer for a given namespace and name.
func (s projectResourcesNamespaceLister) Get(name string) (*v1alpha1.ProjectResources, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("projectresources"), name)
	}
	return obj.(*v1alpha1.ProjectResources), nil
}
