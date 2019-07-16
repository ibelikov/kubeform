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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Domains returns a DomainInformer.
	Domains() DomainInformer
	// DomainRecords returns a DomainRecordInformer.
	DomainRecords() DomainRecordInformer
	// Images returns a ImageInformer.
	Images() ImageInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// Nodebalancers returns a NodebalancerInformer.
	Nodebalancers() NodebalancerInformer
	// NodebalancerConfigs returns a NodebalancerConfigInformer.
	NodebalancerConfigs() NodebalancerConfigInformer
	// NodebalancerNodes returns a NodebalancerNodeInformer.
	NodebalancerNodes() NodebalancerNodeInformer
	// Rdnses returns a RdnsInformer.
	Rdnses() RdnsInformer
	// Sshkeys returns a SshkeyInformer.
	Sshkeys() SshkeyInformer
	// Stackscripts returns a StackscriptInformer.
	Stackscripts() StackscriptInformer
	// Tokens returns a TokenInformer.
	Tokens() TokenInformer
	// Volumes returns a VolumeInformer.
	Volumes() VolumeInformer
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

// Domains returns a DomainInformer.
func (v *version) Domains() DomainInformer {
	return &domainInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DomainRecords returns a DomainRecordInformer.
func (v *version) DomainRecords() DomainRecordInformer {
	return &domainRecordInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Images returns a ImageInformer.
func (v *version) Images() ImageInformer {
	return &imageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Nodebalancers returns a NodebalancerInformer.
func (v *version) Nodebalancers() NodebalancerInformer {
	return &nodebalancerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// NodebalancerConfigs returns a NodebalancerConfigInformer.
func (v *version) NodebalancerConfigs() NodebalancerConfigInformer {
	return &nodebalancerConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// NodebalancerNodes returns a NodebalancerNodeInformer.
func (v *version) NodebalancerNodes() NodebalancerNodeInformer {
	return &nodebalancerNodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Rdnses returns a RdnsInformer.
func (v *version) Rdnses() RdnsInformer {
	return &rdnsInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Sshkeys returns a SshkeyInformer.
func (v *version) Sshkeys() SshkeyInformer {
	return &sshkeyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Stackscripts returns a StackscriptInformer.
func (v *version) Stackscripts() StackscriptInformer {
	return &stackscriptInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Tokens returns a TokenInformer.
func (v *version) Tokens() TokenInformer {
	return &tokenInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Volumes returns a VolumeInformer.
func (v *version) Volumes() VolumeInformer {
	return &volumeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
