package reverse

import "strings"

func AsciiArtReverser(min, max int, processedLines []string, universalMap map[string]string) string {
	result := ""
	totalLines := len(processedLines)
	currentLine := 0
	for currentLine < totalLines {
		if processedLines[currentLine] == "" {
			if currentLine != totalLines-1 {
				result += "\n"
			}
			currentLine++
		} else {
			board := make([][]string, 8)
			for i := 0; i < 8; i++ {
				board[i] = strings.Split(processedLines[currentLine+i], "")
			}
			currentIndex := 0
			for currentIndex < len(board[0]) {
				for width := min; width <= max; width++ {
					if currentIndex > len(board[0]) {
						break
					}

					asciiArt := ExtractAsciiArt(board, currentIndex, width)
					if char, ok := universalMap[asciiArt]; ok {
						result += char
						currentIndex += width
						break
					}
				}
			}
			result += "\n"
			currentLine += 8
		}
	}
	return result
}
