package api

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch/versioned"
)

const (
	GroupName       = "image.openshift.io"
	LegacyGroupName = ""
)

var (
	SchemeGroupVersion       = unversioned.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	LegacySchemeGroupVersion = unversioned.GroupVersion{Group: LegacyGroupName, Version: runtime.APIVersionInternal}

	LegacySchemeBuilder    = runtime.NewSchemeBuilder(addLegacyKnownTypes)
	AddToSchemeInCoreGroup = LegacySchemeBuilder.AddToScheme

	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func LegacyKind(kind string) unversioned.GroupKind {
	return LegacySchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) unversioned.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func LegacyResource(resource string) unversioned.GroupResource {
	return LegacySchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addLegacyKnownTypes(scheme *runtime.Scheme) error {
	types := []runtime.Object{
		&Image{},
		&ImageList{},
		&DockerImage{},
		&ImageSignature{},
		&ImageStream{},
		&ImageStreamList{},
		&ImageStreamMapping{},
		&ImageStreamTag{},
		&ImageStreamTagList{},
		&ImageStreamImage{},
		&ImageStreamImport{},
	}
	scheme.AddKnownTypes(LegacySchemeGroupVersion, types...)
	return nil
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	types := []runtime.Object{
		&Image{},
		&ImageList{},
		&DockerImage{},
		&ImageSignature{},
		&ImageStream{},
		&ImageStreamList{},
		&ImageStreamMapping{},
		&ImageStreamTag{},
		&ImageStreamTagList{},
		&ImageStreamImage{},
		&ImageStreamImport{},
	}
	scheme.AddKnownTypes(SchemeGroupVersion,
		append(types,
			&unversioned.Status{}, // TODO: revisit in 1.6 when Status is actually registered as unversioned
			&kapi.ListOptions{},
			&kapi.SecretList{},
			&kapi.DeleteOptions{},
			&kapi.ExportOptions{},
		)...,
	)
	versioned.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
