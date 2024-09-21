package main

import "testing"

func TestZipstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"bbbbaa", "b4a2"},
		{"aaaa", "a4"},
	}
	tests2 := []struct {
		input    string
		expected string
	}{
		{"bbbbaa", "4b2a"},
		{"aaaa", "4a"},
	}
	t.Run("test Map application", func(t *testing.T) {
		for _, test := range tests {
			prog := ZipString2(test.input)
			if prog != test.expected {
				t.Errorf("Test failed. Got=%s, Want=%s", prog, test.expected)
			}
		}
	})
	t.Run("test basic application", func(t *testing.T) {
		for _, ch := range tests2 {
			zip := ZipString(ch.input)
			if zip != ch.expected {
				t.Errorf("Test Failed. Got=%s, Want=%s", zip, ch.expected)
			}
		}
	})
}
