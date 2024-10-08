package reverse

import "strings"

func ProcessReverseFileLines(fileContent string) []string {
	lines := strings.Split(fileContent, "\n")
	processedLines := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSuffix(line, "$")
		processedLines = append(processedLines, line)
	}

	return processedLines
}

func ExtractAsciiArt(board [][]string, startIndex, width int) string {
	rows := len(board)
	cols := len(board[0])
	result := ""
	for i := 0; i < rows; i++ {
		for j := startIndex; j < startIndex+width; j++ {
			if j < cols {
				result += board[i][j]
			}
		}
	}
	return result
}
