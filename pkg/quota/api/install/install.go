package install

import (
	"k8s.io/kubernetes/pkg/apimachinery/announced"
	"k8s.io/kubernetes/pkg/util/sets"

	"github.com/openshift/origin/pkg/quota/api"
	"github.com/openshift/origin/pkg/quota/api/v1"
)

const importPrefix = "github.com/openshift/origin/pkg/quota/api"

func init() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  api.GroupName,
			VersionPreferenceOrder:     []string{v1.SchemeGroupVersion.Version},
			ImportPrefix:               importPrefix,
			RootScopedKinds:            sets.NewString("ClusterResourceQuota"),
			AddInternalObjectsToScheme: api.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v1.SchemeGroupVersion.Version: v1.AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}
