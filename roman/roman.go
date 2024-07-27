package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	nb := Atoi(os.Args[1])

	if nb == 0 {
		os.Stdout.WriteString("ERROR")
		os.Stdout.WriteString("\n")
	}

	var roman []string
	var romanSum []string

	for nb > 0 {
		if nb >= 1000 {
			roman = append(roman, "M")
			romanSum = append(romanSum, "M")
			nb -= 1000
		} else if nb >= 900 {
			roman = append(roman, "CM")
			romanSum = append(romanSum, "M-C")
			nb -= 900
		} else if nb >= 500 {
			roman = append(roman, "D")
			romanSum = append(romanSum, "D")
			nb -= 500
		} else if nb >= 400 {
			roman = append(roman, "CD")
			romanSum = append(roman, "D-C")
			nb -= 400
		} else if nb >= 100 {
			roman = append(roman, "C")
			romanSum = append(roman, "C")
			nb -= 100
		} else if nb >= 90 {
			roman = append(roman, "XC")
			romanSum = append(romanSum, "C-X")
			nb -= 90
		} else if nb >= 50 {
			roman = append(roman, "L")
			romanSum = append(romanSum, "L")
			nb -= 50
		} else if nb >= 40 {
			roman = append(roman, "XL")
			romanSum = append(roman, "L-X")
			nb -= 40
		} else if nb >= 10 {
			roman = append(roman, "X")
			romanSum = append(romanSum, "X")
			nb -= 10
		} else if nb >= 9 {
			roman = append(roman, "IX")
			romanSum = append(romanSum, "X-I")
			nb -= 9
		} else if nb >= 5 {
			roman = append(roman, "V")
			romanSum = append(romanSum, "V")
		} else if nb >= 4 {
			roman = append(roman, "IV")
			romanSum = append(romanSum, "V-I")
			nb -= 4
		} else if nb >= 1 {
			roman = append(roman, "I")
			romanSum = append(romanSum, "I")
		}
	}
	for _, char := range romanSum {
		for _, rom := range char{
		z01.PrintRune(rom)
	}
}
}

func Atoi(s string) int {
	num := 0

	for _, char := range s {
		num = num*10 + int(char-'0')
	}
	return num
}
