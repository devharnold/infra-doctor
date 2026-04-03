package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Infra-Doctor",
	Short: "Diagnose Infrastructure Issues",
	Long:  "A CLI Tool to diagnose APIs, Kubernetes Clusters and System Integrations",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}