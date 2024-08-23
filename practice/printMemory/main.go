package main

import (
//"fmt"
"github.com/01-edu/z01"
)

func  PrintMemory(arr [10]byte) {
for i, ch := range arr {
PrintStr(hex(ch))
if (i+1)%4 ==0 || i == len(arr)-1 {
z01.PrintRune('\n')
} else {
z01.PrintRune(' ')
}
}
for _, char := range arr {
if char<33 || char>126 {
PrintStr(".")
}else {
PrintStr(string(char))
}
}
z01.PrintRune('\n')
}

func PrintStr(str string) {
for _, ch := range str{
z01.PrintRune(ch)
}
}

func hex(b byte) string{
hexDigit := "0123456789ABCDEF"
return string(hexDigit[b>>4]) + string(hexDigit[b&0x0f])
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
 
