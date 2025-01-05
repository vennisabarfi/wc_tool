package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Printers(cmd *cobra.Command, args []string) {
	// check if flag has been set by user
	if cmd.Flags().Changed("message") {
		message, _ := cmd.Flags().GetString("message")

		fmt.Print("Hi lol lol ", message)
	}

}
