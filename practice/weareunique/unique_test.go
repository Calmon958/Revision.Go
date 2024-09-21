package main

import "testing"

func TestUnique(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected int
	}{
		{"foo", "boo", 2},
		{"", "", -1},
		{"abc", "def", 6},
	}

	t.Run("Test for unique character", func(t *testing.T) {
		for _, tc := range tests {
			got := Unique(tc.str1, tc.str2)
			if got != tc.expected {
				t.Errorf("Test failed. Got=%d, Want=%d", got, tc.expected)
			}
		}
	})
}
