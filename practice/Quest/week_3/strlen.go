package main

// import "fmt"

func StrLen(s string) int {
	counter := 0
	for range s {
		counter++
	}
	return counter
}

// func main() {
// 	l := StrLen("Hello World!")
// 	fmt.Println(l)
// }
