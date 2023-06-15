package cmd

import (
	"fmt"

	"github.com/carlosflorencio/linear-cli/linear"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Prints information about the user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Team ID: ", TeamID)

		linear.User()
	},
}
