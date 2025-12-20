package helper

import (
	"testing"
)

func TestIsValid(t *testing.T) {

	tests := []struct {
		name           string
		input          string
		expectedOutput bool
	}{
		{
			name:           "alphanumeric",
			input:          "hi8d3Dd",
			expectedOutput: true,
		}, {
			name:           "symbols",
			input:          "hi8d3Dd.#",
			expectedOutput: false,
		}, {
			name:           "space",
			input:          "hi8d 3Dd",
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOutput, _ := IsAlphaNumeric(tt.input)
			if tt.expectedOutput != actualOutput {
				t.Errorf("IsAlphaNumeric() expects %v, got %v", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestIsAccessible(t *testing.T) {

	tests := []struct {
		name           string
		input          string
		expectedOutput bool
	}{
		{
			name:           "File exists and has content",
			input:          "hi8d3Dd",
			expectedOutput: true,
		}, {
			name:           "File exists, but is empty",
			input:          "hi8d3Dd.#",
			expectedOutput: false,
		}, {
			name:           "File does not exist",
			input:          "hi8d 3Dd",
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOutput, _ := IsAccessible(tt.input)
			if tt.expectedOutput != actualOutput {
				t.Errorf("IsAccessible() expects %v, got %v", tt.expectedOutput, actualOutput)
			}
		})
	}
}
