package main

import (
"github.com/01-edu/z01"
"os"
)

func main() {
x := len(os.Args)
if x > 1 {
a := os.Args[x-1]
for _, char := range a {
z01.PrintRune(char)
}
}else 
{
return 
}
z01.PrintRune('\n')
}
