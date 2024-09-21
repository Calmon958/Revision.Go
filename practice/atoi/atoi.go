package main

import (
	"github.com/01-edu/z01"
)

func main() {
	s := "278462953"
	nb := 0
	var name []rune

	for _, char := range s {
		nb = nb*10 + int(char-'0')
		// fmt.Println(nb)
	}

	// for nb > 0 {
	// 	// var m int
	// 	m := nb % 10
	// 	nb /= 10
	// 	name = append(name, rune(m+'0'))
	// 	fmt.Println(string(name))

	// }

	for i := len(name) - 1; i >= 0; i-- {
		z01.PrintRune(name[i])
	}
}
