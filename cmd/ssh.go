package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func getServer(container TaskContainer, args []string) (*Server, error) {
	if len(args) == 1 {
		return container.GetServer(args[0])
	}

	if container.HasOneServer() {
		server := container.GetFirstServer()
		return &server, nil
	}

	return nil, fmt.Errorf("%s", "Please provide a server name")
}

var sshCmd = &cobra.Command{
	Use:   "ssh <name>",
	Short: "Connect to an Envoy server.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container := LoadTaskContainer()
		server, err := getServer(container, args)
		if err != nil {
			fmt.Println(err)
			return
		}

		ssh := exec.Command("ssh", server.hosts[0])

		ssh.Stdin = os.Stdin
		ssh.Stderr = os.Stderr
		ssh.Stdout = os.Stdout

		if err := ssh.Run(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
