package main

import (
"os"
"github.com/01-edu/z01"
)

func main() {
//var i rune
//var j rune

if len(os.Args) != 2 {
return
}
/*
arg :=rune(os.Args[1])
for _, char := range arg{
for i := 'a' ; i <= 'z'; i++ {
for j := 'z'; j >= 'a'; j-- {
char[i] = char[j]
}
}
}
z01.PrintRune(char)
}
*/

count := 25
revCount:= -1
for _, char := range os.Args[1] {
if (char>='a' && char <='m') || (char >= 'A' && char <='M') {
z01.PrintRune(char + rune(count))
count=count-2
} else if (char>='n'&& char<='z') || (char >= 'N' && char<='Z') {
z01.PrintRune(char + rune(revCount))
revCount=revCount-2 
}
}
z01.PrintRune('\n')
}

