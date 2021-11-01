/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package get

import (
	"fmt"
	"omc/cmd/get/apps"
	"omc/cmd/get/batch"
	"omc/cmd/get/core"
	"omc/cmd/get/local"
	appz "omc/cmd/get/openshift/apps"
	"omc/cmd/get/openshift/build"
	"omc/cmd/get/openshift/config"
	"omc/cmd/get/openshift/image"
	"omc/cmd/get/openshift/machine"
	"omc/cmd/get/openshift/machineconfiguration"
	"omc/cmd/get/openshift/route"
	"omc/cmd/get/storage"
	"omc/vars"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var outputStringVar string
var allNamespaceBoolVar, showLabelsBoolVar bool

var GetCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called", args)
	},
}

func init() {
	if os.Args[1] == "get" && len(os.Args) > 2 {
		if strings.Contains(os.Args[2], "/") {
			seg := strings.Split(os.Args[2], "/")
			resource, name := seg[0], seg[1]
			os.Args = append([]string{os.Args[0], "get", resource, name}, os.Args[3:]...)
		}
	}

	GetCmd.PersistentFlags().BoolVarP(&vars.AllNamespaceBoolVar, "all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces.")
	GetCmd.PersistentFlags().BoolVarP(&vars.ShowLabelsBoolVar, "show-labels", "", false, "When printing, show all labels as the last column (default hide labels column)")
	GetCmd.PersistentFlags().StringVarP(&vars.OutputStringVar, "output", "o", "", "Output format. One of: json|yaml|wide|jsonpath=...")
	//getCmd.PersistentFlags().StringVarP(&selector, "selector", "l", "", "elector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	GetCmd.AddCommand(
		apps.DaemonSet,
		apps.Deployment,
		apps.ReplicaSet,
		appz.DeploymentConfig,
		batch.Job,
		build.Build,
		build.BuildConfig,
		config.ClusterOperator,
		config.ClusterVersion,
		core.ConfigMap,
		core.Event,
		core.Namespace,
		core.Node,
		core.PersistentVolume,
		core.PersistentVolumeClaim,
		core.Pod,
		core.ReplicationController,
		core.Secret,
		core.Service,
		image.ImageStream,
		machine.Machine,
		machine.MachineSet,
		machineconfiguration.MachineConfig,
		machineconfiguration.MachineConfigPool,
		local.MustGather,
		route.Route,
		storage.StorageClass,
	)
}