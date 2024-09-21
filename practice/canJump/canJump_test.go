package main

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		input []uint
		want bool
	}{
		{[]uint{2, 3, 1, 1, 4}, true},
		{[]uint{3, 2, 1, 0, 4}, false},
	}

	t.Run("Check possibility to reach last index", func(t *testing.T){
		for _, tt := range tests{
			got := CanJump(tt.input)
			if got != tt.want {
				t.Errorf("Test Failed. Got=%t Want=%t", got, tt.want)
			}
		}
	})
}