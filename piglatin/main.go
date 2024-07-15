package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(slatin("apple"))
}

func slatin(str string) (result string) {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for i, v := range str {
		if v >= 'a' && v <= 'z' {
			for j := 0; j < len(vowels); j++ {
				if v == vowels[j] && i == 0 {
					result = str + "ay"
					return
				}
				if v == vowels[j] && i != 0 {
					result = str[i:] + str[:i] + "ay"
					return
				}
				if v != vowels[j]&& i==len(str)-1 {
					os.Stdout.WriteString("Not a vowel")
					return
				}
			}
		}
	}

	return result
}
