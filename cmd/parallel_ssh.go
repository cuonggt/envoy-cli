package cmd

import (
	"fmt"
	"sync"
)

type ParallelSSH struct {
}

func (s ParallelSSH) Run(task Task) {
	processes := task.GetProcesses()

	var wg sync.WaitGroup

	wg.Add(len(processes))

	for _, process := range processes {
		process := process

		go func() {
			defer wg.Done()
			if err := process.Run(DisplayOutput); err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}
