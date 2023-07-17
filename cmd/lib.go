package cmd

import "fmt"

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
