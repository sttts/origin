// This is a generated file. Do not edit directly.

module k8s.io/legacy-cloud-providers

go 1.12

require (
	cloud.google.com/go v0.38.0
	github.com/Azure/azure-sdk-for-go v35.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.0
	github.com/Azure/go-autorest/autorest/adal v0.5.0
	github.com/Azure/go-autorest/autorest/to v0.2.0
	github.com/Azure/go-autorest/autorest/validation v0.1.0 // indirect
	github.com/GoogleCloudPlatform/k8s-cloud-provider v0.0.0-20190822182118-27a4ced34534
	github.com/aws/aws-sdk-go v1.16.26
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/gophercloud/gophercloud v0.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/rubiojr/go-vhd v0.0.0-20160810183302-0bfd3b39853c
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/vmware/govmomi v0.20.3
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.6.1-0.20190607001116-5213b8090861
	gopkg.in/gcfg.v1 v1.2.0
	gopkg.in/warnings.v0 v0.1.1 // indirect
	k8s.io/api v0.0.0-20191204082340-384b28a90b2b
	k8s.io/apimachinery v0.0.0-20191121175448-79c2a76c473a
	k8s.io/apiserver v0.0.0-20191204085103-2ce178ac32b7
	k8s.io/client-go v0.0.0-20191204083517-ea72ff2b5b2f
	k8s.io/cloud-provider v0.0.0-20191204093314-173cef1bc308
	k8s.io/component-base v0.0.0-20191204084121-18d14e17701e
	k8s.io/csi-translation-lib v0.0.0-20191204093550-9ba76651afc2
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
	sigs.k8s.io/yaml v1.1.0
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20191204082340-384b28a90b2b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191121175448-79c2a76c473a
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191204085103-2ce178ac32b7
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191204083517-ea72ff2b5b2f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191204093314-173cef1bc308
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191204084121-18d14e17701e
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191204093550-9ba76651afc2
)
