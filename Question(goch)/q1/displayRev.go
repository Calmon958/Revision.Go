package main

import "github.com/01-edu/z01"

func main() {
for i := 'z'; i>='a'; i-- {
if i%2 == 1 {
z01.PrintRune(i -32)
} else {
z01.PrintRune(i)
}
}
z01.PrintRune('\n')
}

