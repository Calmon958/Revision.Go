package main

import "testing"

func TestSave(t *testing.T) {
	test := []struct{
		input string
		nb int
		want string
	}{
		{"123456789", 3, "123789"},
		{"abcdefghijklmnopqrstuvwyz", 3, "abcghimnostuz"},
		{"Hello there", 0, "Hello there"},
	}

	for _, tc := range test {
		got := Save(tc.input, tc.nb)
		if got != tc.want {
			t.Errorf("Test failed. Got=%s, Want=%s")
		}
	}
}