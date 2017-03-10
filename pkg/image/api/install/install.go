package install

import (
	"fmt"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apimachinery"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	kapi "k8s.io/kubernetes/pkg/api"

	"github.com/openshift/origin/pkg/image/api"
	"github.com/openshift/origin/pkg/image/api/docker10"
	"github.com/openshift/origin/pkg/image/api/dockerpre012"
	"github.com/openshift/origin/pkg/image/api/v1"
)

const importPrefix = "github.com/openshift/origin/pkg/image/api"

var accessor = meta.NewAccessor()

// availableVersions lists all known external versions for this group from most preferred to least preferred
var availableVersions = []schema.GroupVersion{
	v1.LegacySchemeGroupVersion, docker10.LegacySchemeGroupVersion, dockerpre012.LegacySchemeGroupVersion,
}

func init() {
	registered.RegisterVersions(availableVersions)
	externalVersions := []schema.GroupVersion{}
	for _, v := range availableVersions {
		if registered.IsAllowedVersion(v) {
			externalVersions = append(externalVersions, v)
		}
	}
	if len(externalVersions) == 0 {
		glog.Infof("No version is registered for group %v", api.LegacyGroupName)
		return
	}

	if err := registered.EnableVersions(externalVersions...); err != nil {
		panic(err)
	}
	if err := enableVersions(externalVersions); err != nil {
		panic(err)
	}
}

// TODO: enableVersions should be centralized rather than spread in each API
// group.
// We can combine registered.RegisterVersions, registered.EnableVersions and
// registered.RegisterGroup once we have moved enableVersions there.
func enableVersions(externalVersions []schema.GroupVersion) error {
	addVersionsToScheme(externalVersions...)
	preferredExternalVersion := externalVersions[0]

	groupMeta := apimachinery.GroupMeta{
		GroupVersion:  preferredExternalVersion,
		GroupVersions: externalVersions,
		RESTMapper:    newRESTMapper(externalVersions),
		SelfLinker:    runtime.SelfLinker(accessor),
		InterfacesFor: interfacesFor,
	}

	if err := registered.RegisterGroup(groupMeta); err != nil {
		return err
	}
	return nil
}

func addVersionsToScheme(externalVersions ...schema.GroupVersion) {
	// add the internal version to Scheme
	api.AddToSchemeInCoreGroup(kapi.Scheme)
	// add the enabled external versions to Scheme
	for _, v := range externalVersions {
		if !registered.IsEnabledVersion(v) {
			glog.Errorf("Version %s is not enabled, so it will not be added to the Scheme.", v)
			continue
		}
		switch v {
		case v1.LegacySchemeGroupVersion:
			v1.AddToSchemeInCoreGroup(kapi.Scheme)
		case docker10.LegacySchemeGroupVersion:
			docker10.AddToSchemeInCoreGroup(kapi.Scheme)
		case dockerpre012.LegacySchemeGroupVersion:
			dockerpre012.AddToSchemeInCoreGroup(kapi.Scheme)

		default:
			glog.Errorf("Version %s is not known, so it will not be added to the Scheme.", v)
			continue
		}
	}
}

func newRESTMapper(externalVersions []schema.GroupVersion) meta.RESTMapper {
	rootScoped := sets.NewString("Image", "ImageSignature")
	ignoredKinds := sets.NewString()
	return kapi.NewDefaultRESTMapper(externalVersions, interfacesFor, importPrefix, ignoredKinds, rootScoped)
}

func interfacesFor(version schema.GroupVersion) (*meta.VersionInterfaces, error) {
	switch version {
	case v1.LegacySchemeGroupVersion:
		return &meta.VersionInterfaces{
			ObjectConvertor:  kapi.Scheme,
			MetadataAccessor: accessor,
		}, nil

	case docker10.SchemeGroupVersion:
		return &meta.VersionInterfaces{
			ObjectConvertor:  kapi.Scheme,
			MetadataAccessor: accessor,
		}, nil

	case dockerpre012.SchemeGroupVersion:
		return &meta.VersionInterfaces{
			ObjectConvertor:  kapi.Scheme,
			MetadataAccessor: accessor,
		}, nil

	default:
		g, _ := registered.Group(api.LegacyGroupName)
		return nil, fmt.Errorf("unsupported storage version: %s (valid: %v)", version, g.GroupVersions)
	}
}
