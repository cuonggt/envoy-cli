package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <host>",
	Short: "Create a new Envoy file in the current directory.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()

		_, err := os.Stat(cwd + "/Envoyfile")
		if err == nil {
			fmt.Println("Envoy file already exists!")
			return
		}

		f, err := os.Create(cwd + "/Envoyfile")
		if err != nil {
			fmt.Println(err)
		}

		f.WriteString(fmt.Sprintf(`servers:
  web:
    - %s

tasks:
  deploy:
    name: Deploy
    script: |
      cd /path/to/site
      git pull origin master`, args[0]))

		f.Close()

		fmt.Println("Envoy file created!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
