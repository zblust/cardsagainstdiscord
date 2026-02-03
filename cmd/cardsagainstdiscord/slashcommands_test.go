package main

import (
	"testing"
)

func TestStringToInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name:     "Valid Discord ID",
			input:    "123456789012345678",
			expected: 123456789012345678,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Large Discord snowflake",
			input:    "987654321098765432",
			expected: 987654321098765432,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringToInt64(tt.input)
			if result != tt.expected {
				t.Errorf("stringToInt64(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
