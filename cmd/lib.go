package cmd

import (
	"github.com/gookit/color"
)

func InSlice(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func DisplayOutput(host string, line string) {
	color.Yellow.Printf("[%s]", host)
	color.Printf(": %s", line)
}

func OutputError(message string) {
	color.Style{color.BgRed, color.White}.Println(message)
}

func OutputInfo(message string) {
	color.Style{color.Green}.Println(message)
}
