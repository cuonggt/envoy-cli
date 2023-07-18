package cmd

import (
	"os/exec"
)

type SSH struct {
}

func (s SSH) Run(task Task, callback func(string, string, string)) int {
	processes := task.GetProcesses()
	code := 0
	for _, process := range processes {
		if err := process.Run(callback); err != nil {
			exitErr, ok := err.(*exec.ExitError)
			if ok {
				code += exitErr.ExitCode()
			} else {
				code++
			}
		}
	}

	return code
}
