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
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	awsv1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/aws/v1alpha1"
)

// AwsRoute53ResolverRuleAssociationInformer provides access to a shared informer and lister for
// AwsRoute53ResolverRuleAssociations.
type AwsRoute53ResolverRuleAssociationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AwsRoute53ResolverRuleAssociationLister
}

type awsRoute53ResolverRuleAssociationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAwsRoute53ResolverRuleAssociationInformer constructs a new informer for AwsRoute53ResolverRuleAssociation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAwsRoute53ResolverRuleAssociationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAwsRoute53ResolverRuleAssociationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAwsRoute53ResolverRuleAssociationInformer constructs a new informer for AwsRoute53ResolverRuleAssociation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAwsRoute53ResolverRuleAssociationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsRoute53ResolverRuleAssociations().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().AwsRoute53ResolverRuleAssociations().Watch(options)
			},
		},
		&awsv1alpha1.AwsRoute53ResolverRuleAssociation{},
		resyncPeriod,
		indexers,
	)
}

func (f *awsRoute53ResolverRuleAssociationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAwsRoute53ResolverRuleAssociationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *awsRoute53ResolverRuleAssociationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.AwsRoute53ResolverRuleAssociation{}, f.defaultInformer)
}

func (f *awsRoute53ResolverRuleAssociationInformer) Lister() v1alpha1.AwsRoute53ResolverRuleAssociationLister {
	return v1alpha1.NewAwsRoute53ResolverRuleAssociationLister(f.Informer().GetIndexer())
}
