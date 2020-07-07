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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// EKSs returns a EKSInformer.
	EKSs() EKSInformer
	// GoogleServiceAccounts returns a GoogleServiceAccountInformer.
	GoogleServiceAccounts() GoogleServiceAccountInformer
	// RDSs returns a RDSInformer.
	RDSs() RDSInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// EKSs returns a EKSInformer.
func (v *version) EKSs() EKSInformer {
	return &eKSInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GoogleServiceAccounts returns a GoogleServiceAccountInformer.
func (v *version) GoogleServiceAccounts() GoogleServiceAccountInformer {
	return &googleServiceAccountInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RDSs returns a RDSInformer.
func (v *version) RDSs() RDSInformer {
	return &rDSInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
