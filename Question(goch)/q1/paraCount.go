package main

import (
"github.com/01-edu/z01"
"os"
)

func main (){
b := len(os.Args)
z01.PrintRune('0' + rune(b))
z01.PrintRune('\n')
}

