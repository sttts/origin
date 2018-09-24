package clusterup

import (
	"fmt"
	"os"
	"strings"

	"github.com/openshift/origin/pkg/cmd/util/variable"
	dockerutil "github.com/openshift/origin/pkg/oc/clusterup/docker/util"
)

// OpenShiftImages specifies a list of images cluster up require to pull in order to bootstrap a cluster.
var OpenShiftImages = Images{
	// OpenShift Images
	{Name: "cluster-kube-apiserver-operator"},
	{Name: "control-plane"},
	{Name: "cli"},
	{Name: "hyperkube"},
	{Name: "hypershift"},
	{Name: "node"},
	{Name: "pod"},

	// External images
	{Name: "bootkube", PullSpec: "quay.io/coreos/bootkube:v0.13.0"},
	{Name: "etcd", PullSpec: "quay.io/coreos/etcd:v3.2.24"},
}

var defaultTemplate = variable.NewDefaultImageTemplate()

type Image struct {
	Name string

	// PullSpec if specified is used instead of expanding the name via template. Used for non-openshift images.
	PullSpec string
}

func (i *Image) ToPullSpec() string {
	if len(i.PullSpec) > 0 {
		return i.PullSpec
	}
	return defaultTemplate.ExpandOrDie(i.Name)
}

func (i *Image) Pull(puller *dockerutil.Helper) error {
	return puller.CheckAndPull(i.ToPullSpec(), os.Stdout)
}

type Images []Image

func (i Images) EnsurePulled(puller *dockerutil.Helper) error {
	errors := []error{}
	for _, image := range i {
		if err := image.Pull(puller); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) == 0 {
		return nil
	}
	msgs := []string{}
	for _, err := range errors {
		msgs = append(msgs, err.Error())
	}
	return fmt.Errorf("some images failed to pull: %s", strings.Join(msgs, ","))
}

func (i Images) Get(name string) *Image {
	for _, image := range i {
		if image.Name == name {
			return &image
		}
	}
	return nil
}
