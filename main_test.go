package main

import (
	"testing"
)

func TestFinish(t *testing.T) {
	tests := []string{
		"000......",
		"XXX......",
		"0..0..0..",
		"X..X..X..",
		".0..0..0.",
		"0...0...0",
		"..0.0.0..",
		"...XXX...",
		"......000",
		"......XXX",
		"..X..X..X",
		".X..X..X.",
	}
	for _, tt := range tests {
		if !checkFinish(tt) {
			t.Errorf("Got it Wrong: %s\n", tt)
		}
	}
}
