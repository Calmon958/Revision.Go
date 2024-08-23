package main

// import "fmt"

func RecursivePower(nb int, power int) int {
	res := 1
	if power == 0 {
		return 1
	} else if power < 0 || power > 30 {
		return 0
	} else {
		res = nb * RecursivePower(nb, power-1)
	}
	return res
}

// func main() {
// 	fmt.Println(RecursivePower(4, 3))
// }
