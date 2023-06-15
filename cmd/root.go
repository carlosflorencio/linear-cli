package cmd

import (
	"fmt"
	"os"

	"github.com/carlosflorencio/linear-cli/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linear",
	Short: "A CLI tool to interact with the Linear API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		value := config.ApiKey()
		if value == "" {
			fmt.Printf("Error: Required environment variable %s is not set.\n", config.ApiKeyEnvVar)
			os.Exit(1)
		}
	},
}

var TeamID string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&TeamID, "team", "t", "", "Team ID")
	// rootCmd.MarkFlagRequired("team")
}
