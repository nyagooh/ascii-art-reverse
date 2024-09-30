package reverse

import (
	"fmt"
	"regexp"
	"strings"
)

// ArtRetriever returns the ASCII art corresponding to the input string using the provided map.
func ArtRetriever(s, c, l string, m map[rune][]string) (string, error) {
	var result strings.Builder

	// Check for empty input or newline character
	if s == "" {
		result.WriteString("")
		return result.String(), nil
	}

	// Check for newline patterns in the input string
	newline := regexp.MustCompile(`\\n`)

	// Determine the number of newline sequences in the string for the case only newlines are passed as input
	numberOfNewLines := len(newline.FindAllStringIndex(s, -1))

	// Replace all newline sequences with actual newline runes
	str := newline.ReplaceAllString(s, "\n")

	// For the case the string only contains repetition of newline runes , write equal number of newlines to the result Builder
	onlyNewLines := regexp.MustCompile(`^(\n)+$`)
	if onlyNewLines.MatchString(str) {
		for i := 0; i < numberOfNewLines; i++ {
			result.WriteString("\n")
		}
		return result.String(), nil
	}
	lines := strings.Split(str, "\n")
	for ind := 0; ind < len(lines); ind++ {
		if lines[ind] == "" {
			// Add an empty line if the input line is empty
			result.WriteString("\n")
		} else {
			// Add ASCII art for each character in the input line
			for j := 0; j < 8; j++ {
				start := 0
				for start < len(lines[ind]) {
					if l == "" {
						// Add the corresponding ASCII art for each character
						for _, char := range lines[ind] {
							if asciiArt, ok := m[char]; ok {
								if c != "" {
									result.WriteString(Colorize(c, asciiArt[j]))
								} else {
									result.WriteString(asciiArt[j])
								}
							} else {
								// Handle invalid non-printable non-ascii characters in the input
								return "", fmt.Errorf("invalid input: %s", string(char))
							}
						}
						break
					} else if strings.HasPrefix(lines[ind][start:], l) {
						// Colorize the substring l
						for _, char := range l {
							if asciiArt, ok := m[char]; ok {
								result.WriteString(Colorize(c, asciiArt[j]))
							} else {
								// Handle invalid non-printable non-ascii characters in the input
								return "", fmt.Errorf("invalid input: %s", string(char))
							}
						}
						start += len(l)
					} else {
						// Add the ASCII art for the current character without coloring
						char := rune(lines[ind][start])
						if asciiArt, ok := m[char]; ok {
							result.WriteString(asciiArt[j])
						} else {
							// Handle invalid non-printable non-ascii characters in the input
							return "", fmt.Errorf("invalid input: %s", string(char))
						}
						start++
					}
				}
				result.WriteString("\n")
			}
		}
	}
	return result.String(), nil
}
