package cmd

import (
	"fmt"
	"sync"
)

type ParallelSSH struct {
}

func (s ParallelSSH) Run(task Task, callback func(string, string, string)) {
	processes := task.GetProcesses()

	var wg sync.WaitGroup

	wg.Add(len(processes))

	for _, process := range processes {
		process := process

		go func() {
			defer wg.Done()
			if err := process.Run(callback); err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}
