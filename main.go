package main

import (
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 1 {
		banner := "standard"
		input := arg[0]
		fs(input, banner)
	} else if len(arg) == 2 {
		banner := arg[1]
		input := arg[0]
		fs(input, banner)
	} else {
		if len(arg) == 3 {
			alignment := arg[0]
			input := arg[1]
			banner := arg[2]
			printAligned(alignment, input, banner)
		}
	}
}
