package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh <name>",
	Short: "Connect to an Envoy server.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container := LoadTaskContainer()
		var server Server
		if len(args) == 1 {
			server = container.GetServer(args[0])
		} else if container.HasOneServer() {
			server = container.GetFirstServer()
		} else {
			fmt.Println("Please provide a server name.")
			return
		}

		ssh := exec.Command("ssh", server.hosts[0])

		ssh.Stdin = os.Stdin
		ssh.Stderr = os.Stderr
		ssh.Stdout = os.Stdout

		err := ssh.Run()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
