package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]
//	count := 0
	for_, char := range args {
for_, word := range char {
        z01.PrintRune(word)
	}
}
z01.PrintRune('\n')
}

