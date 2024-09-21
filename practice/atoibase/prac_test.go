package main

import "testing"


func TestAtoiBase(t *testing.T) {
	tests:= []struct {
		str string
		base string
		expected int
	}{
		{"125", "0123456789", 125},
		{"1A2B", "0123456789ABCDEF", 6699},
	}

	t.Run("Atoi base", func(t *testing.T){
		for _, tc := range tests {
			got := AtoiBase2(tc.str, tc.base)
			if got != tc.expected{
				t.Errorf("Test failed. Got=%d Want=%d", got, tc.expected)
			}
		}
	})
}