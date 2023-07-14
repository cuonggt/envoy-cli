package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pretend bool

func runTask(container TaskContainer, name string) {
	task, err := container.GetTask(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	if pretend {
		fmt.Printf("%s", task.script)
		return
	}

	for _, host := range task.hosts {
		err := task.Run(host, func(target string, line string) {
			fmt.Printf("[%s]: %s", target, line)
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

var runCmd = &cobra.Command{
	Use:   "run <task>",
	Short: "Run an Envoy task.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container := LoadTaskContainer()

		tasks := []string{args[0]}

		for _, v := range tasks {
			runTask(container, v)
		}
	},
}

func init() {
	runCmd.Flags().BoolVarP(&pretend, "pretend", "p", false, "Dump Bash script for inspection.")
	rootCmd.AddCommand(runCmd)
}
