package cmd

type RemoteProcessor interface {
	Run(task Task, callback func(string, string, string)) int
}
