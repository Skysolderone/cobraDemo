package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MoveService(cmd *cobra.Command, args []string) {
	for _, v := range args {
		switch v {
		case "tgbot":
			fmt.Println("move tgbot")
		case "trading":
			fmt.Println("move trading")
		}
	}
}
