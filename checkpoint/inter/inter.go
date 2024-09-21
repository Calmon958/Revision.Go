package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	res := inter(str1, str2)

	fmt.Println(res)
}

func inter(str1, str2 string) string {
	result := ""

	for _, ch := range str1 {
		for _, ch2 := range str2 {
			if ch == ch2 {
				if !Contains(result, ch){
					result+=string(ch)
				}
			}
		}
	}


	return result
}

func Contains(str string, r rune) bool {
	for _, ch := range str{
		if ch == r {
			return true
		}
	}
	return false
}