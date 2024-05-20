package cmd

import (
	"v1/service"

	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "get version",
	Run:   service.VersionFunc,
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
