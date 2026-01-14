package validator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsValid(t *testing.T) {

	tests := []struct {
		description    string
		input          string
		expectedOutput bool
	}{
		{
			description:    "alphanumeric",
			input:          "hi8d3Dd",
			expectedOutput: true,
		}, {
			description:    "symbols",
			input:          "hi8d3Dd.#",
			expectedOutput: false,
		}, {
			description:    "space",
			input:          "hi8d 3Dd",
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actualOutput, _ := IsAlphaNumeric(tt.input)
			if tt.expectedOutput != actualOutput {
				t.Errorf("IsAlphaNumeric() expects %v, got %v", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestIsAccessible(t *testing.T) {

	testDirectory := t.TempDir()
	fileName := "shouldExist.txt"
	filePath := filepath.Join(testDirectory, fileName)
	content := []byte("this is the content of the file")
	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		t.Fatalf("Unable to set up test case. error: %v", err)
	}

	tests := []struct {
		description    string
		name           string
		expectedOutput bool
	}{
		{
			description:    "File exists",
			name:           fileName,
			expectedOutput: true,
		}, {
			description:    "File does not exist",
			name:           "shouldNotExist.txt",
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			filePath := testDirectory + "/" + tt.name
			actualOutput, _ := IsAccessible(filePath)
			if tt.expectedOutput != actualOutput {
				t.Errorf("IsAccessible() expects %v, got %v", tt.expectedOutput, actualOutput)
			}
		})
	}
}
