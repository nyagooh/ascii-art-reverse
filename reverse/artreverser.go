package reverse

import (
	"strings"
)

// AsciiArtReverser takes the minimum and maximum widths of ASCII art characters, 
// a slice of processed lines from the input file, and a universal map that maps ASCII art to characters.
// It returns the original text by reversing the ASCII art representation.
func AsciiArtReverser(min, max int, processedLines []string, universalMap map[string]string) string {
	result := ""             // Initialize the result string that will hold the reconstructed text
	totalLines := len(processedLines) // Get the total number of lines to process
	currentLine := 0         // Initialize the current line index

	// Loop through each line of processed input
	for currentLine < totalLines {
		// If the current line is empty, add a new line to the result (except for the last line)
		if processedLines[currentLine] == "" {
			if currentLine != totalLines-1 {
				result += "\n"
			}
			currentLine++ // Move to the next line
		} else {
			// Create an 8-line board, where each row represents a line of the ASCII art
			board := make([][]string, 8)
			for i := 0; i < 8; i++ {
				// Split each line into individual characters and store them in the board
				board[i] = strings.Split(processedLines[currentLine+i], "")
			}

			currentIndex := 0 // Initialize the current column index in the board

			// Process the board by extracting ASCII art characters starting at currentIndex
			for currentIndex < len(board[0]) {
				// Try widths from the minimum to the maximum defined width
				for width := min; width <= max; width++ {
					// If the currentIndex exceeds the length of the line, stop the process
					if currentIndex > len(board[0]) {
						break
					}

					// Extract a piece of ASCII art from the board of the given width
					asciiArt := ExtractAsciiArt(board, currentIndex, width)

					// Check if the extracted ASCII art matches a character in the universalMap
					if char, ok := universalMap[asciiArt]; ok {
						// If found, append the corresponding character to the result string
						result += char
						currentIndex += width // Move the index forward by the width of the ASCII art
						break // Stop checking other widths once a match is found
					}
				}
			}

			// After processing 8 lines of ASCII art, add a new line to the result
			result += "\n"
			currentLine += 8 // Move to the next set of 8 lines in the input
		}
	}

	// Return the final reconstructed text
	return result
}
