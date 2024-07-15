package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	hex := "0123456789abcdef"
	// fmt.Println(len(hex))
	num := atoi(os.Args[1])

	result := ""
	for num > 0 {
		digit := num % 16
		result = string(hex[digit]) + result
		num /= 16
	}
	for _, char := range result {
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
	} else if str[0] == '+' {
		start = 1
	}
	strRune := []rune(str)
	result := 0
	for i := start; i < len(strRune); i++ {
		result = result*10 + int(strRune[i]-'0')
	}
	return result * sign
}

// func itoa(nb int) string {
// 	isNegative := false
// 	if nb < 0 {
// 		isNegative = true
// 		nb *= -1
// 	}
// 	result := ""
// 	for nb > 0 {
// 		digit := nb % 10
// 		result = string(digit+'0') + result
// 		nb /= 10
// 	}
// 	if isNegative {
// 		return "-" + result
// 	} else {
// 		return result
// 	}
// }
