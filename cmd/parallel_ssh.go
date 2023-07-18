package cmd

import (
	"os/exec"
	"sync"
)

type ParallelSSH struct{}

func (s ParallelSSH) Run(task Task, callback func(string, string, string)) int {
	processes := task.GetProcesses()

	code := 0

	var wg sync.WaitGroup

	wg.Add(len(processes))

	for _, process := range processes {
		process := process

		go func() {
			defer wg.Done()
			if err := process.Run(callback); err != nil {
				exitErr, ok := err.(*exec.ExitError)
				if ok {
					code += exitErr.ExitCode()
				} else {
					code++
				}
			}
		}()
	}

	wg.Wait()

	return code
}
