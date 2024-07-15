package main

import (
	//"os"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	// str2Rune := []rune(str2)
	// position := 0
	var result []rune

	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			if !confirm(rune(str1[i]), result) {
				result = append(result, rune(str1[i]))
				i++
			}
		}
		j++
	}
	fmt.Println(string(result))
}

func confirm(str rune, str1 []rune) bool {
	for _, char := range str1 {
		if str == char {
			return true
		}
	}
	return false
}
