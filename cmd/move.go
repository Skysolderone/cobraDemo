package cmd

import (
	"v1/service"

	"github.com/spf13/cobra"
)

var moveCommand = &cobra.Command{
	Use:   "move",
	Short: "move exe",
	Run:   service.MoveService,
}

func init() {
	rootCmd.AddCommand(moveCommand)
}
