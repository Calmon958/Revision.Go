package main

import (
"github.com/01-edu/z01"
"os"
)

func main() {
if len(os.Args) != 3 {
return
}

word1 := os.Args[1]
word2 := os.Args[2]
srune := []rune(word2)
var ans string
position := 0

for _, char := range word1 {
for i:=position; i<=len(srune)-1; i++ {
if char == srune[i] {
ans +=string(srune[i])
position = i + 1
break
}
}
}
if word1 == ans {
for _, char := range ans {
z01.PrintRune(char)
}
z01.PrintRune('\n')

}
}
