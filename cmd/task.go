package cmd

import (
	"fmt"
	"os/exec"
)

type Task struct {
	name     string
	hosts    []string
	script   string
	parallel bool
}

func (t Task) GetProcess(host string) Process {
	localhosts := []string{"local", "localhost", "127.0.0.1"}

	if InSlice(host, localhosts) {
		return Process{
			target:  host,
			command: exec.Command("/bin/bash", "-c", t.script),
		}
	}

	command := fmt.Sprintf(`bash -se << \EOF-ENVOY

set -e
%s
EOF-ENVOY`, t.script)

	return Process{
		target:  host,
		command: exec.Command("ssh", host, command),
	}
}

func (t Task) GetProcesses() []Process {
	processes := []Process{}
	for _, v := range t.hosts {
		processes = append(processes, t.GetProcess(v))
	}
	return processes
}
