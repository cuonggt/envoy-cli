package cmd

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <task>",
	Short: "Run an Envoy task.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container := LoadTaskContainer()

		task, ok := container.Tasks[args[0]]

		if !ok {
			fmt.Printf("Task %s is not defined.\n", args[0])
			return
		}

		target := container.GetFirstServer().hosts[0]
		fmt.Println("Connecting to " + target + "...")

		command := fmt.Sprintf(`bash -se \EOF

set -e
%s
EOF`, task.script)

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
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
