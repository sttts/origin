package cm

import (
	"github.com/spf13/pflag"

	kcmapp "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	kcmoptions "k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
)

var (
	// TODO: adapt for origin
	ControllersDisabledByDefault = kcmapp.ControllersDisabledByDefault
)

func OriginControllerManagerAddFlags(cmserver *kcmoptions.CMServer) func(flags *pflag.FlagSet) {
	return func(flags *pflag.FlagSet) {
		cmserver.AddFlags(flags, kcmapp.KnownControllers(), ControllersDisabledByDefault.List())
	}
}
