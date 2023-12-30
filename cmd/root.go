package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "benthic",
	Short:   "A CLI to interact with Benthic Cloud",
	Long:    `A command line interface to interact with Benthic Cloud`,
	Example: "benthic hello",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
