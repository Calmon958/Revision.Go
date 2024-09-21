package main

import (
	"testing"
)

func TestRPN(t *testing.T) {
	tests := []struct{
		input string
		output int
	}{
		{"1 2 * 3 * 4 +", 10},
		{"     1      3 * 2 -", 1},
	}
	

}