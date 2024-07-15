package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	if str == "" {
		z01.PrintRune('\n')
		return
	}
	var result string
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			var index int
			if char >= 'a' && char <= 'z' {
				index = int(char) - 'a' + 1
			} else {
				index = int(char) - 'A' + 1
			}
			for i := 0; i < index; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	for _, word := range result {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
