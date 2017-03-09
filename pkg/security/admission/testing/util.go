package testing

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kapi "k8s.io/kubernetes/pkg/api"

	allocator "github.com/openshift/origin/pkg/security"
)

// CreateSAForTest Build and Initializes a ServiceAccount for tests
func CreateSAForTest() *kapi.ServiceAccount {
	return &kapi.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "default",
			Namespace: "default",
		},
	}
}

// CreateNamespaceForTest builds and initializes a Namespaces for tests
func CreateNamespaceForTest() *kapi.Namespace {
	return &kapi.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
			Annotations: map[string]string{
				allocator.UIDRangeAnnotation:           "1/3",
				allocator.MCSAnnotation:                "s0:c1,c0",
				allocator.SupplementalGroupsAnnotation: "2/3",
			},
		},
	}
}

// UserScc creates a SCC for a given user name
func UserScc(user string) *kapi.SecurityContextConstraints {
	var uid int64 = 9999
	fsGroup := int64(1)
	return &kapi.SecurityContextConstraints{
		ObjectMeta: metav1.ObjectMeta{
			SelfLink: "/api/version/securitycontextconstraints/" + user,
			Name:     user,
		},
		Users: []string{user},
		SELinuxContext: kapi.SELinuxContextStrategyOptions{
			Type: kapi.SELinuxStrategyRunAsAny,
		},
		RunAsUser: kapi.RunAsUserStrategyOptions{
			Type: kapi.RunAsUserStrategyMustRunAs,
			UID:  &uid,
		},
		FSGroup: kapi.FSGroupStrategyOptions{
			Type: kapi.FSGroupStrategyMustRunAs,
			Ranges: []kapi.IDRange{
				{Min: fsGroup, Max: fsGroup},
			},
		},
		SupplementalGroups: kapi.SupplementalGroupsStrategyOptions{
			Type: kapi.SupplementalGroupsStrategyRunAsAny,
		},
	}
}
