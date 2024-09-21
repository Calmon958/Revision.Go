package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConcatAll(t *testing.T) {
	concatenate := []struct {
		input []string
		want  string
	}{
		{[]string{"Hello", "there", "General"}, "HellothereGeneral"},
		{[]string{"Hello Sir", "This is your ID right?"}, "Hello SirThis is your ID right?"},
	}

	concatParams := []struct {
		input []string
		want  string
	}{
		{[]string{"Hello", "there", "General"}, "Hello\nthere\nGeneral"},
		{[]string{"Hello Sir", "This is your ID right?"}, "Hello Sir\nThis is your ID right?"},
	}

	concatAlter := []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{1, 3, 5}, []int{2, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 1, 2}, []int{7, 8, 9, 10}, []int{7, 0, 8, 1, 9, 2, 10}},
	}

	t.Run("Test concatenate", func(t *testing.T) {
		for _, ch := range concatenate {
			got := Concatenate(ch.input)
			if got != ch.want {
				t.Errorf("Test failed. Got=%s, Want=%s", got, ch.want)
			}
		}
	})
	t.Run("Test concatenate", func(t *testing.T) {
		for _, ch := range concatParams {
			got := ConcatParams(ch.input)
			if got != ch.want {
				t.Errorf("Test failed. Got=%s, Want=%s", got, ch.want)
			}
			fmt.Println()
		}
	})
	t.Run("Test concatenate", func(t *testing.T) {
		for _, nb := range concatAlter {
			got := ConcatAlter(nb.input1, nb.input2)
			if !reflect.DeepEqual(got, nb.want) {
				t.Errorf("Test failed. Got=%v, Want=%v", got, nb.want)
			}
		}
	})
}
