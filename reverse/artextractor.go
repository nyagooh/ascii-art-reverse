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
