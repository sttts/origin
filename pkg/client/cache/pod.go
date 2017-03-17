package cache

import (
	"k8s.io/client-go/tools/cache"
	kcorelistersinternal "k8s.io/kubernetes/pkg/client/listers/core/internalversion"
)

type PodStoreLister struct {
	kcorelistersinternal.PodLister
	cache.Indexer
}

func NewPodStoreLister(store cache.Indexer) PodStoreLister {
	return PodStoreLister{
		kcorelistersinternal.NewPodLister(store),
		store,
	}
}
