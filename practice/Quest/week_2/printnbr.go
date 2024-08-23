package main

import "github.com/01-edu/z01"

func PrintNum(nb int) {
	x := '0'
	// if nb == 0 {
	// 	z01.PrintRune('0')
	// 	return
	// }

	for y := 1; y <= nb%10; y++ {
		x++
	}
	for y := -1; y >= nb%10; y-- {
		x++
	}
	if nb/10 != 0 {
		PrintNum(nb / 10)
	}
	z01.PrintRune(x)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}

// func main() {
// 	PrintNbr(-123)
// 	PrintNbr(0)
// 	PrintNbr(123)
// 	z01.PrintRune('\n')
// }
