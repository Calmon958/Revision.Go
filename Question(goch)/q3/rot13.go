package main

import (
"os"
"github.com/01-edu/z01"
)

func main(){
if len(os.Args) != 2 {
return
}

for _, char := range os.Args[1] {
if (char >= 'a' && char <= 'm')  || (char >= 'A' && char <= 'M') { 
z01.PrintRune(char + 13)
} else if (char >= 'n' && char <= 'z') || (char >= 'N' && char <= 'Z') {
z01.PrintRune(char - 13)
}
}
z01.PrintRune('\n')
}
