package reverse

import (
	"errors"
	"os"
	"testing"
)

func TestReadBannerFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test_banner.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	// Write some content to the temporary file
	content := "Test banner content"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	// Close the temporary file
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}
	// Test ReadBannerFile function
	bannerContent, err := ReadBannerFile(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadBannerFile returned error: %v", err)
	}
	if bannerContent != content {
		t.Errorf("ReadBannerFile did not read expected content. Got: %s, Expected: %s", bannerContent, content)
	}
}

func TestMapCreator(t *testing.T) {
	// Test input string
	input := `
         
         
         
         
         
         
         
         
`
	// Call MapCreator function
	result, err := MapCreator(input)

	// Test case: Check if there is an error for space character
	if err == nil {
		t.Error("Expected an error for space character")
	}
	expectedError := errors.New("the bannerfile has been tampered with")
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, Got error: %v", expectedError, err)
	}

	// Test case: Check if the result map is empty for space character
	if len(result) != 0 {
		t.Errorf("Expected empty map for space character, Got: %v", result)
	}
}

func TestArtRetriever(t *testing.T) {
	// Define a sample ASCII art map for testing
	artMap := map[rune][]string{
		'H': {
			" _    _  ",
			"| |  | | ",
			"| |__| | ",
			"|  __  | ",
			"| |  | | ",
			"|_|  |_| ",
			"         ",
			"         ",
		},
		'E': {
			" ______  ",
			"|  ____| ",
			"| |__    ",
			"|  __|   ",
			"| |____  ",
			"|______| ",
			"         ",
			"         ",
		},
	}

	// Create a struct for testcases
	testCases := []struct {
		name     string
		input    string
		color    string
		letters  string
		expected string
		hasError bool
	}{
		// Populate the struct with table driven tests
		{
			"Valid input with color",
			"HE",
			"\033[38;5;196m",
			"",
			"\033[38;5;196m _    _  \u001b[0m\033[38;5;196m ______  \u001b[0m\n\033[38;5;196m| |  | | \u001b[0m\033[38;5;196m|  ____| \u001b[0m\n\033[38;5;196m| |__| | \u001b[0m\033[38;5;196m| |__    \u001b[0m\n\033[38;5;196m|  __  | \u001b[0m\033[38;5;196m|  __|   \u001b[0m\n\033[38;5;196m| |  | | \u001b[0m\033[38;5;196m| |____  \u001b[0m\n\033[38;5;196m|_|  |_| \u001b[0m\033[38;5;196m|______| \u001b[0m\n\033[38;5;196m         \u001b[0m\033[38;5;196m         \u001b[0m\n\033[38;5;196m         \u001b[0m\033[38;5;196m         \u001b[0m\n",
			false,
		},
		{
			"Valid input with color and letters",
			"HE",
			"\033[38;5;196m",
			"H",
			"\033[38;5;196m _    _  \u001b[0m ______  \n\033[38;5;196m| |  | | \u001b[0m|  ____| \n\033[38;5;196m| |__| | \u001b[0m| |__    \n\033[38;5;196m|  __  | \u001b[0m|  __|   \n\033[38;5;196m| |  | | \u001b[0m| |____  \n\033[38;5;196m|_|  |_| \u001b[0m|______| \n\033[38;5;196m         \u001b[0m         \n\033[38;5;196m         \u001b[0m         \n",
			false,
		},
		{
			"Empty input",
			"",
			"",
			"",
			"",
			false,
		},
		{
			"Input with newlines",
			"H\nE",
			"\033[38;5;196m",
			"",
			"\033[38;5;196m _    _  \u001b[0m\n\033[38;5;196m| |  | | \u001b[0m\n\033[38;5;196m| |__| | \u001b[0m\n\033[38;5;196m|  __  | \u001b[0m\n\033[38;5;196m| |  | | \u001b[0m\n\033[38;5;196m|_|  |_| \u001b[0m\n\033[38;5;196m         \u001b[0m\n\033[38;5;196m         \u001b[0m\n\033[38;5;196m ______  \u001b[0m\n\033[38;5;196m|  ____| \u001b[0m\n\033[38;5;196m| |__    \u001b[0m\n\033[38;5;196m|  __|   \u001b[0m\n\033[38;5;196m| |____  \u001b[0m\n\033[38;5;196m|______| \u001b[0m\n\033[38;5;196m         \u001b[0m\n\033[38;5;196m         \u001b[0m\n",
			false,
		},
		{
			"Input with only newlines",
			"\\n\\n",
			"",
			"",
			"\n\n",
			false,
		},
		{
			"Input with invalid character",
			"\r",
			"\033[38;5;196m",
			"",
			"",
			true,
		},
	}

	// Run the tests for each case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ArtRetriever(tc.input, tc.color, tc.letters, artMap)
			if result != tc.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tc.expected, result)
			}
			if (err != nil) != tc.hasError {
				t.Errorf("Expected error: %t, Got error: %v", tc.hasError, err)
			}
		})
	}
}

func TestSetColor(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"Red", "red", "\033[38;5;196m", false},

		// Test hexadecimal colors
		{"3-char Hex", "#abc", "\x1b[38;2;170;187;204m", false},
		{"6-char Hex", "#abcdef", "\x1b[38;2;171;205;239m", false},
		{"Invalid Hex", "#gggggg", "", true},

		// Test RGB colors
		{"RGB", "rgb(10,20,30)", "\x1b[38;2;10;20;30m", false},
		{"RGB out of range", "rgb(300,400,500)", "", true},
		{"Invalid RGB", "rgb(0,20)", "", true},

		// Test invalid input
		{"Invalid color", "invalid", "", true},
	}

	// Run the tests for each case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SetColor(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
			if (err != nil) != tc.hasError {
				t.Errorf("Expected error: %t, Got error: %v", tc.hasError, err)
			}
		})
	}
}

func TestColorize(t *testing.T) {
	color := "\033[38;5;196m"
	message := "Hello, World!"
	expected := "\033[38;5;196mHello, World!\u001b[0m"

	result := Colorize(color, message)
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
