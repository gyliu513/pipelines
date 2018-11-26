// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	tensorboardv1alpha1 "github.com/kubeflow/pipelines/backend/src/crd/pkg/apis/tensorboard/v1alpha1"
	versioned "github.com/kubeflow/pipelines/backend/src/crd/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/pipelines/backend/src/crd/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeflow/pipelines/backend/src/crd/pkg/client/listers/tensorboard/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TensorboardInformer provides access to a shared informer and lister for
// Tensorboards.
type TensorboardInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TensorboardLister
}

type tensorboardInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTensorboardInformer constructs a new informer for Tensorboard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTensorboardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTensorboardInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTensorboardInformer constructs a new informer for Tensorboard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTensorboardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TensorboardV1alpha1().Tensorboards(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TensorboardV1alpha1().Tensorboards(namespace).Watch(options)
			},
		},
		&tensorboardv1alpha1.Tensorboard{},
		resyncPeriod,
		indexers,
	)
}

func (f *tensorboardInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTensorboardInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tensorboardInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tensorboardv1alpha1.Tensorboard{}, f.defaultInformer)
}

func (f *tensorboardInformer) Lister() v1alpha1.TensorboardLister {
	return v1alpha1.NewTensorboardLister(f.Informer().GetIndexer())
}