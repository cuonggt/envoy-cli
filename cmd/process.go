package cmd

import (
	"bufio"
	"os/exec"
)

type Process struct {
	target  string
	command *exec.Cmd
}

func (p Process) Run(callback func(string, string)) error {
	pipe, _ := p.command.StdoutPipe()
	if err := p.command.Start(); err != nil {
		return err
	}
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		callback(p.target, line)
		line, err = reader.ReadString('\n')
	}
	return nil
}
