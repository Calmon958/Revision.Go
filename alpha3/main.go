package main

import "github.com/01-edu/z01"

func main() {
	// alphabet := "abcdefghijklmnopqrstuvwxyz"

	count2 := 0
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		if count < 3 {
			z01.PrintRune(i)
			count++
		} else if count2 < 3 {
			z01.PrintRune(i - 32)
			count2++
		} else {
			count = 0
			count2 = 0
			z01.PrintRune(i)
			count++
		}

	}
	z01.PrintRune('\n')
}
