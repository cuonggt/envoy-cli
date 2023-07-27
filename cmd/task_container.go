package cmd

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

type Server struct {
	Hosts []string
}

type TaskContainer struct {
	Servers map[string]Server
	Tasks   map[string]Task
	Stories map[string][]string
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
		output.Error(fmt.Sprintf("%s", err))
		return taskContainer
	}

	for serverName := range viper.GetStringMap("servers") {
		taskContainer.Servers[serverName] = Server{
			Hosts: viper.GetStringSlice("servers." + serverName),
		}
	}

	for taskName := range viper.GetStringMap("tasks") {
		on := viper.GetStringSlice("tasks." + taskName + ".on")
		hosts := []string{}
		for k, v := range taskContainer.Servers {
			if len(on) == 0 || slices.Contains(on, k) {
				hosts = append(hosts, v.Hosts...)
			}
		}
		taskContainer.Tasks[taskName] = Task{
			Name:     taskName,
			Script:   viper.GetString("tasks." + taskName + ".script"),
			Hosts:    hosts,
			Parallel: viper.GetBool("tasks." + taskName + ".parallel"),
		}
	}

	taskContainer.Stories = viper.GetStringMapStringSlice("stories")

	return taskContainer
}

func (c TaskContainer) GetServer(name string) (*Server, error) {
	server, ok := c.Servers[name]

	if !ok {
		return nil, fmt.Errorf("Server [%s] is not defined", name)
	}

	return &server, nil
}

func (c TaskContainer) GetTask(name string) (*Task, error) {
	task, ok := c.Tasks[name]

	if !ok {
		return nil, fmt.Errorf("Task \"%s\" is not defined", name)
	}

	if task.Script == "" {
		return nil, fmt.Errorf("Task \"%s\" has no script", name)
	}

	return &task, nil
}

func (c TaskContainer) GetStory(name string) []string {
	story, ok := c.Stories[name]
	if !ok {
		return nil
	}
	return story
}
