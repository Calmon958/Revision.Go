package main

import (
"github.com/01-edu/z01"
"os"
)

func main() {
if len(os.Args) != 3 {
return
}

str1 := os.Args[1]
str2 := os.Args[2]
word:=""
i, j := 0,0 
for i<len(str1) && j < len(str2) {
if str1[i] == str2[j] {
word += string(str1[i])
i++
}
j++
}
	if word==str1 {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
z01.PrintRune('\n')
}

