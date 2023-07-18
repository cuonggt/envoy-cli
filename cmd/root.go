package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envoy",
	Short: "Elegant SSH tasks for artisans.",
	Long:  `Envoy is a tool for executing common tasks you run on your remote servers.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var output = ConsoleOutput{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
