// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	controllerv1 "github.com/K-Phoen/dark/internal/pkg/apis/controller/v1"
	versioned "github.com/K-Phoen/dark/internal/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/K-Phoen/dark/internal/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/K-Phoen/dark/internal/pkg/generated/listers/controller/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GrafanaDashboardInformer provides access to a shared informer and lister for
// GrafanaDashboards.
type GrafanaDashboardInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.GrafanaDashboardLister
}

type grafanaDashboardInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGrafanaDashboardInformer constructs a new informer for GrafanaDashboard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGrafanaDashboardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGrafanaDashboardInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGrafanaDashboardInformer constructs a new informer for GrafanaDashboard type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGrafanaDashboardInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1().GrafanaDashboards(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1().GrafanaDashboards(namespace).Watch(context.TODO(), options)
			},
		},
		&controllerv1.GrafanaDashboard{},
		resyncPeriod,
		indexers,
	)
}

func (f *grafanaDashboardInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGrafanaDashboardInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *grafanaDashboardInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&controllerv1.GrafanaDashboard{}, f.defaultInformer)
}

func (f *grafanaDashboardInformer) Lister() v1.GrafanaDashboardLister {
	return v1.NewGrafanaDashboardLister(f.Informer().GetIndexer())
}
