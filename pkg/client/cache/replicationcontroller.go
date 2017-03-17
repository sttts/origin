package cache

import (
	"k8s.io/client-go/tools/cache"
	kcorelistersinternal "k8s.io/kubernetes/pkg/client/listers/core/internalversion"
)

type ReplicationControllerStoreLister struct {
	kcorelistersinternal.ReplicationControllerLister
	cache.Indexer
}

func NewReplicationControllerStoreLister(store cache.Indexer) ReplicationControllerStoreLister {
	return ReplicationControllerStoreLister{
		kcorelistersinternal.NewReplicationControllerLister(store),
		store,
	}
}
