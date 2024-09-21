package main

import (
	"fmt"
	// "github.com/01-edu/z01"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			// result +=Itoa(count) + string(s[i])
		} else {
			result += Itoa(count) + string(s[i-1])
			count = 1
		}
	}
	result += Itoa(count) + string(s[len(s)-1])
	return result
}

func ZipString2(s string) string {
	str := make(map[rune]int)
	for _, ch := range s {
		str[ch]++
	}
	var result string
	for _, ch := range s {
		if str[ch] != 0 {
			result += string(ch) + Itoa(str[ch])
			str[ch] = 0
		}
	}
	// fmt.Println(result)
	return result

	// for _, ch := range result {
	// 	z01.PrintRune(ch)
	// }
	// z01.PrintRune('\n')
}

func Itoa(nb int) string {
	result := ""
	for nb > 0 {
		digit := nb % 10
		result += string(rune(digit+'0')) + result
		nb /= 10
	}
	return result
}

func main() {
	args := []string{"aaaa", "zzzzzZZZZZZ", "", "ziiiiipMeee", "hhellloTthereYouuunggFelllas"}

	for _, arg := range args {
		fmt.Println(ZipString(arg))
		fmt.Println(ZipString2(arg))
	}
}
