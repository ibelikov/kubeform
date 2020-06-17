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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/modules/v1alpha1"
)

type FakeModulesV1alpha1 struct {
	*testing.Fake
}

func (c *FakeModulesV1alpha1) AzureAppServices(namespace string) v1alpha1.AzureAppServiceInterface {
	return &FakeAzureAppServices{c, namespace}
}

func (c *FakeModulesV1alpha1) GoogleServiceAccounts(namespace string) v1alpha1.GoogleServiceAccountInterface {
	return &FakeGoogleServiceAccounts{c, namespace}
}

func (c *FakeModulesV1alpha1) RDSs(namespace string) v1alpha1.RDSInterface {
	return &FakeRDSs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeModulesV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
