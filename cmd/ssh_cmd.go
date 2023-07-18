package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func getServer(container TaskContainer, args []string) (*Server, error) {
	var name string
	if len(args) < 1 {
		name = "web"
	} else {
		name = args[0]
	}
	return container.GetServer(name)
}

var sshCmd = &cobra.Command{
	Use:   "ssh <name>",
	Short: "Connect to an Envoy server.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container := LoadTaskContainer()

		server, err := getServer(container, args)
		if err != nil {
			output.Error(fmt.Sprintf("%s", err))
			return
		}

		ssh := exec.Command("ssh", server.hosts[0])

		ssh.Stdin = os.Stdin
		ssh.Stderr = os.Stderr
		ssh.Stdout = os.Stdout

		if err := ssh.Run(); err != nil {
			output.Error(fmt.Sprintf("%s", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
