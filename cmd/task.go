package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
)

type Task struct {
	name   string
	hosts  []string
	script string
}

func (t *Task) GetProcess(host string) (string, *exec.Cmd) {
	localhosts := []string{"local", "localhost", "127.0.0.1"}

	if InSlice(host, localhosts) {
		return host, exec.Command("/bin/bash", "-c", t.script)
	}

	command := fmt.Sprintf(`bash -se \EOF-ENVOY

set -e
%s
EOF-ENVOY`, t.script)

	return host, exec.Command("ssh", host, command)
}

func (t *Task) Run(host string, callback func(string, string)) error {
	target, process := t.GetProcess(host)

	pipe, _ := process.StdoutPipe()

	if err := process.Start(); err != nil {
		return err
	}

	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		callback(target, line)
		line, err = reader.ReadString('\n')
	}

	return nil
}
