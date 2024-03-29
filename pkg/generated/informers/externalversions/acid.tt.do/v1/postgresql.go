/*
Copyright tt Compose, wangjianc Author
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	acidttdov1 "github.com/wangjianc/pg-operator/pkg/apis/acid.tt.do/v1"
	versioned "github.com/wangjianc/pg-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/wangjianc/pg-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/wangjianc/pg-operator/pkg/generated/listers/acid.tt.do/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PostgresqlInformer provides access to a shared informer and lister for
// Postgresqls.
type PostgresqlInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PostgresqlLister
}

type postgresqlInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPostgresqlInformer constructs a new informer for Postgresql type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPostgresqlInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPostgresqlInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPostgresqlInformer constructs a new informer for Postgresql type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPostgresqlInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AcidV1().Postgresqls(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AcidV1().Postgresqls(namespace).Watch(context.TODO(), options)
			},
		},
		&acidttdov1.Postgresql{},
		resyncPeriod,
		indexers,
	)
}

func (f *postgresqlInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPostgresqlInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *postgresqlInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&acidttdov1.Postgresql{}, f.defaultInformer)
}

func (f *postgresqlInformer) Lister() v1.PostgresqlLister {
	return v1.NewPostgresqlLister(f.Informer().GetIndexer())
}
