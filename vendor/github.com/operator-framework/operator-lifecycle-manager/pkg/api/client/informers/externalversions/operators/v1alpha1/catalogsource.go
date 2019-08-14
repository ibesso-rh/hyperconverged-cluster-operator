/*
Copyright The Kubernetes Authors.

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

	operators_v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	versioned "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	internalinterfaces "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/listers/operators/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CatalogSourceInformer provides access to a shared informer and lister for
// CatalogSources.
type CatalogSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CatalogSourceLister
}

type catalogSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCatalogSourceInformer constructs a new informer for CatalogSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCatalogSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCatalogSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCatalogSourceInformer constructs a new informer for CatalogSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCatalogSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorsV1alpha1().CatalogSources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorsV1alpha1().CatalogSources(namespace).Watch(options)
			},
		},
		&operators_v1alpha1.CatalogSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *catalogSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCatalogSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *catalogSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operators_v1alpha1.CatalogSource{}, f.defaultInformer)
}

func (f *catalogSourceInformer) Lister() v1alpha1.CatalogSourceLister {
	return v1alpha1.NewCatalogSourceLister(f.Informer().GetIndexer())
}
