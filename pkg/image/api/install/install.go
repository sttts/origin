package install

import (
	"k8s.io/kubernetes/pkg/apimachinery/announced"
	"k8s.io/kubernetes/pkg/util/sets"

	"github.com/openshift/origin/pkg/image/api"
	"github.com/openshift/origin/pkg/image/api/docker10"
	"github.com/openshift/origin/pkg/image/api/dockerpre012"
	"github.com/openshift/origin/pkg/image/api/v1"
)

const importPrefix = "github.com/openshift/origin/pkg/image/api"

func init() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  api.GroupName,
			VersionPreferenceOrder:     []string{v1.SchemeGroupVersion.Version, docker10.SchemeGroupVersion.Version, dockerpre012.SchemeGroupVersion.Version},
			ImportPrefix:               importPrefix,
			RootScopedKinds:            sets.NewString("Image", "ImageSignature"),
			AddInternalObjectsToScheme: api.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v1.SchemeGroupVersion.Version:           v1.AddToScheme,
			docker10.SchemeGroupVersion.Version:     docker10.AddToScheme,
			dockerpre012.SchemeGroupVersion.Version: dockerpre012.AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}
