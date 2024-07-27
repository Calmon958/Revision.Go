package main

import (
"fmt"
)

func CountChar(str string, c rune) int {
count := 0
for _, char := range str{
if char == c {
count++
}
}
return count
}

func main() {
	fmt.Println(CountChar("hello", 'e'))
	fmt.Println(CountChar("	 	", ' '))
	fmt.Println(CountChar("the 7 deadly sins", '7'))
	fmt.Println(CountChar("the 7 deadly sins", 's'))
	fmt.Println(CountChar("the 7 deadly sins", 'd'))
	fmt.Println(CountChar("", 'i'))
}
