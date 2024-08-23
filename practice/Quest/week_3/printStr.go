package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for i := range s {
		z01.PrintRune(rune([]byte(s)[i]))
	}
}

// func main() {
// 	PrintStr("Hello World!")
// 	z01.PrintRune('\n')
// }
