package cmd

import (
	"fmt"
	"os/exec"

	"golang.org/x/exp/slices"
)

type Task struct {
	Name     string
	Hosts    []string
	Script   string
	Parallel bool
}

func (t Task) GetProcess(host string) Process {
	localhosts := []string{"local", "localhost", "127.0.0.1"}

	if slices.Contains(localhosts, host) {
		return Process{
			Target:  host,
			Command: exec.Command("/bin/bash", "-c", t.Script),
		}
	}

	command := fmt.Sprintf(`bash -se << \EOF-ENVOY

set -e
%s
EOF-ENVOY`, t.Script)

	return Process{
		Target:  host,
		Command: exec.Command("ssh", host, command),
	}
}

func (t Task) GetProcesses() []Process {
	processes := []Process{}
	for _, v := range t.Hosts {
		processes = append(processes, t.GetProcess(v))
	}
	return processes
}
