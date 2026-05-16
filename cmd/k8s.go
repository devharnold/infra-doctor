package cmd

import (
	"fmt"
	"github.com/devharnold/infra-doctor/internal/k8s"

	"github.com/spf13/cobra"
)

var k8sCmd = &cobra.Command{
	Use: "k8s",
	Short: "Check Kubernetes Cluster",
	Run: func(cmd *cobra.Command, args [] string) {
		err := k8s.CheckPods()
		if err != nil {
			fmt.Println("K8s check failed", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)
}