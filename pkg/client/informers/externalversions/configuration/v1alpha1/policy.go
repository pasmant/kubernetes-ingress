// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	configurationv1alpha1 "github.com/nginxinc/kubernetes-ingress/v3/pkg/apis/configuration/v1alpha1"
	versioned "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/listers/configuration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PolicyInformer provides access to a shared informer and lister for
// Policies.
type PolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PolicyLister
}

type policyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPolicyInformer constructs a new informer for Policy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPolicyInformer constructs a new informer for Policy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1alpha1().Policies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1alpha1().Policies(namespace).Watch(context.TODO(), options)
			},
		},
		&configurationv1alpha1.Policy{},
		resyncPeriod,
		indexers,
	)
}

func (f *policyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *policyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configurationv1alpha1.Policy{}, f.defaultInformer)
}

func (f *policyInformer) Lister() v1alpha1.PolicyLister {
	return v1alpha1.NewPolicyLister(f.Informer().GetIndexer())
}
