package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	nb1 := atoi(os.Args[1])
	nb2 := atoi(os.Args[2])

	for nb2 != 0 {
		nb1, nb2 = nb2, nb1%nb2
	}

	ans := itoa(nb1)
	for _, char := range ans {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func atoi(str string) int {
	sign := 1
	start := 0

	if str[0] == '-' {
		sign = -1
		start = 1
	}
	if str[0] == '+' {
		start = 1
	}
	nb := 0
	strRune := []rune(str)
	for i := start; i < len(strRune); i++ {
		nb = nb*10 + int(strRune[i]-'0')
	}
	return sign * nb
}

func itoa(nb int) string {
	negative := false

	if nb < 0 {
		negative = true
		nb *= -1
	}
	result := ""
	for nb > 0 {
		digit := nb % 10
		// result = string(digit +'0') + result
		f := string(digit + '0')
		result = f + result
		nb /= 10
	}
	if negative {
		return "-" + result
	} else {
		return result
	}
}
