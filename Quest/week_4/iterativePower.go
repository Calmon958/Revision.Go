package main

// import "fmt"

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 || power > 30 {
		return 0
	} else {
		res := 1
		for a := 0; a <= power; a++ {
			res *= nb
		}
		return res
	}
}

// func main() {
// 	fmt.Println(IterativePower(4, 3))
// }
