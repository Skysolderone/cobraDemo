package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionFunc(cmd *cobra.Command, arges []string) {
	fmt.Println("version 0.0.1")
}
