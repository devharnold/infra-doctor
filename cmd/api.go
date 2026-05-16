package cmd

import (
	"fmt"
	"github.com/devharnold/infra-doctor/internal/api"

	"github.com/spf13/cobra"
)

var apiURL string

var apiCmd = &cobra.Command{
	Use: "api",
	Short: "Check API Health",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.Check(apiURL)
		if err != nil {
			fmt.Println("API check failed", err)
			return
		}
		fmt.Println("API is healthy")
	},
}

func init() {
	apiCmd.Flags().StringVar(&apiURL, "url", "", "API to check")
	apiCmd.MarkFlagRequired("url")
}