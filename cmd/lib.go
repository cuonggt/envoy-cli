package cmd

import (
	"fmt"

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
	fmt.Printf("[%s]: %s", host, line)
}

func OutputError(message string) {
	color.Style{color.BgRed, color.White}.Println(message)
}

func OutputInfo(message string) {
	color.Style{color.Green}.Println(message)
}
