package cmd

import (
	"bufio"
	"os/exec"
	"sync"
)

type Process struct {
	target  string
	command *exec.Cmd
}

func (p Process) Run(callback func(string, string, string)) error {
	stdout, err := p.command.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := p.command.StderrPipe()
	if err != nil {
		return err
	}

	if err = p.command.Start(); err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		outScanner := bufio.NewScanner(stdout)
		for outScanner.Scan() {
			line := outScanner.Text()
			callback("out", p.target, line)
		}
	}()

	go func() {
		defer wg.Done()

		errScanner := bufio.NewScanner(stderr)
		for errScanner.Scan() {
			line := errScanner.Text()
			callback("err", p.target, line)
		}
	}()

	wg.Wait()

	return nil
}
