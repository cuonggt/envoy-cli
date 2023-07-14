package cmd

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var pretend bool

func getTasks(container TaskContainer, name string) []string {
	return []string{name}
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

	target := container.GetFirstServer().hosts[0]

	command := fmt.Sprintf(`bash -se \EOF-ENVOY

set -e
%s
EOF-ENVOY`, task.script)

	process := exec.Command("ssh", target, command)

	pipe, _ := process.StdoutPipe()

	if err := process.Start(); err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Printf("[%s]: %s", target, line)
		line, err = reader.ReadString('\n')
	}
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
