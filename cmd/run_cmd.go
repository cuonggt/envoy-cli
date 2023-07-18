package cmd

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
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

	runTaskOverSSH(*task)
}

func runTaskOverSSH(task Task) {
	if pretend {
		fmt.Printf("%s", task.script)
		return
	}

	passToRemoteProcessor(task)
}

func passToRemoteProcessor(task Task) {
	getRemoteProcessor(task).Run(task, func(outType string, host string, line string) {
		if strings.HasPrefix(line, "Warning: Permanently added ") {
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			return
		}
		hostColor := getHostColor(host)
		hostColor.Printf("[%s]", host)
		if outType == "err" {
			color.Red.Printf(": %s\n", line)
		} else {
			color.Printf(": %s\n", line)
		}
	})
}

func getRemoteProcessor(task Task) RemoteProcessor {
	if task.parallel {
		return ParallelSSH{}
	}
	return SSH{}
}

var hostWithColors = []string{}

func getHostColor(host string) color.Color {
	colors := []color.Color{color.Yellow, color.Cyan, color.Magenta, color.Blue}
	if !slices.Contains(hostWithColors, host) {
		hostWithColors = append(hostWithColors, host)
	}

	return colors[slices.Index(hostWithColors, host)%len(colors)]
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
