package cmd

type Task struct {
	name   string
	script string
	hosts  []string
}

// func (t *Task) RunOverSSH()
