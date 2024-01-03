package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fcnCmd = &cobra.Command{
	Use:   "fcn",
	Short: "Commands to interact with fcn",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a function",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Running function...")
}

func init() {
	fcnCmd.AddCommand(runCmd)
	RootCmd.AddCommand(fcnCmd)
}
