package cmd

import (
	"fmt"
	"os"

	"v1/service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "v2ray",
	Short: "v2ray wws process",
	Run:   service.StartV2ray,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
