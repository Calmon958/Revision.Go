package main

import (
"os"
"github.com/01-edu/z01"
)

func main () {
if len(os.Args) != 3 {
return
}

word1 := os.Args[1]
word2 := os.Args[2]
runeWord2 := []rune(word2)
position := 0
var result string

for _, char := range word1{
for i:=position; i<len(runeWord2); i++ {
if char == runeWord2[i] {
position = i +  1
result += string(runeWord2[i])
break
}
}
}
if word1 == result {
for _, char := range result{
z01.PrintRune(char)
}
}
z01.PrintRune('\n')
}

