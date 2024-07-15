package main

import "github.com/01-edu/z01"

func main() {
	// str := "abcdefghijklmnopqrstuvwxyz"
	// abcDEFghiJKLmnoPQRstuVWXyz
	count := 0
	caps := false
	for i := 'a'; i <= 'z'; i++ {
		if !caps {
			z01.PrintRune(i)
			count++
		} else {
			z01.PrintRune(i - 32)
			count++
		}

		if count == 3 {
			count = 0
			caps = !caps
		}
	}
	z01.PrintRune('\n')
}
