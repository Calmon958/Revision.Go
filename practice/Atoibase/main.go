package main

import (
	// "os"
	"fmt"
	"os"
)

// func AtoiBase(s string, base string) int {
// 	if len(base) < 2 || len(base) > 36 {
// 		return 0
// 	}

// 	for i := 0; i < len(base); i++ {
// 		if base[i] == '-' || base[i] == '+' {
// 			return 0
// 		}
// 	}

// 	baseLen := len(base)
// 	result := 0
// 	found := false

// 	for i := 0; i < len(s); i++ {
// 		for j := 0; j < baseLen; j++ {
// 			if s[i] == base[j] {
// 				found = true
// 				result = result*baseLen + j
// 			}
// 		}
// 		if !found {
// 			return 0
// 		}

// 	}
// 	return result
// }

func AtoiBase(str string, base string) int {
	if len(base) < 2 || len(base) > 36 {
		return 0
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return 0
		}
	}

	strRR := []rune{}
	baseRR := []rune{}
	for _, b := range str {
		strRR = append(strRR, indexOf(base, rune(b)))
	}

	for _, b := range base {
		baseRR = append(baseRR, indexOf(base, rune(b)))
	}

	str = string(strRR)
	base = string(baseRR)

	result := 0
	for _, char := range str {
		if char >= '0' && char <= '9' {
			result = result*len(base) + int(char-'0')
		} else {
			num := intIndexOf(base, char)
			if num == -1 {
				os.Exit(1)
			}
			result = result*len(base) + num
		}
	}
	return result
}

func indexOf(str string, r rune) rune {
	for i, cr := range str {
		if cr == r {
			if i <= 9 {
				return rune('0' + i)
			} else {
				diff := i - 10
				return rune('A' + diff)
			}
		}
	}
	return -1
}

func intIndexOf(str string, r rune) int {
	for i, cr := range str {
		if cr == r {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("64F0", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "0123456789choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))

	// math.MaxInt
	fmt.Println(AtoiBase("1728002635214590697", "0123456789A"))
	// fmt.Println(AtoiBase("9223372036854775807", "0123456789A"))
}
