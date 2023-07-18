package cmd

import (
	"github.com/gookit/color"
)

func OutputError(message string) {
	color.Style{color.BgRed, color.White}.Println(message)
}

func OutputInfo(message string) {
	color.Style{color.Green}.Println(message)
}
