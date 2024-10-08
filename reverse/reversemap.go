package reverse

import "strings"

func ReverseMapCreator(s string, Map map[string]string) (int, int) {
	minWidth := int(^uint8(0) >> 1)
	maxWidth := 0
	var lines []string
	printableRune := rune(32)
	if strings.ContainsRune(s, '\r') {
		lines = strings.Split(s, "\r\n")
	} else {
		lines = strings.Split(s, "\n")
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" && len(lines[i]) < minWidth {
			minWidth = len(lines[i])
		}
		if len(lines[i]) > maxWidth {
			maxWidth = len(lines[i])
		}
		// If the current line is empty and there are lines left to process
		if i+1 < len(lines) && lines[i] == "" {
			// Create an empty string to concatenate ASCII art lines for the current character
			artLines := ""
			// Iterate over 8 lines (assuming ASCII art is 8 lines tall)
			for j := 0; j < 8; j++ {
				// Concatenate each line of ASCII art to the string
				artLines += lines[i+1+j]
			}
			// Map the printable rune to its corresponding ASCII art
			Map[artLines] = string(printableRune)
			// Increment the printable rune
			printableRune++
		}
	}
	return minWidth, maxWidth
}

func CreateUniversalMap() (map[string]string, int, int, error) {
    thinkertoy, err := ReadTextFile("./banners/thinkertoy.txt")
    if err != nil {
        return nil, 0, 0, err
    }
    standard, err := ReadTextFile("./banners/standard.txt")
    if err != nil {
        return nil, 0, 0, err
    }
    shadow, err := ReadTextFile("./banners/shadow.txt")
    if err != nil {
        return nil, 0, 0, err
    }

    universalMap := make(map[string]string)

    min, _ := ReverseMapCreator(string(thinkertoy), universalMap)
    _, _ = ReverseMapCreator(string(standard), universalMap)
    _, max := ReverseMapCreator(string(shadow), universalMap)

    return universalMap, min, max, nil
}