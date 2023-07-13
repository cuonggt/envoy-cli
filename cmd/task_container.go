package cmd

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	hosts []string
}

type Task struct {
	name   string
	script string
}

type TaskContainer struct {
	Servers map[string]Server
	Tasks   map[string]Task
}

type Envoyfile struct {
	Servers map[string]interface{}       `yaml:"servers"`
	Tasks   map[string]map[string]string `yaml:"tasks"`
}

func LoadTaskContainer() TaskContainer {
	var taskContainer TaskContainer = TaskContainer{
		Servers: make(map[string]Server),
		Tasks:   make(map[string]Task),
	}

	data, err := os.ReadFile("./Envoyfile")
	if err != nil {
		fmt.Println(err)
		return taskContainer
	}

	var envoyfile Envoyfile

	err = yaml.Unmarshal(data, &envoyfile)
	if err != nil {
		fmt.Println(err)
		return taskContainer
	}

	for k, v := range envoyfile.Servers {
		switch c := v.(type) {
		case string:
			taskContainer.Servers[k] = Server{
				hosts: []string{c},
			}
		case interface{}:
			server := Server{
				hosts: make([]string, len(v.([]interface{}))),
			}
			for index, host := range v.([]interface{}) {
				server.hosts[index] = host.(string)
			}
			taskContainer.Servers[k] = server
		}
	}

	for k, task := range envoyfile.Tasks {
		taskContainer.Tasks[k] = Task{
			name:   k,
			script: task["script"],
		}
	}

	return taskContainer
}

func (c *TaskContainer) GetFirstServer() Server {
	var server Server
	for _, v := range c.Servers {
		server = v
		break
	}
	return server
}

func (c *TaskContainer) HasOneServer() bool {
	return len(c.Servers) == 1
}

func (c *TaskContainer) GetServer(name string) Server {
	return c.Servers[name]
}
