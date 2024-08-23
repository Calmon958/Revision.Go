package main

import "github.com/01-edu/z01"

func main () {
a := "Hello World!"
for _, char:= range a {
z01.PrintRune(char)
}
z01.PrintRune('\n')
}
