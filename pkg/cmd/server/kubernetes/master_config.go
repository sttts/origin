package kubernetes

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/glog"

	apiserveroptions "k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	cmapp "k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	"k8s.io/kubernetes/pkg/admission"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	"k8s.io/kubernetes/pkg/apis/batch"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/auth/authenticator"
	kclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/genericapiserver"
	"k8s.io/kubernetes/pkg/genericapiserver/authorizer"
	"k8s.io/kubernetes/pkg/master"
	"k8s.io/kubernetes/pkg/registry/cachesize"
	"k8s.io/kubernetes/pkg/registry/core/endpoint"
	endpointsetcd "k8s.io/kubernetes/pkg/registry/core/endpoint/etcd"
	"k8s.io/kubernetes/pkg/registry/generic"
	storagefactory "k8s.io/kubernetes/pkg/storage/storagebackend/factory"
	kerrors "k8s.io/kubernetes/pkg/util/errors"
	"k8s.io/kubernetes/pkg/util/intstr"
	knet "k8s.io/kubernetes/pkg/util/net"
	scheduleroptions "k8s.io/kubernetes/plugin/cmd/kube-scheduler/app/options"

	"github.com/openshift/origin/pkg/cmd/flagtypes"
	configapi "github.com/openshift/origin/pkg/cmd/server/api"
	"github.com/openshift/origin/pkg/cmd/server/election"
	cmdflags "github.com/openshift/origin/pkg/cmd/util/flags"
	"github.com/openshift/origin/pkg/controller/shared"
)

// MasterConfig defines the required values to start a Kubernetes master
type MasterConfig struct {
	Options    configapi.KubernetesMasterConfig
	KubeClient kclientset.Interface

	Master            *master.Config
	ControllerManager *cmapp.CMServer
	SchedulerServer   *scheduleroptions.SchedulerServer
	CloudProvider     cloudprovider.Interface

	Informers shared.InformerFactory
}

// BuildDefaultAPIServer constructs the appropriate APIServer and StorageFactory for the kubernetes server.
// It returns an error if no KubernetesMasterConfig was defined.
func BuildDefaultAPIServer(options configapi.MasterConfig) (*apiserveroptions.ServerRunOptions, genericapiserver.StorageFactory, error) {
	if options.KubernetesMasterConfig == nil {
		return nil, nil, fmt.Errorf("no kubernetesMasterConfig defined, unable to load settings")
	}
	_, portString, err := net.SplitHostPort(options.ServingInfo.BindAddress)
	if err != nil {
		return nil, nil, err
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, nil, err
	}

	portRange, err := knet.ParsePortRange(options.KubernetesMasterConfig.ServicesNodePortRange)
	if err != nil {
		return nil, nil, err
	}

	// Defaults are tested in TestAPIServerDefaults
	server := apiserveroptions.NewServerRunOptions()
	// Adjust defaults
	server.EventTTL = 2 * time.Hour
	server.GenericServerRunOptions.ServiceClusterIPRange = net.IPNet(flagtypes.DefaultIPNet(options.KubernetesMasterConfig.ServicesSubnet))
	server.GenericServerRunOptions.ServiceNodePortRange = *portRange
	server.GenericServerRunOptions.EnableProfiling = false
	server.GenericServerRunOptions.SecurePort = port
	server.GenericServerRunOptions.InsecurePort = 0
	server.GenericServerRunOptions.MasterCount = options.KubernetesMasterConfig.MasterCount
	server.GenericServerRunOptions.MaxRequestsInFlight = options.ServingInfo.MaxRequestsInFlight

	// resolve extended arguments
	// TODO: this should be done in config validation (along with the above) so we can provide
	// proper errors
	if err := cmdflags.Resolve(options.KubernetesMasterConfig.APIServerArguments, server.AddFlags); len(err) > 0 {
		return nil, nil, kerrors.NewAggregate(err)
	}

	resourceEncodingConfig := genericapiserver.NewDefaultResourceEncodingConfig()
	resourceEncodingConfig.SetVersionEncoding(
		kapi.GroupName,
		unversioned.GroupVersion{Group: kapi.GroupName, Version: options.EtcdStorageConfig.KubernetesStorageVersion},
		kapi.SchemeGroupVersion,
	)

	resourceEncodingConfig.SetVersionEncoding(
		extensions.GroupName,
		unversioned.GroupVersion{Group: extensions.GroupName, Version: "v1beta1"},
		extensions.SchemeGroupVersion,
	)

	resourceEncodingConfig.SetVersionEncoding(
		batch.GroupName,
		unversioned.GroupVersion{Group: batch.GroupName, Version: "v1"},
		batch.SchemeGroupVersion,
	)

	resourceEncodingConfig.SetVersionEncoding(
		autoscaling.GroupName,
		unversioned.GroupVersion{Group: autoscaling.GroupName, Version: "v1"},
		autoscaling.SchemeGroupVersion,
	)

	storageGroupsToEncodingVersion, err := server.GenericServerRunOptions.StorageGroupsToEncodingVersion()
	if err != nil {
		return nil, nil, err
	}

	// use the stock storage config based on args, but override bits from our config where appropriate
	etcdConfig := server.GenericServerRunOptions.StorageConfig
	etcdConfig.Prefix = options.EtcdStorageConfig.KubernetesStoragePrefix
	etcdConfig.ServerList = options.EtcdClientInfo.URLs
	etcdConfig.KeyFile = options.EtcdClientInfo.ClientCert.KeyFile
	etcdConfig.CertFile = options.EtcdClientInfo.ClientCert.CertFile
	etcdConfig.CAFile = options.EtcdClientInfo.CA

	storageFactory, err := genericapiserver.BuildDefaultStorageFactory(
		etcdConfig,
		server.GenericServerRunOptions.DefaultStorageMediaType,
		kapi.Codecs,
		genericapiserver.NewDefaultResourceEncodingConfig(),
		storageGroupsToEncodingVersion,
		// FIXME: this GroupVersionResource override should be configurable
		[]unversioned.GroupVersionResource{batch.Resource("cronjobs").WithVersion("v2alpha1")},
		master.DefaultAPIResourceConfigSource(), server.GenericServerRunOptions.RuntimeConfig,
	)
	if err != nil {
		return nil, nil, err
	}

	/*storageFactory := genericapiserver.NewDefaultStorageFactory(
		etcdConfig,
		server.DefaultStorageMediaType,
		kapi.Codecs,
		resourceEncodingConfig,
		master.DefaultAPIResourceConfigSource(),
	)*/
	// the order here is important, it defines which version will be used for storage
	storageFactory.AddCohabitatingResources(extensions.Resource("jobs"), batch.Resource("jobs"))
	storageFactory.AddCohabitatingResources(extensions.Resource("horizontalpodautoscalers"), autoscaling.Resource("horizontalpodautoscalers"))

	return server, storageFactory, nil
}

func BuildKubernetesMasterConfig(options configapi.MasterConfig, requestContextMapper kapi.RequestContextMapper, kubeClient kclientset.Interface, informers shared.InformerFactory, admissionControl admission.Interface, originAuthenticator authenticator.Request) (*MasterConfig, error) {
	if options.KubernetesMasterConfig == nil {
		return nil, errors.New("insufficient information to build KubernetesMasterConfig")
	}

	// in-order list of plug-ins that should intercept admission decisions
	// TODO: Push node environment support to upstream in future

	podEvictionTimeout, err := time.ParseDuration(options.KubernetesMasterConfig.PodEvictionTimeout)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PodEvictionTimeout: %v", err)
	}

	// Defaults are tested in TestCMServerDefaults
	cmserver := cmapp.NewCMServer()
	// Adjust defaults
	cmserver.Address = ""                   // no healthz endpoint
	cmserver.Port = 0                       // no healthz endpoint
	cmserver.EnableGarbageCollector = false // disabled until we add the controller
	cmserver.PodEvictionTimeout = unversioned.Duration{Duration: podEvictionTimeout}
	cmserver.VolumeConfiguration.EnableDynamicProvisioning = options.VolumeConfig.DynamicProvisioningEnabled

	// resolve extended arguments
	// TODO: this should be done in config validation (along with the above) so we can provide
	// proper errors
	if err := cmdflags.Resolve(options.KubernetesMasterConfig.ControllerArguments, cmserver.AddFlags); len(err) > 0 {
		return nil, kerrors.NewAggregate(err)
	}

	// resolve extended arguments
	// TODO: this should be done in config validation (along with the above) so we can provide
	// proper errors
	schedulerserver := scheduleroptions.NewSchedulerServer()
	schedulerserver.PolicyConfigFile = options.KubernetesMasterConfig.SchedulerConfigFile
	if err := cmdflags.Resolve(options.KubernetesMasterConfig.SchedulerArguments, schedulerserver.AddFlags); len(err) > 0 {
		return nil, kerrors.NewAggregate(err)
	}

	cloud, err := cloudprovider.InitCloudProvider(cmserver.CloudProvider, cmserver.CloudConfigFile)
	if err != nil {
		return nil, err
	}
	if cloud != nil {
		glog.V(2).Infof("Successfully initialized cloud provider: %q from the config file: %q\n", cmserver.CloudProvider, cmserver.CloudConfigFile)
	}

	var proxyClientCerts []tls.Certificate
	if len(options.KubernetesMasterConfig.ProxyClientInfo.CertFile) > 0 {
		clientCert, err := tls.LoadX509KeyPair(
			options.KubernetesMasterConfig.ProxyClientInfo.CertFile,
			options.KubernetesMasterConfig.ProxyClientInfo.KeyFile,
		)
		if err != nil {
			return nil, err
		}
		proxyClientCerts = append(proxyClientCerts, clientCert)
	}

	server, storageFactory, err := BuildDefaultAPIServer(options)
	if err != nil {
		return nil, err
	}

	// Preserve previous behavior of using the first non-loopback address
	// TODO: Deprecate this behavior and just require a valid value to be passed in
	publicAddress := net.ParseIP(options.KubernetesMasterConfig.MasterIP)
	if publicAddress == nil || publicAddress.IsUnspecified() || publicAddress.IsLoopback() {
		hostIP, err := knet.ChooseHostInterface()
		if err != nil {
			glog.Fatalf("Unable to find suitable network address.error='%v'. Set the masterIP directly to avoid this error.", err)
		}
		publicAddress = hostIP
		glog.Infof("Will report %v as public IP address.", publicAddress)
	}

	genericConfig := genericapiserver.NewConfig().ApplyOptions(server.GenericServerRunOptions)

	// MUST be in synced with the value in CMServer
	genericConfig.EnableGarbageCollection = false // disabled until we add the controller

	genericConfig.PublicAddress = publicAddress
	genericConfig.Authenticator = originAuthenticator // this is used to fulfill the tokenreviews endpoint which is used by node authentication
	genericConfig.Authorizer = authorizer.NewAlwaysAllowAuthorizer()
	genericConfig.AdmissionControl = admissionControl
	genericConfig.RequestContextMapper = requestContextMapper
	genericConfig.APIResourceConfigSource = getAPIResourceConfig(options)
	genericConfig.EnableIndex = false          // TODO(sttts): get rid of our indexAPIPaths and use this
	genericConfig.EnableOpenAPISupport = false // TODO(sttts): use this instead of our OpenAPI support
	genericConfig.EnableSwaggerSupport = false // TODO(sttts): use this instead of our Swagger support

	m := &master.Config{
		GenericConfig: genericConfig,
		MasterCount:   server.GenericServerRunOptions.MasterCount,

		// Set the TLS options for proxying to pods and services
		// Proxying to nodes uses the kubeletClient TLS config (so can provide a different cert, and verify the node hostname)
		ProxyTransport: knet.SetTransportDefaults(&http.Transport{
			TLSClientConfig: &tls.Config{
				// Proxying to pods and services cannot verify hostnames, since they are contacted on randomly allocated IPs
				InsecureSkipVerify: true,
				Certificates:       proxyClientCerts,
			},
		}),

		EnableWatchCache:          server.GenericServerRunOptions.EnableWatchCache,
		KubernetesServiceNodePort: server.GenericServerRunOptions.KubernetesServiceNodePort,
		ServiceIPRange:            server.GenericServerRunOptions.ServiceClusterIPRange,
		ServiceNodePortRange:      server.GenericServerRunOptions.ServiceNodePortRange,

		StorageFactory: storageFactory,

		EventTTL: server.EventTTL,

		KubeletClientConfig: *configapi.GetKubeletClientConfig(options),

		EnableLogsSupport:     false, // don't expose server logs
		EnableCoreControllers: true,

		DeleteCollectionWorkers: server.GenericServerRunOptions.DeleteCollectionWorkers,
	}

	if m.EnableWatchCache {
		cachesize.SetWatchCacheSizes(server.GenericServerRunOptions.WatchCacheSizes)
	}

	if m.EnableCoreControllers {
		glog.V(2).Info("Using the lease endpoint reconciler")
		config, err := m.StorageFactory.NewConfig(kapi.Resource("apiServerIPInfo"))
		if err != nil {
			return nil, err
		}
		leaseStorage, _, err := storagefactory.Create(*config)
		if err != nil {
			return nil, err
		}
		masterLeases := newMasterLeases(leaseStorage)

		endpointConfig, err := m.StorageFactory.NewConfig(kapi.Resource("endpoints"))
		if err != nil {
			return nil, err
		}
		endpointsStorage := endpointsetcd.NewREST(generic.RESTOptions{
			StorageConfig:           endpointConfig,
			Decorator:               generic.UndecoratedStorage,
			DeleteCollectionWorkers: 0,
			ResourcePrefix:          m.StorageFactory.ResourcePrefix(kapi.Resource("endpoints")),
		})

		endpointRegistry := endpoint.NewRegistry(endpointsStorage)

		m.EndpointReconcilerConfig = master.EndpointReconcilerConfig{
			Reconciler: election.NewLeaseEndpointReconciler(endpointRegistry, masterLeases),
			Interval:   master.DefaultEndpointReconcilerInterval,
		}
	}

	if options.DNSConfig != nil {
		_, dnsPortStr, err := net.SplitHostPort(options.DNSConfig.BindAddress)
		if err != nil {
			return nil, fmt.Errorf("unable to parse DNS bind address %s: %v", options.DNSConfig.BindAddress, err)
		}
		dnsPort, err := strconv.Atoi(dnsPortStr)
		if err != nil {
			return nil, fmt.Errorf("invalid DNS port: %v", err)
		}
		m.ExtraServicePorts = append(m.ExtraServicePorts,
			kapi.ServicePort{Name: "dns", Port: 53, Protocol: kapi.ProtocolUDP, TargetPort: intstr.FromInt(dnsPort)},
			kapi.ServicePort{Name: "dns-tcp", Port: 53, Protocol: kapi.ProtocolTCP, TargetPort: intstr.FromInt(dnsPort)},
		)
		m.ExtraEndpointPorts = append(m.ExtraEndpointPorts,
			kapi.EndpointPort{Name: "dns", Port: int32(dnsPort), Protocol: kapi.ProtocolUDP},
			kapi.EndpointPort{Name: "dns-tcp", Port: int32(dnsPort), Protocol: kapi.ProtocolTCP},
		)
	}

	kmaster := &MasterConfig{
		Options:    *options.KubernetesMasterConfig,
		KubeClient: kubeClient,

		Master:            m,
		ControllerManager: cmserver,
		CloudProvider:     cloud,
		SchedulerServer:   schedulerserver,
		Informers:         informers,
	}

	return kmaster, nil
}

// getAPIResourceConfig builds the config for enabling resources
func getAPIResourceConfig(options configapi.MasterConfig) genericapiserver.APIResourceConfigSource {
	resourceConfig := genericapiserver.NewResourceConfig()

	for group := range configapi.KnownKubeAPIGroups {
		for _, version := range configapi.GetEnabledAPIVersionsForGroup(*options.KubernetesMasterConfig, group) {
			gv := unversioned.GroupVersion{Group: group, Version: version}
			resourceConfig.EnableVersions(gv)
		}
	}

	for group := range options.KubernetesMasterConfig.DisabledAPIGroupVersions {
		for _, version := range configapi.GetDisabledAPIVersionsForGroup(*options.KubernetesMasterConfig, group) {
			gv := unversioned.GroupVersion{Group: group, Version: version}
			resourceConfig.DisableVersions(gv)
		}
	}

	return resourceConfig
}
