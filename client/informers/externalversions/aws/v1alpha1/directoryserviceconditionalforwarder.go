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
	"context"
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

// DirectoryServiceConditionalForwarderInformer provides access to a shared informer and lister for
// DirectoryServiceConditionalForwarders.
type DirectoryServiceConditionalForwarderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DirectoryServiceConditionalForwarderLister
}

type directoryServiceConditionalForwarderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDirectoryServiceConditionalForwarderInformer constructs a new informer for DirectoryServiceConditionalForwarder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDirectoryServiceConditionalForwarderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDirectoryServiceConditionalForwarderInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDirectoryServiceConditionalForwarderInformer constructs a new informer for DirectoryServiceConditionalForwarder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDirectoryServiceConditionalForwarderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().DirectoryServiceConditionalForwarders(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().DirectoryServiceConditionalForwarders(namespace).Watch(context.TODO(), options)
			},
		},
		&awsv1alpha1.DirectoryServiceConditionalForwarder{},
		resyncPeriod,
		indexers,
	)
}

func (f *directoryServiceConditionalForwarderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDirectoryServiceConditionalForwarderInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *directoryServiceConditionalForwarderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.DirectoryServiceConditionalForwarder{}, f.defaultInformer)
}

func (f *directoryServiceConditionalForwarderInformer) Lister() v1alpha1.DirectoryServiceConditionalForwarderLister {
	return v1alpha1.NewDirectoryServiceConditionalForwarderLister(f.Informer().GetIndexer())
}
