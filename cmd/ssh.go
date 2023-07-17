package cmd

import "fmt"

type SSH struct {
}

func (s SSH) Run(task Task, callback func(string, string, string)) {
	processes := task.GetProcesses()
	for _, process := range processes {
		if err := process.Run(callback); err != nil {
			fmt.Println(err)
		}
	}
}
