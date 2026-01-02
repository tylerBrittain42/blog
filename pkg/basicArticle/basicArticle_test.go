package basicArticle

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFilePath(t *testing.T) {

	tests := []struct {
		description    string
		inputDir       string
		inputName      string
		expectedOutput string
		shouldError    bool
		errorMessage   string
	}{
		{
			description:    "no changes needed",
			inputDir:       "usr/",
			inputName:      "art1",
			expectedOutput: "usr/art1.md",
			shouldError:    false,
			errorMessage:   "",
		}, {
			description:    "missing slash",
			inputDir:       "usr",
			inputName:      "art1",
			expectedOutput: "usr/art1.md",
			shouldError:    false,
			errorMessage:   "",
		}, {
			description:    "missing name",
			inputDir:       "usr/",
			inputName:      "",
			expectedOutput: "",
			shouldError:    true,
			errorMessage:   "missing file name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			art := BasicArticle{}
			actualOutput, err := art.GetFilePath(tt.inputDir, tt.inputName)
			if tt.shouldError {
				if err == nil {
					t.Fatal("GetFilePath() did not recieve an error when expected")
				}
				if err.Error() != tt.errorMessage {
					t.Errorf("GetFilePath() recieved error '%v', but wanted error '%v'", err, tt.errorMessage)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetFilePath() recieved response %v when an error was expected", actualOutput)
				}
			} else {
				if err != nil {
					t.Errorf("GetFilePath() recieved an unexpected error: %v", err)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetFilePath() expects %v, got %v", tt.expectedOutput, actualOutput)
				}
			}
		})
	}
}
func TestGetTitle(t *testing.T) {

	tests := []struct {
		description    string
		input          string
		expectedOutput string
		shouldError    bool
		errorMessage   string
	}{
		{
			description:    "strip md extension",
			input:          "testArticle.md",
			expectedOutput: "testArticle",
			shouldError:    false,
			errorMessage:   "",
		}, {
			description:    "strip txt extension",
			input:          "testArticle.txt",
			expectedOutput: "testArticle",
			shouldError:    false,
			errorMessage:   "",
		}, {
			description:    "too many dots",
			input:          "test.Article.txt",
			expectedOutput: "",
			shouldError:    true,
			errorMessage:   "filename must contain a single dot",
		}, {
			description:    "no dots",
			input:          "testArticletxt",
			expectedOutput: "",
			shouldError:    true,
			errorMessage:   "filename must contain a single dot",
		}, {
			description:    "nested filepath",
			input:          "subdir/testArticle.md",
			expectedOutput: "testArticle",
			shouldError:    false,
			errorMessage:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			art := BasicArticle{}
			actualOutput, err := art.GetTitle(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Fatal("GetTitle() did not recieve an error when expected")
				}
				if err.Error() != tt.errorMessage {
					t.Errorf("GetTitle() recieved error '%v', but wanted error '%v'", err, tt.errorMessage)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetTitle() recieved response %v when an error was expected", actualOutput)
				}
			} else {
				if err != nil {
					t.Errorf("GetTitle() recieved an unexpected error: %v", err)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetTitle() expects %v, got %v", tt.expectedOutput, actualOutput)
				}
			}
		})
	}
}

func TestGetContent(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		description    string
		fileName       string
		fileContent    string
		createFile     bool
		expectedOutput string
		shouldError    bool
		errorMessage   string
	}{
		{
			description:    "normal content",
			fileName:       "testArticle.md",
			fileContent:    "Hello, this is an article",
			createFile:     true,
			expectedOutput: "Hello, this is an article",
			shouldError:    false,
			errorMessage:   "",
		}, {
			description:    "no content",
			fileName:       "testArticle.txt",
			fileContent:    "",
			createFile:     true,
			expectedOutput: "",
			shouldError:    true,
			errorMessage:   "no content found",
		}, {
			description:    "only whitespace",
			fileName:       "test.Article.txt",
			fileContent:    "   \n   ",
			createFile:     true,
			expectedOutput: "",
			shouldError:    true,
			errorMessage:   "no content found",
		}, {
			description:    "nested file",
			fileName:       "subdir/nested.md",
			fileContent:    "Nested content",
			createFile:     true,
			expectedOutput: "Nested content",
			shouldError:    false,
			errorMessage:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			art := BasicArticle{}
			path := filepath.Join(tmpDir, tt.fileName)
			if tt.createFile {
				if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
					t.Fatalf("failed to create directory: %v", err)
				}
				if err := os.WriteFile(path, []byte(tt.fileContent), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
			}

			actualOutput, err := art.GetContent(path)
			if tt.shouldError {
				if err == nil {
					t.Fatal("GetContent() did not recieve an error when expected")
				}
				if err.Error() != tt.errorMessage {
					t.Errorf("GetContent() recieved error '%v', but wanted error '%v'", err, tt.errorMessage)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetContent() recieved response %v when an error was expected", actualOutput)
				}
			} else {
				if err != nil {
					t.Errorf("GetContent() recieved an unexpected error: %v", err)
				}
				if tt.expectedOutput != actualOutput {
					t.Errorf("GetContent() expects %v, got %v", tt.expectedOutput, actualOutput)
				}
			}
		})
	}
}
