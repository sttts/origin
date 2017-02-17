package rest

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/genericapiserver"

	configapi "github.com/openshift/origin/pkg/cmd/server/api"
	"github.com/openshift/origin/pkg/util/restoptions"
)

// StorageOptions returns the appropriate storage configuration for the origin rest APIs, including
// overiddes.
func StorageOptions(options configapi.MasterConfig) restoptions.Getter {
	return restoptions.NewConfigGetter(
		options,
		&genericapiserver.ResourceConfig{},
		map[unversioned.GroupResource]string{
			{Resource: "clusterpolicies"}:       "authorization/cluster/policies",
			{Resource: "clusterpolicybindings"}: "authorization/cluster/policybindings",
			{Resource: "policies"}:              "authorization/local/policies",
			{Resource: "policybindings"}:        "authorization/local/policybindings",

			{Resource: "oauthaccesstokens"}:         "oauth/accesstokens",
			{Resource: "oauthauthorizetokens"}:      "oauth/authorizetokens",
			{Resource: "oauthclients"}:              "oauth/clients",
			{Resource: "oauthclientauthorizations"}: "oauth/clientauthorizations",

			{Resource: "identities"}: "useridentities",

			// REBASE: add legacy variant for network resources
			{Resource: "clusternetworks"}:       "registry/sdnnetworks",
			{Resource: "egressnetworkpolicies"}: "registry/egressnetworkpolicy",
			{Resource: "hostsubnets"}:           "registry/sdnsubnets",
			{Resource: "netnamespaces"}:         "registry/sdnnetnamespaces",

			{Group: authorizationapi.GroupName, Resource: "clusterpolicies"}:       "authorization/cluster/policies",
			{Group: authorizationapi.GroupName, Resource: "clusterpolicybindings"}: "authorization/cluster/policybindings",
			{Group: authorizationapi.GroupName, Resource: "policies"}:              "authorization/local/policies",
			{Group: authorizationapi.GroupName, Resource: "policybindings"}:        "authorization/local/policybindings",

			{Group: authorizationapi.GroupName, Resource: "oauthaccesstokens"}:         "oauth/accesstokens",
			{Group: authorizationapi.GroupName, Resource: "oauthauthorizetokens"}:      "oauth/authorizetokens",
			{Group: authorizationapi.GroupName, Resource: "oauthclients"}:              "oauth/clients",
			{Group: authorizationapi.GroupName, Resource: "oauthclientauthorizations"}: "oauth/clientauthorizations",

			{Group: authorizationapi.GroupName, Resource: "identities"}: "useridentities",
		},
		map[unversioned.GroupResource]struct{}{
			{Resource: "oauthauthorizetokens"}: {},
			{Resource: "oauthaccesstokens"}:    {},

			{Group: authorizationapi.GroupName, Resource: "oauthauthorizetokens"}: {},
			{Group: authorizationapi.GroupName, Resource: "oauthaccesstokens"}:    {},
		},
	)
}
