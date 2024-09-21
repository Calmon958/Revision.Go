package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
printHex := func(b byte){
	hexString := "0123456789abcdef"
	z01.PrintRune(rune(hexString[b>>4]))
	z01.PrintRune(rune(hexString[b&15]))
}
for i, ch := range arr {
	printHex(ch)
	if (i+1)%4 == 0 || i == len(arr) - 1 {
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(' ')
	}
}

for _, c := range arr {
	if c >= 32 && c <= 126 {
		z01.PrintRune(rune(c))
	} else {
		z01.PrintRune('.')
	}
}
z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}



//had help to define limits
