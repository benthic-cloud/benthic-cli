package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Starts the Device Authorization Flow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is where the authorization flow will start.")
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
