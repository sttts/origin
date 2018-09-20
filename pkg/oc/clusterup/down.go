package clusterup

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"

	"github.com/openshift/origin/pkg/oc/clusterup/docker/openshift"
	"github.com/openshift/origin/pkg/oc/clusterup/docker/util"
)

const CmdDownRecommendedName = "down"

var (
	cmdDownLong = templates.LongDesc(`
		Stops the container running OpenShift on Docker and associated containers.`)

	cmdDownExample = templates.Examples(`
	  # Stop local OpenShift cluster
	  %[1]s`)
)

type ClientStopConfig struct {
}

// NewCmdDown creates a command that stops OpenShift
func NewCmdDown(name, fullName string) *cobra.Command {
	config := &ClientStopConfig{}
	cmd := &cobra.Command{
		Use:     name,
		Short:   "Stop OpenShift on Docker",
		Long:    cmdDownLong,
		Example: fmt.Sprintf(cmdDownExample, fullName),
		Run: func(c *cobra.Command, args []string) {
			kcmdutil.CheckErr(config.Stop())
		},
	}
	return cmd
}

// Stop stops the currently running origin container and any
// containers started by the node.
func (c *ClientStopConfig) Stop() error {
	client, err := util.GetDockerClient()
	if err != nil {
		return err
	}
	helper := util.NewHelper(client)
	glog.V(4).Infof("Killing previous socat tunnel")
	err = openshift.KillExistingSocat()
	if err != nil {
		glog.V(2).Infof("error: cannot kill socat: %v", err)
	}
	glog.V(4).Infof("Stopping and removing origin container")
	if err = helper.StopAndRemoveContainer("origin"); err != nil {
		glog.V(2).Infof("Error stopping origin container: %v", err)
	}

	names, err := helper.ListContainerNames()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, name := range names {
		name = strings.TrimLeft(name, "/")
		if !openshift.ClusterUpContainers.Has(name) {
			if _, err = parseDockerName(name); err != nil {
				continue
			}
		}

		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			glog.V(4).Infof("Stopping container %s", name)
			if err = client.ContainerStop(name, 0); err != nil {
				glog.V(2).Infof("Error stopping container %s: %v", name, err)
			}
			glog.V(4).Infof("Removing container %s", name)
			if err = helper.RemoveContainer(name); err != nil {
				glog.V(2).Infof("Error removing container %s: %v", name, err)
			}
		}(name)
	}
	wg.Wait()

	return nil
}

// Unpacks a container name, returning the pod full name and container name we would have used to
// construct the docker name. If we are unable to parse the name, an error is returned.
func parseDockerName(name string) (hash uint64, err error) {
	const containerNamePrefix = "k8s"
	// For some reason docker appears to be appending '/' to names.
	// If it's there, strip it.
	name = strings.TrimPrefix(name, "/")
	parts := strings.Split(name, "_")
	if len(parts) == 0 || parts[0] != containerNamePrefix {
		err = fmt.Errorf("failed to parse Docker container name %q into parts", name)
		return 0, err
	}
	if len(parts) < 6 {
		// We have at least 5 fields.  We may have more in the future.
		// Anything with less fields than this is not something we can
		// manage.
		glog.Warningf("found a container with the %q prefix, but too few fields (%d): %q", containerNamePrefix, len(parts), name)
		err = fmt.Errorf("Docker container name %q has less parts than expected %v", name, parts)
		return 0, err
	}

	nameParts := strings.Split(parts[1], ".")
	if len(nameParts) > 1 {
		hash, err = strconv.ParseUint(nameParts[1], 16, 32)
		if err != nil {
			glog.Warningf("invalid container hash %q in container %q", nameParts[1], name)
		}
	}

	return hash, nil
}
