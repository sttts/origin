package api

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch/versioned"
)

const (
	LegacyGroupName = ""
	GroupName       = "authorization.openshift.io"
)

// SchemeGroupVersion is group version used to register these objects
var (
	SchemeGroupVersion       = unversioned.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	LegacySchemeGroupVersion = unversioned.GroupVersion{Group: LegacyGroupName, Version: runtime.APIVersionInternal}
)

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) unversioned.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
func LegacyResource(resource string) unversioned.GroupResource {
	return LegacySchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	types := []runtime.Object{
		&Role{},
		&RoleBinding{},
		&Policy{},
		&PolicyBinding{},
		&PolicyList{},
		&PolicyBindingList{},
		&RoleBindingList{},
		&RoleList{},

		&SelfSubjectRulesReview{},
		&SubjectRulesReview{},
		&ResourceAccessReview{},
		&SubjectAccessReview{},
		&LocalResourceAccessReview{},
		&LocalSubjectAccessReview{},
		&ResourceAccessReviewResponse{},
		&SubjectAccessReviewResponse{},
		&IsPersonalSubjectAccessReview{},

		&ClusterRole{},
		&ClusterRoleBinding{},
		&ClusterPolicy{},
		&ClusterPolicyBinding{},
		&ClusterPolicyList{},
		&ClusterPolicyBindingList{},
		&ClusterRoleBindingList{},
		&ClusterRoleList{},

		&RoleBindingRestriction{},
		&RoleBindingRestrictionList{},
	}
	scheme.AddKnownTypes(SchemeGroupVersion,
		append(types,
			&unversioned.Status{}, // TODO: revisit in 1.6 when Status is actually registered as unversioned
			&kapi.ListOptions{},
			&kapi.DeleteOptions{},
			&kapi.ExportOptions{},
		)...,
	)

	versioned.AddToGroupVersion(scheme, SchemeGroupVersion)
	scheme.AddKnownTypes(LegacySchemeGroupVersion, types...)
	return nil
}
