package main

import (
	"fmt"
	"strings"
)

func printAligned(alignment, input, banner string) {
	validAlignments := map[string]bool{
		"--align=center":  true,
		"--align=left":    true,
		"--align=right":   true,
		"--align=justify": true,
	}

	if !validAlignments[alignment] {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	characters, err := loadBanner(banner)
	if err != nil {
		fmt.Println(err)
		return
	}

	termWidth := getTerminalWidth()

	words := strings.Split(input, "\\n")
	if len(words[0]) < 1 {
		words = words[1:]
	}

	for _, w := range words {
		if w == "" {
			fmt.Println()
			continue
		}
		lines := renderWord(w, characters)
		lines = applyAlignment(lines, alignment, termWidth)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

func applyAlignment(lines []string, alignment string, termWidth int) []string {
	for i, line := range lines {
		padding := 0
		switch alignment {
		case "--align=center":
			padding = (termWidth - len(lines)) / 2
		case "--align=right":
			padding = (termWidth - len(lines))
		}
		lines[i] = strings.Repeat(" ", padding) + line
	}

	return lines
}
