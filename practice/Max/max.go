package main

import "fmt"

func Max(a []int) int {
var great int 
for _,char := range a {
if char > great {
great = char
}
}
return great

}




func main() {
	a := []int{23, 123, 1, 1111, 55, 9993}
	max := Max(a)
	fmt.Println(max)
}
