package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Prints information about the user",
	Run: func(cmd *cobra.Command, args []string) {
		// print the team id from the rootCmd
		fmt.Println("Team ID: ", TeamID)
	},
}
