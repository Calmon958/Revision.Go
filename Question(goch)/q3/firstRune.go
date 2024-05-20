package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
//s = "Hello"
char :=[]rune(s)
return char[0]
}


func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

