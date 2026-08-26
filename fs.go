package main

import (
	"fmt"
	"os"
	"strings"
)

func fs(input, banner string) {
	characters, err := loadBanner(banner)
	if err != nil {
		fmt.Println(err)
		return
	}
	words := strings.Split(input, "\\n")
	if len(words[0]) < 1 {
		words = words[1:]
	}
	for _, w := range words {
		if w == "" {
			fmt.Println()
			continue
		}
		for _, line := range renderWord(w, characters) {
			fmt.Println(line)
		}
	}
}

func loadBanner(banner string) ([]string, error) {
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return []string{}, fmt.Errorf("invalid banner: %s", banner)
	}
	// read the data from the file.
	file, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return []string{}, fmt.Errorf("invalid file to open: %w", err)
	}

	// converting the data to split them
	data := string(file)

	characters := strings.Split(data, "\n")
	return characters[:len(characters)-1], nil
}

func renderWord(word string, characters []string) []string {
	if len(word) < 1 {
		return []string{}
	}

	var result []string
	for row := 1; row < 9; row++ {
		line := ""
		for i := 0; i < len(word); i++ {
			c := int(word[i])
			line += characters[(c-32)*9+row]
		}
		result = append(result, line)
	}

	return result
}
