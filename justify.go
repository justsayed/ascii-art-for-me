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
		var lines []string
		if alignment != "--align=justify" {
			lines = renderWord(w, characters)
			lines = applyAlignment(lines, alignment, termWidth)
		} else {
			lines = renderJustify(w, characters, termWidth)
		}
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
			padding = (termWidth - len(line)) / 2
		case "--align=right":
			padding = (termWidth - len(line))
		}
		lines[i] = strings.Repeat(" ", padding) + line
	}

	return lines
}

func renderJustify(word string, characters []string, termWidth int) []string {
	var totalWidth int
	for i := 0; i < len(word); i++ {
		c := int(word[i])
		totalWidth += len(characters[(c-32)*9+1])
	}
	gap := (termWidth - totalWidth) / (len(word) - 1)
	remainder := (termWidth - totalWidth) % (len(word) - 1)

	var result []string
	for row := 1; row < 9; row++ {
		line := ""
		for j := 0; j < len(word); j++ {
			c := int(word[j])
			line += characters[(c-32)*9+row]

			if j < len(word)-1 {
				line += strings.Repeat(" ", gap)

				if j < remainder {
					line += " "
				}
			}
		}
		result = append(result, line)
	}

	return result
}
