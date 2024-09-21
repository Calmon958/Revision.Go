package main

import "fmt"

func AtoiBase2(s, base string) int {
	result := 0

	if !validBase(base) {
		return 0
	} else {
		index := 0
		baseLen := len(base)
		for _, ch := range s {
			for i, v := range base {
				if v == ch {
					index = i
				}
			}
			result = result*baseLen + index
		}
	}

	return result
}

func validBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	baseString := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' {
			return false
		}
		if baseString[ch] {
			return false
		}
		baseString[ch] = true
	}
	return true
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
