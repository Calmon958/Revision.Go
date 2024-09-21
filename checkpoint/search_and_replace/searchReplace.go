package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	a := os.Args[2]
	b := os.Args[3]

	ans := contains(str, a, b)
	fmt.Println(ans)
}

func contains(str, a, b string) string {
	result := ""

	for _, ch := range str {
		if string(ch) == a {
			result += b
		} else {
			result += string(ch)
		}
	}
	return result
}
