package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg := os.Args[1]
	arg2 := os.Args[2]

	var result string
	for _, char := range arg {
		if !contain(char, result) {
			result += string(char)
		}
	}
	for _, char := range arg2 {
		if !contain(char, result) {
			result += string(char)
		}
	}

	for _, word := range result {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}

func contain(strRune rune, str string) bool {
	for _, word := range str {
		if word == strRune {
			return true
		}
	}
	return false
}
