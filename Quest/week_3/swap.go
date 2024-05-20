package main

// import "fmt"

func Swap(a *int, b *int) {
	*b, *a = *a, *b
}

// func main() {
// 	a := 0
// 	b := 1
// 	Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
