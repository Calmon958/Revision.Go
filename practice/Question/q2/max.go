package main

import (
"github.com/01-edu/z01"
"fmt"
)


func Max(a []int) int {
var great int

for _, char := range a {
if char>great {
great = char

}
}
return great


z01.PrintRune('\n')
//using sort to find a solution
for i:=0; i<len(a); i++ {
for j:=0; j<len(a); j++ {
if a[j] < a[i] {
a[j], a[i] = a[i], a[j]
}
}
}
return a[0]
}

func main () {
a := []int{23, 123, 1590, 11, 55, 93}
	//max := Max(a)
	fmt.Println(Max(a))
}
