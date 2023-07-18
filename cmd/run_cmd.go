package cmd

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var pretend bool
var hostWithColors = []string{}

func getTasks(container TaskContainer, taskName string) []string {
	if story := container.GetStory(taskName); story != nil {
		return story
	}
	return []string{taskName}
}

func runTask(container TaskContainer, name string) int {
	task, err := container.GetTask(name)
	if err != nil {
		return 1
	}

	return runTaskOverSSH(*task)
}

func runTaskOverSSH(task Task) int {
	if pretend {
		fmt.Printf("%s", task.script)
		return 1
	}

	return passToRemoteProcessor(task)
}

func passToRemoteProcessor(task Task) int {
	return getRemoteProcessor(task).Run(task, func(outType string, host string, line string) {
		if strings.HasPrefix(line, "Warning: Permanently added ") {
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			return
		}
		hostColor := getHostColor(host).Sprintf("[%s]", host)
		if outType == "err" {
			line = color.Red.Sprintf("%s", line)
		} else {
			line = color.Sprintf("%s", line)
		}
		fmt.Printf("%s: %s\n", hostColor, line)
	})
}

func getRemoteProcessor(task Task) RemoteProcessor {
	if task.parallel {
		return ParallelSSH{}
	}
	return SSH{}
}

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
			thisCode := runTask(container, v)

			if thisCode > 0 {
				fmt.Printf("[%s] %s\n", color.Red.Sprint("✗"), color.Red.Sprint("This task did not complete successfully on one of your servers."))
				break
			}
		}
	},
}

func init() {
	runCmd.Flags().BoolVarP(&pretend, "pretend", "p", false, "Dump Bash script for inspection.")
	rootCmd.AddCommand(runCmd)
}
