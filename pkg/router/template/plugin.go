package templaterouter

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

	"github.com/golang/glog"
	ktypes "k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	kapi "k8s.io/kubernetes/pkg/api"

	routeapi "github.com/openshift/origin/pkg/route/api"
	unidlingapi "github.com/openshift/origin/pkg/unidling/api"
)

// TemplatePlugin implements the router.Plugin interface to provide
// a template based, backend-agnostic router.
type TemplatePlugin struct {
	Router         routerInterface
	IncludeUDP     bool
	ServiceFetcher ServiceLookup
}

func newDefaultTemplatePlugin(router routerInterface, includeUDP bool, lookupSvc ServiceLookup) *TemplatePlugin {
	return &TemplatePlugin{
		Router:         router,
		IncludeUDP:     includeUDP,
		ServiceFetcher: lookupSvc,
	}
}

type TemplatePluginConfig struct {
	WorkingDir             string
	TemplatePath           string
	ReloadScriptPath       string
	ReloadInterval         time.Duration
	DefaultCertificate     string
	DefaultCertificatePath string
	DefaultCertificateDir  string
	StatsPort              int
	StatsUsername          string
	StatsPassword          string
	IncludeUDP             bool
	AllowWildcardRoutes    bool
	PeerService            *ktypes.NamespacedName
	BindPortsAfterSync     bool
	MaxConnections         string
}

// routerInterface controls the interaction of the plugin with the underlying router implementation
type routerInterface interface {
	// Mutative operations in this interface do not return errors.
	// The only error state for these methods is when an unknown
	// frontend key is used; all call sites make certain the frontend
	// is created.

	// CreateServiceUnit creates a new service named with the given id.
	CreateServiceUnit(id string)
	// FindServiceUnit finds the service with the given id.
	FindServiceUnit(id string) (v ServiceUnit, ok bool)

	// AddEndpoints adds new Endpoints for the given id.
	AddEndpoints(id string, endpoints []Endpoint)
	// DeleteEndpoints deletes the endpoints for the frontend with the given id.
	DeleteEndpoints(id string)

	// AddRoute attempts to add a route to the router.
	AddRoute(route *routeapi.Route)
	// RemoveRoute removes the given route
	RemoveRoute(route *routeapi.Route)
	// HasRoute indicates whether the router is configured with the given route
	HasRoute(route *routeapi.Route) bool
	// Reduce the list of routes to only these namespaces
	FilterNamespaces(namespaces sets.String)
	// Commit applies the changes in the background. It kicks off a rate-limited
	// commit (persist router state + refresh the backend) that coalesces multiple changes.
	Commit()
}

func env(name, defaultValue string) string {
	if envValue := os.Getenv(name); envValue != "" {
		return envValue
	}

	return defaultValue
}

// NewTemplatePlugin creates a new TemplatePlugin.
func NewTemplatePlugin(cfg TemplatePluginConfig, lookupSvc ServiceLookup) (*TemplatePlugin, error) {
	templateBaseName := filepath.Base(cfg.TemplatePath)
	globalFuncs := template.FuncMap{
		"endpointsForAlias": endpointsForAlias, //returns the list of valid endpoints
		"env":               env,               //tries to get an environment variable if it can't return a default
		"matchPattern":      matchPattern,      //anchors provided regular expression and evaluates against given string
		"isInteger":         isInteger,         //determines if a given variable is an integer
		"matchValues":       matchValues,       //compares a given string to a list of allowed strings

		"genSubdomainWildcardRegexp": genSubdomainWildcardRegexp, //generates a regular expression matching the subdomain for hosts (and paths) with a wildcard policy
		"generateRouteRegexp":        generateRouteRegexp,        //generates a regular expression matching the route hosts (and paths)
		"genCertificateHostName":     genCertificateHostName,     //generates host name to use for serving/matching certificates
	}
	masterTemplate, err := template.New("config").Funcs(globalFuncs).ParseFiles(cfg.TemplatePath)
	if err != nil {
		return nil, err
	}
	templates := map[string]*template.Template{}

	for _, template := range masterTemplate.Templates() {
		if template.Name() == templateBaseName {
			continue
		}

		templates[template.Name()] = template
	}

	peerKey := ""
	if cfg.PeerService != nil {
		peerKey = peerEndpointsKey(*cfg.PeerService)
	}

	templateRouterCfg := templateRouterCfg{
		dir:                    cfg.WorkingDir,
		templates:              templates,
		reloadScriptPath:       cfg.ReloadScriptPath,
		reloadInterval:         cfg.ReloadInterval,
		defaultCertificate:     cfg.DefaultCertificate,
		defaultCertificatePath: cfg.DefaultCertificatePath,
		defaultCertificateDir:  cfg.DefaultCertificateDir,
		statsUser:              cfg.StatsUsername,
		statsPassword:          cfg.StatsPassword,
		statsPort:              cfg.StatsPort,
		allowWildcardRoutes:    cfg.AllowWildcardRoutes,
		peerEndpointsKey:       peerKey,
		bindPortsAfterSync:     cfg.BindPortsAfterSync,
	}
	router, err := newTemplateRouter(templateRouterCfg)
	return newDefaultTemplatePlugin(router, cfg.IncludeUDP, lookupSvc), err
}

// HandleEndpoints processes watch events on the Endpoints resource.
func (p *TemplatePlugin) HandleEndpoints(eventType watch.EventType, endpoints *kapi.Endpoints) error {
	key := endpointsKey(endpoints)

	glog.V(4).Infof("Processing %d Endpoints for Name: %v (%v)", len(endpoints.Subsets), endpoints.Name, eventType)

	for i, s := range endpoints.Subsets {
		glog.V(4).Infof("  Subset %d : %#v", i, s)
	}

	if _, ok := p.Router.FindServiceUnit(key); !ok {
		p.Router.CreateServiceUnit(key)
	}

	switch eventType {
	case watch.Added, watch.Modified:
		glog.V(4).Infof("Modifying endpoints for %s", key)
		routerEndpoints := createRouterEndpoints(endpoints, !p.IncludeUDP, p.ServiceFetcher)
		key := endpointsKey(endpoints)
		p.Router.AddEndpoints(key, routerEndpoints)
	case watch.Deleted:
		glog.V(4).Infof("Deleting endpoints for %s", key)
		p.Router.DeleteEndpoints(key)
	}

	return nil
}

// HandleNode processes watch events on the Node resource
// The template type of plugin currently does not need to act on such events
// so the implementation just returns without error
func (p *TemplatePlugin) HandleNode(eventType watch.EventType, node *kapi.Node) error {
	return nil
}

// HandleRoute processes watch events on the Route resource.
// TODO: this function can probably be collapsed with the router itself, as a function that
//   determines which component needs to be recalculated (which template) and then does so
//   on demand.
func (p *TemplatePlugin) HandleRoute(eventType watch.EventType, route *routeapi.Route) error {
	switch eventType {
	case watch.Added, watch.Modified:
		p.Router.AddRoute(route)
	case watch.Deleted:
		glog.V(4).Infof("Deleting route %s/%s", route.Namespace, route.Name)
		p.Router.RemoveRoute(route)
	}
	return nil
}

// HandleNamespaces limits the scope of valid routes to only those that match
// the provided namespace list.
func (p *TemplatePlugin) HandleNamespaces(namespaces sets.String) error {
	p.Router.FilterNamespaces(namespaces)
	return nil
}

func (p *TemplatePlugin) Commit() error {
	p.Router.Commit()
	return nil
}

// endpointsKey returns the internal router key to use for the given Endpoints.
func endpointsKey(endpoints *kapi.Endpoints) string {
	return fmt.Sprintf("%s/%s", endpoints.Namespace, endpoints.Name)
}

// peerServiceKey may be used by the underlying router when handling endpoints to identify
// endpoints that belong to its peers.  THIS MUST FOLLOW THE KEY STRATEGY OF endpointsKey.  It
// receives a NamespacedName that is created from the service that is added by the oadm command
func peerEndpointsKey(namespacedName ktypes.NamespacedName) string {
	return fmt.Sprintf("%s/%s", namespacedName.Namespace, namespacedName.Name)
}

// createRouterEndpoints creates openshift router endpoints based on k8s endpoints
func createRouterEndpoints(endpoints *kapi.Endpoints, excludeUDP bool, lookupSvc ServiceLookup) []Endpoint {
	// check if this service is currently idled
	wasIdled := false
	subsets := endpoints.Subsets
	if _, ok := endpoints.Annotations[unidlingapi.IdledAtAnnotation]; ok && len(endpoints.Subsets) == 0 {
		service, err := lookupSvc.LookupService(endpoints)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("unable to find idled service corresponding to idled endpoints %s/%s: %v", endpoints.Namespace, endpoints.Name, err))
			return []Endpoint{}
		}

		if service.Spec.ClusterIP == "" {
			utilruntime.HandleError(fmt.Errorf("headless service %s/%s was marked as idled, but cannot setup unidling without a cluster IP", endpoints.Namespace, endpoints.Name))
			return []Endpoint{}
		}

		svcSubset := kapi.EndpointSubset{
			Addresses: []kapi.EndpointAddress{
				{
					IP: service.Spec.ClusterIP,
				},
			},
		}

		for _, port := range service.Spec.Ports {
			endptPort := kapi.EndpointPort{
				Name:     port.Name,
				Port:     port.Port,
				Protocol: port.Protocol,
			}
			svcSubset.Ports = append(svcSubset.Ports, endptPort)
		}

		subsets = []kapi.EndpointSubset{svcSubset}
		wasIdled = true
	}

	out := make([]Endpoint, 0, len(endpoints.Subsets)*4)

	// TODO: review me for sanity
	for _, s := range subsets {
		for _, p := range s.Ports {
			if excludeUDP && p.Protocol == kapi.ProtocolUDP {
				continue
			}
			for _, a := range s.Addresses {
				ep := Endpoint{
					ID:   fmt.Sprintf("%s:%d", a.IP, p.Port),
					IP:   a.IP,
					Port: strconv.Itoa(int(p.Port)),

					PortName: p.Name,

					NoHealthCheck: wasIdled,
				}
				if a.TargetRef != nil {
					ep.TargetName = a.TargetRef.Name
				} else {
					ep.TargetName = ep.IP
				}

				// IdHash contains an obfuscated internal IP address
				// that is the value passed in the cookie. The IP address
				// is made more difficult to extract by including other
				// internal information in the hash.
				s := ep.ID + ep.TargetName + ep.PortName
				ep.IdHash = fmt.Sprintf("%x", md5.Sum([]byte(s)))

				out = append(out, ep)
			}
		}
	}

	return out
}
