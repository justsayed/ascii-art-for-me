package main

import (
	"fmt"
	"os"
	"strings"
)

func justify(alignment, input, banner string) {
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("invalid banner:", banner)
		return
	}
	file, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Invalid file to open:", err)
		return
	}

	// converting the data to split them
	data := string(file)

	characters := strings.Split(data, "\n")
	characters = characters[:len(characters)-1]

	if len(input) < 1 {
		return
	}

	words := strings.Split(input, "\\n")
	if len(words[0]) < 1 {
		words = words[1:]
	}

	
}

func drawArt(words []string, characters []string, numberOfSpaces int) {
	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		for row := 1; row < 9; row++ {
			line := ""
			for j := 0; j < len(word); j++ {
				c := int(word[j])
				line += characters[(c-32)*9+row]
			}
			fmt.Print(strings.Repeat(" ", numberOfSpaces))
			fmt.Println(line)
		}
	}
}
