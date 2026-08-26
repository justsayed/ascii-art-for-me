package main

import (
	"fmt"
	"os"
	"strings"
)

func fs(input, banner string) {

}

func loadBanner(banner string) ([]string, error) {
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return []string{}, fmt.Errorf("Invalid banner: %s", banner)
	}
	// read the data from the file.
	file, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return []string{}, fmt.Errorf("Invalid file to open: %w", err)
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

	words := strings.Split(word, "\\n")
	var result []string
	for _, w := range words {
		if w == "" {
			result = append(result, "")
			continue
		}
		for row := 1; row < 9; row++ {
			line := renderWord(w, characters)
			for _, arg := range line {
				fmt.Println(arg)
			}
		}
	}
	return result
}
