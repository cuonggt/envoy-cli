package cmd

import (
	"bufio"
	"os/exec"
	"sync"
)

type Process struct {
	Target  string
	Command *exec.Cmd
}

func (p Process) Run(callback func(string, string, string)) error {
	stdout, err := p.Command.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := p.Command.StderrPipe()
	if err != nil {
		return err
	}

	if err = p.Command.Start(); err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		outScanner := bufio.NewScanner(stdout)
		for outScanner.Scan() {
			line := outScanner.Text()
			callback("out", p.Target, line)
		}
	}()

	go func() {
		defer wg.Done()

		errScanner := bufio.NewScanner(stderr)
		for errScanner.Scan() {
			line := errScanner.Text()
			callback("err", p.Target, line)
		}
	}()

	wg.Wait()

	return p.Command.Wait()
}
