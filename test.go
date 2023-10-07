package main

import (
	"testing"
)

func TestGetNextVersion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Basic", "1.2.3", "1.2.4"},
		{"No Minor", "1.2", "1.2"},
		{"No Patch", "1", "1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := getNextVersion(test.input)
			if out != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, out)
			}
		})
	}
}
