package cmd

import (
	"v1/service"

	"github.com/spf13/cobra"
)

var (
	v2rayStart = &cobra.Command{
		Use:   "start",
		Short: "start v2ray",
		Run:   service.StartV2ray,
	}
	v2rayClose = &cobra.Command{
		Use:   "close",
		Short: "close v2ray",
		Run:   service.CloseV2ray,
	}
)

func init() {
	rootCmd.AddCommand(v2rayStart)
	rootCmd.AddCommand(v2rayClose)
}
