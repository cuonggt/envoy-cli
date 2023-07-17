package cmd

import "fmt"

type SSH struct {
}

func (s SSH) Run(task Task) {
	processes := task.GetProcesses()
	for _, process := range processes {
		if err := process.Run(DisplayOutput); err != nil {
			fmt.Println(err)
		}
	}
}
