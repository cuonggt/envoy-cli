package cmd

type RemoteProcessor interface {
	Run(task Task)
}
