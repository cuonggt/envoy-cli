package cmd

import "github.com/gookit/color"

type ConsoleOutput struct{}

func (co ConsoleOutput) Error(message string) {
	color.Style{color.BgRed, color.White}.Println(message)
}

func (co ConsoleOutput) Info(message string) {
	color.Style{color.Green}.Println(message)
}
