package main

import (
"os"
"github.com/01-edu/z01"
)

func main () {
if len(os.Args) != 2 {
return
}
arg := os.Args[1]

s:= ""

for i:= len(arg)-1; i>=0; i-- {
if  arg[i] != ' ' {
s = string(arg[i]) + s
} else {
if s !="" {
break
}
}
}
for _,char := range s {
z01.PrintRune(char)
}
z01.PrintRune('\n')
}
