package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "graphql-codegen version",
	Long:  `graphql-codegen version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
