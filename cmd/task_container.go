package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	hosts []string
}

type TaskContainer struct {
	Servers map[string]Server
	Tasks   map[string]Task
}

func LoadTaskContainer() TaskContainer {
	var taskContainer TaskContainer = TaskContainer{
		Servers: make(map[string]Server),
		Tasks:   make(map[string]Task),
	}

	viper.SetConfigName("Envoyfile")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return taskContainer
	}

	for k := range viper.GetStringMap("servers") {
		taskContainer.Servers[k] = Server{
			hosts: viper.GetStringSlice("servers." + k),
		}
	}

	for k := range viper.GetStringMap("tasks") {
		task := viper.GetStringMapString("tasks." + k)
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

func (c *TaskContainer) GetServer(name string) (*Server, error) {
	server, ok := c.Servers[name]

	if !ok {
		return nil, fmt.Errorf("Server [%s] is not defined", name)
	}

	return &server, nil
}

func (c *TaskContainer) GetTask(name string) (*Task, error) {
	task, ok := c.Tasks[name]

	if !ok {
		return nil, fmt.Errorf("Task \"%s\" is not defined", name)
	}

	if task.script == "" {
		return nil, fmt.Errorf("Task \"%s\" has no script", name)
	}

	return &task, nil
}
