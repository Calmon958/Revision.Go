package main

// import "fmt"

func Fibonacci(index int) int {
	res := 1
	if index < 0 {
		return -1
	} else if index == 0 || index == 1 {
		return index
	} else {
		res = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return res
}

// func main() {
// 	arg1 := 4
// 	fmt.Println(Fibonacci(arg1))
// }
