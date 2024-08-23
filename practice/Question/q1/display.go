package main

import "github.com/01-edu/z01"

func main() {
//even characters to be in uppecase odd to be even
for e := 'a'; e <= 'z'; e++ {
if e%2 == 0 {
z01.PrintRune(e-32)
} else {
z01.PrintRune(e)
}
}
z01.PrintRune('\n')


//given data
str := "abcdefghijklmnopqrstuvwxyz"
for i, char := range str {
if i%2 == 0 {
z01.PrintRune(char)
} else {
z01.PrintRune(rune(char) -32)
}
}
z01.PrintRune('\n')
}
