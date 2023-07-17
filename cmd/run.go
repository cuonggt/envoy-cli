package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pretend bool

func getTasks(container TaskContainer, taskName string) []string {
	if story := container.GetStory(taskName); story != nil {
		return story
	}
	return []string{taskName}
}

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
		if err := task.Run(host, DisplayOutput); err != nil {
			fmt.Println(err)
		}
	}
}

func DisplayOutput(host string, line string) {
	fmt.Printf("[%s]: %s", host, line)
}

var runCmd = &cobra.Command{
	Use:   "run <task>",
	Short: "Run an Envoy task.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container := LoadTaskContainer()

		tasks := getTasks(container, args[0])

		for _, v := range tasks {
			runTask(container, v)
		}
	},
}

func init() {
	runCmd.Flags().BoolVarP(&pretend, "pretend", "p", false, "Dump Bash script for inspection.")
	rootCmd.AddCommand(runCmd)
}
