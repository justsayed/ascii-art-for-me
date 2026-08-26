package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Println("Invalid arguments")
		return
	}
	if len(arg) == 1 {
		banner := "standard"
		input := arg[0]
		fs(input, banner)
	} else if len(arg) == 2 {
		if strings.HasPrefix(arg[0], "--align=") {
			printAligned(arg[0], arg[1], "standard")
		} else {
			fs(arg[0], arg[1])
		}
	} else if len(arg) == 3 {
		if strings.HasPrefix(arg[0], "--align=") {
			printAligned(arg[0], arg[1], arg[2])
		} else {
			fmt.Println("Invalid arguments")
			return
		}
	} else {
		fmt.Println("Invalid arguments")
		return
	}
}
