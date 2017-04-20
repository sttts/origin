package v1

import (
	v1 "github.com/openshift/origin/pkg/oauth/api/v1"
	"github.com/openshift/origin/pkg/oauth/clientset/release_v3_6/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type OauthV1Interface interface {
	RESTClient() rest.Interface
	OAuthClientsGetter
}

// OauthV1Client is used to interact with features provided by the oauth.openshift.io group.
type OauthV1Client struct {
	restClient rest.Interface
}

func (c *OauthV1Client) OAuthClients(namespace string) OAuthClientInterface {
	return newOAuthClients(c, namespace)
}

// NewForConfig creates a new OauthV1Client for the given config.
func NewForConfig(c *rest.Config) (*OauthV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OauthV1Client{client}, nil
}

// NewForConfigOrDie creates a new OauthV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OauthV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OauthV1Client for the given RESTClient.
func New(c rest.Interface) *OauthV1Client {
	return &OauthV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *OauthV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
